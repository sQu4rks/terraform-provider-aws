---
subcategory: "Bedrock"
layout: "aws"
page_title: "AWS: aws_bedrock_foundation_models"
description: |-
  Terraform data source for managing AWS Bedrock Foundation Models.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_bedrock_foundation_models

Terraform data source for managing AWS Bedrock Foundation Models.

## Example Usage

### Basic Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsBedrockFoundationModels } from "./.gen/providers/aws/data-aws-bedrock-foundation-models";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsBedrockFoundationModels(this, "test", {});
  }
}

```

### Filter by Inference Type

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsBedrockFoundationModels } from "./.gen/providers/aws/data-aws-bedrock-foundation-models";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsBedrockFoundationModels(this, "test", {
      byInferenceType: "ON_DEMAND",
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `byCustomizationType` - (Optional) Customization type to filter on. Valid values are `FINE_TUNING`.
* `byInferenceType` - (Optional) Inference type to filter on. Valid values are `ON_DEMAND` and `PROVISIONED`.
* `byOutputModality` - (Optional) Output modality to filter on. Valid values are `TEXT`, `IMAGE`, and `EMBEDDING`.
* `byProvider` - (Optional) Model provider to filter on.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `id` - AWS region.
* `modelSummaries` - List of model summary objects. See [`modelSummaries`](#model_summaries).

### `modelSummaries`

* `customizationsSupported` - Customizations that the model supports.
* `inferenceTypesSupported` - Inference types that the model supports.
* `inputModalities` - Input modalities that the model supports.
* `modelArn` - Model ARN.
* `modelId` - Model identifier.
* `modelName` - Model name.
* `outputModalities` - Output modalities that the model supports.
* `providerName` - Model provider name.
* `responseStreamingSupported` - Indicates whether the model supports streaming.

<!-- cache-key: cdktf-0.20.8 input-227a0115e2e4bd66dcfd4db719eba9ea0fe83af6d032b25579fd0d750d3c565a -->