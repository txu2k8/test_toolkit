package runner

import (
	"time"
)

type SessionRunner struct {
	*testCaseRunner
	startTime time.Time        // record start time of the testcase
	summary   *TestCaseSummary // record test case summary
}
