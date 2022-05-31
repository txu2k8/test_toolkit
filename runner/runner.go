package runner

import (
	"net/http"
	"testing"
)

type TKRunner struct {
	t             *testing.T
	failfast      bool
	httpStatOn    bool
	requestsLogOn bool
	pluginLogOn   bool
	saveTests     bool
	genHTMLReport bool
	httpClient    *http.Client
}

type testCaseRunner struct {
	testCase     *TestCase
	tkRunner     *TKRunner
	parsedConfig *TConfig
	rootDir      string // project root dir
}
