package fastly

import (
	"context"

	"github.com/fastly/go-fastly/v8/fastly"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableFastlyDictionary(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fastly_dictionary",
		Description: "Dictionaries for the service version.",
		List: &plugin.ListConfig{
			Hydrate: listDictionaries,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getDictionary,
			KeyColumns: plugin.SingleColumn("name"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the Dictionary.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the Dictionary.",
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
				Name:        "write_only",
				Type:        proto.ColumnType_BOOL,
				Description: "Determines if items in the dictionary are readable or not.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp (UTC) of when the Dictionary was created.",
			},
			{
				Name:        "deleted_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp (UTC) of when the Dictionary was deleted.",
			},

			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp (UTC) of when the Dictionary was updated.",
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

func listDictionaries(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	serviceClient, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_dictionary.listDictionaries", "connection_error", err)
		return nil, err
	}
	input := fastly.ListDictionariesInput{
		ServiceID:      serviceClient.ServiceID,
		ServiceVersion: serviceClient.ServiceVersion,
	}
	items, err := serviceClient.Client.ListDictionaries(&input)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_dictionary.listDictionaries", "api_error", err)
		return nil, err
	}
	for _, item := range items {
		d.StreamListItem(ctx, item)
	}

	return nil, nil
}

func getDictionary(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	name := d.EqualsQualString("name")

	// check if the name is empty
	if name == "" {
		return nil, nil
	}

	serviceClient, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_dictionary.getDictionary", "connection_error", err)
		return nil, err
	}

	input := &fastly.GetDictionaryInput{
		ServiceID:      serviceClient.ServiceID,
		ServiceVersion: serviceClient.ServiceVersion,
		Name:           name,
	}
	result, err := serviceClient.Client.GetDictionary(input)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_dictionary.getDictionary", "api_error", err)
		return nil, err
	}

	return result, nil
}
