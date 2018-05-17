package tests

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gopkg.in/resty.v1"
)

var (
	APIEndpoint = "http://localhost:8080"
	baseDate    = time.Now().Add(24 * time.Hour).Format("2006-01-02")
	regex       = regexp.MustCompile("\\${(.+?)}")
)

type (
	input struct {
		method string
		url    string
		body   string
	}
	output struct {
		status   int
		body     string
		callback func(resp map[string]interface{})
	}
)

func runTests(t *testing.T, tests []struct {
	name   string
	input  input
	output output
}, binding *map[string]interface{}) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			callAPI(t, test.input, test.output, binding)
		})
	}
}

func callAPI(t *testing.T, in input, out output, binding *map[string]interface{}) {
	resty.SetAllowGetMethodPayload(true)
	resty.SetHostURL(APIEndpoint)
	//resty.SetDebug(true)

	req := resty.R()
	if in.body != "" {
		body := bindValue(in.body, binding)
		req.SetBody(body)
	}

	resp, err := execute(
		req,
		strings.ToUpper(in.method),
		bindValue(in.url, binding),
	)
	assert.NoError(t, err)

	assert.Equal(t, out.status, resp.StatusCode())

	if out.body != "" {
		actual := unmarshalJSON(resp.String())
		if out.callback != nil {
			if keyValues, ok := actual.(map[string]interface{}); ok {
				out.callback(keyValues)
			}
		}
		expected := unmarshalJSON(bindValue(out.body, binding))

		assert.Equal(t, expected, actual)
	}
}

func execute(req *resty.Request, method, url string) (*resty.Response, error) {
	switch method {
	case "GET":
		return req.Get(url)
	case "POST":
		return req.Post(url)
	case "PUT":
		return req.Put(url)
	case "DELETE":
		return req.Delete(url)
	case "PATCH":
		return req.Patch(url)
	default:
		panic(fmt.Errorf("unsupported http method [%s]", method))
	}
}

func unmarshalJSON(body string) (decodeBody interface{}) {
	if err := json.Unmarshal([]byte(body), &decodeBody); err != nil {
		panic(fmt.Errorf("unmarshal error [%s]", body))
	}
	return
}

func bindValue(value string, binding *map[string]interface{}) string {
	for k, v := range *binding {
		value = strings.Replace(
			value,
			fmt.Sprintf("${%s}", k),
			fmt.Sprintf("%v", v),
			-1,
		)
	}
	return value
}
