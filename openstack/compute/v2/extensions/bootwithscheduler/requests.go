package bootwithscheduler

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
	volume "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/bootfromvolume"
	hints "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/schedulerhints"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

// CreateOptsExt is a structure that extends the server `CreateOpts` structure
type CreateOptsExt struct {
	servers.CreateOptsBuilder
	BlockDevice    []volume.BlockDevice `json:"block_device_mapping_v2,omitempty"`
	SchedulerHints hints.SchedulerHints `json:"os:scheduler_hints,omitempty"`
}

func (opts CreateOptsExt) ToServerCreateMap() (map[string]interface{}, error) {
	base, err := opts.CreateOptsBuilder.ToServerCreateMap()
	if err != nil {
		return nil, err
	}

	if len(opts.BlockDevice) == 0 {
		message := fmt.Sprintf(gophercloud.CE_MissingInputMessage, "bootfromvolume.CreateOptsExt.BlockDevice")
		err := gophercloud.NewSystemCommonError(gophercloud.CE_MissingInputCode, message)
		return nil, err
	}

	serverMap := base["server"].(map[string]interface{})

	blockDevice := make([]map[string]interface{}, len(opts.BlockDevice))

	for i, bd := range opts.BlockDevice {
		b, err := gophercloud.BuildRequestBody(bd, "")
		if err != nil {
			return nil, err
		}
		blockDevice[i] = b
	}
	serverMap["block_device_mapping_v2"] = blockDevice
/*
	schedulerHints 只需要“check_resources”一个参数
	if &(opts.SchedulerHints) != nil {
		schedulerHints := make(map[string]interface{})
		if opts.SchedulerHints.CheckResources != "" {
			schedulerHints["check_resources"] = opts.SchedulerHints.CheckResources
		}
		base["os:scheduler_hints"] = schedulerHints
	}
*/

//注释：支持opts.SchedulerHints 可传入多个参数
	schedulerData,err:=opts.SchedulerHints.ToServerSchedulerHintsCreateMap()
	if err!=nil{
		return nil, err
	}

	base["os:scheduler_hints"] = schedulerData
	return base, nil
}
