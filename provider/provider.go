package provider

import (
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
  resjob "github.com/alexfreidah/terraform-nomad-job/resources/nomad_job"
)

func Provider() *schema.Provider {
  return &schema.Provider{
    ResourcesMap: map[string]*schema.Resource{
      "nomad_job": resjob.ResourceNomadJob(), // <- fixed!
    },
  }
}
