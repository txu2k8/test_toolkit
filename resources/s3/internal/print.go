package internal

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/minio/pkg/console"
)

// message interface for all structured messages implementing JSON(), String() methods.
type message interface {
	JSON() string
	String() string
}

// printMsg prints message string or JSON structure depending on the type of output console.
func printMsg(msg message) {
	var msgStr string
	if !globalJSON {
		msgStr = msg.String()
	} else {
		msgStr = msg.JSON()
		if globalJSONLine && strings.ContainsRune(msgStr, '\n') {
			// Reformat.
			var dst bytes.Buffer
			if err := json.Compact(&dst, []byte(msgStr)); err == nil {
				msgStr = dst.String()
			}
		}
	}
	console.Println(msgStr)
}
