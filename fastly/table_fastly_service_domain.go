package fastly

import (
	"context"

	"github.com/fastly/go-fastly/v8/fastly"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableFastlyServiceDomain(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fastly_service_domain",
		Description: "Domains for the service version.",
		List: &plugin.ListConfig{
			Hydrate: listServiceDomains,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the domain or domains associated with this service.",
			},
			{
				Name:        "service_id",
				Type:        proto.ColumnType_STRING,
				Description: "Alphanumeric string identifying the service.",
			},
			{
				Name:        "service_version",
				Type:        proto.ColumnType_INT,
				Description: "Integer identifying a service version.",
			},
			{
				Name:        "comment",
				Type:        proto.ColumnType_STRING,
				Description: "A freeform descriptive note.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp (UTC) of when the domain was created.",
			},
			{
				Name:        "deleted_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp (UTC) of when the domain was deleted.",
			},
			{
				Name:        "locked",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether the service version is locked or not. Objects can not be added or edited on locked versions.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp (UTC) of when the domain was updated.",
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

func listServiceDomains(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, serviceID, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_domain.listServiceDomains", "connection_error", err)
		return nil, err
	}

	items, err := conn.ListServiceDomains(&fastly.ListServiceDomainInput{ServiceID: serviceID})
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_domain.listServiceDomains", "api_error", err)
		return nil, err
	}
	for _, item := range items {
		d.StreamListItem(ctx, item)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
