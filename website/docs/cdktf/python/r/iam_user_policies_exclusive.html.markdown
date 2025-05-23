---
subcategory: "IAM (Identity & Access Management)"
layout: "aws"
page_title: "AWS: aws_iam_user_policies_exclusive"
description: |-
  Terraform resource for maintaining exclusive management of inline policies assigned to an AWS IAM (Identity & Access Management) user.
---

<!-- Please do not edit this file, it is generated. -->
# Resource: aws_iam_user_policies_exclusive

Terraform resource for maintaining exclusive management of inline policies assigned to an AWS IAM (Identity & Access Management) user.

!> This resource takes exclusive ownership over inline policies assigned to a user. This includes removal of inline policies which are not explicitly configured. To prevent persistent drift, ensure any `aws_iam_user_policy` resources managed alongside this resource are included in the `policy_names` argument.

~> Destruction of this resource means Terraform will no longer manage reconciliation of the configured inline policy assignments. It __will not__ delete the configured policies from the user.

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.iam_user_policies_exclusive import IamUserPoliciesExclusive
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        IamUserPoliciesExclusive(self, "example",
            policy_names=[Token.as_string(aws_iam_user_policy_example.name)],
            user_name=Token.as_string(aws_iam_user_example.name)
        )
```

### Disallow Inline Policies

To automatically remove any configured inline policies, set the `policy_names` argument to an empty list.

~> This will not __prevent__ inline policies from being assigned to a user via Terraform (or any other interface). This resource enables bringing inline policy assignments into a configured state, however, this reconciliation happens only when `apply` is proactively run.

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.iam_user_policies_exclusive import IamUserPoliciesExclusive
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name, *, policyNames):
        super().__init__(scope, name)
        IamUserPoliciesExclusive(self, "example",
            policy_names=policy_names,
            user_name=Token.as_string(aws_iam_user_example.name)
        )
```

## Argument Reference

The following arguments are required:

* `user_name` - (Required) IAM user name.
* `policy_names` - (Required) A list of inline policy names to be assigned to the user. Policies attached to this user but not configured in this argument will be removed.

## Attribute Reference

This resource exports no additional attributes.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to exclusively manage inline policy assignments using the `user_name`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.iam_user_policies_exclusive import IamUserPoliciesExclusive
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        IamUserPoliciesExclusive.generate_config_for_import(self, "example", "MyUser")
```

Using `terraform import`, import exclusive management of inline policy assignments using the `user_name`. For example:

```console
% terraform import aws_iam_user_policies_exclusive.example MyUser
```

<!-- cache-key: cdktf-0.20.8 input-f755ace1836f5ffefbfa467551b71213634759995776bb9166a862b503304471 -->