package cdkxgithub


// Experimental.
type UpdateGroups struct {
	// Use to specify a whether the rules in the group apply to version updates or security updates.
	// Experimental.
	AppliesTo UpdateGroupsAppliesTo `field:"optional" json:"appliesTo" yaml:"appliesTo"`
	// Specify a dependency type to be included in the group.
	// Experimental.
	DependencyType UpdateGroupsDependencyType `field:"optional" json:"dependencyType" yaml:"dependencyType"`
	// Exclude certain dependencies from the group.
	//
	// If a dependency is excluded from a group, Dependabot will continue to raise single pull requests to update the dependency to its latest version.
	// Experimental.
	ExcludePatterns *[]*string `field:"optional" json:"excludePatterns" yaml:"excludePatterns"`
	// Configure how dependencies are grouped within this group.
	// Experimental.
	GroupBy *string `field:"optional" json:"groupBy" yaml:"groupBy"`
	// Define strings of characters that match with a dependency name (or multiple dependency names) to include those dependencies in the group.
	// Experimental.
	Patterns *[]*string `field:"optional" json:"patterns" yaml:"patterns"`
	// Specify the semantic versioning level to include in the group.
	// Experimental.
	UpdateTypes *[]UpdateGroupsUpdateTypes `field:"optional" json:"updateTypes" yaml:"updateTypes"`
}

