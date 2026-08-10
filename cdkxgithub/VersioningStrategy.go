package cdkxgithub


// Experimental.
type VersioningStrategy string

const (
	// auto.
	// Experimental.
	VersioningStrategy_AUTO VersioningStrategy = "AUTO"
	// increase.
	// Experimental.
	VersioningStrategy_INCREASE VersioningStrategy = "INCREASE"
	// increase-if-necessary.
	// Experimental.
	VersioningStrategy_INCREASE_HYPHEN_IF_HYPHEN_NECESSARY VersioningStrategy = "INCREASE_HYPHEN_IF_HYPHEN_NECESSARY"
	// lockfile-only.
	// Experimental.
	VersioningStrategy_LOCKFILE_HYPHEN_ONLY VersioningStrategy = "LOCKFILE_HYPHEN_ONLY"
	// widen.
	// Experimental.
	VersioningStrategy_WIDEN VersioningStrategy = "WIDEN"
)

