package cdkxgithub


// Experimental.
type UpdateAllow struct {
	// Experimental.
	DependencyName *string `field:"optional" json:"dependencyName" yaml:"dependencyName"`
	// Experimental.
	DependencyType DependencyType `field:"optional" json:"dependencyType" yaml:"dependencyType"`
	// Use to allow specific types of updates.
	//
	// You can combine this with 'dependency-name: "*"' to allow particular update-types for all dependencies.
	// Experimental.
	UpdateTypes *[]UpdateTypes `field:"optional" json:"updateTypes" yaml:"updateTypes"`
}

