package main

func DebugLog(args ...string) []string {
	result := []string{"[DEBUG]"}
	for _, arg := range args {
		result = append(result, arg)
	}
	return result
}

func InfoLog(args ...string) []string {
	result := []string{"[INFO]"}
	for _, arg := range args {
		result = append(result, arg)
	}
	return result
}

func ErrorLog(args ...string) []string {
	result := []string{"[ERROR]"}
	for _, arg := range args {
		result = append(result, arg)
	}
	return result
}
