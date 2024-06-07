package fastly

import (
	"context"

	"github.com/fastly/go-fastly/v8/fastly"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableFastlyServiceVersion(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fastly_service_version",
		Description: "Service versions in the Fastly account.",
		List: &plugin.ListConfig{
			Hydrate: listServicesVersions,
		},
		Columns: []*plugin.Column{
			{
				Name:        "service_id",
				Type:        proto.ColumnType_STRING,
				Description: "Alphanumeric string identifying the service.",
			},
			{
				Name:        "number",
				Type:        proto.ColumnType_INT,
				Description: "The number of this version.",
			},
			{
				Name:        "active",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether this is the active version or not.",
			},
			{
				Name:        "comment",
				Type:        proto.ColumnType_STRING,
				Description: "A freeform descriptive note.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp (UTC) of when the version was created.",
			},
			{
				Name:        "deleted_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp (UTC) of when the version was deleted.",
			},
			{
				Name:        "locked",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether this version is locked or not. Objects can not be added or edited on locked versions.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp (UTC) of when the version was updated.",
			},
			{
				Name:        "deployed",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether the version is deployed.",
			},
			{
				Name:        "staging",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether the version is in staging.",
			},
			{
				Name:        "testing",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether the version is undergoing testing.",
			},
		},
	}
}

/// HYDRATE FUNCTION

var listServicesVersionHydrateMemoize = plugin.HydrateFunc(listServicesVersionsUncached).Memoize(memoize.WithCacheKeyFunction(listServiceVersionCacheKey))

// Build a cache key for the call to getServiceIdCacheKey.
func listServiceVersionCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "listServiceVersions"
	return key, nil
}

func listServicesVersionHydrate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return listServicesVersionHydrateMemoize(ctx, d, h)
}

func listServicesVersionsUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var servicesV interface{}
	serviceClient, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_version.listServicesVersionsUncached", "connection_error", err)
		return nil, err
	}

	if serviceClient.ServiceID != "" {
		servicesV = []*fastly.Service{
			{ID: serviceClient.ServiceID},
		}
	} else {
		servicesV, err = listServicesHydrate(ctx, d, h)
		if err != nil {
			return nil, err
		}
	}

	services := servicesV.([]*fastly.Service)

	var serviceVersions []*fastly.Version

	for _, s := range services {
		input := &fastly.ListVersionsInput{
			ServiceID: s.ID,
		}

		versions, err := serviceClient.Client.ListVersions(input)
		if err != nil {
			plugin.Logger(ctx).Error("fastly_service_version.getServiceVersion", "api_error", err)
			return nil, err
		}

		serviceVersions = append(serviceVersions, versions...)
	}

	return serviceVersions, nil
}

func listServicesVersions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	resp, err := listServicesVersionHydrate(ctx, d, h)
	if err != nil {
		return nil, err
	}

	versions := resp.([]*fastly.Version)

	for _, item := range versions {
		d.StreamListItem(ctx, item)

		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
