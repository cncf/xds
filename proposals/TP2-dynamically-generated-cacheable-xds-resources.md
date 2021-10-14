TP2: Dynamically Generated Cacheable xDS Resources
----
* Author(s): markdroth, htuch
* Approver: htuch
* Implemented in: <xDS client, ...>
* Last updated: 2021-10-14

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

There are many use-cases where a control plane may need to
dynamically generate the contents of xDS resources to tailor the
resources for individual clients.  One common case is where the
server has a list of routes to configure, but individual routes in
the list may be included or excluded based on the client's dynamic
selection parameters (today, conveyed as node metadata).  Thus,
the server needs to generate a slightly different version of the
`RouteConfiguration` for clients based on the parameters they send.  (See
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
a different client.)

It should be noted that caching xDS proxies, unlike "leaf" clients, will
need to track multiple variants of each resource, since a given caching
proxy may be serving clients that need different variants of a given
resource.  However, since caching xDS proxies deal with resources only
at the transport protocol layer, the resource graph layer is
essentially irrelevant in that case.

### Dynamic Parameters

With the above property in mind, this document proposes the following
data structures:
- **Dynamic parameters**, which are a set of key/value pairs sent by the
  client when subscribing to a resource.
- **Dynamic parameter constraints**, which are a set of criteria that
  can be used to determine whether a set of dynamic parameters matches
  the constraints.  These constraints are part of the cache key for an
  xDS resource (in addition to the resource name itself) on xDS servers,
  xDS clients, and xDS caching proxies.  This provides a mechanism to
  represent multiple variants of a given resource in a cacheable way.

Both of these data structures are used in the xDS transport protocol,
but they are not part of the resource name and therefore do not appear as
part of the resource graph.

When a client subscribes to a resource, it specifies a set of dynamic
parameters.  In response, the server will send a resource whose dynamic
parameter constraints match the dynamic parameters in the subscription
request.  The client will use the dynamic parameter constraints on the
returned resource to determine which of its subscriptions the resource is
associated with.

#### Constraints Representation

Dynamic parameter constraints will be represented in protobuf form as follows:

```proto
message DynamicParameterConstraints {
  // A list of constraints that may be combined with AND or OR semantics.
  message ConstraintList {
    // A constraint for a given key.
    message Constraint {
      message Exists {}
      // The key to match against.
      string key = 1;
      // How to match.
      oneof constraint_type {
        // Matches this exact value.
        string value = 2;
        // Key is present (matches any value except for the key being absent).
        Exists exists = 3;
      }
      // If set to true, the match is inverted -- i.e., the key must NOT
      // match the specified value.
      bool invert = 4;
    }

    enum MatchType {
      // Default value.
      MATCH_TYPE_UNSPECIFIED = 0;
      // Logical AND of constraints.
      MATCH_TYPE_AND = 1;
      // Logical OR of constraints.
      MATCH_TYPE_OR = 2;
    }

    // A list of key/value constraints.
    repeated Constraint constraints = 1;

    // How to match the constraints.
    MatchType match_type = 2;
  }

  // A list of constraint lists. All constraint lists must match (i.e.,
  // logical AND semantics).
  repeated ConstraintList constraints = 1;
}
```

#### Matching Behavior

Note that both xDS servers and clients need to evaluate matching between
a set of dynamic parameters and a set of constraints.  The server does
this when deciding which variant of a given resource to return for a
given subscription request.  When the client receives the resource from
the server, it needs to do the same matching to determine which of its
subscriptions that resource is associated with.  Therefore, the matching
behavior becomes an inherent part of the xDS transport protocol.

(In effect, the resource cache in an xDS client is basically the same
logic as that on an xDS server; the only difference is that in the case
of a client, the resources in the cache come from an xDS stream instead
of from an authoritative database.  Similarly, a caching xDS proxy is
simply an xDS client where the subscriptions come from an incoming xDS
stream.)

For example, let's say that the clients are currently categorized by the
parameter `env`, whose value is either `prod` or `test`.  So any given
client will send one of the following sets of dynamic parameters:
- `{env=prod}`
- `{env=test}`

The resource variants on the server will have the following sets of dynamic
parameter constraints:

```textproto
// For {env=prod}
{constraints:[
  {
    constraints:[{key:"env" value:"prod"}]
    match_type: MATCH_TYPE_AND
  }
]}

// For {env=test}
{constraints:[
  {
    constraints:[{key:"env" value:"test"}]
    match_type: MATCH_TYPE_AND
  }
]}
```

#### Matching Ambiguity

Dynamic parameters, unlike context parameters, will not be
exact-match-only.  Dynamic parameter constraints will be able to represent
certain simple types of flexible matching, such as matching an exact
value or the existance of a key, and simple AND and OR combinations
of constraints.  This flexible matching semantic means that there may be
ambiguities when determining which resources match which subscriptions.
This section defines the matching behavior and a set of best practices for
deployments to follow to avoid this kind of ambiguity.

To illustrate where this comes up in practice, it is useful to consider
what happens in transition scenarios, where a deployment initially
groups its clients on a single key but then wants to add a second key.
The second key needs to be added both in the constraints on the server
side and in the clients' configurations, but those two changes cannot
occur atomically.

Consider the above example where the clients are already divided into
`env=prod` and `env=test`.  Let's say that now the deployment wants to add
an additional key called `version`, whose value will be either `v1` or `v2`,
so that it can further subdivide its clients' configs.

If the new key is added on the server side first, then the server will
have resource variants with constraints like this:

```textproto
// For {env=prod, version=v1}
{constraints:[
  {
    constraints:[
      {key:"env" value:"prod"},
      {key:"version" value:"v1"}
    ]
    match_type: MATCH_TYPE_AND
  }
]}

// For {env=prod, version=v2}
{constraints:[
  {
    constraints:[
      {key:"env" value:"prod"},
      {key:"version" value:"v2"}
    ]
    match_type: MATCH_TYPE_AND
  }
]}
```

But at this point, the clients are continuing to subscribe without
specifying this new key.  So the server or cache would not have any way
to know which of the above variants to use for a subscription specifying
`{env=prod}` but not specifying `version`.

Conversely, if the new key is added on the clients first, then the clients
will start subscribing with dynamic parameters like the following:
- `{env=prod, version=v1}`
- `{env=prod, version=v2}`
- `{env=test, version=v1}`
- `{env=test, version=v2}`

The server or cache has to match those sets of dynamic parameters against
the existing sets of dynamic parameter constraints, which do not specify the
`version` key at all.

We address this transition scenario by allowing a set of constraints
to match a set of dynamic parameters that includes a key that is not
specified by the constraints.  This allows new keys to be added on
clients before the corresponding constraints are added on the resources,
which we expect to be the common case.  (In general, we expect clients
to send a lot of keys that may not actually be used by the server, since
deployments often divide their clients into categories before they have
a need to differentiate the configs for those categories.)

As mentioned above, this approach does introduce the possibility of
matching ambiguity in certain cases, where there may be more than one
variant of a resource that matches the dynamic parameters specified by
the client.  If an xDS transport protocol implementation does encounter
multiple possible matching variants of a resource, its behavior is
undefined.  In the following sections, we evaluate the cases where that
can occur and specify how each one will be addressed.

##### Adding a New Key on the Server First

As stated above, we are optimizing for the case where new keys are added
on clients first, since that is expected to be the common scenario.
However, there may be cases where it is not feasible to have all clients
start sending a new key before the server needs to start making use of
that key.

For example, let's consider the same case as above, where the clients
are initially sending only the `env` key, and the server now wants to
introduce the `version` key.  However, let's say that this is in an
environment where the xDS server is controlled by one team and the clients
are controlled by various other teams, so it's not feasible to force all
clients to start sending the new `version` key all at once.  But there
is one particular client team that is eager to start using the new
`version` key to differentiate the configs of their clients, and they
don't want to wait for all of the other client teams to start sending
the new key.

Consider what happens if the server simply adds a variant of the
resource with the new key:

```textproto
// Existing variant for older clients that are not yet sending the
// version key.
{constraints:[
  {
    constraints:[
      {key:"env" value:"prod"}
    ]
    match_type: MATCH_TYPE_AND
  }
]}

// New variant intended for clients sending the version key.
{constraints:[
  {
    constraints:[
      {key:"env" value:"prod"},
      {key:"version" value:"v1"}
    ]
    match_type: MATCH_TYPE_AND
  }
]}
```

This will work fine for older clients that are not yet sending the
`version` key, because their dynamic parameters will not match the new
variant's constraints.  However, newer clients that are sending dynamic
parameters `{env=prod, version=v1}` will run into ambiguity: those
parameters can match either of the above variants of the resource.

This situation will be avoided by requiring that **all variants of a
given resource must specify constraints for the same set of keys**.

However, in order to make this work for the case where the server starts
sending the constraint on the new key before all clients are sending it,
we provide the `exists` matcher, which will allow the server to specify
a default explicitly for clients that are not yet sending a new key.
In this example, the server would actually have the following two
variants:

```textproto
// Existing variant for older clients that are not yet sending the
// version key.
{constraints:[
  {
    constraints:[
      {key:"env" value:"prod"},
      {key:"version" exists:{} invert:true}
    ]
    match_type: MATCH_TYPE_AND
  }
]}

// New variant for clients sending the version key.
{constraints:[
  {
    constraints:[
      {key:"env" value:"prod"},
      {key:"version" value:"v1"}
    ]
    match_type: MATCH_TYPE_AND
  }
]}
```

This allows maintaining the requirement that all variants of a given
resource have constraints on the same set of keys, while also allowing
the server to explicitly provide a result for older clients that do not
yet send the new key.

##### Variants With Overlapping Constraint Values

There is also a possible ambiguity that can occur if a server provides
multiple variants of a resource whose constraints for a given key
overlap in terms of the values they can match.  For example, let's say
that a server has the following two variants of a resource:

```textproto
// Matches {env=prod} or {env=test}.
{constraints:[
  {
    constraints:[
      {key:"env" value:"prod"},
      {key:"env" value:"test"}
    ]
    match_type: MATCH_TYPE_OR
  }
]}

// Matches {env=qa} or {env=test}.
{constraints:[
  {
    constraints:[
      {key:"env" value:"qa"},
      {key:"env" value:"test"}
    ]
    match_type: MATCH_TYPE_OR
  }
]}
```

Now consider what happens if a client subscribes with dynamic parameters
`{env=test}`.  Those dynamic parameters can match either of the above
variants of the resource.

This situation will be avoided by requiring that **all variants of a given
resource must specify non-overlapping constraints for the same set of keys**.
Control planes must not accept a set of resources that violates this
requirement.

#### Matching Behavior and Best Practices

We advise deployments to avoid ambiguity through the following best practices:
- Whenever there are multiple variants of a resource, all variants must
  list the same set of keys.  This allows the server to ignore constraints
  on keys sent by the client that do not affect the choice of variant
  without causing ambiguity in cache misses.  Servers may use the
  `exists` mechanism to provide backward compatibility for clients that
  are not yet sending a newly added key.
- The constraints on each variant of a given resource must be mutually
  exclusive.  For example, if one variant of a resource matches a given key
  with values "foo" or "bar", and another variant matches that same key
  with values "bar" or "baz", that would cause ambiguity, because both
  variants would match the value "bar".
- There must be a variant of the resource for every value of a key that is
  going to be present.  For example, if clients will send constraints on the
  `env` key requiring the value to be one of `prod`, `test`, or `qa`, then
  you must have each of those three variants of the resource.  (Failure
  to do this will result in the server acting as if the requested
  resource does not exist.)

#### Transport Protocol Changes

The following message will be added to represent a subscription to a
resource by name with associated dynamic parameters:

```proto
// A specification of a resource used when subscribing or unsubscribing.
message ResourceLocator {
  // The resource name to subscribe to.
  string name = 1;

  // A set of dynamic parameters used to match against the dynamic parameter
  // constraints on the resource. This allows clients to select between
  // multiple variants of the same resource.
  map<string, string> dynamic_parameters = 2;
}
```

The following new field will be added to `DiscoveryRequest`, to allow clients
to specify dynamic parameters when subscribing to a resource:

```proto
  // Alternative to resource_names field that allows specifying cache
  // keys along with each resource name. Clients that populate this field
  // must be able to handle responses from the server where resources are
  // wrapped in a Resource message.
  repeated ResourceLocator resource_locators = 7;
```

Similarly, the following fields will be added to `DeltaDiscoveryRequest`:

```proto
  // Alternative to resource_names_subscribe field that allows specifying cache
  // keys along with each resource name.
  repeated ResourceLocator resource_locators_subscribe = 8;

  // Alternative to resource_names_unsubscribe field that allows specifying cache
  // keys along with each resource name.
  repeated ResourceLocator resource_locators_unsubscribe = 9;
```

The following field will be added to the `Resource` message, to allow the
server to return the dynamic parameters associated with each resource:

```proto
  // Dynamic parameter constraints associated with this resource. To be used
  // by client-side caches (including xDS proxies) when matching subscribed
  // resource locators.
  DynamicParameterConstraints dynamic_parameter_constraints = 8;
```

### Migrating From Node Metadata

Today, the equivalent of dynamic parameter constraints is node metadata,
which can be used by servers to determine the set of resources to send
for LDS and CDS wildcard subscriptions or to determine the contents of
other resources (e.g., to select individual routes to be included in an
RDS resource).  For transition purposes, this mechanism can continue
to be supported by the client performing direct translation of node
metadata to exact-match constraints.  For example, if the node metadata
contains the entry `env=prod`, this would be translated to a constraint
`{key_constraints:[{key:"env" value:{constraints:[{value:"prod"}]}}]}`.

Any given xDS client may support either or both of these mechanisms.

### Example

This section shows how the mechanism described in this proposal can be
used to address each the use-case described in the "Background"
section above.

Let's say that every client uses two different dynamic selection
parameters, `env` (which can have one of the values `prod`, `canary`,
or `test`) and `version` (which can have one of the values `v1`, `v2`,
or `v3`).  Now let's say that there is a `RouteConfiguration` with one
route that should be selected via the parameter `env=prod` and another
route that should be selected via the parameter `version=v1`. Without
this design, the server would need to actually provide the cross-product
of these parameter values, so there will be 9 different variants of the
resource, even though there are only 4 unique contents for the resource.
However, this design instead allows the server to provide only the 4
unique variants of the resource, with constraints allowing each client
to get the appropriate one:

<table>
  <tr>
    <th>Dynamic Parameter Constraints on Resource</th>
    <th>Resource Contents</th>
  </tr>

  <tr>
    <td>
<code>{constraints:[
  {
    constraints:[
      {key:"env" value:"prod" invert:true},
      {key:"version" value:"v1" invert:true}
    ]
    match_type: MATCH_TYPE_AND
  }
]}</code>
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
<code>{constraints:[
  {
    constraints:[
      {key:"env" value:"prod"},
      {key:"version" value:"v1" invert:true}
    ]
    match_type: MATCH_TYPE_AND
  }
]}</code>
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
<code>{constraints:[
  {
    constraints:[
      {key:"env" value:"prod" invert:true},
      {key:"version" value:"v1"}
    ]
    match_type: MATCH_TYPE_AND
  }
]}</code>
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
<code>{constraints:[
  {
    constraints:[
      {key:"env" value:"prod"},
      {key:"version" value:"v1"}
    ]
    match_type: MATCH_TYPE_AND
  }
]}</code>
    </td>
    <td>
      <ul>
      <li>does include the route for <code>env=prod</code>
      <li>does include the route for <code>version=v1</code>
      </ul>
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

We could avoid much of the matching ambiguity described above by saying that
a set of constraints must specify all keys present in the subscription
request in order to match.  However, this would mean that if the client
starts subscribing with a new key before the corresponding constraint is
added on the resources on the server, then it will fail to match the
existing resources.  In other words, the process would be:

1. Add a variant of all resources on the server side with a constraint
   for `version=v1` (in addition to all existing constraints).
2. Change clients to start sending the new key.
3. When all clients are updated, remove the resource variants that do
   *not* have the new key.

This will effectively require adding new keys on the server side first,
which seems like a large burden on users.  It also seems fairly tricky
for most users to get the exactly correct set of dynamic parameters on
each resource variant, and if they fail to do it right, they will break
their existing configuration.

Ultimately, although this approach is more semantically precise, it is
also considered too rigid and difficult for users to work with.

## Implementation

TBD (Will probably be implemented in gRPC before Envoy)

## Open issues (if applicable)

N/A
