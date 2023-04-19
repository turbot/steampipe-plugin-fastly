package fastly

import (
	"context"

	"github.com/fastly/go-fastly/v8/fastly"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableFastlyCondition(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fastly_condition",
		Description: "Conditions defined in the service version.",
		List: &plugin.ListConfig{
			ParentHydrate: listServiceVersions,
			Hydrate:       listConditions,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "service_version",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getCondition,
			KeyColumns: plugin.AllColumns([]string{"service_version", "name"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the condition.",
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
				Name:        "comment",
				Type:        proto.ColumnType_STRING,
				Description: "A freeform descriptive note.",
			},
			{
				Name:        "condition_type",
				Type:        proto.ColumnType_STRING,
				Description: "Type of the condition.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp (UTC) of when the condition was created.",
			},
			{
				Name:        "deleted_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp (UTC) of when the condition was deleted.",
			},
			{
				Name:        "priority",
				Type:        proto.ColumnType_INT,
				Description: "Priority determines execution order. Lower numbers execute first.",
			},
			{
				Name:        "statement",
				Type:        proto.ColumnType_STRING,
				Description: "A conditional expression in VCL used to determine if the condition is met.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp (UTC) of when the condition was updated.",
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

func listConditions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	version := h.Item.(*fastly.Version)

	// check if the provided service_version is not matching with the parentHydrate
	if d.EqualsQuals["service_version"] != nil && int(d.EqualsQuals["service_version"].GetInt64Value()) != version.Number {
		return nil, nil
	}

	serviceClient, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_condition.listConditions", "connection_error", err)
		return nil, err
	}

	items, err := serviceClient.Client.ListConditions(&fastly.ListConditionsInput{ServiceID: version.ServiceID, ServiceVersion: version.Number})
	if err != nil {
		plugin.Logger(ctx).Error("fastly_condition.listConditions", "api_error", err)
		return nil, err
	}

	for _, item := range items {
		d.StreamListItem(ctx, item)
	}

	return nil, nil
}

func getCondition(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	serviceVersion := int(d.EqualsQuals["service_version"].GetInt64Value())
	name := d.EqualsQualString("name")

	// check if the name is empty
	if name == "" {
		return nil, nil
	}

	serviceClient, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_condition.getCondition", "connection_error", err)
		return nil, err
	}

	input := &fastly.GetConditionInput{
		ServiceID:      serviceClient.ServiceID,
		ServiceVersion: serviceVersion,
		Name:           name,
	}
	result, err := serviceClient.Client.GetCondition(input)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_condition.getCondition", "api_error", err)
		return nil, err
	}

	return result, nil
}
