package requests

import (
	"fmt"
)

func ErrParamIdRequired(name, typ string) error {
	return fmt.Errorf("%v is required, type: %v", name, typ)
}

// CreateOpening
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
		return ErrParamIdRequired("role", "string")
	}
	if req.Location == "" {
		return ErrParamIdRequired("location", "string")
	}
	if req.Company == "" {
		return ErrParamIdRequired("company", "string")
	}
	if req.Remote == nil {
		return ErrParamIdRequired("remote", "boolean")
	}
	if req.Link == "" {
		return ErrParamIdRequired("link", "string")
	}
	if req.Salary <= 0 {
		return ErrParamIdRequired("salary", "int")
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

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *LoginRequest) Validate() error {
	if req == nil {
		return fmt.Errorf("malformed body request")
	}
	if req.Email == "" {
		return ErrParamIdRequired("email", "string")
	}
	if req.Password == "" {
		return ErrParamIdRequired("password", "string")
	}

	return nil
}
