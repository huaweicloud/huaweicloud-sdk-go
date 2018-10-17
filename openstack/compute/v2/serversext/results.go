package serversext

import (
	"time"
	//"github.com/gophercloud/gophercloud/openstack/ecs/v1/cloudservers"
)

//在server结构上加上了相关bss（Charging）信息及相关volume（VolumeAttached）信息
type ServerExt struct {
	// ID uniquely identifies this server amongst all other servers,
	// including those not accessible to the current tenant.
	ID string `json:"id"`

	// TenantID identifies the tenant owning this server resource.
	TenantID string `json:"tenant_id"`

	// UserID uniquely identifies the user account owning the tenant.
	UserID string `json:"user_id"`

	// Name contains the human-readable name for the server.
	Name string `json:"name"`

	// Status contains the current operational status of the server,
	// such as IN_PROGRESS or ACTIVE.
	Status string `json:"status"`

	// Updated and Created contain ISO-8601 timestamps of when the state of the
	// server last changed, and when it was created.
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`

	// Addresses includes a list of all IP addresses assigned to the server,
	// keyed by pool.
	Addresses map[string]interface{} `json:"addresses"`

	//Addresses  map[string][]cloudservers.Address`json:"addresses"`
	// KeyName indicates which public key was injected into the server on launch.
	KeyName string `json:"key_name"`

	// task state
	TaskState string `json:"OS-EXT-STS:task_state"`

	// vm state
	VMstate string `json:"OS-EXT-STS:vm_state"`

	// SecurityGroups includes the security groups that this instance has applied
	// to it.
	SecurityGroups []map[string]interface{} `json:"security_groups"`
	//SecurityGroups []cloudservers.SecurityGroups `json:"security_groups"`
	// power state
	PowerState int `json:"OS-EXT-STS:power_state"`

	// Metadata includes a list of all user-specified key-value pairs attached
	// to the server.
	Metadata map[string]string `json:"metadata"`
	//Metadata            cloudservers.Metadata             `json:"metadata"`
	// Flavor refers to a JSON object, which itself indicates the hardware
	// configuration of the deployed server.
	Flavor map[string]interface{} `json:"flavor"`
	//Flavor cloudservers.Flavor`json:"flavor"`
	// Image refers to a JSON object, which itself indicates the OS image used to
	// deploy the server.
	Image map[string]interface{} `json:"-"`

	// Links includes HTTP references to the itself, useful for passing along to
	// other APIs that might want a server reference.
	Links []interface{} `json:"links"`

	//availbility zone
	AvailbiltyZone string `json:"OS-EXT-AZ:availability_zone"`

	//volume attached new
	VolumeAttached []VolumeInfo

	//云服务器计费信息 new
	Charging Charging

	//虚拟私有云uuid new
	VpcId   string `json:"vpc_id"`
	ImageId string `json:"image_id"`
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
