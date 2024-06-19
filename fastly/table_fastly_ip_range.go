package fastly

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableFastlyIPRange(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fastly_ip_range",
		Description: "All IP ranges (v4 & v6) in the Fastly account.",
		List: &plugin.ListConfig{
			Hydrate: listIPs,
		},
		Columns: []*plugin.Column{
			{
				Name:        "cidr",
				Type:        proto.ColumnType_CIDR,
				Description: "IP range.",
			},
			{
				Name:        "version",
				Type:        proto.ColumnType_INT,
				Description: "IP version for the range.",
			},
			{
				Name:        "service_id",
				Type:        proto.ColumnType_STRING,
				Description: "Alphanumeric string identifying the service.",
				Hydrate:     getServiceId,
				Transform:   transform.FromValue(),
			},
		},
	}
}

type ipRow struct {
	Cidr    string `json:"cidr"`
	Version int    `json:"version"`
}

/// LIST FUNCTION

func listIPs(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	serviceClient, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_ip_range.listIPs", "connection_error", err)
		return nil, err
	}

	v4Items, v6Items, err := serviceClient.Client.AllIPs()
	if err != nil {
		plugin.Logger(ctx).Error("fastly_ip_range.listIPs", "api_error", err)
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
