package cdkxgithub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk-x/cdkx-github-go/cdkxgithub/jsii"
)

// Prevents a workflow run from failing when a job fails.
//
// Set to true to allow a workflow run to pass when this job fails.
// Experimental.
type NormalJobContinueOnError interface {
	// Experimental.
	Value() interface{}
}

// The jsii proxy struct for NormalJobContinueOnError
type jsiiProxy_NormalJobContinueOnError struct {
	_ byte // padding
}

func (j *jsiiProxy_NormalJobContinueOnError) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func NormalJobContinueOnError_FromBoolean(value *bool) NormalJobContinueOnError {
	_init_.Initialize()

	if err := validateNormalJobContinueOnError_FromBooleanParameters(value); err != nil {
		panic(err)
	}
	var returns NormalJobContinueOnError

	_jsii_.StaticInvoke(
		"@cdk-x/github.NormalJobContinueOnError",
		"fromBoolean",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Experimental.
func NormalJobContinueOnError_FromString(value *string) NormalJobContinueOnError {
	_init_.Initialize()

	if err := validateNormalJobContinueOnError_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns NormalJobContinueOnError

	_jsii_.StaticInvoke(
		"@cdk-x/github.NormalJobContinueOnError",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

