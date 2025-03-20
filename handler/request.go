package handler

import (
	"fmt"
)

func errParamIdRequired(name, typ string) error {
	return fmt.Errorf("%v is required, type: %v", name, typ)
}

//CreateOpening

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (req *CreateOpeningRequest) Validate() error {
	if req == nil {
		return fmt.Errorf("malformed body request")
	}
	if req.Role == "" {
		return errParamIdRequired("role", "string")
	}
	if req.Location == "" {
		return errParamIdRequired("location", "string")
	}
	if req.Company == "" {
		return errParamIdRequired("company", "string")
	}
	if req.Remote == nil {
		return errParamIdRequired("remote", "boolean")
	}
	if req.Link == "" {
		return errParamIdRequired("link", "string")
	}
	if req.Salary <= 0 {
		return errParamIdRequired("salary", "int")
	}

	return nil
}

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (req *UpdateOpeningRequest) Validate() error {
	if req.Role == "" &&
		req.Location == "" &&
		req.Company == "" &&
		req.Remote == nil &&
		req.Link == "" &&
		req.Salary <= 0 {

		return fmt.Errorf("at least one field must be provided to update")
	}

	return nil
}
