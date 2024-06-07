package fastly

import (
	"context"

	"github.com/fastly/go-fastly/v8/fastly"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableFastlyPool(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fastly_pool",
		Description: "Pools in the Fastly account.",
		List: &plugin.ListConfig{
			ParentHydrate: listServiceVersionsByConfig,
			Hydrate:       listPools,
			KeyColumns:    plugin.OptionalColumns([]string{"service_id", "service_version"}),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getPool,
			KeyColumns: plugin.AllColumns([]string{"service_id", "name", "service_version"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "Alphanumeric string identifying a Pool.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Name for the Pool.",
			},
			{
				Name:        "comment",
				Type:        proto.ColumnType_STRING,
				Description: "A freeform descriptive note.",
			},
			{
				Name:        "connect_timeout",
				Type:        proto.ColumnType_INT,
				Description: "How long to wait for a timeout in milliseconds.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp (UTC) of when the pool was created.",
			},
			{
				Name:        "deleted_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp (UTC) of when the pool was deleted.",
			},
			{
				Name:        "first_byte_timeout",
				Type:        proto.ColumnType_INT,
				Description: "How long to wait for the first byte in milliseconds. Optional.",
			},
			{
				Name:        "healthcheck",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the healthcheck to use with this pool. Can be empty and could be reused across multiple backend and pools.",
			},
			{
				Name:        "max_conn_default",
				Type:        proto.ColumnType_INT,
				Description: "Maximum number of connections.",
			},
			{
				Name:        "max_tls_version",
				Type:        proto.ColumnType_STRING,
				Description: "Maximum allowed TLS version on connections to this server. Optional.",
			},
			{
				Name:        "min_tls_version",
				Type:        proto.ColumnType_STRING,
				Description: "Minimum allowed TLS version on connections to this server. Optional.",
			},
			{
				Name:        "override_host",
				Type:        proto.ColumnType_STRING,
				Description: "The hostname to override the Host header. Defaults to null meaning no override of the Host header will occur.",
			},
			{
				Name:        "pool_type",
				Type:        proto.ColumnType_STRING,
				Description: "What type of load balance group to use: random, hash, client.",
			},
			{
				Name:        "quorum",
				Type:        proto.ColumnType_INT,
				Description: "Percentage of capacity (0-100) that needs to be operationally available for a pool to be considered up.",
			},
			{
				Name:        "request_condition",
				Type:        proto.ColumnType_STRING,
				Description: "Condition which, if met, will select this configuration during a request. Optional.",
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
				Name:        "shield",
				Type:        proto.ColumnType_STRING,
				Description: "Selected POP to serve as a shield for the servers. Defaults to null meaning no origin shielding if not set.",
			},
			{
				Name:        "tls_ca_cert",
				Type:        proto.ColumnType_STRING,
				Description: "A secure certificate to authenticate a server with. Must be in PEM format.",
			},
			{
				Name:        "tls_cert_hostname",
				Type:        proto.ColumnType_STRING,
				Description: "The hostname used to verify a server's certificate. It can either be the Common Name (CN) or a Subject Alternative Name (SAN).",
			},
			{
				Name:        "tls_check_cert",
				Type:        proto.ColumnType_BOOL,
				Description: "Be strict on checking TLS certs. Optional.",
			},
			{
				Name:        "tls_ciphers",
				Type:        proto.ColumnType_STRING,
				Description: "List of OpenSSL ciphers.",
			},
			{
				Name:        "tls_client_cert",
				Type:        proto.ColumnType_STRING,
				Description: "The client certificate used to make authenticated requests. Must be in PEM format.",
			},
			{
				Name:        "tls_client_key",
				Type:        proto.ColumnType_STRING,
				Description: "The client private key used to make authenticated requests. Must be in PEM format.",
			},
			{
				Name:        "tls_sni_hostname",
				Type:        proto.ColumnType_STRING,
				Description: "SNI hostname. Optional.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp (UTC) of when the pool was updated.",
			},
			{
				Name:        "use_tls",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether to use TLS.",
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

func listPools(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	serviceVersion := h.Item.(*fastly.Version)

	if d.EqualsQualString("service_id") != "" && d.EqualsQualString("service_id") != serviceVersion.ServiceID {
		return nil, nil
	}

	if d.EqualsQuals["service_version"] != nil && int(d.EqualsQuals["service_version"].GetInt64Value()) != serviceVersion.Number {
		return nil, nil
	}

	serviceClient, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_pool.listPools", "connection_error", err)
		return nil, err
	}
	input := &fastly.ListPoolsInput{
		ServiceID:      serviceVersion.ServiceID,
		ServiceVersion: serviceVersion.Number,
	}
	items, err := serviceClient.Client.ListPools(input)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_pool.listPools", "api_error", err)
		return nil, err
	}

	for _, item := range items {
		d.StreamListItem(ctx, item)
	}

	return nil, nil
}

/// HYDRATE FUNCTION

func getPool(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	name := d.EqualsQualString("name")
	serviceId := d.EqualsQualString("service_id")
	serviceVersion := d.EqualsQuals["service_version"].GetInt64Value()

	// check if the name is empty
	if name == "" || serviceId == "" || serviceVersion == 0 {
		return nil, nil
	}

	serviceClient, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_pool.getPool", "connection_error", err)
		return nil, err
	}

	input := &fastly.GetPoolInput{
		ServiceID:      serviceId,
		ServiceVersion: int(serviceVersion),
		Name:           name,
	}
	result, err := serviceClient.Client.GetPool(input)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_pool.getPool", "api_error", err)
		return nil, err
	}

	return result, nil
}
