TP2: Dynamically Generated Cacheable xDS Resources
----
* Author(s): markdroth, htuch
* Approver: htuch
* Implemented in: <xDS client, ...>
* Last updated: 2021-08-03

## Abstract

This xRFC proposes a new mechanism to allow xDS servers to
dynamically generate the contents of xDS resources for individual
clients while at the same time preserving cacheability.  Unlike the
context parameter mechanism that is part of the new xDS naming scheme (see
[xRFC TP1](TP1-xds-transport-next.md)), the mechanism described in
this proposal is visible only to the transport protocol layer, not to the
data model layer.  This means that if a resource has a parameter that
affects its contents, that parameter is not part of the resource's name,
which means that any other resources that refer to the resource do not
need to encode the parameter.  Therefore, use of these parameters is
not viral, thus making the mechanism much easier to use.

## Background

There are many use-cases where a control plane may need to dynamically
generate the contents of xDS resources to tailor the resources for
individual clients.  Here are some examples:

- **xDS minor/patch version negotiation.** In this case, each client
  supports a given minor and patch version, and the server may choose to
  use newer API features when talking to clients that support newer versions
  of the API.  (See https://github.com/envoyproxy/envoy/issues/8416 for
  details.)
- **Sharding Cluster resources for scalability.**  In this use-case, there
  are a really large number of clusters, too many for any one client to
  handle.  The clusters are divided into shards, and each client is given
  a dynamically changing assignment of which shards to load.  To support
  this, there needs to be a different variant of the `ClusterCollection`
  resource for each combination of shards that may be assigned to a given
  client.  The shard assignments are generally determined dynamically on
  the client but may change at any time.
- **Sharding endpoints for scalability.**  This is similar to the previous
  case, except that there is a single cluster with a large number of
  endpoints.  The goal is that the xDS server will send different subsets
  of endpoints to different clients, thus avoiding unwanted connections
  when there are large numbers of both servers and clients.  (At Google,
  this is referred to as "subsetting", but it's a different feature than
  the one that Envoy uses that term for.)  In this case, it is desirable
  for the xDS server to determine the subset of endpoints to assign to
  each client.
- **Selecting which cluster to send a client to based on an ACL.**  In this
  use-case, there are two different network paths that can be used to
  access the endpoints: one goes directly to the endpoints, with
  client-side load balancing, and the other goes via a reverse proxy.
  The path that goes directly to the endpoints is faster but is
  access-restricted.  The xDS server needs to check an ACL to determine
  whether a given client is authorized to directly access the endpoints.
  If the client is authorized, it will be sent a `RouteConfiguration`
  pointing to the cluster for those endpoints; otherwise, it will be sent
  a different variant of the `RouteConfiguration` that points to a cluster
  containing the reverse proxy endpoint.
- **Dynamic route selection.**  Every client sends a set of dynamic
  selection parameters (today, conveyed as node metadata).  The server
  has a list of routes to configure, but individual routes in the list
  may be included or excluded based on the client's dynamic selection
  parameters.  Thus, the server needs to generate a slightly different
  version of the `RouteConfiguration` for clients based on the parameters
  they send.  (See
  https://cloud.google.com/traffic-director/docs/configure-advanced-traffic-management#config-filtering-metadata
  for an example.)

The new xDS naming scheme described in [xRFC TP1](TP1-xds-transport-next.md)
provides a mechanism called context parameters, which is intended to move all
parameters that affect resource contents into the resource name, thus adding
cacheability to the xDS ecosystem.  However, this approach means that these
parameters become part of the resource graph on an individual client, which
causes a number of problems:
- Dynamic context parameters are viral, spreading from a given resource
  to all earlier resources in the resource graph.  For example, if
  multiple variants of an EDS resource are needed, there need to be two
  different instances of the resource with different names,
  distinguished by a context parameter.  But because the contents of the
  CDS resource include the name of the corresponding EDS resource name,
  that means that we also need two different versions of the CDS
  resource, also distinguished by the same context parameter.  And then
  we need two different versions of the RDS resource, since that needs
  to refer to the CDS resource.  And then two different versions of the
  LDS resource, which refers to the RDS resource.  This causes a
  combinatorial explosion in the number of resources needed, and it adds
  complexity to xDS servers, which need to construct the right variants
  of every resource and make sure that they refer to each other using
  the right names.
- In the new xDS naming scheme, context parameters are exact-match-only.
  This means that if a control plane wants to provide the same resource
  both with and without a given parameter, it needs to publish two
  versions of the resource, each with a different name, even though the
  contents are the same, which can also cause unnecessarily poor cache
  performance.  For example, in the "dynamic route selection" use-case,
  let's say that every client uses two different dynamic selection
  parameters, `env` (which can have one of the values `prod`, `canary`, or
  `test`) and `version` (which can have one of the values `v1`, `v2`, or
  `v3`).  Now let's say that there is a `RouteConfiguration` with one route
  that should be selected via the parameter `env=prod` and another route that
  should be selected via the parameter `version=v1`. This means that there
  are only four variants of the `RouteConfiguration` resource (`{env!=prod,
  version!=v1}`, `{env=prod, version!=v1}`, `{env!=prod, version=v1}`, and
  `{env=prod, version=v1}`).  However, the exact-match semantics means
  that there will have to be nine different versions of this resource,
  one for each combination of values of the two parameters.

### Related Proposals:
* [xRFC TP1: new xDS naming scheme](TP1-xds-transport-next.md)

## Proposal

This document proposes an alternative approach.  We start with the
observation that resource names are used in two places:

- The **transport protocol** layer, which needs to identify the right
  resource contents to send for a given resource name, often obtaining
  those resource contents from a cache.
- The **resource graph** used on an individual client, where there are a
  set of data model resources that refer to each other by name.  For
  example, a `RouteConfiguration` refers to individual `Cluster` resources
  by name.

The use-cases for dynamic resource selection share one important property
that we can take advantage of.  When multiple variants of a given resource
exist, any given client will only ever use one of those variants at a
given time.  That means that the parameters that affect which variant
of the resource is used are required by the transport protocol, but
they are not required by the client's data model.  (For example, in the
"sharding endpoints for scalability" use-case, different clients may see
different variants of the EDS resource, but once a given client has the
right variant, it will be unique on that client, which means that the
CDS resource does not need to refer to different EDS resource names on
different client.)

It should be noted that caching xDS proxies, unlike "leaf" clients, will
need to track multiple variants of each resource, since a given caching
proxy may be serving clients that need different variants of a given
resource.  However, since caching xDS proxies deal with resources only
at the transport protocol layer, the resource graph layer is
essentially irrelevant in that case.

### Dynamic Parameters

With the above property in mind, this document proposes the following
data structures:
- **Dynamic parameters**, which are a set of key/value pairs that are part
  of the cache key for an xDS resource (in addition to the resource name
  itself).  This provides a mechanism to represent multiple variants of a
  given resource in a cacheable way.  These parameters are used to identify
  the specified resource in the transport protocol, but they are not part of
  the resource name and therefore do not appear as part of the resource graph.
- **Dynamic parameter constraints**, which are a set of criteria that
  can be used to determine whether a set of dynamic parameters matches
  the constraints.  When a client subscribes to a resource, it may
  specify a set of dynamic parameter constraints, which will be used to
  select which variant of the resource will be returned by the server.
  In response to a given subscription request from the client containing
  a set of dynamic parameter constraints, the server will send a
  resource whose dynamic parameters match the dynamic parameter
  constraints in the request.  The client will use the dynamic
  parameters on the resource to determine which of its subscriptions the
  resource is associated with.

Dynamic parameters, unlike context parameters, will not be
exact-match-only.  Dynamic parameter constraints will be able to represent
various types of flexible matching, such as range-based matching (which
will be used for the "xDS minor/patch version negotiation" use-case).
This flexible matching semantic means that there are some cases where
ambiguity can occur; we define a set of best practices below to prevent
these cases from occurring in practice.

#### Matching Ambiguity

Flexible matching means that there may be ambiguities when determining
which resources match which subscriptions.  This section defines the matching
behavior and a set of best practices for deployments to follow to avoid this
kind of ambiguity.

To illustrate where this comes up in practice, it is useful to consider
what happens in transition scenarios, where a deployment initially
groups its clients on a single key but then wants to add a second key.
The second key needs to be added both in the constraints on the server
side and in the clients' configurations, but those two changes cannot
occur atomically.

For example, let's say that the clients are currently categorized by the
parameter `env`, whose value is either `prod` or `test`.  The resource
variants on the server will therefore have the following sets of dymamic
parameters:
- `{env=prod}`
- `{env=test}`

Clients will send one of the following two sets of dynamic parameter
constraints, depending on whether they are `prod` or `test` clients:

```textproto
// For {env=prod}
{key_constraints:[
  {key:"env" value:{
    constraints:[{value:"prod"}]
  }}
]}

// For {env=test}
{key_constraints:[
  {key:"env" value:{
    constraints:[{value:"test"}]
  }}
]}
```

Now the deployment wants to add an additional key called `version`,
whose value will be either `v1` or `v2`, so that it can further subdivide
its clients' configs.

If the new key is added on the clients first, then the clients will
start subscribing with dynamic parameters constraints like the following:

```textproto
// For {env=prod, version=v1}
{key_constraints:[
  {key:"env" value:{
    constraints:[{value:"prod"}]
  }},
  {key:"version" value:{
    constraints:[{value:"v1"}]
  }}
]}
```

The server or cache has to match that set of constraints against the
existing sets of dynamic parameters, which do not specify the `version`
key at all.

Conversely, if the new key is added on the server side first, then the
server will have resource variants with parameters like this:
- `{env=prod, version=v1}`
- `{env=prod, version=v2}`
- `{env=test, version=v1}`
- `{env=test, version=v2}`

But at this point, the clients are continuing to subscribe without
constraints on this new key.  So the server or cache needs to figure out
(e.g.) which of the first two sets of constraints to use for constraints
that require `env` to be `prod` but do not specify `version`.

We address this transition scenario by allowing the set of constraints
for a given key to match any resource variant that does not specify that
key at all.  This allows constraints for new keys to be added on clients
before the corresponding keys are added on the resources on the server, but
it does introduce some additional ambiguity into the matching.  For example,
let's say that the server has the following two variants of a resource:
- `{env=prod}`
- `{env=prod, version=v1}`

Now consider what happens if a client subscribes with the following
constraints:

```textproto
{key_constraints:[
  {key:"env" value:{
    constraints:[{value:"prod"}]
  }},
  {key:"version" value:{
    constraints:[{value:"v1"}]
  }}
]}
```

These constraints can match either of the above variants of the resource.
This situation can be avoided by establishing a best practice that all
variants of a given resource must have the same set of keys.

There is still a possible ambiguity that can occur if a server adds
multiple variants of a new key that clients are not yet sending.
For example, let's say that the server has the following two variants
of a resource:
- `{env=prod, version=v1}`
- `{env=prod, version=v2}`

Consider what happens if a client subscribes with the following constraints:

```textproto
{key_constraints:[
  {key:"env" value:{
    constraints:[{value:"prod"}]
  }}
]}
```

These constraints can match either variant of the above resource.
This can be avoided by establishing a best practice of not adding multiple
variants of a new parameter until clients are sending the new parameter.
However, if this does happen, the cache implementation is free to pick
one of the variants at random.

So, the expected order of changes for this kind of transition would be:
1. Change clients to start sending a constraint for `version=v1`.
2. Add the dynamic parameter `version=v1` to all existing resources.
3. Create new variants of each resource with `version=v2`.
4. Change the desired set of clients to send a constraint for
   `version=v2` instead of `version=v1`.

##### Alternatives Considered

We could avoid much of the matching ambiguity described above by saying that
a set of constraints must specify all keys present on the resource in order
to match.  However, this would mean that if the client starts subscribing
with a constraint for the new key before the corresponding key is added on
the resources on the server, then it will fail to match the existing resources.
In other words, the process would be:

1. Add a variant of all resources on the server side with `version=v1`
   (in addition to all existing dynamic parameters).
2. Change clients to start sending constraints with the new key.
3. When all clients are updated, remove the resource variants that do
   *not* have the new key.

This will effectively require adding new keys on the server side first,
which seems like a large burden on users.  It also seems fairly tricky
for most users to get the exactly correct set of dynamic parameters on
each resource variant, and if they fail to do it right, they will break
their existing configuration.

We also considered having the client add the new constraint but mark it
as optional using an `is_optional` field.  That way, it would match
resources both before and after the new key is added on the server.
However, the `is_optional` field would introduce another type of ambiguity
in matching.  Specifically, let's say that the server has the following
two variants of the resource:
- `{env=prod}`
- `{env=prod, version=v1}`

Now a client subscribes with the following set of constraints:

```textproto
{key_constraints:[
  {key:"env" value:{
    constraints:[{value:"prod"}]
  }},
  {key:"version" value:{
    constraints:[{value:"v1"}]
    is_optional: true
  }}
]}
```

These constraints can match either of the above variants of the resource.
For authoritative servers, this could be addressed by establishing
a best practice of not having two variants of a resource that differ
only by keys that the client will send as optional.  However, this
requires coordination between client and server, and requires
machinery on the client to determine when to set the `is_optional` bit.

Ultimately, although this approach is more semantically precise, it is
also considered too rigid and difficult for users to work with.

#### Matching Behavior and Best Practices

We advise deployments to avoid ambiguity through the following best practices:
- Whenever there are multiple variants of a resource, all variants must
  list the same set of keys.  This allows the server to ignore constraints
  on keys sent by the client that do not affect the choice of variant
  without causing ambiguity in cache misses.
- Servers should not create multiple variants of a parameter that is not yet
  being sent by clients.  If they do, clients that do not send that parameter
  will get one of the variants at random.
- There must be a variant of the resource for every value of a key that is
  going to be present.  For example, if clients will send constraints on the
  `env` key requiring the value to be one of `prod`, `test`, or `qa`, then
  you must have each of those three variants of the resource.
  - Note that servers that can make use of the mechanism described under
    [Server-Specified Constraints](#server-specified-constraints) below
    may be able to optimize this in some cases.  See the "Dynamic Route
    Selection" example below for details.
- For cases where a constraint may match multiple values (e.g., a
  range constraint), the largest possible matching value is preferred.
  This means that caches (both on clients and on caching xDS proxies)
  must attempt to fetch a larger value even if they already have a smaller
  matching value already present in the cache.  For example, let's say
  that a cache contains a variant of a resource with the parameter
  `{shard=3}` and a client subscribes with the following constraints:
  ```textproto
  {key_constraints:[
    {key:"shard" value:{
      constraints:[
        {integer_range_list:[
          {range:{min_value:0 max_value:5}}
        ]}
      ]
    }}
  ]}
  ```
  In this case, the cache must attempt to fetch a resource from the
  authoritative server with that constraint before falling back to
  using the one it already has cached, because the preferred value is
  `{shard=5}`, not `{shard=3}`.
  - Note: This is not an issue for glob collections, because in that case
    all matching variants of the resource will be used.

#### API Changes

The API changes necessary to implement this proposal are in
https://github.com/envoyproxy/envoy/pull/17192.

Dynamic parameter constraints will be represented as follows:

```proto
// A set of dynamic parameter constraints used to select the variant of
// a given resource desired by a client. Clients send a set of
// constraints with each subscription request, and servers respond by
// sending a resource with a matching set of dynamic parameters.
message DynamicParameterConstraints {
  // Constraints for a given key.
  message KeyConstraints {
    message Constraint {
      // A list of one or more integer ranges.
      // A value is considered to match if it falls in any of the ranges.
      message IntegerRangeList {
        // At least one of *min_value* or *max_value* must be set.
        message Range {
          // If specified, value may not be less than this.
          uint64 min_value = 1;

          // If specified, value may not be greater than this.
          uint64 max_value = 2;
        }

        repeated Range range = 1;
      }

      oneof constraint {
        // The key must have this specific value.
        string value = 1;

        // The key's value must be integers and within one of the ranges in this list.
        IntegerRangeList integer_range_list = 2;
      }
    }

    // A list of one or more constraints on the value of the key.
    // All constraints must be met.
    repeated Constraint constraints = 2;
  }

  // One entry per key.
  // Note that if a key has a constraint here, it will place restrictions
  // on the key's value if the key is present on a variant of the resource.
  // However, if a key has a constraint here but is not present on the
  // resource, it will match, regardless of what the constraint says.
  map<string, KeyConstraints> key_constraints = 1;
}
```

The following message will be added to represent a subscription to a
resource by name with associated dynamic parameter constraints:

```proto
// A specification of a resource used when subscribing or unsubscribing.
message ResourceLocator {
  // The resource name to subscribe to.
  string name = 1;

  // A set of constraints used to match against the dynamic parameters on the resource. This
  // allows clients to select between multiple variants of the same resource.
  DynamicParameterConstraints dynamic_parameter_constraints = 2;
}
```

The following new field will be added to `DiscoveryRequest`, to allow clients
to specify constraints when subscribing to a resource:

```proto
  // Alternative to resource_names field that allows specifying cache
  // keys along with each resource name. If this is populated in the
  // first request for a resource type on a stream, resource_names is ignored
  // for all subsequent requests for that resource type on that stream.
  // Clients that populate this field must be able to handle responses
  // from the server where resources are wrapped in a Resource message.
  repeated ResourceLocator resource_locators = 7;
```

Similarly, the following fields will be added to `DeltaDiscoveryRequest`:

```proto
  // Alternative to resource_names_subscribe field that allows specifying cache
  // keys along with each resource name. If this is populated in the
  // first request for a resource type on a stream, resource_names_subscribe
  // and resource_names_unsubscribe are ignored for all subsequent requests
  // for that resource type on that stream.
  repeated ResourceLocator resource_locators_subscribe = 8;

  // Alternative to resource_names_unsubscribe field that allows specifying cache
  // keys along with each resource name. If resource_locators_subscribe is
  // populated in the first request for a resource type on a stream,
  // this field is used instead of resource_named_unsubscribe for all
  // subsequent requests for that resource type on that stream.
  repeated ResourceLocator resource_locators_unsubscribe = 9;
```

The following field will be added to the `Resource` message, to allow the
server to return the dynamic parameters associated with each resource:

```proto
  // Dynamic parameters associated with this resource. To be used by client-side caches
  // (including xDS proxies) when matching subscribed resource locators.
  map<string, string> dynamic_parameters = 8;
```

### Server-Specified Constraints

In the "sharding endpoints" and "selecting cluster based on ACL" use-cases,
the constraints need to be dynamically determined by the xDS server, not by
the client.  To support this, we introduce a new xDS resource type called
`DynamicParametersConstraintsMap`, which looks like this:

```proto
package envoy.config.dynamic_parameters.v3;

message DynamicParameterConstraintsMap {
  // Key is resource type name (e.g., "envoy.config.cluster.v3.Cluster").
  map<string, service.discovery.v3.DynamicParameterConstraints> resource_type_constraints =
      1;
}
```

This resource allows the management server to provide the client with a
set of dynamic parameter constraints to be used for each resource type.

The client will obtain this resource from the server either via ADS or
via a new xDS API called Dynamic Parameter Discovery Service (DPDS):

```proto
package envoy.service.dynamic_parameters.v3;

service DynamicParametersDiscoveryService {
  option (envoy.annotations.resource).type = "envoy.config.dynamic_parameters.v3.DynamicParameters";

  rpc StreamDynamicParameterConstraints(stream discovery.v3.DiscoveryRequest)
      returns (stream discovery.v3.DiscoveryResponse) {
  }

  rpc DeltaDynamicParameterConstraints(stream discovery.v3.DeltaDiscoveryRequest)
      returns (stream discovery.v3.DeltaDiscoveryResponse) {
  }

  rpc FetchDynamicParameterConstraints(discovery.v3.DiscoveryRequest)
      returns (discovery.v3.DiscoveryResponse) {
    option (google.api.http).post = "/v3/discovery:dynamic_parameters";
    option (google.api.http).body = "*";
  }
}
```

Use of this resource type is optional and will be configured locally on
the client (e.g., in the bootstrap file).  The client's configuration
will tell it the name of the DPDS resource to subscribe to and what server
to obtain it from.  When the new xdstp: naming scheme is in use, the client
should be able to configure a different DPDS resource to use for each
authority.

If configured, the client will subscribe to the DPDS resource before
subscribing to any other type of resource.  It will then use the
constraints from the DPDS resource when subscribing to all other types
of resources.

If the client cannot obtain the configured DPDS resource, it will ignore
the failure and request the remaining resources with no additional
constraints.  This will likely result in the client sending a request that
does not include constraints for one of the parameters that is used to
distinguish different variants of the resource, and as mentioned in
the [Matching Behavior and Best
Practices](#matching-behavior-and-best-practices) section above, the
control plane is free to return any variant of the resource in that
case.  However, note that the authoritative server cannot control what
choice is made by caching xDS proxies.

Just like any other xDS resource, a DPDS resource can be updated by the
control plane at any time.  When that happens, the constraints to be
used for a given resource type change, which will cause the client to
unsubscribe from all resources of that type using the old constraints and
then resubscribe to all resources using the new constraints.  Note that
this is an eventually consistent model, but the appropriate use of ADS
or distributed coordination can provide stronger consistency.

#### DPDS Example

For example, let's say that the client is configured such that it will
use the DPDS resource 
`xdstp://xds.example.com/envoy.config.context_params.v3.DynamicContextParameters/my_context_params`
for authority `xds.example.com`.  The client is asked to subscribe to the LDS
resource
`xdstp://xds.example.com/envoy.config.listener.v3.Listener/my_listener`.
The client notices that the it has a DPDS resource configured for the
authority of this resource (xds.example.com), so it will first subscribe
to the DPDS resource.  Let's say that it gets back the following
response:

```textproto
{resource_type_constraints:[
  {key:"envoy.config.listener.v3.Listener" value:{
    key_constraints:[
      {key:"listener_type" value:{
        constraints:[{value:"direct"}]
      }}
    ]
  }}
]}
```

The client would then use the following constraints when subscribing to
the LDS resource:

```textproto
{key_constraints:[
  {key:"listener_type" value:{
    constraints:[{value:"direct"}]
  }}
]}
```

Now let's say that the client later gets an update of the DPDS resource
with the following contents:

```textproto
{resource_type_constraints:[
  {key:"envoy.config.listener.v3.Listener" value:{
    key_constraints:[
      {key:"listener_type" value:{
        constraints:[{value:"via_proxy"}]
      }}
    ]
  }}
]}
```

This changes the constrains to be used for LDS resources.
The client would then send a new request that unsubscribes from the
LDS resource with the old constrains and subscribes to the same resource
with the new constraints.

#### Non-Cacheability of DPDS

Because this mechanism is intended to be used in cases where the server
needs to determine the constraints to be used by the client based on
information not included in the resource locator (e.g., node information,
client IP, or client credentials), the DPDS resource itself is always
non-cacheable.  Servers must always set the [`Resource.do_not_cache`
field](https://github.com/envoyproxy/envoy/blob/371099f4f52f94e60f558561e29ce8852e1091da/api/envoy/service/discovery/v3/discovery.proto#L245)
when sending this resource.  Clients that use a caching xDS proxy for
most of their resources will need to obtain this resource directly
from the authoritative server; when using the new xdstp: naming scheme,
this can be done by using a different authority in the DPDS resource name.

#### Possible DPDS Implementation

One possible way for clients to implement this is by adding a transparent
layer between the transport protocol layer and the data model layer.

Let's say that the transport protocol is handled by an XdsClient object
that handles interaction with the xDS server(s) and takes care of all of
the client-side caching.  The XdsClient object has an API that allows
data model code to register a watcher for a particular resource name,
and the XdsClient will invoke a method on the watcher whenever the
resource is updated.

The DPDS functionality can be added as a transparent "wrapper" of the
XdsClient object:
- When a watch is started on the wrapper object for (e.g.) a Listener
  resource, if use of DPDS is not configured or the resource name is an
  old-style resource name, the watch will just be passed down to the real
  XdsClient without modification.  But if DPDS is in use, then the wrapper
  will use the real XdsClient to start a watch for the DPDS resource.
- When the DPDS resource is returned, the wrapper will use it to determine
  what constraints to use when subscribing to the Listener resource, at
  which point it will start a watch for the Listener resource on the real
  XdsClient using those constraints.  Any updates returned by the Listener
  watcher on the real XdsClient will be passed through to the watcher given
  to the wrapper by the data model code.
- Whenever the DPDS resource gets updated, if the constraints for LDS
  resources have changed, the wrapper will stop the watch for the Listener
  resource on the real XdsClient that was using the old constraints and
  start a new watch using the new constraints.  This change will be
  transparent to the data model code that started the watch on the
  XdsClient wrapper.

#### Envoy-Specific Details

(This section applies only to Envoy, not to other xDS clients like gRPC.)

In Envoy, RTDS is used before DPDS, so the DPDS resource cannot be used to
specify constraints for RTDS resources.

The constraints from the DPDS resource will be used to choose the CDS
resource, which means that DPDS resources will be fetched before CDS
resources.  This means that the DPDS resource itself cannot be fetched
from a cluster obtained via CDS; it must either use a static cluster or
a Google gRPC ApiConfigSource.

### Migrating From Node Metadata

Today, the equivalent of dynamic parameter constraints is node metadata,
which can be used by servers to determine the set of resources to send
for LDS and CDS wildcard subscriptions or to determine the contents of
other resources (e.g., to select individual routes to be included in an
RDS resource).  For transition purposes, this mechanism could continue
to be supported in one of two ways:
1. Direct translation of node metadata to exact-match constraints.  For
   example, if the node metadata contains the entry `env=prod`, this
   would be translated to a constraint `{key_constraints:[{key:"env"
   value:{constraints:[{value:"prod"}]}}]}`.
2. Use the mechanism described under [Server-Specified
   Constraints](#server-specified-constraints) above to convert from node
   metadata to dynamic parameter constraints.  (Note that this mechanism
   requires direct access to the authoritative server, because the
   `DynamicParameterConstraintsMap` resource is not cacheable.)

Any given xDS client may support either or both of these mechanisms.

### Examples

This section shows how the mechanism described in this proposal can be
used to address each of the use-cases identified in the "Background"
section above.

#### xDS Minor/Patch Version Negotiation

The client will send the following dynamic parameter constraints, which
indicate the range of versions that it supports:

```textproto
{key_constraints:[
  {key:"xds.version.minor" value:{
    constraints:[
      {integer_range_list:[
        {range:{min_value:0 max_value:5}}
      ]}
    ]
  }}
]}
```

Let's say that a server has a resource that wants to use a new feature
introduced in version 3.0.5 for clients that support that version.  It will
provide two versions of that resource:
- For clients at version 3.0.5 or higher, a resource with keys
  `{"xds.version.patch"=5}`.
- For clients at version 3.0.4 or lower, there will need to be at least one
  variant of the resource for every possible version range that any client
  may request, all with the exact same content.  For example:
  - `{"xds.version.patch"=4}`
  - `{"xds.version.patch"=3}`
  - `{"xds.version.patch"=2}`
  - `{"xds.version.patch"=1}`
  - `{"xds.version.patch"=0}`

#### Sharding Clusters

In this use-case, the client will have a set of shard ranges determined
by some client-side code, resulting in a pair of dynamic parameter
constraints, one for SRDS and another for CDS.  For example, let's
say that a client should use shard ranges [4-6], [11-15], and [46-90].
The dynamic parameter constraints for SRDS and CDS would be:

```textproto
// For SRDS (single resource).
{key_constraints:[
  {key:"shards" value:{
    constraints:[
      {value:"[4-6],[11-15],[46-90]"}
    ]
  }}
]}

// For CDS (glob collection).
{key_constraints:[
  {key:"shards" value:{
    constraints:[
      {integer_range_list:[
        {range:{min_value:4 max_value:6}},
        {range:{min_value:11 max_value:15}},
        {range:{min_value:46 max_value:90}}
      ]}
    ]
  }}
]}
```

The resulting SRDS resource will tell the client what RDS resources
to fetch.  The server can either generate different resource names for
each variant of the RDS resource, or it can choose to apply the same
constraints to RDS as it uses for SRDS.

#### Sharding Endpoints

The client will initially subscribe to the `DynamicParameterConstraintsMap`
resource to get the dynamic parameter constraints to use for each resource
type.  The server will send back the following DPDS resource:

```textproto
{resource_type_constraints:[
   {key:"envoy.config.cluster.v3.ClusterLoadAssignment" value:{
    key_constraints:[
      {key:"subset_id" value:{
        constraints:[{value:"123"}]
      }}
    ]
  }}
]}
```

This tells the client to use the following dynamic parameter constraints
when subscribing to EDS resources:

```textproto
{key_constraints:[
  {key:"subset_id" value:{
    constraints:[{value:"123"}]
  }}
]}
```

The server can provide a different variant of the EDS resources for each
client, each with different dynamic parameter constraints (e.g., one client
would be told to use `shard_id` 123, while another client might be told to
use `shard_id` 456).

#### Selecting Cluster Based on ACL

The client will initially subscribe to the `DynamicParameterConstraintsMap`
resource to get the dynamic parameters to use for each resource type.  The
server will send back the following DPDS resource:

```textproto
{resource_type_constraints:[
  {key: "envoy.config.cluster.v3.RouteConfiguration" value:{
    key_constraints:[
      {key:"use_proxy" value:{
        constraints:[{value:"true"}]  // or "false", depending on the client
      }}
    ]
  }}
]}
```

This tells the client to send a constraint setting the `use_proxy`
parameter to either true or false when subscribing to the RDS resource.

#### Dynamic Route Selection

Let's say that every client uses two different dynamic selection
parameters, `env` (which can have one of the values `prod`, `canary`, or
`test`) and `version` (which can have one of the values `v1`, `v2`, or `v3`).
Now let's say that there is a RouteConfiguration with one route that should
be selected via the parameter `env=prod` and another route that should be
selected via the parameter `version=v1`. Normally, the server will need to
actually provide the cross-product of these parameter values, so there
will be 9 different variants of the resource, even though there are only
4 unique contents for the resource:

<table>
  <tr>
    <th>Dynamic Parameters on Resource</th>
    <th>Resource Contents</th>
  </tr>

  <tr>
    <td>
      <ul>
      <li><code>{env=canary,version=v2}</code>
      <li><code>{env=test,version=v2}</code>
      <li><code>{env=canary,version=v3}</code>
      <li><code>{env=test,version=v3}</code>
      </ul>
    </td>
    <td>
      <ul>
      <li>does <i>not</i> include the route for <code>env=prod</code>
      <li>does <i>not</i> include the route for <code>version=v1</code>
      </ul>
    </td>
  </tr>

  <tr>
    <td>
      <ul>
      <li><code>{env=prod,version=v2}</code>
      <li><code>{env=prod,version=v3}</code>
      </ul>
    </td>
    <td>
      <ul>
      <li>does include the route for <code>env=prod</code>
      <li>does <i>not</i> include the route for <code>version=v1</code>
      </ul>
    </td>
  </tr>

  <tr>
    <td>
      <ul>
      <li><code>{env=canary,version=v1}</code>
      <li><code>{env=test,version=v1}</code>
      </ul>
    </td>
    <td>
      <ul>
      <li>does <i>not</i> include the route for <code>env=prod</code>
      <li>does include the route for <code>version=v1</code>
      </ul>
    </td>
  </tr>

  <tr>
    <td>
      <ul>
      <li><code>{env=prod,version=v1}</code>
      </ul>
    </td>
    <td>
      <ul>
      <li>does include the route for <code>env=prod</code>
      <li>does include the route for <code>version=v1</code>
      </ul>
    </td>
  </tr>

</table>

Note that a server that does not need to operate with caching xDS proxies
could optimize this by using the mechanism described in [Server-Specified
Constraints](#server-specified-constraints) above.  Specifically, it could use
the DPDS resource to set constraints to minimize the number of variants:

<table>
  <tr>
    <th>Node Metadata</th>
    <th>Dynamic Parameter Constraints from DPDS</th>
    <th>Dynamic Parameters on Resource</th>
  </tr>

  <tr>
    <td>
      <ul>
      <li><code>{env=canary,version=v2}</code>
      <li><code>{env=test,version=v2}</code>
      <li><code>{env=canary,version=v3}</code>
      <li><code>{env=test,version=v3}</code>
      </ul>
    </td>
    <td>
<pre>
{key_constraints:[
  {key:"env" value{
    constraints:[{value:"prod"}]
    invert:true
  }},
  {key:"version" value:{
    constraints:[{value:"v1"}]
    invert:true
  }}
]}
</pre>
    </td>
    <td>
      <code>{env=NOT_prod,version=NOT_v1}</code>
    </td>
  </tr>

  <tr>
    <td>
      <ul>
      <li><code>{env=prod,version=v2}</code>
      <li><code>{env=prod,version=v3}</code>
      </ul>
    </td>
    <td>
<pre>
{key_constraints:[
  {key:"env" value{
    constraints:[{value:"prod"}]
  }},
  {key:"version" value:{
    constraints:[{value:"v1"}]
    invert:true
  }}
]}
</pre>
    </td>
    <td>
      <code>{env=prod,version=NOT_v1}</code>
    </td>
  </tr>

  <tr>
    <td>
      <ul>
      <li><code>{env=canary,version=v1}</code>
      <li><code>{env=test,version=v1}</code>
      </ul>
    </td>
    <td>
<pre>
{key_constraints:[
  {key:"env" value{
    constraints:[{value:"prod"}]
    invert:true
  }},
  {key:"version" value:{
    constraints:[{value:"v1"}]
  }}
]}
</pre>
    </td>
    <td>
      <code>{env=NOT_prod,version=v1}</code>
    </td>
  </tr>

  <tr>
    <td>
      <ul>
      <li><code>{env=prod,version=v1}</code>
      </ul>
    </td>
    <td>
<pre>
{key_constraints:[
  {key:"env" value{
    constraints:[{value:"prod"}]
  }},
  {key:"version" value:{
    constraints:[{value:"v1"}]
  }}
]}
</pre>
    </td>
    <td>
      <code>{env=prod,version=v1}</code>
    </td>
  </tr>
</table>

## Rationale

We considered extending the context parameter mechanism from [xRFC
TP1](TP1-xds-transport-next.md) to support flexible matching semantics,
rather that its current exact-match semantics.  However, that approach had
some down-sides:
- It would not have solved the virality problem described in the "Background"
  section above.
- It would have made the new xDS naming scheme a prerequisite for using
  the dynamic resource selection mechanism.  (The mechanism described in
  this doc is completely independent of the new xDS naming scheme; it can
  be used with the legacy xDS naming scheme as well.)

## Implementation

TBD (Will probably be implemented in gRPC before Envoy)

## Open issues (if applicable)

N/A
