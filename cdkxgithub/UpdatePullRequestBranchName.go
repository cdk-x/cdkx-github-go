package cdkxgithub


// Pull request branch name preferences.
// Experimental.
type UpdatePullRequestBranchName struct {
	// Change separator for PR branch name.
	// Experimental.
	Separator UpdatePullRequestBranchNameSeparator `field:"required" json:"separator" yaml:"separator"`
}

