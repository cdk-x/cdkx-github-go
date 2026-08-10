//go:build !no_runtime_type_checking

package cdkxgithub

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk-x/cdkx-core-go/cdkxcore"
)

func (g *jsiiProxy_GhbWorkflow) validateAddDependencyParameters(target cdkxcore.Resource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GhbWorkflow) validateGetAttParameters(attr *string) error {
	if attr == nil {
		return fmt.Errorf("parameter attr is required, but nil was provided")
	}

	return nil
}

func validateGhbWorkflow_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateGhbWorkflow_IsResourceParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateGhbWorkflow_OfParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateNewGhbWorkflowParameters(scope constructs.Construct, id *string, props *WorkflowProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

