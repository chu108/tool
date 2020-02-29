package tool

func GetType(obj interface{}) string {
	switch obj.(type) {
	case string:
		return "string"
	case int:
		return "int"
	case int64:
		return "int64"
	case float32:
		return "float32"
	case float64:
		return "float64"
	case bool:
		return "bool"
	default:
		return ""
	}
}
