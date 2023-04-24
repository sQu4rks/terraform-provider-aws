package vpclattice

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/vpclattice"
	"github.com/aws/aws-sdk-go-v2/service/vpclattice/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// Function annotations are used for datasource registration to the Provider. DO NOT EDIT.
// @SDKDataSource("aws_vpclattice_service_network_service_association")
func DataSourceServiceNetworkServiceAssociation() *schema.Resource {
	return &schema.Resource{
		ReadWithoutTimeout: dataSourceServiceNetworkServiceAssociationRead,

		Schema: map[string]*schema.Schema{
			"arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"custom_domain_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_entry": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"hosted_zone_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"failure_code": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"failure_message": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_network_arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_network_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_network_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_network_service_association_identifier": {
				Type:     schema.TypeString,
				Required: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": tftags.TagsSchemaComputed(), // TIP: Many, but not all, data sources have `tags` attributes.
		},
	}
}

const (
	DSNameServiceNetworkServiceAssociation = "Service Network Service Association Data Source"
)

func dataSourceServiceNetworkServiceAssociationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).VPCLatticeClient()

	service_network_service_association_id := d.Get("service_network_service_association_identifier").(string)

	out, err := findServiceNetworkServiceAssociationById(ctx, conn, service_network_service_association_id)

	if err != nil {
		return create.DiagError(names.VPCLattice, create.ErrActionReading, DSNameServiceNetworkServiceAssociation, service_network_service_association_id, err)
	}

	d.SetId(aws.ToString(out.Id))

	d.Set("arn", out.Arn)
	d.Set("created_at", aws.ToTime(out.CreatedAt).String())
	d.Set("created_by", out.CreatedBy)
	d.Set("custom_domain_name", out.CustomDomainName)
	d.Set("failure_code", out.FailureCode)
	d.Set("failure_message", out.FailureMessage)
	d.Set("id", out.Id)
	d.Set("service_arn", out.ServiceArn)
	d.Set("service_id", out.ServiceId)
	d.Set("service_name", out.ServiceName)
	d.Set("service_network_arn", out.ServiceNetworkArn)
	d.Set("service_network_id", out.ServiceNetworkId)
	d.Set("service_network_name", out.ServiceNetworkName)
	d.Set("status", out.Status)

	if out.DnsEntry != nil {
		if err := d.Set("dns_entry", []interface{}{flattenDNSEntry(out.DnsEntry)}); err != nil {
			return diag.Errorf("setting dns_entry: %s", err)
		}
	} else {
		d.Set("dns_entry", nil)
	}

	ignoreTagsConfig := meta.(*conns.AWSClient).IgnoreTagsConfig
	tags, err := ListTags(ctx, conn, aws.ToString(out.Arn))

	if err != nil {
		return create.DiagError(names.VPCLattice, create.ErrActionReading, DSNameServiceNetworkServiceAssociation, d.Id(), err)
	}

	//lintignore:AWSR002
	if err := d.Set("tags", tags.IgnoreAWS().IgnoreConfig(ignoreTagsConfig).Map()); err != nil {
		return create.DiagError(names.VPCLattice, create.ErrActionSetting, DSNameServiceNetworkServiceAssociation, d.Id(), err)
	}

	return nil
}

func findServiceNetworkServiceAssociationById(ctx context.Context, conn *vpclattice.Client, service_network_service_association_id string) (*vpclattice.GetServiceNetworkServiceAssociationOutput, error) {
	in := &vpclattice.GetServiceNetworkServiceAssociationInput{
		ServiceNetworkServiceAssociationIdentifier: aws.String(service_network_service_association_id),
	}

	out, err := conn.GetServiceNetworkServiceAssociation(ctx, in)

	if err != nil {
		var nfe *types.ResourceNotFoundException
		if errors.As(err, &nfe) {
			return nil, &retry.NotFoundError{
				LastError:   err,
				LastRequest: in,
			}
		}

		return nil, err
	}

	if out == nil || out.Id == nil {
		return nil, tfresource.NewEmptyResultError(in)
	}

	return out, nil
}
