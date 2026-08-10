package cdkxgithub


// Experimental.
type DependabotProps struct {
	// Experimental.
	Updates *[]*Update `field:"required" json:"updates" yaml:"updates"`
	// Dependabot configuration files require this key, and its value must be 2.
	// Experimental.
	Version *float64 `field:"required" json:"version" yaml:"version"`
	// Enable ecosystems that have beta-level support.
	// Experimental.
	EnableBetaEcosystems *bool `field:"optional" json:"enableBetaEcosystems" yaml:"enableBetaEcosystems"`
	// Define groups that span multiple package ecosystems, allowing consolidated pull requests across different ecosystems.
	// Experimental.
	MultiEcosystemGroups *map[string]*MultiEcosystemGroup `field:"optional" json:"multiEcosystemGroups" yaml:"multiEcosystemGroups"`
	// Experimental.
	Registries *map[string]*Registry `field:"optional" json:"registries" yaml:"registries"`
}

