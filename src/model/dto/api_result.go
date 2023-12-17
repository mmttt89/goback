package dto

type JsonSuccess struct {
	Data interface{}
}

type JsonError struct {
	Code    int
	Message string
}
