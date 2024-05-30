package fastly

import (
	"context"

	"github.com/fastly/go-fastly/v8/fastly"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableFastlyACL(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fastly_acl",
		Description: "ACLs for the service version.",
		List: &plugin.ListConfig{
			Hydrate: listACL,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getACL,
			KeyColumns: plugin.AllColumns([]string{"service_id", "name", "service_version"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the ACL.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the ACL.",
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

func listACL(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	resp, err := listServicesVersionHydrate(ctx, d, h)
	if err != nil {
		return nil, err
	}
	versions := resp.([]*fastly.Version)

	serviceClient, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_acl.listACL", "connection_error", err)
		return nil, err
	}

	for _, item := range versions {
		input := fastly.ListACLsInput{
			ServiceID:      item.ServiceID,
			ServiceVersion: item.Number,
		}
		items, err := serviceClient.Client.ListACLs(&input)
		if err != nil {
			plugin.Logger(ctx).Error("fastly_service_acl.listACL", "query_error", err)
			return nil, err
		}
		for _, item := range items {
			d.StreamListItem(ctx, item)
		}
	}

	return nil, nil
}

//// ACL Entry Parent Hydrate

/// HYDRATE FUNCTION

func getACL(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	name := d.EqualsQualString("name")
	serviceId := d.EqualsQualString("service_id")
	serviceVersion := d.EqualsQuals["service_version"].GetInt64Value()

	// check if the name is empty
	if name == "" || serviceId == "" {
		return nil, nil
	}

	serviceClient, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_acl.getACL", "connection_error", err)
		return nil, err
	}
	input := &fastly.GetACLInput{
		ServiceID:      serviceId,
		ServiceVersion: int(serviceVersion),
		Name:           name,
	}

	result, err := serviceClient.Client.GetACL(input)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_acl.getACL", "api_error", err)
		return nil, err
	}

	return result, nil
}
