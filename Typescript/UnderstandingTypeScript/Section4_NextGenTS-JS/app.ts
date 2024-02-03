console.log("Bazinga");

const userName = "Aaaa"; //const cannot be changed
let age = 0; //can be changed
age = 1;

//you should not use var -- its old!
// var = 1


//arrow function syntax

//this is a valid function
// let add =  () => {};

//But then if you try to add a function with a different function signature, it won't work:
// add = (a: number, b: number) => {
//     return a+b;
// };

let add = (a: number, b: number): number => {
    return a+b;
};

//OOOooo if you just have one expression, you can remove the curly braces and write it extra short, like this (the return statement is implicit):
let add2 = (a: number, b: number): number => a+b;

const button = document.querySelector('button');

if (button) {
    // button.addEventListener('click', () => {})
    // I however remain confused about the syntax in this example
    button.addEventListener('click', event => {console.log(event)})
}

//Adds a default value to b in case it is not passed
const add3 = (a: number, b: number = 1) => a+b;

console.log(add(3,3))
console.log(add3(5))