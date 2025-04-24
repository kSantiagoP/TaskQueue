package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type %s) is required", name, typ)
}

type CreateTaskRequest struct {
	Type string `json:"type"`
	//Status      string `json:"status"`
	//Error       string `json:"error"`
	ElapsedTime int64 `json:"elapsedTime"`
}

func (r *CreateTaskRequest) Validate() error {
	if r.Type == "" && r.ElapsedTime <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}

	if r.Type == "" {
		return errParamIsRequired("type", "string")
	}

	if r.ElapsedTime <= 0 {
		return errParamIsRequired("elapsedTime", "int64")
	}

	return nil
}
