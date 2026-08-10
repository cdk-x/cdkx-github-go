package cdkxgithub


// Commit message preferences for the group.
// Experimental.
type MultiEcosystemGroupCommitMessage struct {
	// Specifies that any prefix is followed by a list of the dependencies updated in the commit.
	// Experimental.
	Include *string `field:"optional" json:"include" yaml:"include"`
	// A prefix for all commit messages.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// A separate prefix for all commit messages that update dependencies in the Development dependency group.
	// Experimental.
	PrefixDevelopment *string `field:"optional" json:"prefixDevelopment" yaml:"prefixDevelopment"`
}

