package main

func Repeat(s string, reps int) string {
	var repeated string
	for i := 0; i < reps; i++ {
		repeated += s // didn't know you could use += with strings, neat
	}
	return repeated
}
