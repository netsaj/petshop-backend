package utils

func ErrorToMap(err error) map[string]interface{} {
	return map[string]interface{}{
		"error": err.Error(),
	}
}
