package fastly

import (
	"context"

	"github.com/fastly/go-fastly/v3/fastly"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableFastlyPool(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fastly_service_pool",
		Description: "Pools in the Fastly account.",
		List: &plugin.ListConfig{
			KeyColumns:    plugin.OptionalColumns([]string{"service_id", "service_version"}),
			ParentHydrate: hydrateServiceVersion,
			Hydrate:       listPool,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Alphanumeric string identifying a Pool."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name for the Pool."},
			// Other columns
			{Name: "comment", Type: proto.ColumnType_STRING, Description: "A freeform descriptive note."},
			{Name: "connect_timeout", Type: proto.ColumnType_INT, Description: "How long to wait for a timeout in milliseconds. Optional."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp (UTC) of when the pool was created."},
			{Name: "deleted_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp (UTC) of when the pool was deleted."},
			{Name: "first_byte_timeout", Type: proto.ColumnType_INT, Description: "How long to wait for the first byte in milliseconds. Optional."},
			{Name: "healthcheck", Type: proto.ColumnType_STRING, Description: "Name of the healthcheck to use with this pool. Can be empty and could be reused across multiple backend and pools."},
			{Name: "max_conn_default", Type: proto.ColumnType_INT, Description: "Maximum number of connections."},
			{Name: "max_tls_version", Type: proto.ColumnType_STRING, Description: "Maximum allowed TLS version on connections to this server. Optional."},
			{Name: "min_tls_version", Type: proto.ColumnType_STRING, Description: "Minimum allowed TLS version on connections to this server. Optional."},
			{Name: "override_host", Type: proto.ColumnType_STRING, Description: "The hostname to override the Host header. Defaults to null meaning no override of the Host header will occur."},
			{Name: "pool_type", Type: proto.ColumnType_STRING, Description: "What type of load balance group to use: random, hash, client."},
			{Name: "quorum", Type: proto.ColumnType_INT, Description: "Percentage of capacity (0-100) that needs to be operationally available for a pool to be considered up."},
			{Name: "request_condition", Type: proto.ColumnType_STRING, Description: "Condition which, if met, will select this configuration during a request. Optional."},
			{Name: "service_id", Type: proto.ColumnType_STRING, Description: "Alphanumeric string identifying the service."},
			{Name: "service_version", Type: proto.ColumnType_INT, Description: "Integer identifying a service version."},
			{Name: "shield", Type: proto.ColumnType_STRING, Description: "Selected POP to serve as a shield for the servers. Defaults to null meaning no origin shielding if not set."},
			{Name: "tls_ca_cert", Type: proto.ColumnType_STRING, Description: "A secure certificate to authenticate a server with. Must be in PEM format."},
			{Name: "tls_cert_hostname", Type: proto.ColumnType_STRING, Description: "The hostname used to verify a server's certificate. It can either be the Common Name (CN) or a Subject Alternative Name (SAN)."},
			{Name: "tls_check_cert", Type: proto.ColumnType_BOOL, Description: "Be strict on checking TLS certs. Optional."},
			{Name: "tls_ciphers", Type: proto.ColumnType_STRING, Description: "List of OpenSSL ciphers."},
			{Name: "tls_client_cert", Type: proto.ColumnType_STRING, Description: "The client certificate used to make authenticated requests. Must be in PEM format."},
			{Name: "tls_client_key", Type: proto.ColumnType_STRING, Description: "The client private key used to make authenticated requests. Must be in PEM format."},
			{Name: "tls_sni_hostname", Type: proto.ColumnType_STRING, Description: "SNI hostname. Optional."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp (UTC) of when the pool was updated."},
			{Name: "use_tls", Type: proto.ColumnType_BOOL, Description: "Whether to use TLS."},
		},
	}
}

func listPool(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, _, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_pool.listPool", "connection_error", err)
		return nil, err
	}
	version := h.Item.(*fastly.Version)
	plugin.Logger(ctx).Warn("fastly_service_pool.listPool", "version", version)
	items, err := conn.ListPools(&fastly.ListPoolsInput{ServiceID: version.ServiceID, ServiceVersion: int(version.Number)})
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_pool.listPool", "query_error", err)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func hydrateServiceVersion(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, serviceID, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_version.listServiceVersion", "connection_error", err)
		return nil, err
	}

	if d.KeyColumnQuals["service_id"] != nil {

		serviceID = d.KeyColumnQuals["service_id"].GetStringValue()

		if d.KeyColumnQuals["service_version"] != nil {
			serviceVersion := int(d.KeyColumnQuals["service_version"].GetInt64Value())
			input := fastly.GetVersionInput{ServiceID: serviceID, ServiceVersion: serviceVersion}
			version, err := conn.GetVersion(&input)
			if err != nil {
				plugin.Logger(ctx).Error("fastly_service_version.listServiceVersion", "query_error", err, "input", input)
				return nil, err
			}
			d.StreamListItem(ctx, version)
		} else {
			// List all versions for the service
			items, err := conn.ListVersions(&fastly.ListVersionsInput{ServiceID: serviceID})
			if err != nil {
				plugin.Logger(ctx).Error("fastly_service_version.listServiceVersion", "query_error", err)
				return nil, err
			}
			for _, i := range items {
				d.StreamListItem(ctx, i)
			}
		}

	} else {

		// No service specified, so list all versions of all services
		services, err := conn.ListServices(&fastly.ListServicesInput{})
		if err != nil {
			plugin.Logger(ctx).Error("fastly_service.listService", "query_error", err)
			return nil, err
		}
		for _, i := range services {
			serviceVersion := int(i.ActiveVersion)
			if d.KeyColumnQuals["service_version"] != nil {
				serviceVersion = int(d.KeyColumnQuals["service_version"].GetInt64Value())
			}
			input := fastly.GetVersionInput{ServiceID: i.ID, ServiceVersion: serviceVersion}
			plugin.Logger(ctx).Warn("hydrateServiceVersion", "i.ID", i.ID)
			plugin.Logger(ctx).Warn("hydrateServiceVersion", "serviceVersion", serviceVersion)
			version, err := conn.GetVersion(&input)
			plugin.Logger(ctx).Warn("hydrateServiceVersion", "version", version)
			if err != nil {
				plugin.Logger(ctx).Error("fastly_service_version.listServiceVersion", "query_error", err, "input", input)
				return nil, err
			}
			d.StreamListItem(ctx, version)
		}

	}

	return nil, nil
}
