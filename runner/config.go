package runner

type TConfig struct {
	Name       string                 `json:"name" yaml:"name"` // required
	Verify     bool                   `json:"verify,omitempty" yaml:"verify,omitempty"`
	BaseURL    string                 `json:"base_url,omitempty" yaml:"base_url,omitempty"`
	Headers    map[string]string      `json:"headers,omitempty" yaml:"headers,omitempty"`
	Variables  map[string]interface{} `json:"variables,omitempty" yaml:"variables,omitempty"`
	Parameters map[string]interface{} `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Export     []string               `json:"export,omitempty" yaml:"export,omitempty"`
	Weight     int                    `json:"weight,omitempty" yaml:"weight,omitempty"`
	Path       string                 `json:"path,omitempty" yaml:"path,omitempty"` // testcase file path
}
