package testing

const AddNicsRequest = `
{
    "nics": [
        {
            "subnet_id": "d32019d3-bc6e-4319-9c1d-6722fc136a23",
            "security_groups": [
                {
                    "id": "f0ac4394-7e4a-4409-9701-ba8be283dbc3"
                }
            ]
        }
    ]
}
`

const AddNicsResponse = `
{
    "job_id": "70a599e0-31e7-49b7-b260-868f441e862b"
}
`

const DelNicsRequest = `
{
    "nics": [
         {
            "id": "d32019d3-bc6e-4319-9c1d-6722fc136a23"
        }
    ]
}
`

const DelNicsResponse = `
{
    "job_id": "70a599e0-31e7-49b7-b260-868f441e862b"
}
`

const BindNicRequest = `
{
    "nic": {
           "subnet_id": "d32019d3-bc6e-4319-9c1d-6722fc136a23",
           "ip_address": "192.168.0.7",
           "reverse_binding": true
    }
}
`

const BindNicResponse = `
{
    "port_id": "d32019d3-bc6e-4319-9c1d-6722fc136a23"
}
`

const UnBindNicRequest = `
{
    "nic": {
           "subnet_id": "",
           "ip_address": "",
           "reverse_binding": false
    }
}
`

const UnBindNicResponse = `
{
    "port_id": "d32019d3-bc6e-4319-9c1d-6722fc136a23"
}
`