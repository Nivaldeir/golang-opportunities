package handler

import "fmt"

func errParamsIsRequired(name, typ string) error {
	return fmt.Errorf("params: %s (type: %s) is required", name, typ)
}

type CreateOpeningRequest struct {
	Role     string `json: "role"`
	Company  string `json: "company"`
	Location string `json: "location"`
	Remote   *bool  `json: "remote"`
	Link     string `json: "link"`
	Salary   int64  `json: "salary"`
}

func (c *CreateOpeningRequest) Validate() error {
	if c.Role == "" && c.Company == "" && c.Location == "" && c.Remote == nil && c.Salary <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}
	if c.Role == "" {
		return errParamsIsRequired("role", "string")
	}
	if c.Company == "" {
		return errParamsIsRequired("company", "string")
	}
	if c.Location == "" {
		return errParamsIsRequired("location", "string")
	}
	if c.Link == "" {
		return errParamsIsRequired("link", "string")
	}
	if c.Remote == nil {
		return errParamsIsRequired("remote", "boolen")
	}
	if c.Salary <= 0 {
		return errParamsIsRequired("salary", "int")
	}
	return nil
}

type UpdateOpeningRequest struct {
	Role     string `json: "role"`
	Company  string `json: "company"`
	Location string `json: "location"`
	Remote   *bool  `json: "remote"`
	Link     string `json: "link"`
	Salary   int64  `json: "salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Link != "" || r.Salary > 0 {
		return nil
	}
	return fmt.Errorf("at least one field on request be provided")
}
