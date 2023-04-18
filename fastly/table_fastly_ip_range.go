package fastly

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableFastlyIPRange(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fastly_ip_range",
		Description: "All IP ranges (v4 & v6) in the Fastly account.",
		List: &plugin.ListConfig{
			Hydrate: listIP,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "cidr", Type: proto.ColumnType_CIDR, Description: "IP range."},
			{Name: "version", Type: proto.ColumnType_INT, Description: "IP version for the range."},
		},
	}
}

type ipRow struct {
	Cidr    string `json:"cidr"`
	Version int    `json:"version"`
}

func listIP(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, _, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_ip_range.listIP", "connection_error", err)
		return nil, err
	}
	v4Items, v6Items, err := conn.AllIPs()
	if err != nil {
		plugin.Logger(ctx).Error("fastly_ip_range.listIP", "query_error", err)
		return nil, err
	}
	for _, i := range v4Items {
		d.StreamListItem(ctx, ipRow{i, 4})
	}
	for _, i := range v6Items {
		d.StreamListItem(ctx, ipRow{i, 6})
	}
	return nil, nil
}
