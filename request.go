// Copyright (C) 2023 Stefan Christen <s.christen@dycom.ch>.
//
// Use of this source code is prohibited without
// written permission.

package sdk

import (
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

func (request *Request) handle(callback Callback) (*Response, error) {
	response := Response{
		Version: 1,
		Code:    "OK",
		Title:   request.Config.Title,
	}

	request.Method = strings.ToUpper(request.Method)

	err := callback(request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
