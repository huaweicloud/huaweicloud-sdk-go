package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

func MockListResponse(t *testing.T) {
	th.Mux.HandleFunc("/types", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		r.ParseForm()
		marker := r.Form.Get("marker")
		switch marker {
		case "":
			fmt.Fprintf(w, `
{ 
    "volume_types": [ 
        { 
            "extra_specs": { 
                "volume_backend_name": "SAS",  
                "availability-zone": "az1.dc1" 
            },  
            "name": "SAS",  
            "qos_specs_id": null,  
            "id": "6c81c680-df58-4512-81e7-ecf66d160638",  
            "is_public": true,  
            "description": null 
        },  
        { 
            "extra_specs": { 
                "volume_backend_name": "SATA",  
                "availability-zone": "az1.dc1" 
            },  
            "name": "SATA",  
            "qos_specs_id": "585f29d6-7147-42e7-bfb8-ca214f640f6f",  
            "is_public": true,  
            "id": "ea6e3c13-aac5-46e0-b280-745ed272e662",  
            "description": null 
        }
    ] 
}

  `)
		case "1":
			fmt.Fprintf(w, `{"volume_types": []}`)
		default:
			t.Fatalf("Unexpected marker: [%s]", marker)
		}
	})
}

func MockGetResponse(t *testing.T) {
	th.Mux.HandleFunc("/types/ea6e3c13-aac5-46e0-b280-745ed272e662", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `
{ 
    "volume_type": { 
        "extra_specs": { 
            "volume_backend_name": "SATA",  
            "availability-zone": "az1.dc1" 
        },  
        "name": "SATA",  
        "qos_specs_id": null,  
        "is_public": true,  
        "id": "ea6e3c13-aac5-46e0-b280-745ed272e662",  
        "description": null 
    } 
}

`)
	})
}
