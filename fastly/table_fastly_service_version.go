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
			Hydrate: getServiceVersion,
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

func getServiceVersion(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	serviceClient, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_version.getServiceVersion", "connection_error", err)
		return nil, err
	}

	input := &fastly.GetVersionInput{
		ServiceID:      serviceClient.ServiceID,
		ServiceVersion: serviceClient.ServiceVersion,
	}

	version, err := serviceClient.Client.GetVersion(input)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_version.getServiceVersion", "api_error", err)
		return nil, err
	}
	d.StreamListItem(ctx, version)

	return nil, nil
}
