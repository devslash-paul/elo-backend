package models

type ModelValidation struct {
	Field string
	Error string
}

type ModelError struct {
	Errors []ModelValidation
}
