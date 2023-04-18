package fastly

import (
	"context"

	"github.com/fastly/go-fastly/v8/fastly"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableFastlyDictionary(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fastly_dictionary",
		Description: "Dictionaries for the service version.",
		List: &plugin.ListConfig{
			Hydrate:    listDictionary,
			KeyColumns: plugin.AllColumns([]string{"service_version"}),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getDictionary,
			KeyColumns: plugin.AllColumns([]string{"service_version", "name"}),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the Dictionary."},
			{Name: "write_only", Type: proto.ColumnType_BOOL, Description: "Determines if items in the dictionary are readable or not."},
			{Name: "data", Type: proto.ColumnType_JSON, Hydrate: getDictionaryItems, Description: "Key pair items in the dictionary."},
			{Name: "items", Type: proto.ColumnType_JSON, Hydrate: getDictionaryItems, Description: "Key pair items in the dictionary."},
			// Other columns
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp (UTC) of when the Dictionary was created."},
			{Name: "deleted_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp (UTC) of when the Dictionary was deleted."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The ID of the Dictionary."},
			{Name: "service_id", Type: proto.ColumnType_STRING, Description: "Alphanumeric string identifying the service."},
			{Name: "service_version", Type: proto.ColumnType_INT, Description: "Integer identifying a service version."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp (UTC) of when the Dictionary was updated."},
		},
	}
}

type dictionaryItemRow struct {
	Items []*fastly.DictionaryItem
	Data  map[string]string
}

func listDictionary(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, serviceID, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_dictionary.listDictionary", "connection_error", err)
		return nil, err
	}
	input := fastly.ListDictionariesInput{
		ServiceID:      serviceID,
		ServiceVersion: int(d.EqualsQuals["service_version"].GetInt64Value()),
	}
	items, err := conn.ListDictionaries(&input)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_dictionary.listDictionary", "query_error", err)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getDictionary(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, serviceID, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_dictionary.getDictionary", "connection_error", err)
		return nil, err
	}
	serviceVersion := int(d.EqualsQuals["service_version"].GetInt64Value())
	name := d.EqualsQuals["name"].GetStringValue()
	input := fastly.GetDictionaryInput{ServiceID: serviceID, ServiceVersion: serviceVersion, Name: name}
	result, err := conn.GetDictionary(&input)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_dictionary.getDictionary", "query_error", err, "input", input)
		return nil, err
	}
	return result, nil
}

type dictionaryData struct {
	Items []*fastly.DictionaryItem
	Data  map[string]string
}

func getDictionaryItems(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, _, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_dictionary.getDictionaryItems", "connection_error", err)
		return nil, err
	}
	dict := h.Item.(*fastly.Dictionary)
	input := fastly.ListDictionaryItemsInput{
		ServiceID:    dict.ServiceID,
		DictionaryID: dict.ID,
	}
	items, err := conn.ListDictionaryItems(&input)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_dictionary.getDictionaryItems", "query_error", err)
		return nil, err
	}
	data := map[string]string{}
	for _, item := range items {
		data[item.ItemKey] = item.ItemValue
	}
	return dictionaryData{Items: items, Data: data}, nil
}
