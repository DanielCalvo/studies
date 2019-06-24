package main

import (
	//"fmt"
	//"github.com/go-vgo/robotgo"
	//"time"
	//"bytes"
	//"log"
	"fmt"
	"github.com/go-vgo/robotgo"
)

type rune struct {
	name        string
	magic_words string
	manacost    int
	number_made int
	value       int
}

var HMM = rune{
	name:        "HMM",
	magic_words: "adori vis",
	manacost:    350,
	number_made: 10,
	value:       12,
}
var GFB = rune{
	name:        "GFB",
	magic_words: "adori mas flam",
	manacost:    530,
	number_made: 4,
	value:       57,
}
var Thunderstorm = rune{
	name:        "Thunderstorm",
	magic_words: "adori mas vis",
	manacost:    430,
	number_made: 4,
	value:       47,
}
var SD = rune{
	name:        "SD",
	magic_words: "adori mas vis",
	manacost:    985,
	number_made: 3,
	value:       135,
}

//You can use robotgo.MPoint for this if you want to
//type coordinates struct {
//	position_x int
//	position_y int
//}

var use_life_ring bool = true
var use_mana_pot bool = false
var mana_pot_interval = 0

var manabar_x int = 0
var manabar_y int = 0

var food_slot_x int = 0
var food_slot_y int = 0

func main() {

	//session_mana_spent := 0
	//session_value_generated := 0
	//manabar_full_counter := 0

	magic_words := SD.magic_words
	fmt.Println(magic_words)

	//Things done: Set up the counters
	//Set up logging functionalities

	//###
	//Functionality 1: Find image on screen
	//###
	//manabar_full := robotgo.OpenBitmap("/home/daniel/PycharmProjects/studies/Golang/sample_things/manabar_full.png")
	//manabar_square_12x12 := robotgo.OpenBitmap("/home/daniel/PycharmProjects/studies/Golang/sample_things/manabar_square_12x12.png")
	//fmt.Println(robotgo.FindEveryBitmap(manabar_square_12x12, manabar_full, 0.1))
	////Move mouse to a part of the screen:
	//xfce_x, xfce_y := robotgo.FindBitmap(xfce_menu_bitmap, robotgo.CaptureScreen(), 0.1)
	//robotgo.MovesClick(xfce_x, xfce_y)

	//###
	//Functionality 2: Logging with timestamps
	//###
	//log.Println("Hello world!)

	//Feature ideas:
	//Find a way to use mana pots with defined intervals (ex: one each minute)

}

//FEATURES FOR LATER:

// ----- 1
//Let's specify the data in a .yaml file and read it from there!
//Things to specify:
//Runes
//Positions on screen (manabar, food item, center of screen if using potions on self)
//if using rings
//if using potions, and the interval!
