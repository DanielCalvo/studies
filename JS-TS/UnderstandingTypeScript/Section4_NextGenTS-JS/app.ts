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
//TS will still adhere to the argument order though, it won't "guess" that you're trying to pass a parameter to the argument with the non-defined argument
const add3 = (a: number, b: number = 1) => a+b;

console.log(add(3,3))
console.log(add3(5))

// Arrays are objects and objects are reference values
const hobbies = ['sports', 'cooking']
const activeHobbies = ['hiking']

//Works:
// activeHobbies.push(hobbies[0], hobbies[1])

//But there's a better way of doing it
activeHobbies.push(...hobbies)

const person = {
    name: 'joe',
    age: 0
};

// You're not creating a copy of this object, you're just copying the pointer
const copiedPerson = person;

// This however does copy the kv pairs from person into copiedPerson2
const copiedPerson2 = {... person};

//Rest parameters: I want to be able to handle as many values as the user passes in
// This will merge the arguments into an array
const add4 = (...numbers: number[]) => {
    // Very little idea of what's going on here, I'll need to investigate this later
    numbers.reduce((curResult, curValue) => {
        return curResult + curValue
    }, 0);
};

const addedNumbers = add4(1,5,3,5)
console.log(addedNumbers)

// Destructuring an array
const [hobby1, hobby2, ...remainingHobbies] = hobbies;
// const {name, age} = person; //neat
//Or you can also do this by:
const {name: nameP, age: ageP} = person;
console.log(nameP, ageP)