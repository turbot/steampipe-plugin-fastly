package fastly

import (
	"context"
	"strings"

	"github.com/fastly/go-fastly/v3/fastly"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableFastlyToken(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fastly_token",
		Description: "Tokens in the Fastly account.",
		List: &plugin.ListConfig{
			Hydrate: listToken,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the token."},
			// Other columns
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Time-stamp (UTC) of when the token was created."},
			{Name: "expires_at", Type: proto.ColumnType_TIMESTAMP, Description: "Time-stamp (UTC) of when the token will expire (optional)."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Alphanumeric string identifying a token."},
			{Name: "ip", Type: proto.ColumnType_IPADDR, Description: "IP Address of the client that last used the token."},
			{Name: "last_used_at", Type: proto.ColumnType_TIMESTAMP, Description: "Time-stamp (UTC) of when the token was last used."},
			{Name: "scopes", Type: proto.ColumnType_JSON, Transform: transform.FromField("Scope").Transform(scopesStringToArray), Description: "List of authorization scopes."},
			{Name: "services", Type: proto.ColumnType_JSON, Description: "List of alphanumeric strings identifying services (optional). If no services are specified, the token will have access to all services on the account."},
			{Name: "user_id", Type: proto.ColumnType_STRING, Description: "Alphanumeric string identifying the user."},
		},
	}
}

func listToken(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, _, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_token.listToken", "connection_error", err)
		return nil, err
	}
	items, err := conn.ListTokens()
	if err != nil {
		plugin.Logger(ctx).Error("fastly_token.listToken", "query_error", err)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
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
