package trigger

import (
    "github.com/gophercloud/gophercloud"
    "github.com/gophercloud/gophercloud/pagination"
)


type CreateOptsBuilder interface {
    ToCreateTriggerMap() (map[string]interface{}, error)
}

//Trigger struct
type CreateOpts struct {
    TriggerTypeCode string                 `json:"trigger_type_code" required:"true"`
    EventTypeCode   string                 `json:"event_type_code" required:"true"`
    EventData       map[string]interface{} `json:"event_data" required:"true"`
}

func (opts CreateOpts) ToCreateTriggerMap() (map[string]interface{}, error) {
    return gophercloud.BuildRequestBody(opts, "")
}

//Creating a Trigger
func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder, functionUrn string) (r CreateResult) {
    b, err := opts.ToCreateTriggerMap()
    if err != nil {
        r.Err = err
        return
    }
    _, r.Err = c.Post(createURL(c, functionUrn), b, &r.Body, nil)
    return
}

//Querying All Triggers of a Function
func List(c *gophercloud.ServiceClient, functionUrn string) pagination.Pager {
    url := listURL(c, functionUrn)
    return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
        return TriggerPage{pagination.SinglePageBase(r)}
    })
}

//Querying the Information About a Trigger
func Get(c *gophercloud.ServiceClient, functionUrn, triggerTypeCode, triggerId string) (r GetResult) {
    _, r.Err = c.Get(getURL(c, functionUrn, triggerTypeCode, triggerId), &r.Body, &gophercloud.RequestOpts{OkCodes: []int{200}})
    return
}

//Deleting a Trigger
func Delete(c *gophercloud.ServiceClient, functionUrn, triggerTypeCode, triggerId string) (r DeleteResult) {
    _, r.Err = c.Delete(deleteURL(c, functionUrn, triggerTypeCode, triggerId), &gophercloud.RequestOpts{OkCodes: []int{204, 200}})
    return
}

//Deleting All Triggers of a Function
func DeleteAll(c *gophercloud.ServiceClient, functionUrn string) (r DeleteResult) {
    _, r.Err = c.Delete(deleteAllURL(c, functionUrn), nil)
    return
}
