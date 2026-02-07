//Classic function. This function is hoisted, meaning you can call it before it is defined.
//This is also known as a standard function, or a named function
function square(number) {
    return number * number;
}

console.log(square(4))

//This is a function expression, also known as function as a variable, or anonymous function. These functions are not hoisted,
const mysquare = function(x) {
    return x * x;
}
console.log(mysquare(5))

//And this is called an arrow function
const arrowsquare = (x) => {
    return x * x;
}
console.log(arrowsquare(6))

//There is also a concise syntax for a single return expression
const shortsquare = num => num * num;
console.log(arrowsquare(7))

//Shock, absolute beginner shock: Javascript only lets you return one value! This does not work like it would in go!
//const salutepeople = (a,b) => {
//    return `Hello ${a}`, `Hello ${b}`
//}

//Here's a fixed version, or the JS way of doing it:
//a,b are arguments, followed by the function body. In this case the body is a single line.
const salutepeople = (a,b) => {
    return [`Hello ${a}`, `Hello ${b}`]
}

const [s1, s2] = salutepeople('alice', 'bob')
console.log(s1, s2)

//Lets do another function: A function that receives an array of numbers and returns the highest and lowest values!
const getminmax = (num) => {
    const min = Math.min(...num)
    const max = Math.max(...num)
    return {min, max} //Huh, it seems I stumbled upon variadic parameters in go. In js this seems to be called the rest parameter.
}
//Lets do a function with the rest parameter -- this function receives an array name and sums alls of its elements
const

const mycar = {
    make: "Honda",
    yeah: 1999
}

//When you pass an object, if the function changes the objects properties, that change is visible outside the function:
function myFunc(theObject) {
    theObject.make = "Toyota";
}
console.log("My car was a Honda, but now its a " + mycar.make)

//Same deal happens with an arrays
const arr = [100]
function arrFunc (arr) {
    arr[0] = 99
}
arrFunc(arr)
console.log(arr)

//A function can be anonymous -- it does not have to have a name
const sq = function (number) {
    return number * number;
}
console.log(sq(5))

//But you can provide a name with the function expression, this allows the function to refer itself
const factorial1 = function fac(n) {
    return n < 2 ? 1 : n * fac(n - 1);
}

//You can also define functions based on an if condition. Interestting!
let someFunc;
num = 0;
if (num === 0) {
    someFunc = function (theObject) {
        theObject.make = "Toyota";
    };
}

//On calling functions
//Protip: Defining a function does not execute it
//A funtion can call itself

function factorial2(n) {
    if (n === 0 || n === 1) {
        return 1;
    } else {
        return n * factorial(n - 1);
    }
}

//Turns out functions themselves are objects (gonna need an example on this later I think
// You can also do function hoisting, which is calling a function and defining it later, as the JS interpreter "hoists" the entire function declaration to the top of the scope

console.log(addTwo(6,6))
function addTwo(a, b) {
    return a+b
}
//Though hoisting only works with function declarations, not function declarations. This doesn't work:
// console.log(multiplyByTwo(8))
// const multiplyByTwo = function (n) {
//     return n*2
// }

//On scoping: Variables defined inside a function cannot be accessed anywhere outside a function
//However, a function can access all variables and functions defined inside the scope in which it is defined

const num1 = 20;
const num2 = 3;
function multiply() {
    return num1 * num2;
}
console.log(multiply()); //returns 60, good lord this is a bad idea

//A function can refer to and call itself. Consider this function:
const foo = function bar() {
    // statements go here
};

// You can call the function by calling:
// bar()
// arguments.callee()
// foo()

//A function that calls itself is a recursive function. In some ways a recursive function is analogous to a loop
let x = 0;
while (x < 10) {
    x++
}

//Can be converted into a recursive function
function loop(x){
    if (x >= 10) {
        return;
    }
    loop(x+1)
}
loop(2)

//Closures: In JS you can nest functions, and inner functions can access variables from the outer function
// This seems to get really elaborated -- I wonder if using an object with functions just isn't a better a idea?
const pet = function (name) {
    const getName = function() {
        return name
    }
    return getName; //Returns the getName functio
}

const myPet = pet("Joey")
console.log(myPet())

//there's an arguments object!
function someArgs(something) {
    console.log(arguments) //Huh, what's this?
    console.log("Argument is:",typeof arguments)
    console.log(arguments[0])
}
someArgs("banana", "apple", "orange", "melon")

//You can also have default parameters:
function multiply(a, b = 1) {
    return a*b
}
console.log(multiply(5), multiply(5,5))

//Rest parameters: Allows you to represent an indefinite number of arguments as an array
function multiply2(multiplier, ...myArgs) {
    return myArgs.map((x) => multiplier * x)
}
myArr = multiply2(2, 1,2,3,4)

