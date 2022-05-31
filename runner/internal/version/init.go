package version

import (
	_ "embed"
)

// go:embed VERSION
var VERSION string

const RunnerMinVersion = "v1.0.0-beta"
