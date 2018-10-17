package cloudservers

import (
	"encoding/base64"
	"strings"

	"github.com/gophercloud/gophercloud"
)

type CreateOpts struct {
	//待创建云服务器的系统镜像，需要指定已创建镜像的ID。
	ImageRef string `json:"imageRef" required:"true"`

	//待创建云服务器的系统规格的ID。
	FlavorRef string `json:"flavorRef" required:"true"`

	//云服务器名称。
	Name string `json:"name" required:"true"`

	//创建云服务器过程中注入用户数据。支持注入文本、文本文件或gzip文件。更多关于待注入用户数据的信息，请参见用户数据注入。
	UserData []byte `json:"-"`

	// AdminPass sets the root user password. If not set, a randomly-generated
	// password will be created and returned in the response.
	AdminPass string `json:"adminPass,omitempty"`

	//如果需要使用SSH密钥方式登录云服务器，请指定已创建密钥的名称。
	KeyName string `json:"key_name,omitempty"`

	//待创建云服务器所属虚拟私有云（简称VPC），需要指定已创建VPC的ID。
	VpcId string `json:"vpcid" required:"true"`

	//待创建云服务器的网卡信息。
	Nics []Nic `json:"nics" required:"true"`

	//配置云服务器的弹性IP信息，弹性IP有三种配置方式。
	PublicIp *PublicIp `json:"publicip,omitempty"`

	//创建云服务器数量。
	Count int `json:"count,omitempty"`

	//云服务器名称是否允许重名。
	IsAutoRename *bool `json:"isAutoRename,omitempty"`

	//云服务器对应系统盘相关配置。
	RootVolume RootVolume `json:"root_volume" required:"true"`

	//云服务器对应数据盘相关配置。每一个数据结构代表一块待创建的数据盘。
	DataVolumes []DataVolume `json:"data_volumes,omitempty"`

	//云服务器对应安全组信息。
	SecurityGroups []SecurityGroup `json:"security_groups,omitempty"`

	//待创建云服务器所在的可用分区，需要指定可用分区（AZ）的名称。
	AvailabilityZone string `json:"availability_zone" required:"true"`

	//创建云服务器附加信息。
	ExtendParam *ServerExtendParam `json:"extendparam,omitempty"`

	//创建云服务器元数据。
	MetaData *MetaData `json:"metadata,omitempty"`

	//云服务器调度信息。
	SchedulerHints *SchedulerHints `json:"os:scheduler_hints,omitempty"`

	//弹性云服务器的标签。
	Tags []string `json:"tags,omitempty"`

	//弹性云服务器的标签。
	ServerTags []ServerTags `json:"server_tags,omitempty"`
}

// Create requests a server to be provisioned to the user in the current tenant.
func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (jobId, orderId string, err error) {
	var r CreateResult
	reqBody, err := opts.ToServerCreateMap()
	if err != nil {
		return
	}

	_, err = client.Post(createURL(client), reqBody, &r.Body, &gophercloud.RequestOpts{OkCodes: []int{200}})
	if err != nil {
		return
	}

	job, errJob := r.ExtractJob()
	order, errOrder := r.ExtractOrder()
	if errJob != nil && errOrder != nil {
		return
	}

	jobId = job.Id
	orderId = order.Id
	return
}

// Get job result.
func GetJobResult(client *gophercloud.ServiceClient, id string) (JobResult, error) {
	var r JobExecResult
	url := jobURL(client, id)

	//把v1.1替换成v1
	url2 := strings.Replace(url, "/v1.1/", "/v1/", 1)

	_, err := client.Get(url2, &r.Body, nil)
	if err != nil {
		return JobResult{}, err
	}

	return r.ExtractJobResult()
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToServerCreateMap() (map[string]interface{}, error)
}

// ToServerCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToServerCreateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	if opts.UserData != nil {
		var userData string
		if _, err := base64.StdEncoding.DecodeString(string(opts.UserData)); err != nil {
			userData = base64.StdEncoding.EncodeToString(opts.UserData)
		} else {
			userData = string(opts.UserData)
		}
		b["user_data"] = &userData
	}

	return map[string]interface{}{"server": b}, nil
}

type Nic struct {
	//待创建云服务器的网卡信息。
	SubnetId string `json:"subnet_id" required:"true"`

	//待创建云服务器网卡的IP地址，IPv4格式。
	IpAddress string `json:"ip_address,omitempty"`
}

type PublicIp struct {
	//为待创建云服务器分配已有弹性IP时，分配的弹性IP的ID，UUID格式。
	Id string `json:"id,omitempty"`

	//配置云服务器自动分配弹性IP时，创建弹性IP的配置参数。
	Eip *Eip `json:"eip,omitempty"`
}

type Eip struct {
	//弹性IP地址类型。
	IpType string `json:"iptype" required:"true"`

	//弹性IP地址带宽参数。
	BandWidth *BandWidth `json:"bandwidth" required:"true"`

	//创建弹性IP的附加信息。
	ExtendParam *EipExtendParam `json:"extendparam,omitempty"`
}

type BandWidth struct {
	//带宽（Mbit/s），取值范围为[1,300]。
	Size int `json:"size,omitempty"`

	//带宽的共享类型。PER，表示独享，WHOLE，表示独享。
	ShareType string `json:"sharetype" required:"true"`

	//带宽的计费类型。
	ChargeMode string `json:"chargemode,omitempty"`

	//带宽ID，创建WHOLE类型带宽的弹性IP时可以指定之前的共享带宽创建。
	Id string `json:"id,omitempty"`
}

type EipExtendParam struct {
	//公网IP的计费模式。prePaid-预付费，即包年包月；postPaid-后付费，即按需付费；
	ChargingMode string `json:"chargingMode,omitempty"`
}

type RootVolume struct {
	//云服务器系统盘对应的磁盘类型，需要与系统所提供的磁盘类型相匹配。
	VolumeType string `json:"volumetype" required:"true"`

	//系统盘大小，容量单位为GB， 输入大小范围为[1,1024]。
	Size int `json:"size,omitempty"`

	//磁盘的产品信息。
	ExtendParam *VolumeExtendParam `json:"extendparam,omitempty"`
}

type DataVolume struct {
	//云服务器数据盘对应的磁盘类型，需要与系统所提供的磁盘类型相匹配。
	VolumeType string `json:"volumetype" required:"true"`

	//数据盘大小，容量单位为GB，输入大小范围为[10,32768]。
	Size int `json:"size" required:"true"`

	//创建共享磁盘的信息。
	MultiAttach *bool `json:"multiattach,omitempty"`

	//数据卷是否使用SCSI锁。
	PassThrough *bool `json:"hw:passthrough,omitempty"`

	//磁盘的产品信息。
	Extendparam *VolumeExtendParam `json:"extendparam,omitempty"`
}

type VolumeExtendParam struct {
	//整机镜像中自带的原始数据盘ID，用于指定整机镜像自带的数据盘信息。
	SnapshotId string `json:"snapshotId,omitempty"`
}

type ServerExtendParam struct {
	//计费模式。
	ChargingMode string `json:"chargingMode,omitempty"`

	//云服务器所在区域ID。
	RegionID string `json:"regionID,omitempty"`

	//订购周期类型。
	PeriodType string `json:"periodType,omitempty"`

	//订购周期数。
	PeriodNum int `json:"periodNum,omitempty"`

	//是否自动续订。
	IsAutoRenew string `json:"isAutoRenew,omitempty"`

	//下单订购后，是否自动从客户的账户中支付，而不需要客户手动去进行支付。
	IsAutoPay string `json:"isAutoPay,omitempty"`

	//是否配置虚拟机自动恢复的功能。
	SupportAutoRecovery string `json:"support_auto_recovery,omitempty"`
}

type MetaData struct {
	//用户ID。
	OpSvcUserId string `json:"op_svc_userid,omitempty"`
}

type SecurityGroup struct {
	//云服务器组ID，UUID格式。
	ID string `json:"id" required:"true"`
}

type SchedulerHints struct {
	//云服务器组ID，UUID格式。
	Group string `json:"group,omitempty"`
}

type ServerTags struct {
	Key   string `json:"key" required:"true"`
	Value string `json:"value,omitempty"`
}
