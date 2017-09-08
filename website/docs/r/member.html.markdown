---
layout: "etcd"
page_title: "etcd: etcd_member"
sidebar_current: "docs-etcd-resource-member"
description: |-
  Configures a new etcd node with an existing etcd cluster.
---

# etcd\_member

Configures a new etcd node with an existing etcd cluster.

## Example Usage

```hcl
# Configure a node with an etcd cluster
resource "etcd_member" "node3" {
  name = "infra3"
  peer_urls = [
    "http://infra3.etcd.${vars.domain}:2380"
  ]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the domain
* `peer_urls` - (Required) A list of peer IP addresses or URLs.

## Attributes Reference

The following attributes are exported:

* `id` - The base10 etcd member ID.
