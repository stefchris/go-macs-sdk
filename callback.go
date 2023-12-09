// Copyright (C) 2023 Stefan Christen <s.christen@dycom.ch>.
//
// Use of this source code is prohibited without
// written permission.

package sdk

type Callback func(request *Request, response *Response) error
