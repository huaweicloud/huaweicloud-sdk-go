package testing

const FlavorListResponse = `
{
    "flavors": [
        {
            "id": "104",
            "name": "m1.large",
            "vcpus": "4",
            "ram": 8192,
            "disk": "0",
            "swap": "",
            "OS-FLV-EXT-DATA:ephemeral": 0,
            "rxtx_factor": null,
            "OS-FLV-DISABLED:disabled": null,
            "rxtx_quota": null,
            "rxtx_cap": null,
            "os-flavor-access:is_public": null,
            "os_extra_specs": {
                "ecs:performancetype": "normal"
            }
        },
        {
            "id": "105",
            "name": "m2.large",
            "vcpus": "8",
            "ram": 16384,
            "disk": "0",
            "swap": "",
            "OS-FLV-EXT-DATA:ephemeral": 0,
            "rxtx_factor": null,
            "OS-FLV-DISABLED:disabled": null,
            "rxtx_quota": null,
            "rxtx_cap": null,
            "os-flavor-access:is_public": null,
            "os_extra_specs": {
                "ecs:performancetype": "normal"
            }
        }
    ]
}
`

const FlavorListWithFormResponse = `
{
    "flavors": [
        {
            "id": "104",
            "name": "m1.large",
            "vcpus": "4",
            "ram": 8192,
            "disk": "0",
            "swap": "",
            "OS-FLV-EXT-DATA:ephemeral": 0,
            "rxtx_factor": null,
            "OS-FLV-DISABLED:disabled": null,
            "rxtx_quota": null,
            "rxtx_cap": null,
            "os-flavor-access:is_public": null,
            "os_extra_specs": {
                "ecs:performancetype": "normal"
            }
        }
    ]
}
`

const ResizeRequest = `
{
    "resize": {
        "flavorRef": "c3.15xlarge.2",
        "dedicated_host_id": "459a2b9d-804a-4745-ab19-a113bb1b4ddc"
    }
}
`

const ResizeResponse = `
{
    "job_id": "70a599e0-31e7-49b7-b260-868f441e862b"
}
`