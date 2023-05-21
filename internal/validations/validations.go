package validations

import (
	"github.com/gin-gonic/gin/binding"
	validator "github.com/go-playground/validator/v10"
)

func RegisterVideoValidations() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("are-response-interfaces-valid", ValidateResponseInterfaces)
		v.RegisterValidation("is-video-capturable", ValidateVideoCapturable)
	}
}
