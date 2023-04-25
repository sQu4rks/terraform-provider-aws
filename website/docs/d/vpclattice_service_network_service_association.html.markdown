---
subcategory: "VPC Lattice"
layout: "aws"
page_title: "AWS: aws_vpclattice_service_network_service_association"
description: |-
  Terraform data source for managing an AWS VPC Lattice Service Network Service Association.
---

# Data Source: aws_vpclattice_service_network_service_association

Terraform data source for managing an AWS VPC Lattice Service Network Service Association.

## Example Usage

### Basic Usage

```terraform
data "aws_vpclattice_service_network_service_association" "example" {
}
```

## Argument Reference

The following arguments are required:

* `service_network_service_association_identifier` - (Required) Id or Amazon Resource Name (ARN) of the association.


## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `arn` - ARN of the Service Network Service Association. 
* `created_at` Date and time that the association was created, specified in ISO-8601 format.
* `created_by` Account that created the association.
* `custom_domain_name` Custom domain name of the service
* `failure_code` Failure code
* `failure_message` Message of the failure
* `id` ID of the service network and service association
* `service_arn` ARN of the associated service
* `service_id` ID of the associated service
* `service_name` Name of the associated service
* `service_network_arn` ARN of the service network
* `service_network_id` ID of the service network
* `service_network_name` Name of the service network
* `status` Status of the association. Possible values: CREATE_IN_PROGRESS, ACTIVE, DELETE_IN_PROGRESS, CREATE_FAILED, DELETE_FAILED
* `dns_entry` Name of the DNS Service

dns entry (`dns_entry`) has the following attributes:

* `domain_name` Domain name of the service
* `hosted_zone_id` ID of the hosted zone