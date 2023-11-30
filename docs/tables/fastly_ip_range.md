---
title: "Steampipe Table: fastly_ip_range - Query Fastly IP Ranges using SQL"
description: "Allows users to query Fastly IP Ranges, specifically the start and end of the IP range, providing insights into the network configuration."
---

# Table: fastly_ip_range - Query Fastly IP Ranges using SQL

Fastly is a cloud computing services provider that offers a content delivery network, edge computing, cloud security, and streaming and video delivery services. Fastly's IP ranges are used by their content delivery network to serve content to end users. These IP ranges are critical in understanding the network configuration and ensuring optimal content delivery.

## Table Usage Guide

The `fastly_ip_range` table provides insights into the IP ranges used by Fastly's content delivery network. As a network engineer, explore the start and end of each IP range through this table, which can help in understanding the network configuration and optimizing content delivery. Utilize it to uncover information about the network, such as the allocation of IP ranges and the reach of the content delivery network.

## Examples

### Basic info
Explore which IP ranges are being used and their corresponding versions in your Fastly network. This can help optimize network performance and security by ensuring you're using the most up-to-date IP versions.

```sql
select
  cidr,
  version
from
  fastly_ip_range;
```

### List IP ranges within a particular CIDR range
Explore which IP ranges fall within a specific CIDR range to better manage and monitor network traffic and security. This is particularly useful in identifying potential network vulnerabilities and planning for network expansion.

```sql
select
  cidr,
  version
from
  fastly_ip_range
where
  cidr << '199.0.0.0/8';
```

### Show IPv6 ranges
Explore the segments of your network that are using IPv6. This can help identify areas for potential upgrades or troubleshooting.

```sql
select
  cidr,
  version
from
  fastly_ip_range
where
  version = 6;
```