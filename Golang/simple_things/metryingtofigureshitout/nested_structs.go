package main

import (
	"fmt"
)

//yaml-to-go sometimes generates a bunch of nested slices of structs that I'm having a difficult time with.
//Let's try figuring this out in incremental steps!

type Person struct {
	Name string
}

type People []struct {
	Name string
	Pets []string
}

type Host struct {
	hostname   string
	Controller struct {
		Name string
		IP   string
	}
}

type Hosts struct {
	Controllers []struct {
		Name string `yaml:"name"`
		IP   string `yaml:"ip"`
	} `yaml:"controllers"`
	Workers []struct {
		Name string `yaml:"name"`
		IP   string `yaml:"ip"`
	} `yaml:"workers"`
}

func main() {
	s := []int{1, 2, 3}
	fmt.Println(s)

	p1 := Person{
		Name: "Joe",
	}
	fmt.Println(p1)

	p2 := []Person{
		{
			Name: "Joe",
		},
	}
	fmt.Println(p2)

	p3 := People{
		{
			Name: "Joe",
			Pets: []string{"Cat", "Dog"},
		},
	}
	fmt.Println(p3)

	myhost := Host{
		hostname: "banana",
		Controller: struct {
			Name string
			IP   string
		}{Name: "asd", IP: "asd"},
	}
	fmt.Println(myhost)

	myhosts := Hosts{ //Damn this isn't very pretty...
		Controllers: []struct {
			Name string `yaml:"name"`
			IP   string `yaml:"ip"`
		}{
			{
				Name: "banana",
				IP:   "123",
			},
			{
				Name: "apple",
				IP:   "456",
			},
		},
		Workers: []struct {
			Name string `yaml:"name"`
			IP   string `yaml:"ip"`
		}{
			{
				Name: "orange",
				IP:   "789",
			},
			{
				Name: "grape",
				IP:   "000",
			},
		},
	}
	fmt.Println(myhosts)

}
