package models

type FormError struct {
	Field string
	Error string
}

type FormResponse struct {
	Errors []FormError
	Data   any
}

func NewFormResponse() FormResponse {
	var errors []FormError
	response := FormResponse{
		Errors: errors,
		Data:   nil,
	}
	return response
}

func (fr *FormResponse) AddError(field string, message string) {
	err := FormError{
		Field: field,
		Error: message,
	}
	fr.Errors = append(fr.Errors, err)
}
