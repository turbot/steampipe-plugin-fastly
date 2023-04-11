package fastly

import (
	"context"

	"github.com/fastly/go-fastly/v3/fastly"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableFastlyServiceVersion(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fastly_service_version",
		Description: "Service versions in the Fastly account.",
		List: &plugin.ListConfig{
			Hydrate: listServiceVersion,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getServiceVersion,
			KeyColumns: plugin.AllColumns([]string{"number"}),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "service_id", Type: proto.ColumnType_STRING, Description: "Alphanumeric string identifying the service."},
			{Name: "number", Type: proto.ColumnType_INT, Description: "The number of this version."},
			// Other columns
			{Name: "active", Type: proto.ColumnType_BOOL, Description: "Whether this is the active version or not."},
			{Name: "comment", Type: proto.ColumnType_STRING, Description: "A freeform descriptive note."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp (UTC) of when the version was created."},
			{Name: "deleted_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp (UTC) of when the version was deleted."},
			// Not used - {Name: "deployed", Type: proto.ColumnType_BOOL, Description: "Unused at this time."},
			{Name: "locked", Type: proto.ColumnType_BOOL, Description: "Whether this version is locked or not. Objects can not be added or edited on locked versions."},
			// Not used - {Name: "staging", Type: proto.ColumnType_BOOL, Description: "Unused at this time."},
			// Not used - {Name: "testing", Type: proto.ColumnType_BOOL, Description: "Unused at this time."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp (UTC) of when the version was updated."},
		},
	}
}

func listServiceVersion(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, serviceID, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_version.listServiceVersion", "connection_error", err)
		return nil, err
	}
	items, err := conn.ListVersions(&fastly.ListVersionsInput{ServiceID: serviceID})
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_version.listServiceVersion", "query_error", err)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getServiceVersion(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, serviceID, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_version.listServiceVersion", "connection_error", err)
		return nil, err
	}
	number := int(d.KeyColumnQuals["number"].GetInt64Value())
	input := fastly.GetVersionInput{ServiceID: serviceID, ServiceVersion: number}
	version, err := conn.GetVersion(&input)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_version.listServiceVersion", "query_error", err, "input", input)
		return nil, err
	}
	return version, nil
}
