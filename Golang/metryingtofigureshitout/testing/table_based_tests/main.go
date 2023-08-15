package main

// Lets do something silly for an example here. Does a string contain a given char 3 times?
func DoesItHaveChar3Times(s, char string) bool {
	counter := 0
	for _, v := range s {
		if string(v) == char {
			counter++
		}
	}
	if counter == 3 {
		return true
	}
	return false
}
