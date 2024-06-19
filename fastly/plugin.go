package fastly

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-fastly",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		ConnectionKeyColumns: []plugin.ConnectionKeyColumn{
			{
				Name:    "service_id",
				Hydrate: getServiceId,
			},
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		DefaultIgnoreConfig: &plugin.IgnoreConfig{
			ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"404"}),
		},
		DefaultRetryConfig: &plugin.RetryConfig{
			ShouldRetryErrorFunc: shouldRetryError([]string{"429"}),
		},
		TableMap: map[string]*plugin.Table{
			"fastly_acl":             tableFastlyACL(ctx),
			"fastly_acl_entry":       tableFastlyACLEntry(ctx),
			"fastly_backend":         tableFastlyBackend(ctx),
			"fastly_condition":       tableFastlyCondition(ctx),
			"fastly_data_center":     tableFastlyDataCenter(ctx),
			"fastly_dictionary":      tableFastlyDictionary(ctx),
			"fastly_health_check":    tableFastlyHealthCheck(ctx),
			"fastly_ip_range":        tableFastlyIPRange(ctx),
			"fastly_pool":            tableFastlyPool(ctx),
			"fastly_service":         tableFastlyService(ctx),
			"fastly_service_domain":  tableFastlyServiceDomain(ctx),
			"fastly_service_version": tableFastlyServiceVersion(ctx),
			"fastly_token":           tableFastlyToken(ctx),
		},
	}
	return p
}
