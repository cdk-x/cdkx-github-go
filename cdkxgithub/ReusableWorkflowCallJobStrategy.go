package cdkxgithub


// A strategy creates a build matrix for your jobs.
//
// You can define different variations of an environment to run each job in.
// Experimental.
type ReusableWorkflowCallJobStrategy struct {
	// Experimental.
	Matrix interface{} `field:"required" json:"matrix" yaml:"matrix"`
	// When set to true, GitHub cancels all in-progress jobs if any matrix job fails.
	//
	// Default: true.
	// Experimental.
	FailFast interface{} `field:"optional" json:"failFast" yaml:"failFast"`
	// The maximum number of jobs that can run simultaneously when using a matrix job strategy.
	//
	// By default, GitHub will maximize the number of jobs run in parallel depending on the available runners on GitHub-hosted virtual machines.
	// Experimental.
	MaxParallel interface{} `field:"optional" json:"maxParallel" yaml:"maxParallel"`
}

