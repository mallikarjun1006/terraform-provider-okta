package okta

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccOktaDataSourceEveryoneGroup_read(t *testing.T) {
	ri := acctest.RandInt()
	config := testAccDataSourceEveryoneGroupConfig(ri)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.okta_everyone_group.everyone-"+strconv.Itoa(ri), "id"),
				),
			},
		},
	})
}

func testAccDataSourceEveryoneGroupConfig(rInt int) string {
	return fmt.Sprintf(`
data "okta_everyone_group" "everyone-%d" {}
`, rInt)
}
