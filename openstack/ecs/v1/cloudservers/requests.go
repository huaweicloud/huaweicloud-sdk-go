package cloudservers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// Get requests details on a single server, by ID.
func Get(client *gophercloud.ServiceClient, serverID string) (r GetResult) {
	if serverID == "" {
		message := fmt.Sprintf(gophercloud.CE_MissingInputMessage, "serverID")
		err := gophercloud.NewSystemCommonError(gophercloud.CE_MissingInputCode, message)
		r.Err = err
		return r
	}
	_, r.Err = client.Get(getURL(client, serverID), &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 203},
	})
	return
}

func GetServerRecoveryStatus(client *gophercloud.ServiceClient, serverID string) (r RecoveryResult) {
	if serverID == "" {
		message := fmt.Sprintf(gophercloud.CE_MissingInputMessage, "serverID")
		err := gophercloud.NewSystemCommonError(gophercloud.CE_MissingInputCode, message)
		r.Err = err
		return r
	}
	_, r.Err = client.Get(autorecoveryURL(client, serverID), &r.Body, nil)
	return
}

func ConfigServerRecovery(client *gophercloud.ServiceClient, serverID string, opts string) (r ErrResult) {
	if serverID == "" {
		message := fmt.Sprintf(gophercloud.CE_MissingInputMessage, "serverID")
		err := gophercloud.NewSystemCommonError(gophercloud.CE_MissingInputCode, message)
		r.Err = err
		return r
	}

	b := map[string]string{
		"support_auto_recovery": opts,
	}

	_, r.Err = client.Put(autorecoveryURL(client, serverID), &b, nil, &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})
	return
}

//func AddServerOnMonitorList(client *gophercloud.ServiceClient, serverID string) (r ErrResult) {
//	if serverID == "" {
//		message := fmt.Sprintf(gophercloud.CE_MissingInputMessage, "serverID")
//		err := gophercloud.NewSystemCommonError(gophercloud.CE_MissingInputCode, message)
//		r.Err = err
//		return r
//	}
//
//	b := map[string]interface{}{
//		"monitorMetrics": nil,
//	}
//
//	_, r.Err = client.Post(actionURL(client, serverID), &b, nil, &gophercloud.RequestOpts{
//		OkCodes: []int{200},})
//	return
//}

//BatchChangeOptsBuilder allows extensions to add additional parameters to the BatchChangeOpts request.
type BatchChangeOptsBuilder interface {
	// Returns value that can be passed to json.Marshal
	ToServerBatchChangeMap() (map[string]interface{}, error)
}

//BatchChangeOpts defining the configuration to batch change OS of servers
type BatchChangeOpts struct {
	AdminPass string    `json:"adminpass,omitempty"`
	KeyName   string    `json:"keyname,omitempty"`
	UserID    string    `json:"userid,omitempty"`
	ImageID   string    `json:"imageid" required:"true"`
	Servers   []Server  `json:"servers" required:"true"`
	MetaData  *MetaData `json:"metadata,omitempty"`
}

//Server defining the server configuration in BatchChangeOpts struct.
type Server struct {
	ID string `json:"id" required:"true"`
}

//MetaData defining the metadata configuration in BatchChangeOpts struct.
type MetaData struct {
	UserData string `json:"user_data,omitempty"`
}

//ToServerBatchChangeMap builds a request body from BatchChangeOpts.
func (opts BatchChangeOpts) ToServerBatchChangeMap() (map[string]interface{}, error) {
	body, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"os-change": body}, nil
}

//BatchChangeOS batch change OS of servers based on the configuration defined in the BatchChangeOpts struct.
func BatchChangeOS(client *gophercloud.ServiceClient, opts BatchChangeOptsBuilder) (r BatchChangeResult) {
	body, err := opts.ToServerBatchChangeMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(batchChangeURL(client), body, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToServerListDetailQuery() (string, error)
}

// ListOpts allows the filtering and sorting of collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the server attributes you want to see returned.
type ListOpts struct {
	// Offset is the current page number.
	Offset int `q:"offset"`

	// Flavor is the ID of the flavor.
	Flavor string `q:"flavor"`

	// Name is the name of the server.
	Name string `q:"name"`

	// Status is the value of the status of the server so that you can filter on
	// "ACTIVE" for example.
	Status string `q:"status"`

	// Limit is an integer value for the limit of values to return.
	Limit int `q:"limit"`

	// Tags is used to filter out the servers with the specified tags
	Tags string `q:"tags"`

	// NotTags queries the cloud server that does not contain this value in the tag field.
	NotTags string `q:"not-tags"`

	// When you create an elastic cloud server in batches, you can specify the returned ID to query the elastic cloud server created in batches.
	ReservationID string `q:"reservation_id"`

	// EnterpriseProjectID specifies the server that is bound to an enterprise project.
	EnterpriseProjectID string `q:"enterprise_project_id"`

	// ipv4 address filtering results
	Ip string `q:"ip"`
}

// ToServerListDetailQuery formats a ListOpts into a query string.
func (opts ListOpts) ToServerListDetailQuery() (string, error) {
	if opts.Offset == 0 {
		opts.Offset = 1
	}
	if opts.Limit == 0 {
		opts.Limit = 25
	}
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// ListDetail makes a request against the API to list servers accessible to you.
func ListDetail(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listDetailURL(client)
	if opts != nil {
		query, err := opts.ToServerListDetailQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return CloudServerPage{pagination.OffsetPage{PageResult: r}}
	})
}

//BatchStartOptsBuilder allows extensions to add additional parameters to the BatchStartOpts request.
type BatchStartOptsBuilder interface {
	// Returns value that can be passed to json.Marshal
	ToServerBatchStartMap() (map[string]interface{}, error)
}

//BatchStartOpts defining the configuration to batch start servers
type BatchStartOpts struct {
	Servers []Server `json:"servers" required:"true"`
}

//ToServerBatchStartMap builds a request body from BatchStartOpts.
func (opts BatchStartOpts) ToServerBatchStartMap() (map[string]interface{}, error) {
	body, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{"os-start": body}, nil
}

//BatchStart makes a request against the API to batch start servers.
func BatchStart(client *gophercloud.ServiceClient, opts BatchStartOpts) (r JobResult) {
	body, err := opts.ToServerBatchStartMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(batchActionURL(client), body, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

//BatchRebootOptsBuilder allows extensions to add additional parameters to the BatchRebootOpts request.
type BatchRebootOptsBuilder interface {
	// Returns value that can be passed to json.Marshal
	ToServerBatchRebootMap() (map[string]interface{}, error)
}

//BatchRebootOpts defining the configuration to batch reboot servers
type BatchRebootOpts struct {
	// Type is the type of reboot to perform on the server.
	Type    Type     `json:"type" required:"true"`
	Servers []Server `json:"servers" required:"true"`
}

// Type describes the mechanisms by which a server reboot or stop can be requested.
type Type string

// These constants determine how a server should be rebooted or stopped.
const (
	Soft Type = "SOFT"
	Hard Type = "HARD"
)

//ToServerBatchRebootMap builds a request body from BatchRebootOpts.
func (opts BatchRebootOpts) ToServerBatchRebootMap() (map[string]interface{}, error) {
	body, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{"reboot": body}, nil
}

//BatchReboot makes a request against the API to batch reboot servers.
func BatchReboot(client *gophercloud.ServiceClient, opts BatchRebootOpts) (r JobResult) {
	body, err := opts.ToServerBatchRebootMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(batchActionURL(client), body, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

//BatchStopOptsBuilder allows extensions to add additional parameters to the BatchStopOpts request.
type BatchStopOptsBuilder interface {
	// Returns value that can be passed to json.Marshal
	ToServerBatchStopMap() (map[string]interface{}, error)
}

//BatchStopOpts defining the configuration to batch stop servers
type BatchStopOpts struct {
	// Type is the type of stop to perform on the server.
	Type    Type     `json:"type,omitempty"`
	Servers []Server `json:"servers" required:"true"`
}

//ToServerBatchStopMap builds a request body from BatchStopOpts.
func (opts BatchStopOpts) ToServerBatchStopMap() (map[string]interface{}, error) {
	body, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{"os-stop": body}, nil
}

//BatchStop makes a request against the API to batch stop servers.
func BatchStop(client *gophercloud.ServiceClient, opts BatchStopOpts) (r JobResult) {
	body, err := opts.ToServerBatchStopMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(batchActionURL(client), body, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

//BatchUpdateOptsBuilder allows extensions to add additional parameters to the BatchUpdateOpts request.
type BatchUpdateOptsBuilder interface {
	// Returns value that can be passed to json.Marshal
	ToServerBatchUpdateMap() (map[string]interface{}, error)
}

//BatchUpdateOpts defining the configuration to batch update servers
type BatchUpdateOpts struct {
	Name    string   `json:"name" required:"true"`
	DryRun  bool     `json:"dry_run,omitempty"`
	Servers []Server `json:"servers" required:"true"`
}

//ToServerBatchUpdateMap builds a request body from BatchUpdateOpts.
func (opts BatchUpdateOpts) ToServerBatchUpdateMap() (map[string]interface{}, error) {
	body, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return body, nil
}

//BatchUpdate makes a request against the API to batch update servers.
func BatchUpdate(client *gophercloud.ServiceClient, opts BatchUpdateOpts) (r BatchUpdateResult) {
	body, err := opts.ToServerBatchUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Put(batchUpdateURL(client), body, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
		HandleError: func(httpStatus int, responseContent string) error {
			var batchUpdateError BatchOperateError
			message := responseContent
			err := json.Unmarshal([]byte(responseContent), &batchUpdateError)
			if err == nil {
				return &batchUpdateError
			}
			return &gophercloud.UnifiedError{
				ErrCode:    gophercloud.MatchErrorCode(httpStatus, message),
				ErrMessage: message,
			}
		},
	})

	return
}

//TagCreate defining the key and value of a tag for creating
type TagCreate struct {
	Key   string `json:"key" required:"true"`
	Value string `json:"value" required:"true"`
}

//TagDelete defining the key and value of a tag for deleting
type TagDelete struct {
	Key   string `json:"key" required:"true"`
	Value string `json:"value,omitempty"`
}

//BatchTagCreateOpts defining the configuration for batch server tags action
type BatchTagCreateOpts struct {
	Tags []TagCreate `json:"tags" required:"true"`
}

//BatchTagDeleteOpts defining the configuration for batch server tags action
type BatchTagDeleteOpts struct {
	Tags []TagDelete `json:"tags" required:"true"`
}

//BatchTagCreateOptsBuilder allows extensions to add additional parameters to the BatchTagActionOpts request.
type BatchTagCreateOptsBuilder interface {
	// Returns value that can be passed to json.Marshal
	ToTagBatchCreateMap() (map[string]interface{}, error)
}

//BatchTagDeleteOptsBuilder allows extensions to add additional parameters to the BatchTagActionOpts request.
type BatchTagDeleteOptsBuilder interface {
	// Returns value that can be passed to json.Marshal
	ToTagBatchDeleteMap() (map[string]interface{}, error)
}

//ToTagBatchCreateMap builds a request body from BatchTagActionOpts.
func (opts BatchTagCreateOpts) ToTagBatchCreateMap() (map[string]interface{}, error) {
	body, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return body, nil
}

//ToTagBatchDeleteMap builds a request body from BatchTagActionOpts.
func (opts BatchTagDeleteOpts) ToTagBatchDeleteMap() (map[string]interface{}, error) {
	body, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return body, nil
}

//BatchCreateServerTags requests to batch create server tags
func BatchCreateServerTags(client *gophercloud.ServiceClient, serverID string, opts BatchTagCreateOptsBuilder) (r ErrResult) {
	if serverID == "" {
		message := fmt.Sprintf(gophercloud.CE_MissingInputMessage, "serverID")
		err := gophercloud.NewSystemCommonError(gophercloud.CE_MissingInputCode, message)
		r.Err = err
		return r
	}

	b, err := opts.ToTagBatchCreateMap()
	b["action"] = "create"
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(batchTagActionURL(client, serverID), &b, nil, &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})
	return
}

//BatchDeleteServerTags requests to batch delete server tags
func BatchDeleteServerTags(client *gophercloud.ServiceClient, serverID string, opts BatchTagDeleteOptsBuilder) (r ErrResult) {
	if serverID == "" {
		message := fmt.Sprintf(gophercloud.CE_MissingInputMessage, "serverID")
		err := gophercloud.NewSystemCommonError(gophercloud.CE_MissingInputCode, message)
		r.Err = err
		return r
	}

	b, err := opts.ToTagBatchDeleteMap()
	b["action"] = "delete"
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(batchTagActionURL(client, serverID), &b, nil, &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})
	return
}

// ListProjectTags makes a request against the API to list project tags accessible to you.
func ListProjectTags(client *gophercloud.ServiceClient) (r ProjectTagsResult) {
	_, r.Err = client.Get(listProjectTagsURL(client), &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// ListServerTags makes a request against the API to list server tags accessible to you.
func ListServerTags(client *gophercloud.ServiceClient, serverID string) (r ServerTagsResult) {
	if serverID == "" {
		message := fmt.Sprintf(gophercloud.CE_MissingInputMessage, "serverID")
		err := gophercloud.NewSystemCommonError(gophercloud.CE_MissingInputCode, message)
		r.Err = err
		return r
	}
	_, r.Err = client.Get(listServerTagsURL(client, serverID), &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type Nic struct {
	//待创建云服务器的网卡信息。
	SubnetId string `json:"subnet_id" required:"true"`

	//待创建云服务器网卡的IP地址，IPv4格式。
	IpAddress string `json:"ip_address,omitempty"`
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
	AvailabilityZone string `json:"availability_zone"`

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
func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (jobId, serverIds string, err error) {
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
	if errJob != nil {
		err = errJob
		return
	}

	server, errServer := r.ExtractServer()
	if errServer != nil {
		err = errServer
		return
	}

	jobId = job.ID
	for i := 0; i < len(server.IDs); i++ {
		serverIds += server.IDs[i] + ","
	}
	return
}

type VolumeExtendParam struct {
	//整机镜像中自带的原始数据盘ID，用于指定整机镜像自带的数据盘信息。
	SnapshotId string `json:"snapshotId,omitempty"`
}

type PublicIp struct {
	//为待创建云服务器分配已有弹性IP时，分配的弹性IP的ID，UUID格式。
	Id string `json:"id,omitempty"`

	//配置云服务器自动分配弹性IP时，创建弹性IP的配置参数。
	Eip *Eip `json:"eip,omitempty"`
}

type SecurityGroup struct {
	//云服务器组ID，UUID格式。
	ID string `json:"id" required:"true"`
}

type SchedulerHints struct {
	//云服务器组ID，UUID格式。
	Group string `json:"group,omitempty"`
}

type ServerExtendParam struct {
	//计费模式。
	ChargingMode int `json:"chargingMode,omitempty"`

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

	//企业项目ID。
	EnterpriseProjectID string `json:"enterprise_project_id,omitempty"`

	//是否配置虚拟机自动恢复的功能。
	SupportAutoRecovery string `json:"support_auto_recovery,omitempty"`

	// 创建竞价实例时，需指定该参数的值为“spot”。
	MarketType string `json:"marketType,omitempty"`

	// 用户愿意为竞价实例每小时支付的最高价格。
	SpotPrice string `json:"spotPrice,omitempty"`

	// 购买的竞价实例时长。
	SpotDurationHours int `json:"spot_duration_hours,omitempty"`

	// 表示购买的“竞价实例时长”的个数。
	SpotDurationCount int `json:"spot_duration_count,omitempty"`

	// 竞价实例中断策略，当前支持immediate。
	InterruptionPolicy string `json:"interruption_policy,omitempty"`
}

type CreateOptsBuilder interface {
	ToServerCreateMap() (map[string]interface{}, error)
}

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

type Eip struct {
	//弹性IP地址类型。
	IpType string `json:"iptype" required:"true"`

	//弹性IP地址带宽参数。
	BandWidth *BandWidth `json:"bandwidth" required:"true"`

	//创建弹性IP的附加信息。
	ExtendParam *EipExtendParam `json:"extendparam,omitempty"`
}

type EipExtendParam struct {
	//公网IP的计费模式。prePaid-预付费，即包年包月；postPaid-后付费，即按需付费；
	ChargingMode string `json:"chargingMode,omitempty"`
}

type BandWidth struct {
	//带宽（Mbit/s），取值范围为[1,300]。
	Size int `json:"size,omitempty"`

	//带宽的共享类型。PER，表示独享，WHOLE，表示共享。
	ShareType string `json:"sharetype" required:"true"`

	//带宽的计费类型。
	ChargeMode string `json:"chargemode,omitempty"`

	//带宽ID，创建WHOLE类型带宽的弹性IP时可以指定之前的共享带宽创建。
	Id string `json:"id,omitempty"`
}
