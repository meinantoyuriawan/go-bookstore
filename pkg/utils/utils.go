package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// func ParseBody(r *http.Request, x interface{}) {
// 	if body, err := ioutil.ReadAll(r.Body); err != nil {
// 		if err := json.Unmarshal([]byte(body), x); err != nil {
// 			return
// 		}
// 	}
// }

func ParseBody(r *http.Request, x interface{}) {
	body, err := io.ReadAll(r.Body)

	if err == nil {
		if er := json.Unmarshal([]byte(body), x); er != nil {
			return
		}
	}

}
