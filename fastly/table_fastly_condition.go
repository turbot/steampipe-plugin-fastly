package fastly

import (
	"context"

	"github.com/fastly/go-fastly/v3/fastly"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableFastlyCondition(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fastly_condition",
		Description: "Conditions defined in the service version.",
		List: &plugin.ListConfig{
			Hydrate:    listCondition,
			KeyColumns: plugin.AllColumns([]string{"service_version"}),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "service_id", Type: proto.ColumnType_STRING, Description: "Alphanumeric string identifying the service."},
			{Name: "service_version", Type: proto.ColumnType_INT, Transform: transform.FromQual("service_version"), Description: "Integer identifying a service version."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the condition."},
			// Other columns
			{Name: "comment", Type: proto.ColumnType_STRING, Description: "A freeform descriptive note."},
			{Name: "condition_type", Type: proto.ColumnType_STRING, Description: "Type of the condition."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp (UTC) of when the condition was created."},
			{Name: "deleted_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp (UTC) of when the condition was deleted."},
			{Name: "priority", Type: proto.ColumnType_INT, Description: "Priority determines execution order. Lower numbers execute first."},
			{Name: "statement", Type: proto.ColumnType_STRING, Description: "A conditional expression in VCL used to determine if the condition is met."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp (UTC) of when the condition was updated."},
		},
	}
}

func listCondition(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, serviceID, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_version.listCondition", "connection_error", err)
		return nil, err
	}
	items, err := conn.ListConditions(&fastly.ListConditionsInput{ServiceID: serviceID, ServiceVersion: int(d.KeyColumnQuals["service_version"].GetInt64Value())})
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_version.listCondition", "query_error", err)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}
