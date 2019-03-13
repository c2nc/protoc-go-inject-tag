package main

import "os"

func PathNotExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return true
	}
	return false
}

// FileExists checks to see if a file exists
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func EnvOrDefault(key, value string) string {
	var res string
	if res = os.Getenv(key); res == "" {
		return value
	}
	return res
}