package cdkxgithub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk-x/cdkx-github-go/cdkxgithub/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk-x/cdkx-core-go/cdkxcore"
	"github.com/cdk-x/cdkx-github-go/cdkxgithub/internal"
)

// Experimental.
type GhbDependabot interface {
	cdkxcore.Resource
	// The cdkx-owned API group and version this Resource's generated shape belongs to, formatted as `<provider>.cdk-x.com/v<N>` (e.g. `"github.cdk-x.com/v1"`) — Kubernetes-CRD-style. The version is owned by cdkx's generator, not the upstream provider's own API/schema version: it increments only when the *generated shape* changes in a breaking way.
	// Experimental.
	ApiVersion() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Props() *DependabotProps
	// The PascalCase resource type identifier, e.g. `"Workflow"`. Combined with `apiVersion`, forms this Resource's stable (group, version, type) identifier — analogous to a Kubernetes GroupVersionKind.
	// Experimental.
	ResourceType() *string
	// Typed convenience wrapper over constructs' own node.addDependency(). Only accepts another Resource (by type) — enforces the cross-stack veto here. Dependencies declared on a Component via the raw node.addDependency() API are resolved to their owning Resource later, in synth() (see Resource.of()), not here.
	// Experimental.
	AddDependency(target cdkxcore.Resource)
	// Returns a lazy reference to one of this Resource's runtime attributes.
	//
	// Returns: an IResolvable that resolves to the synthesized reference form.
	//
	// Example:
	//   const arn = bucket.getAtt('arn');
	//
	// Experimental.
	GetAtt(attr *string) cdkxcore.IResolvable
	// Returns this Resource's own properties, serialized.
	//
	// Subclasses are
	// responsible for walking their own `node.children` to collect and
	// inline any Component descendants under whatever property key makes
	// sense for their output format (e.g. Steps under "steps") — core has
	// no opinion on this.
	//
	// Typed as `Record<string, unknown>` rather than a stricter JSON-tree union:
	// jsii cannot represent a recursive union type across its target languages
	// (Java/.NET/Go) — `unknown` is what jsii falls back to for an untyped
	// value (it compiles to the same `any`-equivalent per language as `any`
	// would), but keeps callers within this TS codebase honest, since they
	// still have to narrow before reading through it (see
	// Synthesizer.synthesizeStack's cast to the internal PropertyValue tree
	// for where that trust boundary is).
	// Experimental.
	ToProperties() *map[string]interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	// Experimental.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for GhbDependabot
type jsiiProxy_GhbDependabot struct {
	internal.Type__cdkxcoreResource
}

func (j *jsiiProxy_GhbDependabot) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GhbDependabot) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GhbDependabot) Props() *DependabotProps {
	var returns *DependabotProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GhbDependabot) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}


// Experimental.
func NewGhbDependabot(scope constructs.Construct, id *string, props *DependabotProps) GhbDependabot {
	_init_.Initialize()

	if err := validateNewGhbDependabotParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_GhbDependabot{}

	_jsii_.Create(
		"@cdk-x/github.GhbDependabot",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewGhbDependabot_Override(g GhbDependabot, scope constructs.Construct, id *string, props *DependabotProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdk-x/github.GhbDependabot",
		[]interface{}{scope, id, props},
		g,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func GhbDependabot_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGhbDependabot_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdk-x/github.GhbDependabot",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Type guard for Resource.
//
// Returns: true if `x` is a Resource, narrowing its type.
//
// Example:
//   if (Resource.isResource(construct)) {
//     // construct is now typed as Resource
//   }
//
// Experimental.
func GhbDependabot_IsResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGhbDependabot_IsResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdk-x/github.GhbDependabot",
		"isResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Resolves the nearest Resource for a given construct — itself, if it already is one, otherwise its nearest Resource ancestor (stopping at the Stack boundary).
//
// Used by synth() to resolve dependencies declared
// on a Component to the Resource that actually owns it.
//
// Returns: the owning Resource.
//
// Example:
//   const resource = Resource.of(someComponent);
//
// Experimental.
func GhbDependabot_Of(construct constructs.IConstruct) cdkxcore.Resource {
	_init_.Initialize()

	if err := validateGhbDependabot_OfParameters(construct); err != nil {
		panic(err)
	}
	var returns cdkxcore.Resource

	_jsii_.StaticInvoke(
		"@cdk-x/github.GhbDependabot",
		"of",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GhbDependabot) AddDependency(target cdkxcore.Resource) {
	if err := g.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addDependency",
		[]interface{}{target},
	)
}

func (g *jsiiProxy_GhbDependabot) GetAtt(attr *string) cdkxcore.IResolvable {
	if err := g.validateGetAttParameters(attr); err != nil {
		panic(err)
	}
	var returns cdkxcore.IResolvable

	_jsii_.Invoke(
		g,
		"getAtt",
		[]interface{}{attr},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GhbDependabot) ToProperties() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"toProperties",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GhbDependabot) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GhbDependabot) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		g,
		"with",
		args,
		&returns,
	)

	return returns
}

