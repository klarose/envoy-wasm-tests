package staticheader

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"istio.io/proxy/test/envoye2e/driver"
	"istio.io/proxy/test/envoye2e/env"
	"istio.io/proxy/testdata"

	test "github.com/klarose/envoy-wasm-tests/e2etest"
)

var TestCases = []struct {
	Name            string
	Method          string
	Host            string
	Path            string
	RequestHeaders  map[string]string
	ResponseHeaders map[string]string
	ResponseCode    int
}{
	{
		Name:            "TestDoNothing",
		Method:          "GET",
		Path:            "/api",
		RequestHeaders:  map[string]string{"x-test-input": ""},
		ResponseHeaders: map[string]string{"Dummy-Header": "some value"},
		ResponseCode:    200,
	},
}

func TestNop(t *testing.T) {
	istioVersion := os.Getenv("ISTIO_TEST_VERSION")
	if istioVersion == "" {
		t.Fatal("set ISTIO_TEST_VERSION environment variable")
	}
	for _, testCase := range TestCases {
		t.Run(testCase.Name, func(t *testing.T) {
			params := driver.NewTestParams(t, map[string]string{
				"WasmFile": filepath.Join(env.GetBazelBinOrDie(), "extensions/static-header/static-header.wasm"),
			}, test.ExtensionE2ETests)
			if testCase.Host != "" {
				params.Vars["Host"] = testCase.Host
			}
			params.Vars["ServerHTTPFilters"] = params.LoadTestData("e2etest/staticheader/testdata/server_filter.yaml.tmpl")
			if err := (&driver.Scenario{
				Steps: []driver.Step{
					&driver.XDS{},
					&driver.Update{
						Node: "server", Version: "0", Listeners: []string{string(testdata.MustAsset("listener/server.yaml.tmpl"))},
					},
					&driver.Envoy{
						Bootstrap:       params.FillTestData(string(testdata.MustAsset("bootstrap/server.yaml.tmpl"))),
						DownloadVersion: os.Getenv("ISTIO_TEST_VERSION"),
					},
					&driver.Sleep{Duration: 1 * time.Second},
					&driver.HTTPCall{
						Port:            params.Ports.ServerPort,
						Method:          testCase.Method,
						Path:            testCase.Path,
						RequestHeaders:  testCase.RequestHeaders,
						ResponseHeaders: testCase.ResponseHeaders,
						ResponseCode:    testCase.ResponseCode,
					},
				},
			}).Run(params); err != nil {
				t.Fatal(err)
			}
		})
	}
}
