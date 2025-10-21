package search

import "strings"

func SearchUser(users[]string, keyword string) []string{
	var result []string
	keyword = strings.ToLower(keyword)

	for _, name := range users {
		if strings.ToLower(name) == keyword { 
			result = append(result, name)
		}
	}
	return result
}