package tmdb

import (
	"bytes"
	"encoding/json"
)

func prettyJSON(body []byte) string {
	var buf bytes.Buffer
	err := json.Indent(&buf, body, " ", "\t")
	if err != nil {
		return string(body)
	}
	return string(buf.Bytes())
}
