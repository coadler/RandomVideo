package lights

import (
    "fmt"
    "net/http"
    "math/rand"
    "time"
)

var VideoList = []string{

}

func init() {
    http.HandleFunc("/", Index)
    http.HandleFunc("/.*", Index)
}
