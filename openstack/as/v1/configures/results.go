package configures

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
	"strconv"
)

type commonResult struct {
	gophercloud.Result
}

type PublicIp struct {
	// Specifies the configuration parameter for creating an EIP that
	// will be automatically assigned to the ECS.
	Eip Eip `json:"eip"`
}

type Eip struct {
	// Specifies the IP address type.Enumerated values of the IP
	// address type:5_bgp: indicates the dynamic BGP.5_lxbgp: BGP,5_telcom: indicates China
	// Telecom.5_union: indicates China Unicom.
	IpType string `json:"ip_type"`

	// Specifies the bandwidth.
	Bandwidth Bandwidth `json:"bandwidth"`
}

type Bandwidth struct {
	ID string `json:"id"`
	// Specifies the bandwidth (Mbit/s). The value range is 1 to
	// 100.
	Size int `json:"size"`

	// Specifies the bandwidth sharing type.Enumerated value: PER
	// (indicates exclusive bandwidth).Only exclusive bandwidth is available.
	ShareType string `json:"share_type"`

	// Specifies the bandwidth charging mode.If the field value is
	// bandwidth, the ECS service is charged by bandwidth.If the field value is traffic, the
	// ECS service is charged by traffic.If the field value is others, the ECS creation will
	// fail.
	ChargingMode string `json:"charging_mode"`
}

type InstanceConfig struct {
	// This field is reserved.
	InstanceName string `json:"instance_name"`

	// Specifies the ECS ID. When using the existing ECS
	// specifications as the template to create AS configurations, specify this parameter.
	// In this case, flavorRef, imageRef, and disk fields do not take effect.If the
	// instance_id field is not specified, flavorRef, imageRef, and disk fields are
	// mandatory.
	InstanceId string `json:"instance_id"`

	// Specifies the ECS specifications ID, which defines the
	// specifications of the CPU and memory for the ECS. You can obtain its value from the
	// API used to query specifications and expansion details about ECSs. For details, see
	// section Querying Specifications and Expansion Details About ECSs in the Elastic Cloud
	// Server API Reference.
	FlavorRef string `json:"flavorRef"`

	// Specifies the image ID. It is the same as image_id. You can
	// obtain its value from the API used to query IMS images. For details, see section
	// Querying Images in the Image Management Service API Reference.
	ImageRef string `json:"imageRef"`

	// Specifies the disk group information. System disks are
	// mandatory and data disks are optional.
	Disk []Disk `json:"disk"`

	// This field is reserved.
	AdminPass string `json:"adminPass"`

	// Specifies the name of the SSH key pair used to log in to the
	// ECS.
	KeyName string `json:"key_name"`

	// Specifies information about the injected file. Only text files
	// can be injected. A maximum of five files can be injected at a time and the maximum
	// size of each file is 1 KB.
	Personality []Personality `json:"personality"`

	// Specifies the EIP of the ECS. The EIP can be configured in the
	// following two ways:Not configured (delete this field),Assigned automatically
	PublicIp PublicIp `json:"public_ip"`

	// Specifies the cloud-init user data.Text, text files, and gzip
	// files can be injected. The file content must be encoded with Base64, and the maximum
	// allowed file size is 32 KB.
	UserData string `json:"user_data"`

	// Specifies the metadata of ECSs to be created.
	Metadata map[string]interface{} `json:"metadata"`

	SecurityGroups []SecurityGroup `json:"security_groups"`

	ServerGroupID string `json:"server_group_id"`

	Tenancy string `json:"tenancy"`

	DedicatedHostID string `json:"dedicated_host_id"`

	MarketType string `json:"market_type"`

	KeyFingerPrint string `json:"key_fingerprint"`
}

type ScalingConfiguration struct {
	// Specifies the AS configuration ID. This parameter is globally
	// unique.
	ScalingConfigurationId string `json:"scaling_configuration_id"`

	// Specifies the tenant ID.
	Tenant string `json:"tenant"`

	// Specifies the AS configuration name.
	ScalingConfigurationName string `json:"scaling_configuration_name"`

	// Specifies the information about instance configurations.
	InstanceConfig InstanceConfig `json:"instance_config"`

	// Specifies the time when AS configurations are created. The time
	// format complies with UTC.
	CreateTime string `json:"create_time"`
}

type CreateResult struct {
	commonResult
}

func (r CreateResult) Extract() (*CreateResponse, error) {
	var response CreateResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type CreateResponse struct {
	// Specifies the AS configuration ID.
	ScalingConfigurationId string `json:"scaling_configuration_id"`
}

type DeleteResult struct {
	gophercloud.ErrResult
}

type DeleteWithBatchResult struct {
	gophercloud.ErrResult
}

type GetResult struct {
	commonResult
}

func (r GetResult) Extract() (*ScalingConfiguration, error) {
	var s struct {
		ScalingConfiguration ScalingConfiguration `json:"scaling_configuration"`
	}
	err := r.ExtractInto(&s)
	return &s.ScalingConfiguration, err
}

type ListResponse struct {
	// Specifies the total number of query records.
	TotalNumber int `json:"total_number"`

	// Specifies the start line number.
	StartNumber int `json:"start_number"`

	// Specifies the number of query records.
	Limit int `json:"limit"`

	// Specifies the AS configuration list.
	ScalingConfigurations []ScalingConfiguration `json:"scaling_configurations"`
}

type ConfigPage struct {
	pagination.NumberPageBase
}

// IsEmpty returns true if a ListResult contains no services.
func (r ConfigPage) IsEmpty() (bool, error) {
	data, err := ExtractConfigs(r)
	return data.StartNumber > data.TotalNumber, err
}

// LastMarker returns the last service in a ListResult.
func (r ConfigPage) LastStartNumber() (string, error) {
	data, err := ExtractConfigs(r)
	if err != nil {
		return "", err
	}
	nextStartNumber := data.Limit + data.StartNumber
	if nextStartNumber >= data.TotalNumber {
		return "", nil
	}
	return strconv.Itoa(nextStartNumber), nil
}

// ExtractServices is a function that takes a ListResult and returns the services' information.
func ExtractConfigs(r pagination.Page) (ListResponse, error) {
	var s ListResponse
	err := (r.(ConfigPage)).ExtractInto(&s)
	return s, err
}
