TP3: xds-error-propagation
----
* Author(s): anicr7
* Approvers: htuch, markdroth, adisuissa
* Implemented in: <xDS client, ...>
* Last updated: 2024-10-25

## Abstract

This proposal introduces enhancements to the [xDS transport protocol](https://www.envoyproxy.io/docs/envoy/latest/api-docs/xds_protocol), focusing on improved debugging and observability. Specifically, it outlines a mechanism for the control plane to communicate error information directly to xDS clients.

The objective of this proposal is to suggest a way for clients to receive feedback from xDS Management Servers in case of partial/drastic failures without closing any streams or connections.

This proposal includes a new field for each subscribed Resource, called `ResourceError`. This field will provide detailed information for resource specific issues. The client must use this additional field to obtain notification for resources the xDS Management server couldn’t procure and provide necessary notification to the application. 

## Background

A frequent use case for xDS involves client subscription to multiple resources from the xDS management server. Currently, if the xDS management server cannot provide a subset of these resources, the client experiences a complete loss of visibility. This can occur for various reasons, including but not limited to potential permission issues. This lack of granularity in the response can pose significant operational challenges. 

The xDS protocol does have a way for xDS clients to NACK responses back from the xDS Management server. The NACK response also contains an `error_detail` field which the Management Server can use to extract further information about the rejection. But this NACK’ing mechanism is restricted to the client i.e. there is no way for Management Server to actually convey a notification(Ex: unavailability, permission issues etc) regarding some or all the resource requests that are being subscribed by the client. This eventually leads to the client timeouts which, although it conveys an error back to the application, it misses the actual context for the issue. In most cases the xDS Management Servers might not even be accessible for the applications to debug these issues without escalation. 

### Related Proposals:

N/A

## Proposal

The proposal introduces two new fields to the DiscoveryResponse proto. The first, `resource_errors`, will detail Resource-specific errors and be included in both SotW and Incremental xDS protocols. The second, `removed_resources`, will specifically identify resources that no longer exist, addressing a current gap in the SotW protocol's signaling capabilities. 

```textproto
// New Proto Message
// Contains detailed error detail for a single resource
message ResourceError {
 ResourceName resource_name = 1;

 google.rpc.Status error_detail = 2;
}

message DiscoveryResponse {

// The version of the response data.
string version_info = 1;

….
…


 // The control plane instance that sent the response.
  config.core.v3.ControlPlane control_plane = 6;

// NEW_FIELD
// An optional repeated list of error details that the control plane 
// can provide to notify the client of issues for the resources that 
// are subscribed.
// This allows the xDS management server to provide optional 
// notifications in case of unavailability, permissions errors 
// without the client having to wait for the `config fetch` timeout.
repeated ResourceError resource_errors = N;

// NEW FIELD
// For non-existing resources.
repeated ResourceName removed_resources = N;
}

message DeltaDiscoveryResponse {
  …

  // NEW FIELD
  repeated ResourceError resource_errors = N;
}
```

### Protocol Behavior
The client must use this additional field to obtain notification for resources the xDS Management server couldn’t procure. The xDS client should cancel any resource timers once this message is received and convey the error message to the application. With the addition of this field, when a client receives an explicit error or does-not-exist indicator from the management server, it should react the same way it would have if its does-not-exist timer fired. 

The xDS Management server is only expected to return the error message once rather than throughout for future responses. The client is expected to remember the error message until either a new error message is returned or the resource is returned. 

### Wildcard Resources

It is possible to subscribe to all resources by a client using a wildcard or "" resource name. The control plane in this case can provide error details for two different use cases. One when the issue is with the glob itself or later on when the issue is specific to individual resources. 

1. For errors associated with wildcard("" for legacy or "*") and for xDS-TP glob collections; the control plane `error_details` resource name will match the relevant wildcard request("" or "*" or xDS-TP glob collections). This can be used by the control plane to indicate an error with the collection as a whole.
2. For errors associated with specific individual resources that match the glob,
    * The resource name should be the specific resource name associated with the error  OR
    * The control could just not use this mechanism for wildcard subscriptions, because if the client doesn't have permission to access a resource, then it probably shouldn't be considered to match the wildcard subscription to begin with.

## Rationale

The major alternative to this proposal is to use Wrapped Resources by using Resource Containers defined in https://www.envoyproxy.io/docs/envoy/latest/xds/core/v3/resource.proto#xds-core-v3-resource. 

### Wrapped Resources

 xDS resource containers are the default protos used for the Incremental xDS protocol and it's usage in SoTW is controlled via the client feature `xds.config.supports-resource-in-sotw`. 

In this proposal the error information is directly passed as part of the resource field in the Resource Container, using the artifact that the field is a protobuf.Any. This enables us to designate the resource as either a `ResourceError` if the xDS management server encountered problems, or as the actual `Resource` if no errors occurred. 

#### Backward Compatibility

To avoid possible confusion with this behavior, it must be protected with a client feature similar to `supports-resource-in-sotw` called `supports-resource-error-unwrapping`. 

Note: This should also be documented here: https://www.envoyproxy.io/docs/envoy/latest/api/client_features#currently-defined-client-features

#### Why not this approach?

This approach has two major drawbacks compared to the chosen approach:

* This alternative would not work for non-wrapped resources
* It introduces a backwards compatibility issue, which adding a new field wouldn’t have as clients would just ignore them. 

## Implementation

This will probably be implemented in gRPC before Envoy.

## Open issues (if applicable)

### Possible improvements to NACKs

Currently, the xDS protocol does not provide a clear mechanism for partial NACKs i.e. a way for the client to accept some of the resources that are sent by the management server.

There have been few discussions of this behavior mainly in https://github.com/grpc/proposal/blob/master/A46-xds-nack-semantics-improvement.md and https://github.com/envoyproxy/envoy/issues/32880. Currently clients have a general behavior of accepting certain resources even if they are NACKed but this information is not clearly communicated via the xDS protocol.

Eventually using a similar field as this proposal gives us an opportunity to fix this issue in xDS for both the client and the management server, as the DiscoveryRequest could also use the new message `ResourceError` to provide explicit details about the resources not accepted. 
