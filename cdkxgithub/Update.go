package cdkxgithub


// Experimental.
type Update struct {
	// Package manager to use.
	// Experimental.
	PackageEcosystem *string `field:"required" json:"packageEcosystem" yaml:"packageEcosystem"`
	// Customize which updates are allowed.
	// Experimental.
	Allow *[]*UpdateAllow `field:"optional" json:"allow" yaml:"allow"`
	// Assignees to set on pull requests.
	// Experimental.
	Assignees *[]*string `field:"optional" json:"assignees" yaml:"assignees"`
	// Dependabot attempts to detect your commit message preferences and use similar patterns.
	//
	// Use this option to specify your preferences explicitly.
	// Experimental.
	CommitMessage *UpdateCommitMessage `field:"optional" json:"commitMessage" yaml:"commitMessage"`
	// Defines a cooldown period for dependency updates, allowing updates to be delayed for a configurable number of days.
	//
	// This feature enables users to customize how often Dependabot generates new version updates, offering greater control over update frequency.
	// Experimental.
	Cooldown *UpdateCooldown `field:"optional" json:"cooldown" yaml:"cooldown"`
	// Locations of package manifests.
	// Experimental.
	Directories *[]*string `field:"optional" json:"directories" yaml:"directories"`
	// Location of package manifests.
	// Experimental.
	Directory *string `field:"optional" json:"directory" yaml:"directory"`
	// List of file paths to exclude from dependency updates.
	// Experimental.
	ExcludePaths *[]*string `field:"optional" json:"excludePaths" yaml:"excludePaths"`
	// Configure groups for dependencies.
	//
	// Each 'groups' property is arbitrary will appear in pull request titles and branch names. For example, the code snippet '{"groups": {"NPM dependencies": {"patterns": ["*"]}}}' sets the group name to 'NPM dependencies'.
	// Experimental.
	Groups *map[string]*UpdateGroups `field:"optional" json:"groups" yaml:"groups"`
	// Ignore certain dependencies or versions.
	// Experimental.
	Ignore *[]*UpdateIgnore `field:"optional" json:"ignore" yaml:"ignore"`
	// Allow or deny code execution in manifest files.
	// Experimental.
	InsecureExternalCodeExecution InsecureExternalCodeExecution `field:"optional" json:"insecureExternalCodeExecution" yaml:"insecureExternalCodeExecution"`
	// Labels to set on pull requests.
	// Experimental.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// Associate all pull requests raised for a package manager with a milestone.
	//
	// You need to specify the numeric identifier of the milestone and not its label.
	// Experimental.
	Milestone *float64 `field:"optional" json:"milestone" yaml:"milestone"`
	// String identifier linking this ecosystem to a multi-ecosystem group.
	// Experimental.
	MultiEcosystemGroup *string `field:"optional" json:"multiEcosystemGroup" yaml:"multiEcosystemGroup"`
	// A name for the update configuration.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Limit number of open pull requests for version updates.
	// Experimental.
	OpenPullRequestsLimit *float64 `field:"optional" json:"openPullRequestsLimit" yaml:"openPullRequestsLimit"`
	// Array of dependency patterns to include in a multi-ecosystem group.
	//
	// Required when using multi-ecosystem-group. Use '*' to include all dependencies.
	// Experimental.
	Patterns *[]*string `field:"optional" json:"patterns" yaml:"patterns"`
	// Pull request branch name preferences.
	// Experimental.
	PullRequestBranchName *UpdatePullRequestBranchName `field:"optional" json:"pullRequestBranchName" yaml:"pullRequestBranchName"`
	// Disable automatic rebasing.
	//
	// 'auto' is the default and Dependabot will rebase open pull requests when changes are detected. 'disabled' will disable automatic rebasing.
	// Experimental.
	RebaseStrategy UpdateRebaseStrategy `field:"optional" json:"rebaseStrategy" yaml:"rebaseStrategy"`
	// Experimental.
	Registries interface{} `field:"optional" json:"registries" yaml:"registries"`
	// Schedule preferences.
	// Experimental.
	Schedule *UpdateSchedule `field:"optional" json:"schedule" yaml:"schedule"`
	// Specify a different branch for manifest files and for pull requests.
	// Experimental.
	TargetBranch *string `field:"optional" json:"targetBranch" yaml:"targetBranch"`
	// Tell Dependabot to vendor dependencies when updating them.
	//
	// Don't use this option if you're using 'gomod'.
	// Experimental.
	Vendor *bool `field:"optional" json:"vendor" yaml:"vendor"`
	// How to update manifest version requirements.
	// Experimental.
	VersioningStrategy VersioningStrategy `field:"optional" json:"versioningStrategy" yaml:"versioningStrategy"`
}

