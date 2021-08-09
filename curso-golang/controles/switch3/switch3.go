package main

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "Não sei"
	}
}
