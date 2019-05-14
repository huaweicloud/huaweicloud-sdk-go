package instances

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type ActionOpts struct {
	// Specifies the ECS ID.
	InstancesId []string `json:"instances_id" required:"true"`

	// Specifies whether to delete ECS instances when they are removed
	// from an AS group. The value can be no (default) or yes.This parameter takes effect
	// only when the action is set to REMOVE.
	InstanceDelete string `json:"instance_delete,omitempty"`

	// Specifies an action to be performed on instances in batches.
	// The options are as follows:ADD: adds instances to the AS group.REMOVE: removes
	// instances from the AS group.PROTECT: enables instance protection.UNPROTECT: disables
	// instance protection.
	Action string `json:"action" required:"true"`
}

type ActionOptsBuilder interface {
	ToInstancesActionMap() (map[string]interface{}, error)
}

func (opts ActionOpts) ToInstancesActionMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Action(client *gophercloud.ServiceClient, scalingGroupId string, opts ActionOptsBuilder) (r ActionResult) {
	b, err := opts.ToInstancesActionMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(ActionURL(client, scalingGroupId), b, nil, &gophercloud.RequestOpts{
		JSONResponse: nil,
		OkCodes:      []int{204},
	})
	return
}

type DeleteOpts struct {
	// Specifies whether the instances are deleted when they are
	// removed from the AS group. The value can be yes or no (default).
	InstanceDelete string `q:"instance_delete" required:"true"`
}

type DeleteOptsBuilder interface {
	ToDeleteQuery() (string, error)
}

func (opts DeleteOpts) ToDeleteQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

func Delete(client *gophercloud.ServiceClient, instanceId string, opts DeleteOptsBuilder) (r DeleteResult) {
	url := DeleteURL(client, instanceId)
	if opts != nil {
		query, err := opts.ToDeleteQuery()
		if err != nil {
			r.Err = err
			return
		}
		url += query
	}

	_, r.Err = client.Delete(url, &gophercloud.RequestOpts{
		JSONResponse: nil,
		OkCodes:      []int{204},
	})
	return
}

type ListOpts struct {
	// Specifies the instance lifecycle status in the AS
	// group.INSERVICE: The instance in the AS group is in use.PENDING: The instance is
	// being added to the AS group.REMOVING: The instance is being removed from the AS
	// group.PENDING_WAIT: The instance is waiting to be added to the AS
	// group.REMOVING_WAIT: The instance is waiting to be removed from the AS group.
	LifeCycleState string `q:"life_cycle_state"`

	// Specifies the instance health status.INITIALIZING: The instance
	// is initializing.NORMAL: The instance is normal.ERROR: The instance is abnormal.
	HealthStatus string `q:"health_status"`

	// Specifies the start line number. The default value is 0.
	StartNumber int `q:"start_number"`

	// Specifies the total number of query records. The default is 20
	// and the maximum is 100.
	Limit                  int    `q:"limit"`
	ProtectFromScalingDown string `q:"protect_from_scaling_down"`
}

type ListOptsBuilder interface {
	ToListQuery() (string, error)
}

func (opts ListOpts) ToListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

func List(client *gophercloud.ServiceClient, scalingGroupId string, opts ListOptsBuilder) pagination.Pager {
	url := ListURL(client, scalingGroupId)
	if opts != nil {
		query, err := opts.ToListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		p := InstancePage{pagination.NumberPageBase{PageResult: r}}
		p.NumberPageBase.Owner = p
		return p
	})
}
