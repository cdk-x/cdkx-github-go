package cdkxgithub


// Identifies the type of registry.
// Experimental.
type RegistryType string

const (
	// cargo-registry.
	// Experimental.
	RegistryType_CARGO_HYPHEN_REGISTRY RegistryType = "CARGO_HYPHEN_REGISTRY"
	// composer-repository.
	// Experimental.
	RegistryType_COMPOSER_HYPHEN_REPOSITORY RegistryType = "COMPOSER_HYPHEN_REPOSITORY"
	// docker-registry.
	// Experimental.
	RegistryType_DOCKER_HYPHEN_REGISTRY RegistryType = "DOCKER_HYPHEN_REGISTRY"
	// git.
	// Experimental.
	RegistryType_GIT RegistryType = "GIT"
	// goproxy-server.
	// Experimental.
	RegistryType_GOPROXY_HYPHEN_SERVER RegistryType = "GOPROXY_HYPHEN_SERVER"
	// hex-organization.
	// Experimental.
	RegistryType_HEX_HYPHEN_ORGANIZATION RegistryType = "HEX_HYPHEN_ORGANIZATION"
	// hex-repository.
	// Experimental.
	RegistryType_HEX_HYPHEN_REPOSITORY RegistryType = "HEX_HYPHEN_REPOSITORY"
	// helm-registry.
	// Experimental.
	RegistryType_HELM_HYPHEN_REGISTRY RegistryType = "HELM_HYPHEN_REGISTRY"
	// maven-repository.
	// Experimental.
	RegistryType_MAVEN_HYPHEN_REPOSITORY RegistryType = "MAVEN_HYPHEN_REPOSITORY"
	// npm-registry.
	// Experimental.
	RegistryType_NPM_HYPHEN_REGISTRY RegistryType = "NPM_HYPHEN_REGISTRY"
	// nuget-feed.
	// Experimental.
	RegistryType_NUGET_HYPHEN_FEED RegistryType = "NUGET_HYPHEN_FEED"
	// pub-repository.
	// Experimental.
	RegistryType_PUB_HYPHEN_REPOSITORY RegistryType = "PUB_HYPHEN_REPOSITORY"
	// python-index.
	// Experimental.
	RegistryType_PYTHON_HYPHEN_INDEX RegistryType = "PYTHON_HYPHEN_INDEX"
	// rubygems-server.
	// Experimental.
	RegistryType_RUBYGEMS_HYPHEN_SERVER RegistryType = "RUBYGEMS_HYPHEN_SERVER"
	// terraform-registry.
	// Experimental.
	RegistryType_TERRAFORM_HYPHEN_REGISTRY RegistryType = "TERRAFORM_HYPHEN_REGISTRY"
)

