package xpprovider

import (
	"github.com/oracle/terraform-provider-oci/internal/provider"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func GetProvider() *schema.Provider {
	return provider.Provider()
}
