package Common

import (
    "fmt"
    "encoding/json"
    "github.com/gophercloud/gophercloud"
)

//打印結果
func Show(v interface{}) {
    f, err := json.Marshal(v)
    if err != nil {
        fmt.Printf("Show %s has err : %s", v, err)
    }
    fmt.Println("Result:", string(f))
}

func CheckErr(err error) bool {
    if err != nil {
        if ue, ok := err.(*gophercloud.UnifiedError); ok {
            fmt.Printf("Fail! \n[ ErrCode:%s \n Message:%s ] \n", ue.ErrorCode(), ue.Message())
            return true
        } else {
            fmt.Printf("Err::%s \n", err)
            return true
        }
    }
    return false
}
