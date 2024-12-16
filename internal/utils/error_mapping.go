package utils

import "gorm.io/gorm"

func MapError(err error) (int, string) {
	switch err {
	case gorm.ErrRecordNotFound:
		return 404, "registro não encontrado."
	default:
		return 500, err.Error()
	}
}
