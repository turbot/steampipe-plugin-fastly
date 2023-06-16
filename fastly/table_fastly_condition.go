package fastly

import (
	"context"

	"github.com/fastly/go-fastly/v8/fastly"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// 	TABLE DEFINITION

func tableFastlyCondition(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fastly_condition",
		Description: "Conditions defined in the service version.",
		List: &plugin.ListConfig{
			Hydrate: listConditions,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getCondition,
			KeyColumns: plugin.SingleColumn("name"),
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
				Name:        "type",
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

/// LIST FUNCTION

func listConditions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	serviceClient, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_condition.listConditions", "connection_error", err)
		return nil, err
	}

	input := &fastly.ListConditionsInput{
		ServiceID:      serviceClient.ServiceID,
		ServiceVersion: serviceClient.ServiceVersion,
	}
	items, err := serviceClient.Client.ListConditions(input)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_condition.listConditions", "api_error", err)
		return nil, err
	}

	for _, item := range items {
		d.StreamListItem(ctx, item)
	}

	return nil, nil
}

/// HYDRATE FUNCTION

func getCondition(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
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
		ServiceVersion: serviceClient.ServiceVersion,
		Name:           name,
	}

	result, err := serviceClient.Client.GetCondition(input)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_condition.getCondition", "api_error", err)
		return nil, err
	}

	return result, nil
}
