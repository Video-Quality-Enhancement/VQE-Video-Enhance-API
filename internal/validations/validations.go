package validations

import (
	"github.com/gin-gonic/gin/binding"
	validator "github.com/go-playground/validator/v10"
	"golang.org/x/exp/slog"
)

func RegisterVideoValidations() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// v.RegisterValidation("are-response-interfaces-valid", ValidateResponseInterfaces)
		slog.Info("Video validations registered", "validator", v) // remove this later if u add more validations
	}
}
