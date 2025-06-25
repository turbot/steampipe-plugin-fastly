package fastly

import (
	"context"

	"github.com/fastly/go-fastly/v8/fastly"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableFastlyACLEntry(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fastly_acl_entry",
		Description: "ACL entries for the service version.",
		List: &plugin.ListConfig{
			ParentHydrate: listACL,
			Hydrate:       listACLEntries,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "acl_id",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getACLEntry,
			KeyColumns: plugin.AllColumns([]string{"acl_id", "id"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the ACL entry.",
			},
			{
				Name:        "acl_id",
				Type:        proto.ColumnType_STRING,
				Description: "Alphanumeric string identifying a ACL.",
			},
			{
				Name:        "ip",
				Type:        proto.ColumnType_IPADDR,
				Description: "An IP address.",
			},
			{
				Name:        "negated",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether to negate the match. Useful primarily when creating individual exceptions to larger subnets.",
			},
			{
				Name:        "comment",
				Type:        proto.ColumnType_STRING,
				Description: "A freeform descriptive note.",
			},
			{
				Name:        "service_id",
				Type:        proto.ColumnType_STRING,
				Description: "Alphanumeric string identifying the service.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp (UTC) of when the ACL was created.",
			},
			{
				Name:        "deleted_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp (UTC) of when the ACL was deleted.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp (UTC) of when the ACL was updated.",
			},
			{
				Name:        "subnet",
				Type:        proto.ColumnType_INT,
				Description: "Subnet associated with the ACL entry.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
		},
	}
}

/// LIST FUNCTION

func listACLEntries(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	acl := h.Item.(*fastly.ACL)

	// check if the provided acl_id is not matching with the parentHydrate
	if d.EqualsQuals["acl_id"] != nil && d.EqualsQualString("acl_id") != acl.ID {
		return nil, nil
	}

	serviceClient, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_acl_entry.listACLEntries", "connection_error", err)
		return nil, err
	}

	// Limiting the results
	maxLimit := 1000
	if d.QueryContext.Limit != nil {
		limit := int(*d.QueryContext.Limit)
		if limit < maxLimit {
			maxLimit = limit
		}
	}

	input := &fastly.ListACLEntriesInput{
		ServiceID: acl.ServiceID,
		ACLID:     acl.ID,
		PerPage:   maxLimit,
	}

	paginator := serviceClient.Client.NewListACLEntriesPaginator(input)
	for {
		if paginator.HasNext() {
			items, err := paginator.GetNext()
			if err != nil {
				plugin.Logger(ctx).Error("fastly_acl_entry.listACLEntries", "api_error", err)
				return nil, err
			}

			for _, item := range items {
				d.StreamListItem(ctx, item)

				// Context can be cancelled due to manual cancellation or the limit has been hit
				if d.RowsRemaining(ctx) == 0 {
					return nil, nil
				}
			}
		} else {
			break
		}
	}

	return nil, nil
}

/// HYDRATE FUNCTION

func getACLEntry(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	aclId := d.EqualsQuals["acl_id"].GetStringValue()
	id := d.EqualsQuals["id"].GetStringValue()
	serviceId := d.EqualsQuals["service_id"].GetStringValue()

	// check if aclId or id is empty
	if aclId == "" || id == "" {
		return nil, nil
	}

	serviceClient, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_acl_entry.getACLEntry", "connection_error", err)
		return nil, err
	}

	input := &fastly.GetACLEntryInput{
		ServiceID: serviceId,
		ACLID:     aclId,
		ID:        id,
	}
	result, err := serviceClient.Client.GetACLEntry(input)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_acl_entry.getACLEntry", "api_error", err)
		return nil, err
	}

	return result, nil
}
