package vsphere

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/find"
)

func resourceVSphereResourcePool() *schema.Resource {
	s := map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString
			Required:    true,
			Description: "TODO",
		},
		"cpu_allocation": {
			Type:        schema.TypeInt,
			Description: "TODO",
		},
		"memory_allocation": {
			Type:        schema.TypeInt,
			Description: "TODO",
		},
		"entity": {
			Type:        schema.TypeString,
			Description: "TODO",
		},
	}

	return &schema.Resource {
		Create: resourceVSphereResourcePoolCreate,
		Read: resourceVSphereResourcePoolRead,
		Update: resourceVSphereResourcePoolUpdate,
		Delete: resourceVSphereResourcePoolDelete,
		Importer: &schema.ResourceImporter{
			State: resourceVSphereResourcePoolImport,
		},
		SchemaVersion: 3,
		MigrateState: resourceVSphereResourcePoolMigrateState,
		Schema: s,
	}
}

func resourceVSphereResourcePoolCreate(d *schema.ResourceData, meta interface {}) error {
	client := meta.(*VSphereClient).vimClient
	rp := *object.NewResourcePool(c, )

}

func resourceVSphereResourcePoolRead(d *schema.ResourceData, meta interface {}) error {

}

func resourceVSphereResourcePoolUpdate(d *schema.ResourceData, meta interface {}) error {

}

func resourceVSphereResourcePoolDelete(d *schema.ResourceData, meta interface {}) error {

}

func resourceVSphereResourcePoolImport(d *schema.ResourceData, meta interface {}) error {

}

func resourceVSphereResourcePoolMigrateState(d *schema.ResourceData, meta interface {}) error {

}