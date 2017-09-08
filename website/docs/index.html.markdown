---
layout: "etcd"
page_title: "Provider: Etcd"
sidebar_current: "docs-etcd-index"
description: |-
  The etcd provider is used to interact with an existing etcd cluster.
---

# Etcd Provider

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
provider "etcd" {
  endpoints = [
    "https://${vars.node1_url}:2379",
    "https://${vars.node2_url}:2379",
  ]

  dial_timeout       = "5s"
  auto_sync_interval = "5s"

  tls_trusted_ca  = "${file(cluster-ca-cert.pem)}"
  tls_cert        = "${file(client-cert.pem}"
  tls_key         = "${flie(client-key.pem)}"

  username = "username"
  password = "password"

  reject_old_cluster = true
}

# Configure an additional member 
resource "digitalocean_droplet" "web" {
  name = "node3"
  peer_urls = [
    "https://${vars.new_node_url}:2380"
  ]
}
```

## Argument Reference

The following arguments are supported:

* `endpoints` - (Required) A list of client endpoints for an existing etcd cluster.
* `dial_timeout` - (Required) Etcd config dial timeout.
* `auto_sync_interval` - (Optional) Etcd auto sync interval.
* `tls_trusted_ca` - (Optional) Raw client auth certificate authority.
* `tls_cert` - (Optional) Raw client auth certificate.
* `tls_key` - (Optional) Raw client auth key.
* `username` - (Optional) Cluster authentication username.
* `password` - (Optional) Cluster authentication password.
* `reject_old_cluster` - (Optional) Boolean to reject old clusters.

For more information about these options, see the [etcd.clienv3 `type Config` godoc](https://godoc.org/github.com/coreos/etcd/clientv3#Config).
