package fastly

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-fastly",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError,
		},
		TableMap: map[string]*plugin.Table{
			"fastly_acl":       tableFastlyACL(ctx),
			"fastly_acl_entry": tableFastlyACLEntry(ctx),
			"fastly_backend":   tableFastlyBackend(ctx),
			//"fastly_bulk_certificate": tableFastlyBulkCertificate(ctx),
			//"fastly_cache_settings": tableFastlyCacheSettings(ctx),
			"fastly_condition": tableFastlyCondition(ctx),
			//"fastly_current_user": tableFastlyCurrentUser(ctx),
			"fastly_data_center":  tableFastlyDataCenter(ctx),
			"fastly_dictionary":   tableFastlyDictionary(ctx),
			"fastly_health_check": tableFastlyHealthCheck(ctx),
			"fastly_ip_range":     tableFastlyIPRange(ctx),
			"fastly_pool":         tableFastlyPool(ctx),
			//"fastly_region":          tableFastlyRegion(ctx),
			"fastly_service":         tableFastlyService(ctx),
			"fastly_service_domain":  tableFastlyServiceDomain(ctx),
			"fastly_service_version": tableFastlyServiceVersion(ctx),
			"fastly_token":           tableFastlyToken(ctx),
		},
	}
	return p
}
