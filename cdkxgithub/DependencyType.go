package cdkxgithub


// Experimental.
type DependencyType string

const (
	// direct.
	// Experimental.
	DependencyType_DIRECT DependencyType = "DIRECT"
	// indirect.
	// Experimental.
	DependencyType_INDIRECT DependencyType = "INDIRECT"
	// all.
	// Experimental.
	DependencyType_ALL DependencyType = "ALL"
	// production.
	// Experimental.
	DependencyType_PRODUCTION DependencyType = "PRODUCTION"
	// development.
	// Experimental.
	DependencyType_DEVELOPMENT DependencyType = "DEVELOPMENT"
)

