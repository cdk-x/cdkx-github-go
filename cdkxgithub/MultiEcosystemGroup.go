package cdkxgithub


// Define a group that spans multiple package ecosystems, allowing consolidated pull requests across different ecosystems.
// Experimental.
type MultiEcosystemGroup struct {
	// Schedule preferences for the group.
	// Experimental.
	Schedule *MultiEcosystemGroupSchedule `field:"required" json:"schedule" yaml:"schedule"`
	// Assignees to set on pull requests (additive - merges with ecosystem-level assignees).
	// Experimental.
	Assignees *[]*string `field:"optional" json:"assignees" yaml:"assignees"`
	// Commit message preferences for the group.
	// Experimental.
	CommitMessage *MultiEcosystemGroupCommitMessage `field:"optional" json:"commitMessage" yaml:"commitMessage"`
	// Specify a dependency type to be included in the group.
	// Experimental.
	DependencyType MultiEcosystemGroupDependencyType `field:"optional" json:"dependencyType" yaml:"dependencyType"`
	// Exclude certain dependencies from the group.
	// Experimental.
	ExcludePatterns *[]*string `field:"optional" json:"excludePatterns" yaml:"excludePatterns"`
	// Labels to set on pull requests (additive - merges with ecosystem-level labels).
	// Experimental.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// Associate all pull requests raised for this group with a milestone.
	//
	// You need to specify the numeric identifier of the milestone and not its label.
	// Experimental.
	Milestone *float64 `field:"optional" json:"milestone" yaml:"milestone"`
	// Limit number of open pull requests for version updates.
	// Experimental.
	OpenPullRequestsLimit *float64 `field:"optional" json:"openPullRequestsLimit" yaml:"openPullRequestsLimit"`
	// Pull request branch name preferences for the group.
	// Experimental.
	PullRequestBranchName *MultiEcosystemGroupPullRequestBranchName `field:"optional" json:"pullRequestBranchName" yaml:"pullRequestBranchName"`
	// Specify a different branch for manifest files and for pull requests.
	// Experimental.
	TargetBranch *string `field:"optional" json:"targetBranch" yaml:"targetBranch"`
	// Specify the semantic versioning update types for the group.
	// Experimental.
	UpdateTypes *[]MultiEcosystemGroupUpdateTypes `field:"optional" json:"updateTypes" yaml:"updateTypes"`
}

