package common

type ErrorResponse struct {
	ErrorCode    string
	ErrorMessage string
}

func (errorResponse *ErrorResponse) IsValid() bool {
	return errorResponse.ErrorCode == ""
}

type DtoResponse struct {
	Data  interface{}
	Error ErrorResponse
}
