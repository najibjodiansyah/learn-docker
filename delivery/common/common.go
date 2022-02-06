package common

// return c.JSON(http.StatusOK, map[string]interface{}{
// 			"message": "success get all users",
// 			"users":   users,
// 		})

//DefaultResponse default payload response
type DefaultResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//NewInternalServerErrorResponse default internal server error response
func SuccessOperation() DefaultResponse {
	return DefaultResponse{
		200,
		"Successful Operation : OK",
	}

}
func SuccesOperationWithData(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code" : 200,
		"message" : "Successful Operation : OK",
		"data" : data,
	}
}


// NewInternalServerErrorResponse default internal server error response
func SuccessOperationCreated() DefaultResponse {
	return DefaultResponse{
		201,
		"Successful Operation : Created",
	}
}

//NewInternalServerErrorResponse default internal server error response
func InternalServerError() DefaultResponse {
	return DefaultResponse{
		500,
		"Internal Server Error",
	}
}

//NewNotFoundResponse default not found error response
func NotFound() DefaultResponse {
	return DefaultResponse{
		404,
		"Not Found",
	}
}

//NewBadRequestResponse default not found error response
func BadRequest() DefaultResponse {
	return DefaultResponse{
		400,
		"Bad Request",
	}
}

//ForbiddedRequest default not found error response
func ForbiddedRequest() DefaultResponse {
	return DefaultResponse{
		403,
		"Forbidded Request",
	}
}

//NewConflictResponse default not found error response
func Unauthorized() DefaultResponse {
	return DefaultResponse{
		401,
		"Unauthorized",
	}
}

//NewConflictResponse default not found error response
func Conflict() DefaultResponse {
	return DefaultResponse{
		409,
		"Data Has Been Modified",
	}
}
