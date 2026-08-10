package cdkxgithub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk-x/cdkx-github-go/cdkxgithub/jsii"
)

// The maximum number of minutes to run the step before killing the process.
// Experimental.
type StepTimeoutMinutes interface {
	// Experimental.
	Value() interface{}
}

// The jsii proxy struct for StepTimeoutMinutes
type jsiiProxy_StepTimeoutMinutes struct {
	_ byte // padding
}

func (j *jsiiProxy_StepTimeoutMinutes) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func StepTimeoutMinutes_FromNumber(value *float64) StepTimeoutMinutes {
	_init_.Initialize()

	if err := validateStepTimeoutMinutes_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns StepTimeoutMinutes

	_jsii_.StaticInvoke(
		"@cdk-x/github.StepTimeoutMinutes",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Experimental.
func StepTimeoutMinutes_FromString(value *string) StepTimeoutMinutes {
	_init_.Initialize()

	if err := validateStepTimeoutMinutes_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns StepTimeoutMinutes

	_jsii_.StaticInvoke(
		"@cdk-x/github.StepTimeoutMinutes",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

