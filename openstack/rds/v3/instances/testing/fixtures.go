package testing

import (
	"fmt"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"net/http"
	"testing"
)

const ListOutput = `
{
	"instances": [{
		"id": "ed7cc6166ec24360a5ed5c5c9c2ed726in01",
		"status": "ACTIVE",
		"name": "mysql-0820-022709-01",
		"port": 3306,
		"type": "Single",
		"region": "aaa",
		"datastore": {
			"type": "MySQL",
			"version": "5.7"
		},
		"created": "2018-08-20T02:33:49+0800",
		"updated": "2018-08-20T02:33:50+0800",
		"volume": {
			"type": "ULTRAHIGH",
			"size": 100
		},
		"nodes": [{
			"id": "06f1c2ad57604ae89e153e4d27f4e4b8no01",
			"name": "mysql-0820-022709-01_node0",
			"role": "master",
			"status": "ACTIVE",
			"availability_zone": "bbb"
		}],
		"private_ips": ["192.168.0.142"],
		"public_ips": ["10.154.219.187", "10.154.219.186"],
		"db_user_name": "root",
		"vpc_id": "b21630c1-e7d3-450d-907d-39ef5f445ae7",
		"subnet_id": "45557a98-9e17-4600-8aec-999150bc4eef",
		"security_group_id": "38815c5c-482b-450a-80b6-0a301f2afd97",
		"flavor_ref": "rds.mysql.s1.large",
		"switch_strategy": "",
                "charge_info": {
                    "charge_mode": "postPaid"
                },
		"backup_strategy": {
			"start_time": "19:00-20:00",
			"keep_days": 7
		},
		"maintenance_window": "02:00-06:00",
		"related_instance": [],
		"disk_encryption_id": "",
		"enterprise_project_id": "0",
		"time_zone": ""
	}, {
		"id": "ed7cc6166ec24360a5ed5c5c9c2ed726in02",
		"status": "ACTIVE",
		"name": "mysql-0820-022709-02",
		"port": 3306,
		"type": "Single",
		"region": "aaa",
		"datastore": {
			"type": "MySQL",
			"version": "5.7"
		},
		"created": "2018-08-20T02:33:49+0800",
		"updated": "2018-08-20T02:33:50+0800",
		"volume": {
			"type": "ULTRAHIGH",
			"size": 100
		},
		"nodes": [{
			"id": "06f1c2ad57604ae89e153e4d27f4e4b8no01",
			"name": "mysql-0820-022709-01_node0",
			"role": "master",
			"status": "ACTIVE",
			"availability_zone": "bbb"
		}],
		"private_ips": ["192.168.0.142"],
		"public_ips": ["10.154.219.187", "10.154.219.186"],
		"db_user_name": "root",
		"vpc_id": "b21630c1-e7d3-450d-907d-39ef5f445ae7",
		"subnet_id": "45557a98-9e17-4600-8aec-999150bc4eef",
		"security_group_id": "38815c5c-482b-450a-80b6-0a301f2afd97",
		"flavor_ref": "rds.mysql.s1.large",
		"switch_strategy": "",
                "charge_info": {
                    "charge_mode": "postPaid"
                },
		"backup_strategy": {
			"start_time": "19:00-20:00",
			"keep_days": 7
		},
		"maintenance_window": "02:00-06:00",
		"related_instance": [],
		"disk_encryption_id": "",
		"enterprise_project_id": "0",
		"time_zone": ""
	}],
	"total_count": 2
}
`
const CreatSingleRdsRequet  =`
{
	"name": "trove-instance-rep2",
	"datastore": {
		"type": "MySQL",
		"version": "5.6"
	},
	"flavor_ref": "rds.mysql.s1.large",
	"volume": {
		"type": "ULTRAHIGH",
		"size": 100
	},	
	"region": "cn-north-4",
	"availability_zone": "cn-north-4",
	"vpc_id": "490a4a08-ef4b-44c5-94be-3051ef9e4fce",
	"subnet_id": "0e2eda62-1d42-4d64-a9d1-4e9aa9cd994f",
	"security_group_id": "2a1f7fc8-3307-42a7-aa6f-42c8b9b8f8c5",
	"port": "8635",
	"backup_strategy": {
		"start_time": "08:15-09:15",
		"keep_days": 12
	},
	"password": "YpurPassword"
}
`
const CreateSingleRdsResp  = `{
	"instance": {
		"id": "dsfae23fsfdsae3435in01",
		"name": "trove-instance-rep2",
		"datastore": {
			"type": "MySQL",
			"version": "5.6"
		},
		"flavor_ref": "rds.mysql.s1.large",
		"volume": {
			"type": "ULTRAHIGH",
			"size": 100
		},
		"disk_encryption_id": "2gfdsh-844a-4023-a776-fc5c5fb71fb4",
        "region": "cn-north-4",
		"availability_zone": "cn-north-4",
		"vpc_id": "490a4a08-ef4b-44c5-94be-3051ef9e4fce",
		"subnet_id": "0e2eda62-1d42-4d64-a9d1-4e9aa9cd994f",
		"security_group_id": "2a1f7fc8-3307-42a7-aa6f-42c8b9b8f8c5",
		"port": "8635",
		"backup_strategy": {
			"start_time": "08:15-09:15",
			"keep_days": 12
		},
		"configuration_id": "452408-44c5-94be-305145fg",
		"charge_info": {
			"charge_mode": "postPaid"
		}
	},
	"job_id": "dff1d289-4d03-4942-8b9f-463ea07c000d"
}
`
const RdsJobidResp  = `
{
	"job_id": "dff1d289-4d03-4942-8b9f-463ea07c000d"
}
`
const ErrorLogResp  = `
{
	"error_log_list": [{
		"time": "2018-12-04T14:24:42",
		"level": "ERROR",
		"content": "Slave I/O for channel '': error connecting to master 'rdsRepl@172.16.30.111:3306' - retry-time: 60  retries: 1, Error_code: 203"
	}, {
		"time": "2018-12-04T14:24:42",
		"level": "ERROR",
		"content": "Slave I/O for channel '': error connecting to master 'rdsRepl@172.11.11.111:8081' - retry-time: 60  retries: 1, Error_code: 203"
	}],
	"total_record": 2
}
`
const  SlowLogResp  = `
{ 
    "slow_log_list":[
        {
            "count":"1",
            "time":"1.04899 s",
            "lock_time": "0.00003 s",
            "rows_sent": "0",
            "rows_examined": "0",
            "database": "mysql",
            "users": "root",
            "query_sample": "INSERT INTO time_zone_name (Name, Time_zone_id) VALUES (N, @time_zone_id);",
            "type": "INSERT"
        },
		{
            "count":"2",
            "time":"2.04899 s",
            "lock_time": "0.00003 s",
            "rows_sent": "2",
            "rows_examined": "0",
            "database": "mysql",
            "users": "root",
            "query_sample": "DELETE time_zone_name (Name, Time_zone_id) VALUES (N, @time_zone_id);",
            "type": "DELETE"
        }
    ],
   "total_record":2
}
`
func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/instances", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}
func HandleCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/instances", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, CreatSingleRdsRequet)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CreateSingleRdsResp)
	})
}
func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/instances/3e93d3eb20b34bfbbdcc81a79c1c3045", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		//th.TestJSONRequest(t, r, CreatSingleRdsRequet)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, RdsJobidResp)
	})
}
func HandleResizeFlavorSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/instances/3e93d3eb20b34bfbbdcc81a79c1c3045/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		//th.TestJSONRequest(t, r, CreatSingleRdsRequet)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, RdsJobidResp)
	})
}

func HandleEnlargeVolumeSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/instances/3e93d3eb20b34bfbbdcc81a79c1c3045/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		//th.TestJSONRequest(t, r, CreatSingleRdsRequet)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, RdsJobidResp)
	})
}

func HandleRestarRdsInstanceSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/instances/3e93d3eb20b34bfbbdcc81a79c1c3045/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		//th.TestJSONRequest(t, r, CreatSingleRdsRequet)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, RdsJobidResp)
	})
}

func HandleSingleToHaRdsInstanceSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/instances/3e93d3eb20b34bfbbdcc81a79c1c3045/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		//th.TestJSONRequest(t, r, CreatSingleRdsRequet)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, RdsJobidResp)
	})
}
func HandleListErrorLogSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/instances/3e93d3eb20b34bfbbdcc81a79c1c3045/errorlog", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		//th.TestJSONRequest(t, r, CreatSingleRdsRequet)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, ErrorLogResp)
	})
}

func HandleListSlowLogSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/instances/3e93d3eb20b34bfbbdcc81a79c1c3045/slowlog", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		//th.TestJSONRequest(t, r, CreatSingleRdsRequet)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, SlowLogResp)
	})
}

func HandleListSlowLogOptSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/instances/3e93d3eb20b34bfbbdcc81a79c1c3045/slowlog", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, SlowLogResp)
		params := r.URL.Query()
		limit := params.Get("limit")
		startDate := params.Get("start_date")
		endDate := params.Get("end_date")
		fmt.Println(startDate,endDate,limit)
		th.CheckEquals(t,"10",limit)
	})
}