package cdkxgithub


// Experimental.
type DefaultsRun struct {
	// Experimental.
	Shell Shell `field:"optional" json:"shell" yaml:"shell"`
	// Experimental.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

