package testing

import (
	"fmt"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"net/http"
	"testing"
)

const DbUserResp = `
{
    "resp": "successful"
}
`
const ListDbUserResp  =`
{
	"total_count": 10,
	"users": [{
		"name": "mysql.infoschema"
	}, {
		"name": "mysql.session"
	}, {
		"name": "mysql.sys"
	}, {
		"name": "rdsAdmin"
	}, {
		"name": "rdsBackup"
	}, {
		"name": "rdsMetric"
	}]
}
`
const DeleteFailResp  =`
{
	"ErrorCode":"DBS.200824",
	"Message":"The database account does not exist."

}
`
func HandleCreateDbUserSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/instances/dsfae23fsfdsae3435in01/db_user", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, DbUserResp)
	})
}

func HandleListDbUserSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/instances/dsfae23fsfdsae3435in01/db_user/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, ListDbUserResp)
	})
}
func HandleDeleteDbUserFail(t *testing.T) {
	th.Mux.HandleFunc("/instances/dsfae23fsfdsae3435in01/db_user/rds_009", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, DeleteFailResp)
	})
}
func HandleDeleteDbUserSuccess(t *testing.T) {
	th.Mux.HandleFunc("/instances/dsfae23fsfdsae3435in01/db_user/rds_009", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, DbUserResp)
	})
}
