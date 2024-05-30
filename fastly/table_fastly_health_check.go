package fastly

import (
	"context"

	"github.com/fastly/go-fastly/v8/fastly"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableFastlyHealthCheck(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fastly_health_check",
		Description: "Health checks for the service version.",
		List: &plugin.ListConfig{
			ParentHydrate: listServicesVersions,
			Hydrate:       listHealthChecks,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getHealthCheck,
			KeyColumns: plugin.AllColumns([]string{"name", "service_id", "service_version"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the health check.",
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
				Name:        "method",
				Type:        proto.ColumnType_STRING,
				Description: "Which HTTP method to use.",
			},
			{
				Name:        "host",
				Type:        proto.ColumnType_STRING,
				Description: "Which host to check.",
			},
			{
				Name:        "path",
				Type:        proto.ColumnType_STRING,
				Description: "The path to check.",
			},
			{
				Name:        "check_interval",
				Type:        proto.ColumnType_INT,
				Description: "How often to run the health check in milliseconds.",
			},
			{
				Name:        "comment",
				Type:        proto.ColumnType_STRING,
				Description: "A freeform descriptive note.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp (UTC) of when the health check was created.",
			},
			{
				Name:        "deleted_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp (UTC) of when the health check was deleted.",
			},
			{
				Name:        "expected_response",
				Type:        proto.ColumnType_INT,
				Description: "The status code expected from the host.",
			},
			{
				Name:        "http_version",
				Type:        proto.ColumnType_STRING,
				Description: "Whether to use version 1.0 or 1.1 HTTP.",
			},
			{
				Name:        "initial",
				Type:        proto.ColumnType_INT,
				Description: "When loading a config, the initial number of probes to be seen as OK.",
			},
			{
				Name:        "threshold",
				Type:        proto.ColumnType_INT,
				Description: "How many health checks must succeed to be considered healthy.",
			},
			{
				Name:        "timeout",
				Type:        proto.ColumnType_INT,
				Description: "Timeout in milliseconds.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp (UTC) of when the health check was updated.",
			},
			{
				Name:        "window",
				Type:        proto.ColumnType_INT,
				Description: "The number of most recent health check queries to keep for this health check.",
			},
			{
				Name:        "headers",
				Type:        proto.ColumnType_JSON,
				Description: "A list of headers corresponding to this health check.",
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

func listHealthChecks(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	serviceVersion := h.Item.(*fastly.Version)

	serviceClient, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_health_check.listHealthChecks", "connection_error", err)
		return nil, err
	}
	input := &fastly.ListHealthChecksInput{
		ServiceID:      serviceVersion.ServiceID,
		ServiceVersion: serviceVersion.Number,
	}
	items, err := serviceClient.Client.ListHealthChecks(input)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_health_check.listHealthChecks", "api_error", err)
		return nil, err
	}

	for _, item := range items {
		d.StreamListItem(ctx, item)
	}

	return nil, nil
}

/// HYDRATE FUNCTION

func getHealthCheck(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	name := d.EqualsQualString("name")
	serviceId := d.EqualsQualString("service_id")
	serviceVersion := d.EqualsQuals["service_version"].GetInt64Value()

	// check if the name is empty
	if name == "" || serviceId== "" || serviceVersion == 0 {
		return nil, nil
	}

	serviceClient, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_health_check.getHealthCheck", "connection_error", err)
		return nil, err
	}

	input := &fastly.GetHealthCheckInput{
		ServiceID:      serviceId,
		ServiceVersion: int(serviceVersion),
		Name:           name,
	}
	result, err := serviceClient.Client.GetHealthCheck(input)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_health_check.getHealthCheck", "api_error", err)
		return nil, err
	}

	return result, nil
}
