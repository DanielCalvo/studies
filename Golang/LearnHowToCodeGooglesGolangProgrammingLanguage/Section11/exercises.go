package main

import "fmt"

func main() {
	//Exercise 1
	var x [5]int
	for i := 0; i < 5; i++ {
		x[i] = i * 10
	}
	for j := 0; j < len(x); j++ {
		fmt.Println(x[j])
	}
	fmt.Printf("X is of type: %T", x)
	fmt.Println()

	//Exercise 2
	y := []int{100, 101, 102, 103, 104, 105, 106, 107, 108, 109}
	fmt.Printf("y is of type: %T", y)

	for j := 0; j < len(y); j++ {
		fmt.Println(y[j])
	}

	//Exercise 4
	fmt.Println("\nExercise 4")
	b := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(b)
	b = append(b, 52)
	fmt.Println(b)
	b = append(b, 53, 54, 55)
	c := []int{56, 57, 58, 59, 60}
	b = append(b, c...)
	fmt.Println(b)

	//Exercise 5
	fmt.Println("\nExercise 5")
	b1 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	b1 = append(b1[:3], b1[6:]...)
	fmt.Println(b1)

	//Exercise 6
	fmt.Println("\nExercise 6")
	fruits := make([]string, 5, 5)
	//Oddly enough, if you append, the lenght and cap of the slice go up to 10
	//fruits = append(fruits, "banana", "apple", "orange", "grape", "mango")
	fruits = []string{"banana", "apple", "orange", "grape", "mango"}
	fmt.Println("Lenght of the fruit slice:", len(fruits))
	fmt.Println("Capacity of the fruit slice:", cap(fruits))

	//Exercise 7
	fmt.Println("\nExercise 7")
	slice1 := []string{"James", "Bond", "Shaken, not stirred"}
	slice2 := []string{"Miss,", "Moneypenny", "Hello"}
	slice3 := [][]string{slice1, slice2}
	//fmt.Println(slice3)

	for i := 0; i < len(slice3); i++ {
		fmt.Println(slice3[i])
		for j := 0; j < len(slice3[i]); j++ {
			fmt.Println(slice3[i][j])
		}
	}

	//Exercise 8
	fmt.Println("\nExercise 8")
	m := map[string][]string{
		"Dani": []string{"Bikes", "Shiny things", "Occasional beer"},
		"Jade": []string{"Food", "Pool water", "Hunting things"},
	}

	for k, v := range m {
		fmt.Print(k, " likes: ")
		for i := 0; i < len(v); i++ {
			fmt.Print(v[i], " ")
		}
		fmt.Println()
	}

	//Exercise 9
	fmt.Println("\nExercise 9")
	m["Shanti"] = []string{"Food", "Jumping around"}
	for k, v := range m {
		fmt.Print(k, " likes: ")
		for i := 0; i < len(v); i++ {
			fmt.Print(v[i], " ")
		}
		fmt.Println()
	}

	//Exercise 10
	fmt.Println("\nExercise 10")
	delete(m, "Dani")
	fmt.Println(m)

}
