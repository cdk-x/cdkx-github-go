package cdkxgithub


// Pull request branch name preferences for the group.
// Experimental.
type MultiEcosystemGroupPullRequestBranchName struct {
	// Change separator for PR branch name.
	// Experimental.
	Separator MultiEcosystemGroupPullRequestBranchNameSeparator `field:"required" json:"separator" yaml:"separator"`
}

