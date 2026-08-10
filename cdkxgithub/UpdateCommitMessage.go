package cdkxgithub


// Dependabot attempts to detect your commit message preferences and use similar patterns.
//
// Use this option to specify your preferences explicitly.
// Experimental.
type UpdateCommitMessage struct {
	// Specifies that any prefix is followed by a list of the dependencies updated in the commit.
	// Experimental.
	Include *string `field:"optional" json:"include" yaml:"include"`
	// A prefix for all commit messages.
	//
	// When you specify a prefix for commit messages, GitHub will automatically add a colon between the defined prefix and the commit message provided the defined prefix ends with a letter, number, closing parenthesis, or closing bracket. This means that, for example, if you end the prefix with a whitespace, there will be no colon added between the prefix and the commit message.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// A separate prefix for all commit messages that update dependencies in the Development dependency group.
	//
	// When you specify a value for this option, the prefix is used only for updates to dependencies in the Production dependency group. This is not supported by all package ecosystems.
	// Experimental.
	PrefixDevelopment *string `field:"optional" json:"prefixDevelopment" yaml:"prefixDevelopment"`
}

