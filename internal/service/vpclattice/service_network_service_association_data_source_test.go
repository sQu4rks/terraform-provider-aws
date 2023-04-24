package vpclattice_test

import (
	"fmt"
	"regexp"
	"testing"

	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"

	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccVPCLatticeServiceNetworkServiceAssociationDataSource_basic(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	//var servicenetworkserviceassociation vpclattice.DescribeServiceNetworkServiceAssociationResponse
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dataSourceName := "data.aws_vpclattice_service_network_service_association.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.VPCLatticeEndpointID)
			testAccPreCheck(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.VPCLatticeEndpointID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckServiceNetworkServiceAssociationDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccServiceNetworkServiceAssociationDataSourceConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					acctest.MatchResourceAttrRegionalARN(dataSourceName, "arn", "vpc-lattice", regexp.MustCompile(`servicenetworkserviceassociation/+.`)),
				),
			},
		},
	})
}

func testAccServiceNetworkServiceAssociationDataSourceConfig_basic(rName string) string {
	return fmt.Sprintf(`
	resource "aws_vpclattice_service" "test" {
		name = %[1]q
	  }
	  
	  resource "aws_vpclattice_service_network" "test" {
		name = %[1]q
	  }
	  
	  resource "aws_vpclattice_service_network_service_association" "test" {
		service_identifier         = aws_vpclattice_service.test.id
		service_network_identifier = aws_vpclattice_service_network.test.id
	  }

	  data "aws_vpclattice_service_network_service_association" "test" {
		service_network_service_association_identifier  = aws_vpclattice_service_network_service_association.test.id
	  }
`, rName)
}
