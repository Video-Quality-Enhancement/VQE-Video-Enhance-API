package validations

import (
	validator "github.com/go-playground/validator/v10"
	"gocv.io/x/gocv"
	"golang.org/x/exp/slog"
)

func ValidateVideoCapturable(fl validator.FieldLevel) bool {
	cap, err := gocv.OpenVideoCapture(fl.Field().String())
	if err != nil {
		slog.Error("Error opening video capture file using url", "videoUrl", fl.Field().String(), "error", err.Error())
		return false
	}
	defer cap.Close()
	return cap.IsOpened()
}
