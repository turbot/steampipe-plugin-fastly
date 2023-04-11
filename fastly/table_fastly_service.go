package fastly

import (
	"context"

	"github.com/fastly/go-fastly/v3/fastly"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableFastlyService(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fastly_service",
		Description: "Services in the Fastly account.",
		List: &plugin.ListConfig{
			Hydrate: listService,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getService,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Alphanumeric string identifying a service."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the service."},
			// Other columns
			{Name: "active_version", Type: proto.ColumnType_INT, Transform: transform.FromValue().Transform(getActiveVersion), Description: "Configuration for the active version of this service."},
			{Name: "comment", Type: proto.ColumnType_STRING, Description: "A freeform descriptive note."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Time-stamp (UTC) of when the service was created."},
			{Name: "customer_id", Type: proto.ColumnType_STRING, Description: "Alphanumeric string identifying the customer."},
			{Name: "deleted_at", Type: proto.ColumnType_TIMESTAMP, Description: "Time-stamp (UTC) of when the service was deleted."},
			{Name: "service_type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Type"), Description: "The type of this service."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "Time-stamp (UTC) of when the service was updated."},
			// Use fastly_service_version instead - {Name: "versions", Type: proto.ColumnType_JSON, Description: "A list of versions associated with the service."},
		},
	}
}

func listService(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, _, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service.listService", "connection_error", err)
		return nil, err
	}
	items, err := conn.ListServices(&fastly.ListServicesInput{})
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service.listService", "query_error", err)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getService(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, _, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service.getService", "connection_error", err)
		return nil, err
	}
	id := d.KeyColumnQuals["id"].GetStringValue()
	result, err := conn.GetServiceDetails(&fastly.GetServiceInput{ID: id})
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service.getService", "query_error", err, "id", id)
		return nil, err
	}
	// Note: Returns a fastly.ServiceDetail
	return result, nil
}

func getActiveVersion(_ context.Context, d *transform.TransformData) (interface{}, error) {
	switch v := d.Value.(type) {
	case *fastly.Service:
		return v.ActiveVersion, nil
	case *fastly.ServiceDetail:
		return v.ActiveVersion.Number, nil
	}
	return nil, nil
}
