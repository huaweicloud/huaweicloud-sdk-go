package cloudserversext

import (
	"github.com/gophercloud/gophercloud/openstack/ecs/v1/cloudservers"
)

//在server结构上加上了相关bss（Charging）信息及相关volume（VolumeAttached）信息
type CloudServerExt struct {

	CloudServer *cloudservers.CloudServer

	//volume attached new
	VolumeAttached []VolumeInfo

	//云服务器计费信息 new
	Charging Charging
}


type Charging struct {
	//0代表按需计费、1代表包周期计费
	ChargingMode string

	//ChargingMode为1时有效 new
	ValidTime string

	//ChargingMode为1时有效 new
	ExpireTime string
}

type VolumeInfo struct {
	//卷的uuid
	ID string

	//云硬盘类型
	VolumeType string

	//云硬盘大小
	Size int
}
