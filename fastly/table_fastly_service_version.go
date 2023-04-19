package fastly

import (
	"context"

	"github.com/fastly/go-fastly/v8/fastly"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableFastlyServiceVersion(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fastly_service_version",
		Description: "Service versions in the Fastly account.",
		List: &plugin.ListConfig{
			Hydrate: listServiceVersions,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getServiceVersion,
			KeyColumns: plugin.SingleColumn("number"),
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
		},
	}
}

func listServiceVersions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	serviceClient, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_version.listServiceVersions", "connection_error", err)
		return nil, err
	}

	items, err := serviceClient.Client.ListVersions(&fastly.ListVersionsInput{ServiceID: serviceClient.ServiceID})
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_version.listServiceVersions", "api_error", err)
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

func getServiceVersion(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	serviceClient, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_version.getServiceVersion", "connection_error", err)
		return nil, err
	}

	number := int(d.EqualsQuals["number"].GetInt64Value())
	input := fastly.GetVersionInput{ServiceID: serviceClient.ServiceID, ServiceVersion: number}

	version, err := serviceClient.Client.GetVersion(&input)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_version.getServiceVersion", "api_error", err)
		return nil, err
	}

	return version, nil
}
