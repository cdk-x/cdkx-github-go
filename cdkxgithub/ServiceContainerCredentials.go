package cdkxgithub


// If the image's container registry requires authentication to pull the image, you can use credentials to set a map of the username and password.
//
// The credentials are the same values that you would provide to the `docker login` command.
// Experimental.
type ServiceContainerCredentials struct {
	// Experimental.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Experimental.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

