package fastly

import (
	"context"

	"github.com/fastly/go-fastly/v8/fastly"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableFastlyService(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fastly_service",
		Description: "Services in the Fastly account.",
		Get: &plugin.GetConfig{
			Hydrate: getService,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "id", Require: plugin.Required},
			},
		},
		List: &plugin.ListConfig{
			Hydrate: listServices,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "Alphanumeric string identifying a service.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the service.",
			},
			{
				Name:        "active_version",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("ActiveVersion.Number", "ActiveVersion"),
				Description: "Configuration for the active version of this service.",
			},
			{
				Name:        "comment",
				Type:        proto.ColumnType_STRING,
				Description: "A freeform descriptive note.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Time-stamp (UTC) of when the service was created.",
			},
			{
				Name:        "customer_id",
				Type:        proto.ColumnType_STRING,
				Description: "Alphanumeric string identifying the customer.",
			},
			{
				Name:        "deleted_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Time-stamp (UTC) of when the service was deleted.",
			},
			{
				Name:        "type",
				Type:        proto.ColumnType_STRING,
				Description: "The type of this service.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Time-stamp (UTC) of when the service was updated.",
			},
			{
				Name:        "version",
				Type:        proto.ColumnType_JSON,
				Description: "Versions associated with the service.",
				Hydrate:     getService,
				Transform:   transform.FromField("Version"),
			},
			{
				Name:        "versions",
				Type:        proto.ColumnType_JSON,
				Description: "A list of versions associated with the service.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

/// LIST FUNCTION

func listServices(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	serviceClient, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service.listServices", "connection_error", err)
		return nil, err
	}

	input := &fastly.ListServicesInput{
		PerPage: 100,
	}

	paginator := serviceClient.Client.NewListServicesPaginator(input)
	for {
		if paginator.HasNext() {
			items, err := paginator.GetNext()
			if err != nil {
				plugin.Logger(ctx).Error("fastly_service.listServices", "api_error", err)
				return nil, err
			}

			for _, item := range items {
				d.StreamListItem(ctx, item)

				// Context can be cancelled due to manual cancellation or the limit has been hit
				if d.RowsRemaining(ctx) == 0 {
					return nil, nil
				}
			}
		} else {
			break
		}
	}

	return nil, nil
}

var listServicesHydrateMemoize = plugin.HydrateFunc(listServicesHydrateUncached).Memoize(memoize.WithCacheKeyFunction(listServiceCacheKey))

// Build a cache key for the call to getServiceIdCacheKey.
func listServiceCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "listServices"
	return key, nil
}

func listServicesHydrate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return listServicesHydrateMemoize(ctx, d, h)
}

func listServicesHydrateUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	serviceClient, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service.listServices", "connection_error", err)
		return nil, err
	}

	var services []*fastly.Service

	input := &fastly.ListServicesInput{
		PerPage: 100,
	}

	paginator := serviceClient.Client.NewListServicesPaginator(input)
	for {
		if paginator.HasNext() {
			items, err := paginator.GetNext()
			if err != nil {
				plugin.Logger(ctx).Error("fastly_service.listServices", "api_error", err)
				return nil, err
			}
			services = append(services, items...)
		} else {
			break
		}
	}

	return services, nil
}

func getService(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("id")
	if h.Item != nil {
		sc := h.Item.(*fastly.Service)
		id = sc.ID
	}

	// Empty check
	if id == "" {
		return nil, nil
	}
	serviceClient, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service.getService", "connection_error", err)
		return nil, err
	}

	service, err := serviceClient.Client.GetServiceDetails(&fastly.GetServiceInput{ID: id})
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service.getService", "api_error", err)
		return nil, err
	}

	return service, nil
}
