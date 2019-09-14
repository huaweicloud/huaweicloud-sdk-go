package testing

import "github.com/gophercloud/gophercloud/openstack/fgs/v2/trigger"

var Triggers = `
[
  {
    "trigger_id": "CNNNTN9RNSOQR17Q0000016267F1DCCF0406ACCA317B4246",
    "trigger_type_code": "OBS",
    "event_type_code": "MessageCreated",
    "trigger_status": "ACTIVE",
    "event_data": {
      "bucket": "obs-384b",
      "events": [
        "s3:ObjectCreated:*"
      ],
      "prefix": "",
      "suffix": ""
    },
    "last_updated_time": "2019-03-27T22:51:25+08:00",
    "created_time": "2019-03-27T22:51:25+08:00"
  }
]
`
var Trigger = `
{
  "trigger_id": "191e425a-924a-45d9-b730-38b5ffe339ad",
  "trigger_type_code": "SMN",
  "event_type_code": "MessageCreated",
  "trigger_status": "ACTIVE",
  "event_data": {
    "topic_urn": "abc",
    "subscription_status": "Unconfirmed"
  },
  "last_updated_time": "2019-06-27 23:33:57 +0000 UTC",
  "created_time": "2019-06-27 23:33:57 +0000 UTC"
}
`

var (
    date    = make(map[string]interface{})
    twoData = make(map[string]interface{})
)

func init() {
    date["bucket"] = "obs-384b"
    date["events"] = []string{"s3:ObjectCreated:*"}
    date["prefix"] = ""
    date["suffix"] = ""

    twoData["topic_urn"] = "abc"
    twoData["subscription_status"] = "Unconfirmed"
}

var TriggerOne = trigger.Trigger{
    TriggerId:       "CNNNTN9RNSOQR17Q0000016267F1DCCF0406ACCA317B4246",
    TriggerTypeCode: "OBS",
    EventTypeCode:   "MessageCreated",
    Status:          "ACTIVE",
    EventData:       date,
    LastUpdatedTime: "2019-03-27T22:51:25+08:00",
    CreatedTime:     "2019-03-27T22:51:25+08:00",
}

var TriggerTwo = trigger.Trigger{
    TriggerId:       "191e425a-924a-45d9-b730-38b5ffe339ad",
    TriggerTypeCode: "SMN",
    EventTypeCode:   "MessageCreated",
    Status:          "ACTIVE",
    EventData:       twoData,
    LastUpdatedTime: "2019-06-27 23:33:57 +0000 UTC",
    CreatedTime:     "2019-06-27 23:33:57 +0000 UTC",
}
