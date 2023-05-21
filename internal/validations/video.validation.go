package validations

import (
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/constants"
	validator "github.com/go-playground/validator/v10"
	"gocv.io/x/gocv"
)

func ValidateResponseInterfaces(fl validator.FieldLevel) bool {
	responseInterfaceSet := constants.GetResponseInterfaceSet()
	responseInterfaces := fl.Field().Interface().([]string)
	for _, responseInterface := range responseInterfaces {
		if _, ok := responseInterfaceSet[constants.ResponseInterface(responseInterface)]; !ok {
			return false
		}
	}
	return true
}

func ValidateVideoCapturable(fl validator.FieldLevel) bool {
	cap, err := gocv.OpenVideoCapture(fl.Field().String())
	if err != nil {
		return false
	}
	defer cap.Close()
	return cap.IsOpened()
}
