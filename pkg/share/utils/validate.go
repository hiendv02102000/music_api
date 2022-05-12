package utils

import "be_soc/pkg/share/validators"

func CheckValidate(x interface{}) error {
	return validators.GetValidator().Struct(x)
}
