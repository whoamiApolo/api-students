package api

import "fmt"

type StudentRequest struct {
	Name   string `json:"name"`
	CPF    int    `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active *bool  `json:"active"` //forçar o ponteiro ser true ou false
}

func errParamRequired(param, typ string) error {
	return fmt.Errorf("param: '%s' of type '%s' is required", param, typ)
}

func (s *StudentRequest) Validate() error {
	if s.Name == "" {
		return errParamRequired("name", "string")
	}

	if s.Email == "" {
		return errParamRequired("name", "string")
	}

	if s.CPF == 0 {
		return errParamRequired("cpf", "int")
	}

	if s.Age == 0 {
		return errParamRequired("name", "int")
	}

	if s.Active == nil {
		return errParamRequired("name", "bool")
	}

	return nil
}
