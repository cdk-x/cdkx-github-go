package cdkxgithub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk-x/cdkx-github-go/cdkxgithub/jsii"
)

// Prevents a job from failing when a step fails.
//
// Set to true to allow a job to pass when this step fails.
// Experimental.
type StepContinueOnError interface {
	// Experimental.
	Value() interface{}
}

// The jsii proxy struct for StepContinueOnError
type jsiiProxy_StepContinueOnError struct {
	_ byte // padding
}

func (j *jsiiProxy_StepContinueOnError) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func StepContinueOnError_FromBoolean(value *bool) StepContinueOnError {
	_init_.Initialize()

	if err := validateStepContinueOnError_FromBooleanParameters(value); err != nil {
		panic(err)
	}
	var returns StepContinueOnError

	_jsii_.StaticInvoke(
		"@cdk-x/github.StepContinueOnError",
		"fromBoolean",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Experimental.
func StepContinueOnError_FromString(value *string) StepContinueOnError {
	_init_.Initialize()

	if err := validateStepContinueOnError_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns StepContinueOnError

	_jsii_.StaticInvoke(
		"@cdk-x/github.StepContinueOnError",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

