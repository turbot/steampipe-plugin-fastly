package fastly

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableFastlyDataCenter(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fastly_data_center",
		Description: "Data centers in the Fastly network.",
		List: &plugin.ListConfig{
			Hydrate: listDataCenter,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "code", Type: proto.ColumnType_STRING, Description: "Data center location code, e.g. BNE."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the data center."},
			{Name: "location_group", Type: proto.ColumnType_STRING, Transform: transform.FromField("Group"), Description: "Location group, e.g. Europe."},
			{Name: "longitude", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("Coordinates.Longtitude"), Description: "Location longitude."},
			{Name: "latitude", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("Coordinates.Latitude"), Description: "Location latitude."},
			{Name: "shield", Type: proto.ColumnType_STRING, Description: "Data center shield."},
		},
	}
}

func listDataCenter(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, _, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_data_center.listDataCenter", "connection_error", err)
		return nil, err
	}
	items, err := conn.AllDatacenters()
	if err != nil {
		plugin.Logger(ctx).Error("fastly_data_center.listDataCenter", "query_error", err)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}
