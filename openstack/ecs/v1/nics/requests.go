package nics

import (
	"github.com/gophercloud/gophercloud"
)

type Nic struct {
	// Specifies the subnet ID of the NIC to be added.
	SubnetId string `json:"subnet_id" required:"true"`

	// Specifies the IP address.
	// If this parameter is unavailable, the IP address is automatically assigned.
	IpAddress string `json:"ip_address,omitempty"`

	// Specifies the security groups for NICs.
	SecurityGroups []SecurityGroup `json:"security_groups"`
}
type SecurityGroup struct {
	//Specifies the ID of the security group.
	ID string `json:"id" required:"true"`
}

// AddOpts represents options for add nics.
type AddOpts struct {
	Nics []Nic `json:"nics"`
}

// ToAddNicsOptsMap builds a request body from AddOpts.
func (opts AddOpts) ToAddNicsOptsMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// AddOptsBuilder allows extensions to add additional parameters to the
// AddNics request.
type AddOptsBuilder interface {
	ToAddNicsOptsMap() (map[string]interface{}, error)
}

// Adding nics to an ecs in batches based on the
// configuration defined in the AddOpts struct.
func AddNics(client *gophercloud.ServiceClient, seviceId string, opts AddOptsBuilder) (jobId string, err error) {
	var r AddResult
	reqBody, err := opts.ToAddNicsOptsMap()
	if err != nil {
		return
	}
	_, err = client.Post(addUrl(client, seviceId), reqBody, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200,202,204}})
	if err != nil {
		return
	}

	job, err := r.ExtractJob()
	if err != nil {
		return
	}
	jobId = job.Id
	return
}

// DelOptsBuilder allows extensions to add additional parameters to the
// DeleteNics request.
type DelOptsBuilder interface {
	ToDelNicsOptsMap() (map[string]interface{}, error)
}

// AddOpts represents options for delete nics.
type DelOpts struct {
	Nics []Nics `json:"nics"`
}

type Nics struct {
	//Specifies the port ID of the NIC.
	ID string `json:"id"`
}

// ToDelNicsOptsMap builds a request body from DelOpts.
func (opts DelOpts) ToDelNicsOptsMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Deleting nics from an ecs in batches based on the
// configuration defined in the DelOpts struct.
func DeleteNics(client *gophercloud.ServiceClient, seviceId string, opts DelOptsBuilder) (jobId string, err error) {
	var r DelResult
	reqBody, err := opts.ToDelNicsOptsMap()
	if err != nil {
		return
	}
	_, err = client.Post(deleteUrl(client, seviceId), reqBody, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200,202,204}})
	if err != nil {
		return
	}

	job, err := r.ExtractJob()
	if err != nil {
		return
	}
	jobId = job.Id
	return
}

// BindOpts represents options for binding a virtual ip address to an ecs nic.
type BindOpts struct {
	// Specifies the subnet ID of the NIC.
	SubnetId string `json:"subnet_id" required:"true"`

	// Specifies the virtual IP address to be bound to a NIC.
	IpAddress string `json:"ip_address" required:"true"`

	// Indicates the allowed_address_pairs attribute of a virtual IP address,
	// specifying whether the NIC IP/MAC address pair is added.
	ReverseBinding *bool `json:"reverse_binding" required:"true"`

	// Specifies the DHCP, router, LB, or Nova to which a device belongs.
	DeviceOwner string `json:"device_owner,omitempty"`
}

// BindOptsBuilder allows extensions to add additional parameters to the
// BindOpts request.
type BindOptsBuilder interface {
	ToBindOptsMap() (map[string]interface{}, error)
}

// ToBindOptsMap builds a request body from BindOpts.
func (opts BindOpts) ToBindOptsMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "nic")
}

// Binding a virtual ip address to an ecs nic based on the
// configuration defined in the BindOpts struct.
func BindNic(c *gophercloud.ServiceClient, nicId string, opts BindOptsBuilder) (r BindResult) {
	b, err := opts.ToBindOptsMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(putURL(c, nicId), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200,202,204},
	})
	return
}

// UnBindOpts represents options for unbinding a virtual ip address to an ecs nic.
type UnBindOpts struct {
	// Specifies the subnet ID of the NIC.
	// This parameter must be left blank when you unbind the virtual IP address from an ECS NIC.
	SubnetId string `json:"subnet_id"`

	// Specifies the virtual IP address to be unbound from a NIC.
	// This parameter must be left blank when you unbind the virtual IP address from an ECS NIC.
	IpAddress string `json:"ip_address"`

	// Indicates the allowed_address_pairs attribute of a virtual IP address,
	// specifying whether the NIC IP/MAC address pair is added.
	ReverseBinding *bool `json:"reverse_binding,omitempty"`
}

// UnBindOptsBuilder allows extensions to add additional parameters to the
// UnBindOpts request.
type UnBindOptsBuilder interface {
	ToUnBindOptsMap() (map[string]interface{}, error)
}

// ToUnBindOptsMap builds a request body from UnBindOpts.
func (opts UnBindOpts) ToUnBindOptsMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "nic")
}

// Unbinding a virtual ip address to an ecs nic based on the
// configuration defined in the UnBindOpts struct.
func UnBindNic(c *gophercloud.ServiceClient, nicId string, opts UnBindOptsBuilder) (r BindResult) {
	b, err := opts.ToUnBindOptsMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = c.Put(putURL(c, nicId), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200,202,204},
	})
	return
}
