package cdkxgithub


// Disable automatic rebasing.
//
// 'auto' is the default and Dependabot will rebase open pull requests when changes are detected. 'disabled' will disable automatic rebasing.
// Experimental.
type UpdateRebaseStrategy string

const (
	// auto.
	// Experimental.
	UpdateRebaseStrategy_AUTO UpdateRebaseStrategy = "AUTO"
	// disabled.
	// Experimental.
	UpdateRebaseStrategy_DISABLED UpdateRebaseStrategy = "DISABLED"
)

