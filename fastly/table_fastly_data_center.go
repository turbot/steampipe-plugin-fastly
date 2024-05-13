package fastly

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableFastlyDataCenter(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fastly_data_center",
		Description: "Data centers in the Fastly network.",
		List: &plugin.ListConfig{
			Hydrate: listDataCenters,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the data center.",
			},
			{
				Name:        "code",
				Type:        proto.ColumnType_STRING,
				Description: "Data center location code, e.g. BNE.",
			},
			// group is a keyword in PostgreSQL, so here transform function has been used
			{
				Name:        "location_group",
				Type:        proto.ColumnType_STRING,
				Description: "Location group, e.g. Europe.",
				Transform:   transform.FromField("Group"),
			},
			{
				Name:        "longitude",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Coordinates.Longtitude"),
				Description: "Location longitude.",
			},
			{
				Name:        "latitude",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Coordinates.Latitude"),
				Description: "Location latitude.",
			},
			{
				Name:        "shield",
				Type:        proto.ColumnType_STRING,
				Description: "Data center shield.",
			},
			{
				Name:        "service_id",
				Type:        proto.ColumnType_STRING,
				Description: "Alphanumeric string identifying the service.",
				Hydrate:     getServiceId,
				Transform:   transform.FromValue(),
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

func listDataCenters(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	serviceClient, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_data_center.listDataCenters", "connection_error", err)
		return nil, err
	}

	items, err := serviceClient.Client.AllDatacenters()
	if err != nil {
		plugin.Logger(ctx).Error("fastly_data_center.listDataCenters", "api_error", err)
		return nil, err
	}

	for _, item := range items {
		d.StreamListItem(ctx, item)
	}

	return nil, nil
}
