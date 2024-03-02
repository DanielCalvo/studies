
//if object has certain fields, do certain things
const person = {
    name: "Max",
    age: 30,
    hobbies: ["Rock Climbing", "Being cool"]
}

if (person.name) {
    console.log("person.name is true")
}

if (person.haspets) { //This field does not exist on the object, I wonder what gives
    console.log("person.haspets is true")
} else {
    console.log("person.haspets is false. I wonder if I can add an attribute to an object it?")
    person.haspets = true; //I can!
}
console.log(person)

/*
The == operator compares the values of two variables after performing type conversion if necessary.
On the other hand, the === operator compares the values of two variables without performing type conversion
* */

const mynum = 10;
const mystring = "10";
console.log(mynum == mystring) //true
console.log(mynum === mystring) //false

// Don't confuse the boolean object with the boolean primitive. You should rarely find yourself using the boolean object
const b = new Boolean(false)
console.log(b)

//Super duper basic, but remember let is an initializer, can't do it twice:
let mynumberino = 0
// let mynumberino = 0 //errors!
//change your variable:
mynumberino = 1

// Lets familiarize ourselves with some falsy values
//Falsy values are:

/*
false
undefined
null
0
NaN
the empty string ("")
* */

mybool = false
if (! mybool) {
    console.log("Yep my bool is false")
}
// This is undefined:
if (! person.invalid_attribute) {
    console.log("An attribute of an object that doesn't exist counts as undefined. Nice!")
}

// lets say we want to try to find an element, but if we can't, we return null
function findElement(elementId){
    // Lets imagine we only have an element with an elementId of 1
    if (elementId === 1) {
        return {}
    } else {
        return null
    }

}

elem = findElement(2)
if (findElement(2)){
    console.log("this will never run")
} else if (findElement(2 === null)) {
    console.log("Yep findElement is null")
}

//is empty object null?
console.log({} === null) //no
//is a variable that is initialized but not given a value null?
let somevar
console.log(somevar === null) //no
console.log(somevar === undefined) //yes
console.log(typeof somevar) //undefined, interesting

//0 is also false, seems straightforward enough
let myzero = 0
console.log(myzero === false) //its false, buuuut
if (0) {
    console.log("This line will never run")
} else {
    console.log("this will as 0 is false")
}

dividedBy0 = 1 / "asd"
if (dividedBy0) {
    console.log("Oughta be false")
} else {
    console.log("If you try to divide a number by a string you get Nan, which is falsy but actually a number:", typeof dividedBy0)
}

//but if you try divide by 0 you can an infinity type:
console.log(1 / 0) //Infinity! Though its considered a member of the number type
// Can I compare types?
console.log(typeof 1/0 === typeof 123) //its false though

console.log("" == false)
if ("") {
    console.log("Oughta be false")
} else {
    console.log("Yeah empty string is false!")
}

//Interesting though, there oughta be some type conversion going above, because:
console.log("" == false) //evaluates to true, which means empty string is false in this context
console.log("" === false) //evaluates to false if you don't convert the types, which makes sense

//uh-oh, an alert function -- doesn't work out of the box though?
// alert(
//     console.log("This is inside an alert, something's wrong")
// )
