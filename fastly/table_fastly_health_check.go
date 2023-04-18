package fastly

import (
	"context"

	"github.com/fastly/go-fastly/v8/fastly"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableFastlyHealthCheck(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fastly_health_check",
		Description: "Health checks for the service version.",
		List: &plugin.ListConfig{
			Hydrate:    listHealthCheck,
			KeyColumns: plugin.AllColumns([]string{"service_version"}),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getHealthCheck,
			KeyColumns: plugin.AllColumns([]string{"service_version", "name"}),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the health check."},
			{Name: "method", Type: proto.ColumnType_STRING, Description: "Which HTTP method to use."},
			{Name: "host", Type: proto.ColumnType_STRING, Description: "Which host to check."},
			{Name: "path", Type: proto.ColumnType_STRING, Description: "The path to check."},
			// Other columns
			{Name: "check_interval", Type: proto.ColumnType_INT, Description: "How often to run the healthcheck in milliseconds."},
			{Name: "comment", Type: proto.ColumnType_STRING, Description: "A freeform descriptive note."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp (UTC) of when the health check was created."},
			{Name: "deleted_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp (UTC) of when the health check was deleted."},
			{Name: "expected_response", Type: proto.ColumnType_INT, Description: "The status code expected from the host."},
			{Name: "http_version", Type: proto.ColumnType_STRING, Description: "Whether to use version 1.0 or 1.1 HTTP."},
			{Name: "initial", Type: proto.ColumnType_INT, Description: "When loading a config, the initial number of probes to be seen as OK."},
			{Name: "service_id", Type: proto.ColumnType_STRING, Description: "Alphanumeric string identifying the service."},
			{Name: "service_version", Type: proto.ColumnType_INT, Description: "Integer identifying a service version."},
			{Name: "threshold", Type: proto.ColumnType_INT, Description: "How many healthchecks must succeed to be considered healthy."},
			{Name: "timeout", Type: proto.ColumnType_INT, Description: "Timeout in milliseconds."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp (UTC) of when the health check was updated."},
			{Name: "window", Type: proto.ColumnType_INT, Description: "The number of most recent healthcheck queries to keep for this healthcheck."},
		},
	}
}

func listHealthCheck(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, serviceID, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_health_check.listHealthCheck", "connection_error", err)
		return nil, err
	}

	input := fastly.ListHealthChecksInput{
		ServiceID: serviceID,
	}
	if d.EqualsQuals["service_version"] != nil {
		input.ServiceVersion = int(d.EqualsQuals["service_version"].GetInt64Value())
	}

	items, err := conn.ListHealthChecks(&input)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_health_check.listHealthCheck", "query_error", err)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getHealthCheck(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, serviceID, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_health_check.getHealthCheck", "connection_error", err)
		return nil, err
	}
	serviceVersion := int(d.EqualsQuals["service_version"].GetInt64Value())
	name := d.EqualsQuals["name"].GetStringValue()
	input := fastly.GetHealthCheckInput{ServiceID: serviceID, ServiceVersion: serviceVersion, Name: name}
	result, err := conn.GetHealthCheck(&input)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_health_check.getHealthCheck", "query_error", err, "input", input)
		return nil, err
	}
	return result, nil
}
