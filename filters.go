package main

import "strings"

func isInteresting(extension string, runpkg runconf) bool {
	return strIsAnyOf(strings.ToLower(extension), runpkg.LookFor)
}

func fileIsToBeFilteredOut(path string, runpkg runconf) bool {
	return stringContainsAnyOfSubstrings(path, runpkg.FilterOut)
}

func stringContainsAnyOfSubstrings(fullString string, list []string) bool {
	for _, subString := range list {
		if strings.Contains(fullString, subString) {
			return true
		}
	}
	return false
}

func strIsAnyOf(str string, list []string) bool {
	for _, elem := range list {
		if elem == str {
			return true
		}
	}
	return false
}
