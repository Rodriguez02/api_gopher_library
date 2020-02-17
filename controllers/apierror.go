/*
	apierror.go is a script that shows the error in the API
	but customised to identify easily the problem, thanks to
	error control in services.go
*/

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
		case services.ErrorNoName, services.ErrorNoSurname, services.ErrorInvalidID, 
				services.ErrorUserExists, services.ErrorSpecialCharInBooks, services.ErrorNoAvailability, 
				services.ErrorInvalidDueDate, services.ErrorInvalidFormatDate, services.ErrorLoanExists,
				services.ErrorNoIDBook, services.ErrorNoIDUser, services.ErrorNoDueDate, services.ErrorExpiredLoans:
			return ApiError{400, e.Error()}
		case services.ErrorUsersNotFound, services.ErrorUserNotFound, services.ErrorLoanNotFound, 
				services.ErrorLoansNotFound, services.ErrorBookNotFound:
			return ApiError{404, e.Error()}
		default:
			return ApiError{500, e.Error()}
	}
}
