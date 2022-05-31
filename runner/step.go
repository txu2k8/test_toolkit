package runner

type StepType string

// 测试步骤类型
const (
	stepTypeUC        StepType = "uc"
	stepTypeSDK       StepType = "sdk"
	stepTypeAPI       StepType = "api"
	stepTypeUI        StepType = "ui"
	stepTypeTestCase  StepType = "testcase"
	stepTypeThinkTime StepType = "thinktime"
	stepTypeOther     StepType = "other"
)

// 测试步骤结果
type StepResult struct {
	Name       string      `json:"name" yaml:"name"`                                 // step name
	StepType   StepType    `json:"step_type" yaml:"step_type"`                       // step type, testcase/uc/sdk/api/ui
	Success    bool        `json:"success" yaml:"success"`                           // step execution result
	Elapsed    int64       `json:"elapsed_ms" yaml:"elapsed_ms"`                     // step execution time in millisecond(ms)
	Data       interface{} `json:"data,omitempty" yaml:"data,omitempty"`             // session data or slice of step data
	Attachment string      `json:"attachment,omitempty" yaml:"attachment,omitempty"` // step error information
}

// 测试步骤
type TStep struct {
	Name     string   `json:"name" yaml:"name"`           // required
	StepType StepType `json:"step_type" yaml:"step_type"` // required
}

// IStep represents interface for all types for teststeps, includes:
// StepRequest, StepRequestWithOptionalArgs, StepRequestValidation, StepRequestExtraction,
// StepTestCaseWithOptionalArgs,
// StepTransaction, StepRendezvous, StepWebSocket.
type IStep interface {
	Name() string
	Type() StepType
	Struct() *TStep
	Run(*SessionRunner) (*StepResult, error)
}
