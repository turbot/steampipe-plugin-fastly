package fastly

import (
	"context"

	"github.com/fastly/go-fastly/v3/fastly"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableFastlyACL(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fastly_acl",
		Description: "ACLs for the service version.",
		List: &plugin.ListConfig{
			Hydrate:    listACL,
			KeyColumns: plugin.AllColumns([]string{"service_version"}),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getACL,
			KeyColumns: plugin.AllColumns([]string{"service_version", "name"}),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The ID of the ACL."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the ACL."},
			// Other columns
			{Name: "service_id", Type: proto.ColumnType_STRING, Description: "Alphanumeric string identifying the service."},
			{Name: "service_version", Type: proto.ColumnType_INT, Description: "Integer identifying a service version."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp (UTC) of when the ACL was created."},
			{Name: "deleted_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp (UTC) of when the ACL was deleted."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp (UTC) of when the ACL was updated."},
		},
	}
}

func listACL(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, serviceID, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_acl.listACL", "connection_error", err)
		return nil, err
	}
	input := fastly.ListACLsInput{
		ServiceID:      serviceID,
		ServiceVersion: int(d.KeyColumnQuals["service_version"].GetInt64Value()),
	}
	items, err := conn.ListACLs(&input)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_acl.listACL", "query_error", err)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getACL(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, serviceID, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_acl.getACL", "connection_error", err)
		return nil, err
	}
	serviceVersion := int(d.KeyColumnQuals["service_version"].GetInt64Value())
	name := d.KeyColumnQuals["name"].GetStringValue()
	input := fastly.GetACLInput{ServiceID: serviceID, ServiceVersion: serviceVersion, Name: name}
	result, err := conn.GetACL(&input)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_acl.getACL", "query_error", err, "input", input)
		return nil, err
	}
	return result, nil
}
