package testing

const BatchQueryRequest = `
{
	"metrics": [{
		"namespace": "MINE.APP",
		"dimensions": [{
			"name": "instance_id",
			"value": "33328f02-3814-422e-b688-bfdba93d4050"
		}],
		"metric_name": "cpu_util"
	},
	{
		"namespace": "MINE.APP",
		"dimensions": [{
			"name": "instance_id",
			"value": "33328f02-3814-422e-b688-bfdba93d4051"
		}],
		"metric_name": "cpu_util"
	}],
	"from": 1484153313000,
	"to": 1484653313000,
	"period": "1",
	"filter": "average"
}
`

const BatchQueryResponse = `
{
	"metrics": [{
			"unit": "request/s",
			"datapoints": [{
					"average": 0,
					"timestamp": 1484401920000
				},
				{
					"average": 1,
					"timestamp": 1484407920000
				}
			],
			"namespace": "MINE.APP",
			"dimensions": [{
				"name": "instance_id",
				"value": "33328f02-3814-422e-b688-bfdba93d4050"
			}],
			"metric_name": "cpu_util"
		},
		{
			"unit": "request/s",
			"datapoints": [{
				"average": 2.3,
				"timestamp": 1484401920000
			}, {
				"average": 1.2,
				"timestamp": 1484407920000
			}],
			"namespace": "MINE.APP",
			"dimensions": [{
				"name": "instance_id",
				"value": "33328f02-3814-422e-b688-bfdba93d4051"
			}],
			"metric_name": "cpu_util"
		}
	]
}
`