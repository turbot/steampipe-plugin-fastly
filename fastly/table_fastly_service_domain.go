package fastly

import (
	"context"

	"github.com/fastly/go-fastly/v3/fastly"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableFastlyServiceDomain(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fastly_service_domain",
		Description: "ServiceDomains for the service version.",
		List: &plugin.ListConfig{
			Hydrate: listServiceDomain,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the domain or domains associated with this service."},
			{Name: "comment", Type: proto.ColumnType_STRING, Description: "A freeform descriptive note."},
			// Other columns
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp (UTC) of when the domain was created."},
			{Name: "deleted_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp (UTC) of when the domain was deleted."},
			//{Name: "locked", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "service_id", Type: proto.ColumnType_STRING, Description: "Alphanumeric string identifying the service."},
			{Name: "service_version", Type: proto.ColumnType_INT, Description: "Integer identifying a service version."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp (UTC) of when the domain was updated."},
		},
	}
}

func listServiceDomain(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, serviceID, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_service_domain.listServiceDomain", "connection_error", err)
		return nil, err
	}
	items, err := conn.ListServiceDomains(&fastly.ListServiceDomainInput{ID: serviceID})
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_service_domain.listServiceDomain", "query_error", err)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}
