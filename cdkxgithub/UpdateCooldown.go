package cdkxgithub


// Defines a cooldown period for dependency updates, allowing updates to be delayed for a configurable number of days.
//
// This feature enables users to customize how often Dependabot generates new version updates, offering greater control over update frequency.
// Experimental.
type UpdateCooldown struct {
	// Default cooldown period for dependencies without specific rules (optional).
	// Experimental.
	DefaultDays *float64 `field:"optional" json:"defaultDays" yaml:"defaultDays"`
	// List of dependencies excluded from cooldown.
	//
	// Supports wildcards (`*`).
	// Experimental.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// List of dependencies to apply cooldown.
	//
	// Supports wildcards (`*`).
	// Experimental.
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
	// Cooldown period for major version updates (optional, applies only to package managers supporting SemVer).
	// Experimental.
	SemverMajorDays *float64 `field:"optional" json:"semverMajorDays" yaml:"semverMajorDays"`
	// Cooldown period for minor version updates (optional, applies only to package managers supporting SemVer).
	// Experimental.
	SemverMinorDays *float64 `field:"optional" json:"semverMinorDays" yaml:"semverMinorDays"`
	// Cooldown period for patch version updates (optional, applies only to package managers supporting SemVer).
	// Experimental.
	SemverPatchDays *float64 `field:"optional" json:"semverPatchDays" yaml:"semverPatchDays"`
}

