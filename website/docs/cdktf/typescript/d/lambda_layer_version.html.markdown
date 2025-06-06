---
subcategory: "Lambda"
layout: "aws"
page_title: "AWS: aws_lambda_layer_version"
description: |-
  Provides a Lambda Layer Version data source.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_lambda_layer_version

Provides information about a Lambda Layer Version.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { VariableType, TerraformVariable, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsLambdaLayerVersion } from "./.gen/providers/aws/data-aws-lambda-layer-version";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    /*Terraform Variables are not always the best fit for getting inputs in the context of Terraform CDK.
    You can read more about this at https://cdk.tf/variables*/
    const layerName = new TerraformVariable(this, "layer_name", {
      type: VariableType.STRING,
    });
    new DataAwsLambdaLayerVersion(this, "existing", {
      layerName: layerName.stringValue,
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `layerName` - (Required) Name of the lambda layer.
* `version` - (Optional) Specific layer version. Conflicts with `compatibleRuntime` and `compatibleArchitecture`. If omitted, the latest available layer version will be used.
* `compatibleRuntime` (Optional) Specific runtime the layer version must support. Conflicts with `version`. If specified, the latest available layer version supporting the provided runtime will be used.
* `compatibleArchitecture` (Optional) Specific architecture the layer version could support. Conflicts with `version`. If specified, the latest available layer version supporting the provided architecture will be used.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `codeSha256` - Base64-encoded representation of raw SHA-256 sum of the zip file.
* `description` - Description of the specific Lambda Layer version.
* `licenseInfo` - License info associated with the specific Lambda Layer version.
* `compatibleRuntimes` - List of [Runtimes][1] the specific Lambda Layer version is compatible with.
* `compatibleArchitectures` - A list of [Architectures][2] the specific Lambda Layer version is compatible with.
* `arn` - ARN of the Lambda Layer with version.
* `layerArn` - ARN of the Lambda Layer without version.
* `createdDate` - Date this resource was created.
* `signingJobArn` - ARN of a signing job.
* `signingProfileVersionArn` - The ARN for a signing profile version.
* `sourceCodeHash` - (**Deprecated** use `codeSha256` instead) Base64-encoded representation of raw SHA-256 sum of the zip file.
* `sourceCodeSize` - Size in bytes of the function .zip file.
* `version` - This Lambda Layer version.

[1]: https://docs.aws.amazon.com/lambda/latest/dg/API_GetLayerVersion.html#SSS-GetLayerVersion-response-CompatibleRuntimes
[2]: https://docs.aws.amazon.com/lambda/latest/dg/API_GetLayerVersion.html#SSS-GetLayerVersion-response-CompatibleArchitectures

<!-- cache-key: cdktf-0.20.8 input-7bb4746c172a77d961c97787816afef2700b5b13fec8261bf752baf3f6f7c4be -->