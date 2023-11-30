---
title: "Steampipe Table: fastly_data_center - Query Fastly Data Centers using SQL"
description: "Allows users to query Fastly Data Centers, providing insights into their geographical distribution and operational status."
---

# Table: fastly_data_center - Query Fastly Data Centers using SQL

Fastly is a cloud computing service provider that offers an edge cloud platform, which moves data and applications closer to users to improve their digital experiences. Fastly's edge cloud platform enables customers to create great digital experiences quickly, securely, and reliably by processing, serving, and securing its customers' applications as close to their end-users as possible. Fastly Data Centers are the physical locations where these edge servers are placed, and they play a crucial role in delivering fast, secure, and scalable online experiences.

## Table Usage Guide

The `fastly_data_center` table provides insights into Fastly's data center locations and their operational status. As a network engineer or site reliability engineer, this table can be used to understand the geographical distribution of Fastly data centers, their operational status, and other relevant details. This information can be crucial in planning for network expansion, managing network traffic, and ensuring optimal performance of your digital experiences.

## Examples

### Basic info
Discover the segments that utilize specific data centers by pinpointing their names, codes, and geographical locations. This is useful for understanding the distribution and reach of your data infrastructure across different regions.

```sql
select
  name,
  code,
  location_group,
  longitude,
  latitude,
  shield
from
  fastly_data_center;
```

### Show data center detail for a particular code
Explore the specific details of a data center based on its unique code. This can help in understanding its geographical location and security settings, which can be useful for resource allocation and risk assessment.

```sql
select
  name,
  code,
  location_group,
  longitude,
  latitude,
  shield
from
  fastly_data_center
where
  code = 'BNE';
```

### List data centers in Europe
Explore which data centers are located in Europe, allowing you to understand their geographical distribution and plan your data management strategy accordingly.

```sql
select
  name,
  code,
  location_group,
  longitude,
  latitude,
  shield
from
  fastly_data_center
where
  location_group = 'Europe';
```

### Show data center count by location group
Analyze the distribution of data centers across various location groups. This offers a clear perspective on where resources are concentrated, helping to strategize future infrastructure expansion or optimization.

```sql
select
  location_group,
  count(*)
from
  fastly_data_center
group by
  location_group
order by
  count desc;
```