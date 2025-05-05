package nomad_job

import (
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Resource() *schema.Resource {
    return &schema.Resource{
        Create: resourceNomadJobCreate,
        Read:   resourceNomadJobRead,
        Update: resourceNomadJobUpdate,
        Delete: resourceNomadJobDelete,

        Schema: map[string]*schema.Schema{
            "job_hcl": {
                Type:     schema.TypeString,
                Required: true,
                Description: "HCL contents of the Nomad job file",
            },
        },
    }
}

func ResourceNomadJob() *schema.Resource {
  return &schema.Resource{
    // Define Create/Read/Update/Delete here
  }
}

func resourceNomadJobCreate(d *schema.ResourceData, meta interface{}) error {
    hcl := d.Get("job_hcl").(string)
    job, err := api.ParseHCL(hcl, true)
    if err != nil {
        return err
    }

    _, _, err = client.Jobs().Register(job, nil)
    if err != nil {
        return err
    }

    d.SetId(*job.ID)
    return resourceNomadJobRead(d, meta)
}

func resourceNomadJobRead(d *schema.ResourceData, meta interface{}) error {
    job, _, err := client.Jobs().Info(d.Id(), nil)
    if err != nil {
        if api.IsErrNotFound(err) {
            d.SetId("") // mark as deleted
            return nil
        }
        return err
    }

    d.Set("job_name", job.Name)
    return nil
}

func resourceNomadJobUpdate(d *schema.ResourceData, meta interface{}) error {
    return resourceNomadJobCreate(d, meta) // idempotent
}

func resourceNomadJobDelete(d *schema.ResourceData, meta interface{}) error {
    _, _, err := client.Jobs().Deregister(d.Id(), false, nil)
    return err
}
