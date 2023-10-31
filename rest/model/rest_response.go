package rest_model

type RestError struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

type RestResponse struct {
	RestError *RestError `json:"error"`
	Data      any        `json:"data"`
}

func NewRestResponse(data any, restError *RestError) *RestResponse {
	return &RestResponse{Data: data, RestError: restError}
}

func NewRestError(errorMessage string, errorCode string) *RestError {
	return &RestError{ErrorMessage: errorMessage, ErrorCode: errorCode}
}

func RestErrorResponse(errorMessage string, errorCode string) *RestResponse {
	return NewRestResponse(
		nil,
		NewRestError(errorMessage, errorCode),
	)
}

func RestSuccessResponse(data any) *RestResponse {
	return NewRestResponse(
		data,
		NewRestError("Success", "Success"),
	)
}
