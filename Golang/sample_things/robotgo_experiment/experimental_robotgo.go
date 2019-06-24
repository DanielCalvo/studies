package main

import (
	//"fmt"
	//"github.com/go-vgo/robotgo"
	//"time"
	//"bytes"
	//"log"
	"fmt"
	"github.com/go-vgo/robotgo"
	"math/rand"
	"os"
	"time"
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
var Energy_bomb = rune{
	name:        "Energy Bomb",
	magic_words: "adevo mas vis",
	manacost:    880,
	number_made: 2,
	value:       203,
}

//You can use robotgo.MPoint for this if you want to
//type coordinates struct {
//	position_x int
//	position_y int
//}

func eat(x int, y int) {
	rand.Seed(time.Now().UnixNano())

	robotgo.MovesClick(x+random_int(-5, 5), y+random_int(-5, 5), "right")
	time.Sleep(300 * time.Millisecond)
	robotgo.MovesClick(x+random_int(-5, 5), y+random_int(-5, 5), "right")
	time.Sleep(300 * time.Millisecond)
	robotgo.MovesClick(x+random_int(-5, 5), y+random_int(-5, 5), "right")
	time.Sleep(300 * time.Millisecond)
	robotgo.MovesClick(x+random_int(-5, 5), y+random_int(-5, 5), "right")
	time.Sleep(300 * time.Millisecond)
}

func put_life_ring_on() {

	empty_ring_slot := robotgo.OpenBitmap(image_directory + "empty_ring_slot_v2.png")
	empty_ring_slot_x, empty_ring_slot_y := robotgo.FindBitmap(empty_ring_slot, robotgo.CaptureScreen(), 0.1)

	if empty_ring_slot_x == -1 || empty_ring_slot_y == -1 {
		fmt.Println("Life ring slot in use")
		return
	}

	life_ring := robotgo.OpenBitmap(image_directory + "life_ring.png")
	life_ring_x, life_ring_y := robotgo.FindBitmap(life_ring, robotgo.CaptureScreen(), 0.0)

	if life_ring_x == -1 || life_ring_y == -1 {
		fmt.Println("No life ring found")
		return
	}

	fmt.Println("Life ring: ", life_ring_x, life_ring_y)
	fmt.Println("empty ring slot: ", empty_ring_slot_x, empty_ring_slot_y)

	robotgo.MoveMouseSmooth(life_ring_x, life_ring_y)
	time.Sleep(500 * time.Millisecond)

	robotgo.DragSmooth(empty_ring_slot_x, empty_ring_slot_y)
}

func random_int(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}

func make_rune(rune string) {
	local_chat := robotgo.OpenBitmap(image_directory + "local_chat.png")
	local_chat_x, local_chat_y := robotgo.FindBitmap(local_chat, robotgo.CaptureScreen(), 0.0)
	robotgo.MovesClick(local_chat_x+random_int(0, 500), local_chat_y+random_int(20, 40))
	//Apparently the longer thee delay is, the faster the string is typed, wtf lol
	robotgo.TypeStrDelay(rune, 1000)
	robotgo.KeyTap("enter")
}

func check_blank_rune() {
	blank_rune := robotgo.OpenBitmap(image_directory + "blank_rune.png")
	blank_rune_x, blank_rune_y := robotgo.FindBitmap(blank_rune, robotgo.CaptureScreen(), 0.1)

	if blank_rune_x == -1 || blank_rune_y == -1 {
		fmt.Println("Blank rune not found, runemaking not possible, exiting")
		os.Exit(1)
	} else {
		fmt.Println("Found a blank rune at", blank_rune_x, blank_rune_y)
	}
}

func drink_manapot() {

}

var use_life_ring bool = true
var use_mana_pot bool = false
var mana_pot_interval = 0

//Set these up before running the program
var manabar_x int = 1024
var manabar_y int = 32

var food_slot_x int = 1530
var food_slot_y int = 428

var character_center_screen_x int = 841
var character_center_screen_y int = 488

//Don't forget the bar at the end!
var image_directory string = "/home/daniel/PycharmProjects/studies/Golang/sample_things/robotgo_experiment/images/"

func main() {
	//session_mana_spent := 0
	//session_value_generated := 0
	manabar_full_counter := 0

	for {
		check_blank_rune()

		if use_life_ring == true {
			put_life_ring_on()
		}

		manabar_color := robotgo.GetPixelColor(manabar_x, manabar_y)
		if manabar_color == "00469b" {
			fmt.Println("Manabar full")
			eat(food_slot_x, food_slot_y)
			make_rune(Energy_bomb.magic_words)
			if manabar_full_counter > 3 {
				drink_manapot() //Shouldn't you drink the manaport when the manaber is empty?
			}

			manabar_full_counter += 1
		} else if manabar_color == "2a2b2a" {
			fmt.Println("Manabar empty")
			//drink_potion()
			manabar_full_counter = 0
		} else {
			fmt.Println("Tibia not open or misaligned")
		}

		if manabar_full_counter > 10 {
			fmt.Println("Mana full for more than 10 times in a row, something's wrong")
			os.Exit(1)
		}
		fmt.Println("manabar_full_counter:", manabar_full_counter)

		fmt.Println("Sleeping 10")
		time.Sleep(10 * time.Second)
	}

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
