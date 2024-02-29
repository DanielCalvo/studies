//Here's a simple function:

function square(number) {
    return number * number;
}
console.log(square(4))

//When you pass an object, if the function changes the objects properties, that change is visible outside the function:
function myFunc(theObject) {
    theObject.make = "Toyota";
}

const mycar = {
    make: "Honda",
    yeah: 1999
}
console.log(mycar)
myFunc(mycar)
console.log(mycar)

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

//A function that calls itself is a recursive function. In some ways a recurse function is analogous to a loop
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