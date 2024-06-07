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
			KeyColumns: plugin.AllColumns([]string{"service_id", "service_version", "name"}),
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

	// We have not used the parent hydrate relation here because:
	// 1. Steampipe does not yet support parent hydrate chaining up to three levels deep (Service > ACL > ACL Entry).
	// 2. Since this table is used as the parent of `fastly_acl_entry`, we avoided using the parent hydrate concept here.

	// Check if the service details have been configured in connection config.
	var versions []*fastly.Version
	cfgVersions, err := configServiceVersionHydrate(ctx, d, h)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_acl.listACL.configServiceVersionHydrate", "api_error", err)
		return nil, err
	}

	if cfgVersions != nil {
		versions = cfgVersions.([]*fastly.Version)
	} else {
		// Fetch all the available services and its versions
		resp, err := listServicesVersionHydrate(ctx, d, h)
		if err != nil {
			return nil, err
		}

		versions = resp.([]*fastly.Version)
	}

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

/// HYDRATE FUNCTION

func getACL(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	name := d.EqualsQualString("name")
	serviceId := d.EqualsQualString("service_id")
	serviceVersion := d.EqualsQuals["service_version"].GetInt64Value()

	// check if the name is empty
	if name == "" || serviceId == "" || serviceVersion == 0 {
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
