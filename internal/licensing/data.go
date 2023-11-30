package licensing

// The list of plans.
const (
	// PlanOldEnterpriseStarter is the old "Enterprise Starter" plan.
	// Deprecated: PlanOldEnterpriseStarter has been deprecated and we will stop issuing licenses for it.
	PlanOldEnterpriseStarter Plan = "old-starter-0"
	// PlanOldEnterprise is the old "Enterprise" plan.
	// Deprecated: PlanOldEnterprise has been deprecated and we will stop issuing licenses for it.
	PlanOldEnterprise Plan = "old-enterprise-0"

	// PlanTeam0 is the "Team" plan pre-4.0.
	// Deprecated: PlanTeam0 has been deprecated and we will stop issuing licenses for it.
	PlanTeam0 Plan = "team-0"
	// PlanEnterprise0 is the "Enterprise" plan pre-4.0.
	// Deprecated: PlanEnterprise0 has been deprecated and we will stop issuing licenses for it.
	PlanEnterprise0 Plan = "enterprise-0"

	// PlanBusiness0 is the "Business" plan for 4.0.
	// Deprecated: PlanBusiness0 has been deprecated and we will stop issuing licenses for it.
	PlanBusiness0 Plan = "business-0"
	// PlanEnterprise1 is the "Enterprise" plan for 4.0.
	// Deprecated: PlanEnterprise1 has been deprecated and we will stop issuing licenses for it.
	PlanEnterprise1 Plan = "enterprise-1"

	// PlanEnterpriseExtension is for customers who require an extended trial on a new Sourcegraph 4.4.2 instance.
	// Deprecated: PlanEnterpriseExtension has been deprecated and we will stop issuing licenses for it.
	PlanEnterpriseExtension Plan = "enterprise-extension"

	// PlanFree0 is the default plan if no license key is set before 4.5.
	// Deprecated: PlanFree0 has been deprecated and we will stop issuing licenses for it.
	PlanFree0 Plan = "free-0"

	// PlanFree1 is the default plan if no license key is set from 4.5 onwards.
	// Deprecated: PlanFree1 has been deprecated and we will stop issuing licenses for it.
	PlanFree1 Plan = "free-1"

	// PlanAirGappedEnterprise is the same PlanEnterprise1 but with FeatureAllowAirGapped, and works starting from 5.1.
	// Deprecated: PlanAirGappedEnterprise has been deprecated and we will stop issuing licenses for it.
	PlanAirGappedEnterprise Plan = "enterprise-air-gap-0"

	PlanCodeSearch           Plan = "code-search"
	PlanCIP                  Plan = "cip"
	PlanCodeSearchPro        Plan = "code-search-pro"
	PlanCodeAIPro            Plan = "code-ai-pro"
	PlanCodyEnterprise       Plan = "cody-enterprise"
	PlanCodeSearchEnterprise Plan = "code-search-enterprise"
	PlanCodeAIEnterprise     Plan = "code-ai-enterprise"
)

var AllPlans = []Plan{
	// Deprecated plans:
	PlanOldEnterpriseStarter,
	PlanOldEnterprise,
	PlanTeam0,
	PlanEnterprise0,
	PlanBusiness0,
	PlanEnterprise1,
	PlanEnterpriseExtension,
	PlanFree0,
	PlanFree1,
	PlanAirGappedEnterprise,

	// Current plans:
	PlanCodeSearch,
	PlanCIP,
	PlanCodeSearchPro,
	PlanCodeAIPro,
	PlanCodyEnterprise,
	PlanCodeSearchEnterprise,
	PlanCodeAIEnterprise,
}

// The list of features. For each feature, add a new const here and the checking logic in
// isFeatureEnabled.
const (
	// FeatureSSO is whether non-builtin authentication may be used, such as GitHub
	// OAuth, GitLab OAuth, SAML, and OpenID.
	FeatureSSO BasicFeature = "sso"

	// FeatureACLs is whether the Background Permissions Syncing may be be used for
	// setting repository permissions.
	FeatureACLs BasicFeature = "acls"

	// FeatureExplicitPermissionsAPI is whether the Explicit Permissions API may be be used for
	// setting repository permissions.
	FeatureExplicitPermissionsAPI BasicFeature = "explicit-permissions-api"

	// FeatureExtensionRegistry is whether publishing extensions to this Sourcegraph instance has been
	// purchased. If not, then extensions must be published to Sourcegraph.com. All instances may use
	// extensions published to Sourcegraph.com.
	FeatureExtensionRegistry BasicFeature = "private-extension-registry"

	// FeatureRemoteExtensionsAllowDisallow is whether explicitly specify a list of allowed remote
	// extensions and prevent any other remote extensions from being used has been purchased. It
	// does not apply to locally published extensions.
	FeatureRemoteExtensionsAllowDisallow BasicFeature = "remote-extensions-allow-disallow"

	// FeatureBranding is whether custom branding of this Sourcegraph instance has been purchased.
	FeatureBranding BasicFeature = "branding"

	// FeatureCampaigns is whether campaigns (now: batch changes) on this Sourcegraph instance has been purchased.
	//
	// DEPRECATED: See FeatureBatchChanges.
	FeatureCampaigns BasicFeature = "campaigns"

	// FeatureMonitoring is whether monitoring on this Sourcegraph instance has been purchased.
	FeatureMonitoring BasicFeature = "monitoring"

	// FeatureBackupAndRestore is whether builtin backup and restore on this Sourcegraph instance
	// has been purchased.
	FeatureBackupAndRestore BasicFeature = "backup-and-restore"

	// FeatureCodeInsights is whether Code Insights on this Sourcegraph instance has been purchased.
	FeatureCodeInsights BasicFeature = "code-insights"

	// FeatureSCIM is whether SCIM User Management has been purchased on this instance.
	FeatureSCIM BasicFeature = "SCIM"

	// FeatureCody is whether or not Cody and embeddings has been purchased on this instance.
	FeatureCody BasicFeature = "cody"

	// FeatureAllowAirGapped is whether or not air gapped mode is allowed on this instance.
	FeatureAllowAirGapped BasicFeature = "allow-air-gapped"

	// FeatureCodeMonitors is whether code monitors is allowed on this Sourcegraph instance.
	FeatureCodeMonitors BasicFeature = "code-monitors"

	// FeatureNotebooks is whether the notebooks feature is allowed on this Sourcegraph instance.
	FeatureNotebooks BasicFeature = "notebooks"

	// FeatureCodeSearch is whether the code search feature suite is allowed on this Sourcegraph instance.
	FeatureCodeSearch BasicFeature = "code-search"
)

type PlanDetails struct {
	DisplayName string
	Features    []Feature
}

// planDetails defines the features that are enabled for each plan.
var planDetails = map[Plan]PlanDetails{
	PlanOldEnterpriseStarter: {
		DisplayName: "Sourcegraph Enterprise Starter",
		Features: []Feature{
			&FeatureBatchChanges{MaxNumChangesets: 10},
			&FeaturePrivateRepositories{Unrestricted: true},
			FeatureCodeMonitors,
			FeatureNotebooks,
			FeatureCodeSearch,
		},
	},
	PlanOldEnterprise: {
		DisplayName: "Sourcegraph Enterprise",
		Features: []Feature{
			FeatureSSO,
			FeatureACLs,
			FeatureExplicitPermissionsAPI,
			FeatureExtensionRegistry,
			FeatureRemoteExtensionsAllowDisallow,
			FeatureBranding,
			FeatureCampaigns,
			&FeatureBatchChanges{Unrestricted: true},
			&FeaturePrivateRepositories{Unrestricted: true},
			FeatureMonitoring,
			FeatureBackupAndRestore,
			FeatureCodeInsights,
			FeatureSCIM,
			FeatureCody,
			FeatureCodeMonitors,
			FeatureNotebooks,
			FeatureCodeSearch,
		},
	},
	PlanTeam0: {
		DisplayName: "Sourcegraph Team",
		Features: []Feature{
			FeatureACLs,
			FeatureExplicitPermissionsAPI,
			FeatureSSO,
			&FeatureBatchChanges{MaxNumChangesets: 10},
			&FeaturePrivateRepositories{Unrestricted: true},
			FeatureCodeMonitors,
			FeatureNotebooks,
			FeatureCodeSearch,
		},
	},
	PlanEnterprise0: {
		DisplayName: "Sourcegraph Enterprise",
		Features: []Feature{
			FeatureACLs,
			FeatureExplicitPermissionsAPI,
			FeatureSSO,
			&FeatureBatchChanges{MaxNumChangesets: 10},
			&FeaturePrivateRepositories{Unrestricted: true},
			FeatureSCIM,
			FeatureCody,
			FeatureCodeMonitors,
			FeatureNotebooks,
			FeatureCodeSearch,
		},
	},

	PlanBusiness0: {
		DisplayName: "Sourcegraph Business",
		Features: []Feature{
			FeatureACLs,
			FeatureCampaigns,
			&FeatureBatchChanges{Unrestricted: true},
			&FeaturePrivateRepositories{Unrestricted: true},
			FeatureCodeInsights,
			FeatureSSO,
			FeatureSCIM,
			FeatureCody,
			FeatureCodeMonitors,
			FeatureNotebooks,
			FeatureCodeSearch,
		},
	},
	PlanEnterprise1: {
		DisplayName: "Sourcegraph Enterprise",
		Features: []Feature{
			FeatureACLs,
			FeatureCampaigns,
			FeatureCodeInsights,
			&FeatureBatchChanges{Unrestricted: true},
			&FeaturePrivateRepositories{Unrestricted: true},
			FeatureExplicitPermissionsAPI,
			FeatureSSO,
			FeatureSCIM,
			FeatureCody,
			FeatureCodeMonitors,
			FeatureNotebooks,
			FeatureCodeSearch,
		},
	},
	PlanEnterpriseExtension: {
		DisplayName: "Sourcegraph Enterprise",
		Features: []Feature{
			FeatureACLs,
			FeatureCampaigns,
			FeatureCodeInsights,
			&FeatureBatchChanges{Unrestricted: true},
			&FeaturePrivateRepositories{Unrestricted: true},
			FeatureExplicitPermissionsAPI,
			FeatureSSO,
			FeatureSCIM,
			FeatureCody,
			FeatureCodeMonitors,
			FeatureNotebooks,
			FeatureCodeSearch,
		},
	},
	PlanFree0: {
		DisplayName: "Sourcegraph Free",
		Features: []Feature{
			FeatureSSO,
			FeatureMonitoring,
			&FeatureBatchChanges{MaxNumChangesets: 10},
			&FeaturePrivateRepositories{Unrestricted: true},
			FeatureCodeMonitors,
			FeatureNotebooks,
			FeatureCodeSearch,
		},
	},
	PlanFree1: {
		DisplayName: "Sourcegraph Free",
		Features: []Feature{
			FeatureMonitoring,
			&FeatureBatchChanges{MaxNumChangesets: 10},
			&FeaturePrivateRepositories{MaxNumPrivateRepos: 1},
			FeatureCodeMonitors,
			FeatureNotebooks,
			FeatureCodeSearch,
		},
	},
	PlanAirGappedEnterprise: {
		DisplayName: "Sourcegraph Enterprise",
		Features: []Feature{
			FeatureACLs,
			FeatureCampaigns,
			FeatureCodeInsights,
			&FeatureBatchChanges{Unrestricted: true},
			&FeaturePrivateRepositories{Unrestricted: true},
			FeatureExplicitPermissionsAPI,
			FeatureSSO,
			FeatureSCIM,
			FeatureCody,
			FeatureAllowAirGapped,
			FeatureCodeMonitors,
			FeatureNotebooks,
			FeatureCodeSearch,
		},
	},
	PlanCodeSearch: {
		DisplayName: "PlanCodeSearch",
		Features: []Feature{
			FeatureSSO,
			FeatureACLs,
			FeatureExplicitPermissionsAPI,
			FeatureSCIM,
			FeatureCodeMonitors,
			FeatureNotebooks,
			FeatureCodeSearch,
			&FeatureBatchChanges{MaxNumChangesets: 10}, // TODO: Need a way to override this with unlimited.
		},
	},
	PlanCIP: {
		DisplayName: "PlanCIP",
		Features: []Feature{
			FeatureSSO,
			FeatureACLs,
			FeatureExplicitPermissionsAPI,
			FeatureCodeInsights,
			FeatureSCIM,
			FeatureCodeMonitors,
			FeatureNotebooks,
			FeatureCodeSearch,
			&FeatureBatchChanges{Unrestricted: true},
		},
	},
	PlanCodeSearchPro: {
		DisplayName: "Code Search Pro",
		Features: []Feature{
			FeatureSSO,
			FeatureACLs,
			FeatureSCIM,
			FeatureCodeSearch,
			&FeaturePrivateRepositories{MaxNumPrivateRepos: 250}, // TODO: Max total repos, not max private.
			// TODO: Max code host connections 1
		},
	},
	PlanCodeAIPro: {
		DisplayName: "PlanCodeAIPro",
		Features: []Feature{
			FeatureSSO,
			FeatureACLs,
			FeatureSCIM,
			FeatureCody,
			&FeaturePrivateRepositories{MaxNumPrivateRepos: 250}, // TODO: Max total repos, not max private.
			// TODO: Max code host connections 1
			// TODO: Max cody user count through tag.
		},
	},
	PlanCodyEnterprise: {
		DisplayName: "PlanCodyEnterprise",
		Features: []Feature{
			FeatureSSO,
			FeatureACLs,
			FeatureExplicitPermissionsAPI,
			FeatureSCIM,
			FeatureCody,
		},
	},
	PlanCodeSearchEnterprise: {
		DisplayName: "PlanCodeSearchEnterprise",
		Features: []Feature{
			FeatureSSO,
			FeatureACLs,
			FeatureExplicitPermissionsAPI,
			FeatureCodeInsights,
			FeatureSCIM,
			FeatureCodeMonitors,
			FeatureNotebooks,
			FeatureCodeSearch,
			&FeatureBatchChanges{Unrestricted: true},
		},
	},
	PlanCodeAIEnterprise: {
		DisplayName: "PlanCodeAIEnterprise",
		Features: []Feature{
			FeatureSSO,
			FeatureACLs,
			FeatureExplicitPermissionsAPI,
			FeatureCodeInsights,
			FeatureSCIM,
			FeatureCody,
			FeatureCodeMonitors,
			FeatureNotebooks,
			FeatureCodeSearch,
			&FeatureBatchChanges{Unrestricted: true},
		},
	},
}
