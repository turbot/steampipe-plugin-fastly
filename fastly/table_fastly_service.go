package fastly

import (
	"context"

	"github.com/fastly/go-fastly/v8/fastly"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableFastlyService(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fastly_service",
		Description: "Services in the Fastly account.",
		List: &plugin.ListConfig{
			Hydrate: getService,
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
				Transform:   transform.FromField("ActiveVersion.Number"),
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
				Type:        proto.ColumnType_STRING,
				Description: "Versions associated with the service.",
			},
			{
				Name:        "service_id",
				Type:        proto.ColumnType_STRING,
				Description: "Alphanumeric string identifying the service.",
				Hydrate:     getServiceId,
				Transform:   transform.FromValue(),
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

func getService(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	serviceClient, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service.getService", "connection_error", err)
		return nil, err
	}

	service, err := serviceClient.Client.GetServiceDetails(&fastly.GetServiceInput{ID: serviceClient.ServiceID})
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service.getService", "api_error", err)
		return nil, err
	}
	d.StreamListItem(ctx, service)

	return nil, nil
}
