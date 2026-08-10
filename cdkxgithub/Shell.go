package cdkxgithub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk-x/cdkx-github-go/cdkxgithub/jsii"
)

// You can override the default shell settings in the runner's operating system using the shell keyword.
//
// You can use built-in shell keywords, or you can define a custom set of shell options.
// Experimental.
type Shell interface {
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Shell
type jsiiProxy_Shell struct {
	_ byte // padding
}

func (j *jsiiProxy_Shell) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func Shell_FromString(value *string) Shell {
	_init_.Initialize()

	if err := validateShell_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns Shell

	_jsii_.StaticInvoke(
		"@cdk-x/github.Shell",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

