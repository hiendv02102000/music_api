package utils

import "music-api/pkg/share/validators"

func CheckValidate(x interface{}) error {
	return validators.GetValidator().Struct(x)
}
