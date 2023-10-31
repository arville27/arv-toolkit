package controllers

import (
	"arville27/arv-toolkit/modules/auth"
	"arville27/arv-toolkit/modules/helper"
	"arville27/arv-toolkit/modules/splyr"
	rest_model "arville27/arv-toolkit/rest/model"
	"errors"
	"net/http"
)

// Error response
func ResolveError(err error) (*rest_model.RestError, int) {
	if err != nil {
		var serviceError helper.ServiceError
		if errors.As(err, &serviceError) {
			switch {
			// Auth module
			case errors.Is(serviceError.ErrorType(), auth.InvalidCredential):
				return rest_model.NewRestError(serviceError.Reason(), serviceError.ErrorType().Error()), http.StatusUnauthorized
			// Splyr module
			case errors.Is(serviceError.ErrorType(), splyr.LyricsNotFound):
				return rest_model.NewRestError(serviceError.Reason(), serviceError.ErrorType().Error()), http.StatusNotFound
			case errors.Is(serviceError.ErrorType(), splyr.FailedFetchLyrics):
				return rest_model.NewRestError(serviceError.Reason(), serviceError.ErrorType().Error()), http.StatusBadRequest
			case errors.Is(serviceError.ErrorType(), splyr.FailedRequestingAccessToken):
				return rest_model.NewRestError(serviceError.Reason(), serviceError.ErrorType().Error()), http.StatusUnauthorized
			case errors.Is(serviceError.ErrorType(), splyr.InvalidTrackId):
				return rest_model.NewRestError(serviceError.Reason(), serviceError.ErrorType().Error()), http.StatusBadRequest
			default:
				return rest_model.NewRestError(serviceError.Reason(), serviceError.ErrorType().Error()), http.StatusInternalServerError
			}
		}
	}

	return rest_model.NewRestError("Internal server error", "Internal server error"), http.StatusInternalServerError
}
