// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package ssmincidents

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssmincidents"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
)

// ListTags lists ssmincidents service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func ListTags(ctx context.Context, conn *ssmincidents.Client, identifier string) (tftags.KeyValueTags, error) {
	input := &ssmincidents.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(ctx, input)

	if err != nil {
		return tftags.New(ctx, nil), err
	}

	return KeyValueTags(ctx, output.Tags), nil
}

// ListTags lists ssmincidents service tags and set them in Context.
// It is called from outside this package.
func (p *servicePackage) ListTags(ctx context.Context, meta any, identifier string) error {
	tags, err := ListTags(ctx, meta.(*conns.AWSClient).SSMIncidentsClient(), identifier)

	if err != nil {
		return err
	}

	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = types.Some(tags)
	}

	return nil
}

// map[string]string handling

// Tags returns ssmincidents service tags.
func Tags(tags tftags.KeyValueTags) map[string]string {
	return tags.Map()
}

// KeyValueTags creates KeyValueTags from ssmincidents service tags.
func KeyValueTags(ctx context.Context, tags map[string]string) tftags.KeyValueTags {
	return tftags.New(ctx, tags)
}

// GetTagsIn returns ssmincidents service tags from Context.
// nil is returned if there are no input tags.
func GetTagsIn(ctx context.Context) map[string]string {
	if inContext, ok := tftags.FromContext(ctx); ok {
		if tags := Tags(inContext.TagsIn.UnwrapOrDefault()); len(tags) > 0 {
			return tags
		}
	}

	return nil
}

// SetTagsOut sets ssmincidents service tags in Context.
func SetTagsOut(ctx context.Context, tags map[string]string) {
	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = types.Some(KeyValueTags(ctx, tags))
	}
}