package cdkxgithub


// Use to specify a whether the rules in the group apply to version updates or security updates.
// Experimental.
type UpdateGroupsAppliesTo string

const (
	// version-updates.
	// Experimental.
	UpdateGroupsAppliesTo_VERSION_HYPHEN_UPDATES UpdateGroupsAppliesTo = "VERSION_HYPHEN_UPDATES"
	// security-updates.
	// Experimental.
	UpdateGroupsAppliesTo_SECURITY_HYPHEN_UPDATES UpdateGroupsAppliesTo = "SECURITY_HYPHEN_UPDATES"
)

