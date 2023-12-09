// Copyright (C) 2023 Stefan Christen <s.christen@dycom.ch>.
//
// Use of this source code is prohibited without
// written permission.

package sdk

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
)

func RunOnce(callback Callback) {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		writeError(os.Stdout, err.Error())
		os.Exit(1)
	}

	data = bytes.TrimRight(data, "\x00")

	var request Request
	err = json.Unmarshal(data, &request)
	if err != nil {
		writeError(os.Stdout, err.Error())
		os.Exit(1)
	}

	response, err := request.handle(callback)
	if err != nil {
		writeError(os.Stdout, err.Error())
		os.Exit(1)
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		writeError(os.Stdout, err.Error())
		os.Exit(1)
	}

	os.Stdout.Write(bytes)
	os.Exit(0)
}
