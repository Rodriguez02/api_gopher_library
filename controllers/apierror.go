package controllers

import "api_gopher_library/services"

type ApiError struct {
	Status  int
	Message string
}

func (e *ApiError) Error() string {
	return e.Message
}

func parseError(e error) ApiError {
	switch e {
	case services.ErrorNoName, services.ErrorNoSurname, services.ErrorInvalidID, services.ErrorUserExists, services.ErrorSpecialCharInBooks:
		return ApiError{400, e.Error()}
	case services.ErrorUsersNotFound, services.ErrorUserNotFound:
		return ApiError{404, e.Error()}
	default:
		return ApiError{500, e.Error()}
	}
}
