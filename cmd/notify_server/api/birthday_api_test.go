package api

import (
	"bytes"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/test_data"

	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/localconf"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type apiTestCase struct {
	tag              string
	method           string
	urlToServe       string
	urlToHit         string
	body             string
	function         gin.HandlerFunc
	status           int
	responseFilePath string
}

func newRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	localconf.Config.DB = test_data.ResetDB()

	return router
}

// Used to run single API test case. It makes HTTP request and returns its response
func testAPI(router *gin.Engine, method string, urlToServe string, urlToHit string, function gin.HandlerFunc, body string) *httptest.ResponseRecorder {
	router.Handle(method, urlToServe, function)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(method, urlToHit, bytes.NewBufferString(body))
	router.ServeHTTP(res, req)
	return res
}
func TestAPI(t *testing.T) {
	path := test_data.GetTestCaseFolder()
	tests := []apiTestCase{
		{"t1 - get a User", "GET", "/users/:id", "/users/1", "", GetBirthday, http.StatusOK, path + "/user_t1.json"},
		{"t2 - get a User not Present", "GET", "/users/:id", "/users/9999", "", GetBirthday, http.StatusNotFound, ""},
	}
	for _, test := range tests {
		router := newRouter()
		res := testAPI(router, test.method, test.urlToServe, test.urlToHit, test.function, test.body)
		assert.Equal(t, test.status, res.Code, test.tag)
		if test.responseFilePath != "" {
			response, _ := ioutil.ReadFile(test.responseFilePath)
			assert.JSONEq(t, string(response), res.Body.String(), test.tag)
		}
	}

}
