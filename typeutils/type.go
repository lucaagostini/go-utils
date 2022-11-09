package typeutils

func IsString(obj interface{}) bool {
	switch obj.(type) {
	case string:
		return true
	default:
		return false
	}
}
