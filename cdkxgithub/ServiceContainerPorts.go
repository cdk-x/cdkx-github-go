package cdkxgithub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk-x/cdkx-github-go/cdkxgithub/jsii"
)

// Experimental.
type ServiceContainerPorts interface {
	// Experimental.
	Value() interface{}
}

// The jsii proxy struct for ServiceContainerPorts
type jsiiProxy_ServiceContainerPorts struct {
	_ byte // padding
}

func (j *jsiiProxy_ServiceContainerPorts) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func ServiceContainerPorts_FromNumber(value *float64) ServiceContainerPorts {
	_init_.Initialize()

	if err := validateServiceContainerPorts_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ServiceContainerPorts

	_jsii_.StaticInvoke(
		"@cdk-x/github.ServiceContainerPorts",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Experimental.
func ServiceContainerPorts_FromString(value *string) ServiceContainerPorts {
	_init_.Initialize()

	if err := validateServiceContainerPorts_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ServiceContainerPorts

	_jsii_.StaticInvoke(
		"@cdk-x/github.ServiceContainerPorts",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

