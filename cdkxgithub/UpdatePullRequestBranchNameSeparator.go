package cdkxgithub


// Change separator for PR branch name.
// Experimental.
type UpdatePullRequestBranchNameSeparator string

const (
	// -.
	// Experimental.
	UpdatePullRequestBranchNameSeparator_VALUE_HYPHEN UpdatePullRequestBranchNameSeparator = "VALUE_HYPHEN"
	// _.
	// Experimental.
	UpdatePullRequestBranchNameSeparator_VALUE_UNDERSCORE UpdatePullRequestBranchNameSeparator = "VALUE_UNDERSCORE"
	// /.
	// Experimental.
	UpdatePullRequestBranchNameSeparator_VALUE_FORWARD_SLASH UpdatePullRequestBranchNameSeparator = "VALUE_FORWARD_SLASH"
)

