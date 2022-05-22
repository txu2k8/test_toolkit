package internal

var (
	// Version - version time.RFC3339.
	Version = "DEVELOPMENT.GOGET"
	// ReleaseTag - release tag in TAG.%Y-%m-%dT%H-%M-%SZ.
	ReleaseTag = "DEVELOPMENT.GOGET"
	// CommitID - latest commit id.
	CommitID = "DEVELOPMENT.GOGET"
	// ShortCommitID - first 12 characters from CommitID.
	ShortCommitID = CommitID[:12]
)
