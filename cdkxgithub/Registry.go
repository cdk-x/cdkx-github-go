package cdkxgithub


// Experimental.
type Registry struct {
	// Identifies the type of registry.
	// Experimental.
	Type RegistryType `field:"required" json:"type" yaml:"type"`
	// The URL to use to access the dependencies in this registry.
	//
	// The protocol is optional. If not specified, 'https://' is assumed. Dependabot adds or ignores trailing slashes as required.
	// Experimental.
	Url *string `field:"required" json:"url" yaml:"url"`
	// The AWS account ID for AWS CodeArtifact authentication.
	// Experimental.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// The audience for OIDC or AWS authentication.
	// Experimental.
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
	// Experimental.
	AuthKey *string `field:"optional" json:"authKey" yaml:"authKey"`
	// The AWS region for AWS CodeArtifact authentication.
	// Experimental.
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// The client ID for Azure OIDC authentication.
	// Experimental.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The domain for AWS CodeArtifact authentication.
	// Experimental.
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The domain owner for AWS CodeArtifact authentication.
	// Experimental.
	DomainOwner *string `field:"optional" json:"domainOwner" yaml:"domainOwner"`
	// The identity mapping name for JFrog OIDC authentication.
	// Experimental.
	IdentityMappingName *string `field:"optional" json:"identityMappingName" yaml:"identityMappingName"`
	// The JFrog OIDC provider name for authentication.
	// Experimental.
	JfrogOidcProviderName *string `field:"optional" json:"jfrogOidcProviderName" yaml:"jfrogOidcProviderName"`
	// A reference to a Dependabot secret containing an access key for this registry.
	// Experimental.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Experimental.
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// A reference to a Dependabot secret containing the password for the specified user.
	// Experimental.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Experimental.
	PublicKeyFingerprint *string `field:"optional" json:"publicKeyFingerprint" yaml:"publicKeyFingerprint"`
	// The name of the cargo registry.
	// Experimental.
	Registry *string `field:"optional" json:"registry" yaml:"registry"`
	// For registries with type: python-index, if the boolean value is true, pip resolves dependencies by using the specified URL rather than the base URL of the Python Package Index (by default https://pypi.org/simple).
	// Experimental.
	ReplacesBase *bool `field:"optional" json:"replacesBase" yaml:"replacesBase"`
	// Experimental.
	Repo *string `field:"optional" json:"repo" yaml:"repo"`
	// The AWS role name for AWS CodeArtifact authentication.
	// Experimental.
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
	// For registries with type: npm-registry, the npm scope or scopes served by this registry, for example '@my-org'.
	//
	// Dependabot binds only the listed scopes to this registry when generating the .npmrc, so packages outside those scopes continue to resolve from the base registry. This value takes precedence over scope inference from an existing .npmrc or from the lockfile.
	// Experimental.
	Scope interface{} `field:"optional" json:"scope" yaml:"scope"`
	// The tenant ID for Azure OIDC authentication.
	// Experimental.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
	// A reference to a Dependabot secret containing an access token for this registry.
	// Experimental.
	Token *string `field:"optional" json:"token" yaml:"token"`
	// The username that Dependabot uses to access the registry.
	// Experimental.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

