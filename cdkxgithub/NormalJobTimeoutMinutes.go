package cdkxgithub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk-x/cdkx-github-go/cdkxgithub/jsii"
)

// The maximum number of minutes to let a workflow run before GitHub automatically cancels it.
//
// Default: 360.
// Experimental.
type NormalJobTimeoutMinutes interface {
	// Experimental.
	Value() interface{}
}

// The jsii proxy struct for NormalJobTimeoutMinutes
type jsiiProxy_NormalJobTimeoutMinutes struct {
	_ byte // padding
}

func (j *jsiiProxy_NormalJobTimeoutMinutes) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func NormalJobTimeoutMinutes_FromNumber(value *float64) NormalJobTimeoutMinutes {
	_init_.Initialize()

	if err := validateNormalJobTimeoutMinutes_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns NormalJobTimeoutMinutes

	_jsii_.StaticInvoke(
		"@cdk-x/github.NormalJobTimeoutMinutes",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Experimental.
func NormalJobTimeoutMinutes_FromString(value *string) NormalJobTimeoutMinutes {
	_init_.Initialize()

	if err := validateNormalJobTimeoutMinutes_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns NormalJobTimeoutMinutes

	_jsii_.StaticInvoke(
		"@cdk-x/github.NormalJobTimeoutMinutes",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

