// Copyright (C) 2023 Stefan Christen <s.christen@dycom.ch>.
//
// Use of this source code is prohibited without
// written permission.

package sdk

import (
	"fmt"
	"strings"
)

type Request struct {
	Version int
	Method  string
	Module  string
	Config  struct {
		Title string
		Prefs map[string]string
	}
	Args map[string]string
	Form map[string]string
}

func (request *Request) handle(callbacks map[string]Callback) (*Response, error) {
	response := Response{
		Version: 1,
		Code:    "OK",
		Title:   request.Config.Title,
	}

	request.Method = strings.ToUpper(request.Method)

	key := strings.ToUpper(request.Module)
	if keyFromPrefs := strings.ToUpper(request.Config.Prefs["MODULE"]); keyFromPrefs != "" {
		key = keyFromPrefs
	}

	callback, found := callbacks[key]
	if !found {
		return nil, fmt.Errorf("module not found: %s", key)
	}

	err := callback(request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
