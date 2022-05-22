package models

// aliasConfig ...
type AliasConfig struct {
	Name      string `json:"name"`
	URL       string `json:"url"`
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}

// s3Config define S3 config
type S3Config struct {
	S3Ip        string // endpoint: https://<S3Ip>:<S3Port>, eg: https://10.25.119.86:443
	S3Port      int    // port (default: 443)
	Endpoint    string // Parse(S3Ip,S3Port) --> Endpoint
	S3AccessID  string
	S3SecretKey string
}

type FileCreateInput struct {
	LocalDataDir  string   // The local data Dir
	FileArgs      []string // Files args array,eg: {"txt:20:1k-10k", "dd:1:100mb"}
	RandomPercent int      // percent of files with random data
	EmptyPercent  int      // percent of files with empty data
	RenameFile    bool     // rename files name each time if true
}

// fileCreateConfig define how to create files
type FileCreateConfig struct {
	Type       string // txt or other data
	Num        int    // file number
	SizeMin    int64  // the min size of file
	SizeMax    int64  // the max size of the file
	NamePrefix string // the file name prefix
	ParentDir  string // the file parent dir path
}

type FileCreateConfigMap struct {
	CreateInput  FileCreateInput    // The FileCreateInput
	CreateConfig []FileCreateConfig // Parse(FileCreateInput) --> FileCreateConfig
}
