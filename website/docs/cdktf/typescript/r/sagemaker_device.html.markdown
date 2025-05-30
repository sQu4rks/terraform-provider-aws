---
subcategory: "SageMaker AI"
layout: "aws"
page_title: "AWS: aws_sagemaker_device"
description: |-
  Provides a SageMaker AI Device resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_sagemaker_device

Provides a SageMaker AI Device resource.

## Example Usage

### Basic usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SagemakerDevice } from "./.gen/providers/aws/sagemaker-device";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new SagemakerDevice(this, "example", {
      device: {
        deviceName: "example",
      },
      deviceFleetName: Token.asString(
        awsSagemakerDeviceFleetExample.deviceFleetName
      ),
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `deviceFleetName` - (Required) The name of the Device Fleet.
* `device` - (Required) The device to register with SageMaker AI Edge Manager. See [Device](#device) details below.

### Device

* `description` - (Required) A description for the device.
* `deviceName` - (Optional) The name of the device.
* `iotThingName` - (Optional) Amazon Web Services Internet of Things (IoT) object name.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The id is constructed from `device-fleet-name/device-name`.
* `arn` - The Amazon Resource Name (ARN) assigned by AWS to this Device.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import SageMaker AI Devices using the `device-fleet-name/device-name`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SagemakerDevice } from "./.gen/providers/aws/sagemaker-device";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    SagemakerDevice.generateConfigForImport(
      this,
      "example",
      "my-fleet/my-device"
    );
  }
}

```

Using `terraform import`, import SageMaker AI Devices using the `device-fleet-name/device-name`. For example:

```console
% terraform import aws_sagemaker_device.example my-fleet/my-device
```

<!-- cache-key: cdktf-0.20.8 input-fdd2ff8d47aba6fbcc83e6b503865508ca01826e328fbf1c740e252689e014e5 -->