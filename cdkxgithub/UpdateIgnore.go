package cdkxgithub


// Experimental.
type UpdateIgnore struct {
	// Use to ignore updates for dependencies with matching names, optionally using * to match zero or more characters.
	// Experimental.
	DependencyName *string `field:"optional" json:"dependencyName" yaml:"dependencyName"`
	// Use to ignore types of updates.
	//
	// You can combine this with 'dependency-name: "*"' to ignore particular update-types for all dependencies.
	// Experimental.
	UpdateTypes *[]UpdateTypes `field:"optional" json:"updateTypes" yaml:"updateTypes"`
	// Use to ignore specific versions or ranges of versions.
	//
	// If you want to define a range, use the standard pattern for the package manager.
	// Experimental.
	Versions interface{} `field:"optional" json:"versions" yaml:"versions"`
}

