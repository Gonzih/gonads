package result

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestBasicHttpCall(t *testing.T) {
	From(http.Get("https://gonzih.me")).
		FMap(func(in F) Result {
			resp := in.(*http.Response)
			defer resp.Body.Close()
			return From(ioutil.ReadAll(resp.Body))
		}).
		Map(func(in F) T {
			body := in.([]byte)
			return string(body)
		}).
		Map(func(in F) T {
			s := in.(string)
			log.Println(len(s))
			return s
		})
}
