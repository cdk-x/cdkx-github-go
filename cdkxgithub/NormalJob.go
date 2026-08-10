package cdkxgithub


// Each job must have an id to associate with the job.
//
// The key job_id is a string and its value is a map of the job's configuration data. You must replace <job_id> with a string that is unique to the jobs object. The <job_id> must start with a letter or _ and contain only alphanumeric characters, -, or _.
// Experimental.
type NormalJob struct {
	// The type of machine to run the job on.
	//
	// The machine can be either a GitHub-hosted runner, or a self-hosted runner.
	// Experimental.
	RunsOn interface{} `field:"required" json:"runsOn" yaml:"runsOn"`
	// Concurrency ensures that only a single job or workflow using the same concurrency group will run at a time.
	//
	// A concurrency group can be any string or expression. The expression can use any context except for the secrets context.
	// You can also specify concurrency at the workflow level.
	// When a concurrent job or workflow is queued, if another job or workflow using the same concurrency group in the repository is in progress, the queued job or workflow will be pending. By default any previously pending job or workflow in the concurrency group will be canceled; this behavior can be changed with `queue`. To also cancel any currently running job or workflow in the same concurrency group, specify cancel-in-progress: true.
	// Experimental.
	Concurrency interface{} `field:"optional" json:"concurrency" yaml:"concurrency"`
	// A container to run any steps in a job that don't already specify a container.
	//
	// If you have steps that use both script and container actions, the container actions will run as sibling containers on the same network with the same volume mounts.
	// If you do not set a container, all steps will run directly on the host specified by runs-on unless a step refers to an action configured to run in a container.
	// Experimental.
	Container interface{} `field:"optional" json:"container" yaml:"container"`
	// Prevents a workflow run from failing when a job fails.
	//
	// Set to true to allow a workflow run to pass when this job fails.
	// Experimental.
	ContinueOnError NormalJobContinueOnError `field:"optional" json:"continueOnError" yaml:"continueOnError"`
	// A map of default settings that will apply to all steps in the job.
	// Experimental.
	Defaults *Defaults `field:"optional" json:"defaults" yaml:"defaults"`
	// A map of environment variables that are available to all steps in the job.
	// Experimental.
	Env interface{} `field:"optional" json:"env" yaml:"env"`
	// The environment that the job references.
	// Experimental.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// You can use the if conditional to prevent a job from running unless a condition is met.
	//
	// You can use any supported context and expression to create a conditional.
	// Expressions in an if conditional do not require the ${{ }} syntax. For more information, see https://help.github.com/en/articles/contexts-and-expression-syntax-for-github-actions.
	// Experimental.
	If interface{} `field:"optional" json:"if" yaml:"if"`
	// The name of the job displayed on GitHub.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Experimental.
	Needs interface{} `field:"optional" json:"needs" yaml:"needs"`
	// A map of outputs for a job.
	//
	// Job outputs are available to all downstream jobs that depend on this job.
	// Experimental.
	Outputs *map[string]*string `field:"optional" json:"outputs" yaml:"outputs"`
	// Experimental.
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// Additional containers to host services for a job in a workflow.
	//
	// These are useful for creating databases or cache services like redis. The runner on the virtual machine will automatically create a network and manage the life cycle of the service containers.
	// When you use a service container for a job or your step uses container actions, you don't need to set port information to access the service. Docker automatically exposes all ports between containers on the same network.
	// When both the job and the action run in a container, you can directly reference the container by its hostname. The hostname is automatically mapped to the service name.
	// When a step does not use a container action, you must access the service using localhost and bind the ports.
	// Experimental.
	Services *map[string]*ServiceContainer `field:"optional" json:"services" yaml:"services"`
	// Experimental.
	Snapshot interface{} `field:"optional" json:"snapshot" yaml:"snapshot"`
	// A job contains a sequence of tasks called steps.
	//
	// Steps can run commands, run setup tasks, or run an action in your repository, a public repository, or an action published in a Docker registry. Not all steps run actions, but all actions run as a step. Each step runs in its own process in the virtual environment and has access to the workspace and filesystem. Because steps run in their own process, changes to environment variables are not preserved between steps. GitHub provides built-in steps to set up and complete a job.
	// Must contain either `uses` or `run`.
	// Experimental.
	Steps *[]*Step `field:"optional" json:"steps" yaml:"steps"`
	// A strategy creates a build matrix for your jobs.
	//
	// You can define different variations of an environment to run each job in.
	// Experimental.
	Strategy *NormalJobStrategy `field:"optional" json:"strategy" yaml:"strategy"`
	// The maximum number of minutes to let a workflow run before GitHub automatically cancels it.
	//
	// Default: 360.
	// Experimental.
	TimeoutMinutes NormalJobTimeoutMinutes `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
}

