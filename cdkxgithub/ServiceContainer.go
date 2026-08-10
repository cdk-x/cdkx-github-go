package cdkxgithub


// Experimental.
type ServiceContainer struct {
	// The Docker image to use as the service container to run the action.
	//
	// The value can be the Docker Hub image name or a registry name.
	// Experimental.
	Image *string `field:"required" json:"image" yaml:"image"`
	// Overrides the Docker image's default command (`CMD`).
	//
	// The value is passed as arguments after the image name in the `docker create` command. If you also specify `entrypoint`, `command` provides the arguments to that entrypoint.
	// Experimental.
	Command *string `field:"optional" json:"command" yaml:"command"`
	// If the image's container registry requires authentication to pull the image, you can use credentials to set a map of the username and password.
	//
	// The credentials are the same values that you would provide to the `docker login` command.
	// Experimental.
	Credentials *ServiceContainerCredentials `field:"optional" json:"credentials" yaml:"credentials"`
	// Overrides the Docker image's default `ENTRYPOINT`.
	//
	// The value is a single string defining the executable to run. Use this when you need to replace the image's entrypoint entirely. You can combine `entrypoint` with `command` to pass arguments to the custom entrypoint.
	// Experimental.
	Entrypoint *string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// Sets a map of environment variables in the service container.
	// Experimental.
	Env interface{} `field:"optional" json:"env" yaml:"env"`
	// Additional Docker container resource options.
	//
	// For a list of options, see https://docs.docker.com/engine/reference/commandline/create/#options.
	// Experimental.
	Options *string `field:"optional" json:"options" yaml:"options"`
	// Sets an array of ports to expose on the service container.
	// Experimental.
	Ports *[]ServiceContainerPorts `field:"optional" json:"ports" yaml:"ports"`
	// Sets an array of volumes for the service container to use.
	//
	// You can use volumes to share data between services or other steps in a job. You can specify named Docker volumes, anonymous Docker volumes, or bind mounts on the host.
	// To specify a volume, you specify the source and destination path: <source>:<destinationPath>
	// The <source> is a volume name or an absolute path on the host machine, and <destinationPath> is an absolute path in the container.
	// Experimental.
	Volumes *[]*string `field:"optional" json:"volumes" yaml:"volumes"`
}

