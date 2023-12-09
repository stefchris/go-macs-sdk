package sdk

import (
	"encoding/json"
	"io"
)

func writeError(w io.Writer, text string) {
	response := Response{
		Version: 1,
		Code:    "FAILURE",
		Text:    text,
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		bytes = []byte(`{"Version":1,"Code":"FAILURE","Text":"` + text + `"}`)
	}

	w.Write(bytes)
}
