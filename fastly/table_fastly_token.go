package fastly

import (
	"context"
	"strings"

	"github.com/fastly/go-fastly/v8/fastly"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableFastlyToken(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fastly_token",
		Description: "Tokens in the Fastly account.",
		List: &plugin.ListConfig{
			Hydrate: listTokens,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "Alphanumeric string identifying a token.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the token.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Time-stamp (UTC) of when the token was created.",
			},
			{
				Name:        "expires_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Time-stamp (UTC) of when the token will expire (optional).",
			},
			{
				Name:        "ip",
				Type:        proto.ColumnType_IPADDR,
				Description: "IP Address of the client that last used the token.",
				Transform:   transform.FromField("IP"),
			},
			{
				Name:        "last_used_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Time-stamp (UTC) of when the token was last used.",
			},
			{
				Name:        "scopes",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Scope").Transform(scopesStringToArray),
				Description: "List of authorization scopes.",
			},
			{
				Name:        "services",
				Type:        proto.ColumnType_JSON,
				Description: "List of alphanumeric strings identifying services (optional). If no services are specified, the token will have access to all services on the account.",
			},
			{
				Name:        "user_id",
				Type:        proto.ColumnType_STRING,
				Description: "Alphanumeric string identifying the user.",
				Transform:   transform.FromField("UserID"),
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

func listTokens(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	serviceClient, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_token.listTokens", "connection_error", err)
		return nil, err
	}

	items, err := serviceClient.Client.ListTokens()
	if err != nil {
		plugin.Logger(ctx).Error("fastly_token.listTokens", "query_error", err)
		return nil, err
	}

	for _, item := range items {
		d.StreamListItem(ctx, item)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func scopesStringToArray(_ context.Context, d *transform.TransformData) (interface{}, error) {
	if d.Value == nil {
		return []string{}, nil
	}
	scopes := d.Value.(fastly.TokenScope)
	return strings.Split(string(scopes), " "), nil
}
