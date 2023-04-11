package fastly

import (
	"context"

	"github.com/fastly/go-fastly/v3/fastly"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableFastlyACLEntry(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fastly_acl",
		Description: "ACL entries for the service version.",
		List: &plugin.ListConfig{
			Hydrate:    listACLEntry,
			KeyColumns: plugin.AllColumns([]string{"acl_id"}),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getACLEntry,
			KeyColumns: plugin.AllColumns([]string{"acl_id", "id"}),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "acl_id", Type: proto.ColumnType_STRING, Description: "Alphanumeric string identifying a ACL."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The ID of the ACL entry."},
			{Name: "ip", Type: proto.ColumnType_IPADDR, Description: "An IP address."},
			{Name: "negated", Type: proto.ColumnType_BOOL, Description: "Whether to negate the match. Useful primarily when creating individual exceptions to larger subnets."},
			{Name: "comment", Type: proto.ColumnType_STRING, Description: "A freeform descriptive note."},
			// Other columns
			{Name: "service_id", Type: proto.ColumnType_STRING, Description: "Alphanumeric string identifying the service."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp (UTC) of when the ACL was created."},
			{Name: "deleted_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp (UTC) of when the ACL was deleted."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp (UTC) of when the ACL was updated."},
		},
	}
}

func listACLEntry(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, serviceID, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_acl.listACLEntry", "connection_error", err)
		return nil, err
	}
	input := fastly.ListACLEntriesInput{
		ServiceID: serviceID,
		ACLID:     d.KeyColumnQuals["acl_id"].GetStringValue(),
	}
	items, err := conn.ListACLEntries(&input)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_acl.listACLEntry", "query_error", err, "input", input)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getACLEntry(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, serviceID, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_acl.getACLEntry", "connection_error", err)
		return nil, err
	}
	aclID := d.KeyColumnQuals["acl_id"].GetStringValue()
	id := d.KeyColumnQuals["id"].GetStringValue()
	input := fastly.GetACLEntryInput{ServiceID: serviceID, ACLID: aclID, ID: id}
	result, err := conn.GetACLEntry(&input)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_acl.getACLEntry", "query_error", err, "input", input)
		return nil, err
	}
	return result, nil
}
