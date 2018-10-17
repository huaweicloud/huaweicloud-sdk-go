package cloudimages

import (
	"fmt"
	"strings"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type ListOptsBuilder interface {
	ToImageListQuery() (string, error)
}

type ListOpts struct {
	Isregistered           string `q:"__isregistered"`
	Imagetype              string `q:"__imagetype"`
	Protected              *bool  `q:"protected"`
	Visibility             string `q:"visibility"`
	Owner                  string `q:"owner"`
	ID                     string `q:"id"`
	Status                 string `q:"status"`
	Name                   string `q:"name"`
	ContainerFormat        string `q:"container_format"`
	DiskFormat             string `q:"disk_format"`
	MinRam                 *int   `q:"min_ram"`
	MinDisk                int    `q:"min_disk"`
	OsBit                  string `q:"__os_bit"`
	Platform               string `q:"__platform"`
	Marker                 string `q:"marker"`
	Limit                  int    `q:"limit"`
	SortKey                string `q:"sort_key"`
	SortDir                string `q:"sort_dir"`
	OsType                 string `q:"__os_type"`
	Tag                    string `q:"tag"`
	MemberStatus           string `q:"member_status"`
	SupportKvm             string `q:"__support_kvm"`
	SupportXen             string `q:"__support_xen"`
	SupportDiskintensive   string `q:"__support_diskintensive"`
	SupportHighperformance string `q:"__support_highperformance"`
	SupportXenGpuType      string `q:"__support_xen_gpu_type"`
	VirtualEnvType         string `q:"virtual_env_type"`
}

func (opts ListOpts) ToImageListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToImageListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return ImagePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// CreateOptsBuilder allows extensions to add parameters to the Create request.
type CreateOptsBuilder interface {
	// Returns value that can be passed to json.Marshal
	ToImageCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents options used to create an image.
type CreateByServerOpts struct {
	// Name is the name of the new image.
	Name string `json:"name" required:"true"`

	// Description of image.
	Description string `json:"description"`

	// server id to be converse
	InstanceId string `json:"instance_id"`

	// image label "key.value"
	Tags []string `json:"tags"`
}

// CreateOpts represents options used to create an image.
type CreateByFileOpts struct {
	// Name is the name of the new image.
	Name string `json:"name" required:"true"`

	// Description of image.
	Description string `json:"description"`

	//OBS桶中外部镜像文件地址。
	ImageUrl string `json:"image_url"`

	//操作系统版本。
	OsVersion string `json:"os_version"`

	//是否自动配置，取值为true或false。
	IsConfig bool `json:"is_config"`

	//是否完成了初始化配置。取值为true或false。
	IsConfigInit bool `json:"is_config_init"`

	//最小系统盘大小。
	MinDisk int `json:"min_disk"`

	//创建加密镜像的用户主密钥
	CmkId string `json:"cmk_id"`

	//image label "key.value"
	Tags []string `json:"tags"`
}

// ToImageCreateMap assembles a request body based on the contents of
// a CreateOpts.
func (opts CreateByServerOpts) ToImageCreateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (opts CreateByFileOpts) ToImageCreateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Create implements create image request.
func CreateImageByServer(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToImageCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(createURL(client), b, &r.Body, &gophercloud.RequestOpts{OkCodes: []int{200}})
	return
}

// Create implements create image request.
func CreateImageByFile(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToImageCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(createURL(client), b, &r.Body, &gophercloud.RequestOpts{OkCodes: []int{200}})
	return
}

// Get job result.
func GetJobResult(client *gophercloud.ServiceClient, id string) (r JobExecResult) {
	//把v2替换成v1,并且拼接上projectId
	newStr := fmt.Sprintf("/v1/%s/", client.ProviderClient.GetProjectID())
	newUrl := strings.Replace(jobURL(client, id), "/v2/", newStr, 1)

	_, r.Err = client.Get(newUrl, &r.Body, nil)
	return
}
