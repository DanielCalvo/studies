var name = 'asd';
console.log(name);

var age = 22
console.log(age);
age = "111"
console.log(age); //Anarchy. Complete anarchy.

function summarizeuser(userName){
    return ("User is: " + userName)
}

console.log(summarizeuser(name))

//var is a bit outdated -- let is preferred
let name1 = "joe"
const name2 = "joejoe" //Cannot be changed. Use const as frequently as possible

//Alritey lets do an anonymous function, aka "arrow functions"
const sum = (num1, num2) => {
    return num1 + num2
}

// Functions!
// You can even make the function shorter:
const sum1 = (a, b) => a + b;
const add1 = a => a + 1; //Also valid if you have a single argument, interesting
const addWhatever = () => 1+2; //If you have no arguments, you need to have the ()

console.log(sum(2,2));
console.log(sum(2,4));
console.log(sum1(22,4));

//Objects, properties and methods!
const person = {
    name: "Joe",
    age: 0,
    happy: true,
    greet: () => {
        console.log("Hi I am " + this.name) //This refers to the surrounding object. Returns undefined however
    },
    greet1: function() {
        console.log("Hi I am " + this.name) //This refers to the surrounding object. This works
    },
    greet2() {
        console.log("Hi I am " + this.name) //This refers to the surrounding object. This works. Author likes this one.
    },
};

console.log(person);
person.greet(); //undefined, interesting. This refers to the global JS scope and not the object
person.greet1(); //This works though
person.greet2(); //This also works

//Arrays and array methods!
const hobbies = ["Studying", "Reading", "Walking"];
const anarchy = ["asd", 0, true, {}];

for (let hobby of hobbies) {
    console.log(hobby)
}

//Array methods!
// Maps executes a function in each element of the array and returns array
console.log(hobbies.map(hobby => {
    return hobby.toUpperCase();
}))

//you could also do it like this:
console.log(hobbies.map(hobby => hobby.toUpperCase()));

//Arrays, objects: They are reference types
hobbies.push('Cooking');
// Reference types only store an address pointing to a place in memory where that array is stored. That pointed has not changed by us adding an element
console.log(hobbies);

//Spread and rest operators
const copiedArray = hobbies.slice();
console.log(copiedArray);

// doesn't work, nested arrays
const copiedArray2 = [hobbies];
console.log(copiedArray2);

const copiedArray3 = [...hobbies]; //Spread operator -- pulls out all the elements of an array or all the properties of an object, and puts it in whatever is around that operator
console.log(copiedArray3);

const copiedPerson = {...person}
console.log(copiedPerson);

//rest operator
//this will bundle all the function parameters in an array for you
const toArray = (...args) => {
    return args;
}
console.log(toArray(1,2,3,4));
//Both operators look the same -- its the place where you use it that defines how you call it
//To pull elements or properties out of arrays or objects: Then its the spread operator
//Are you using it to merge multiple arguments into an array, and you use it in the argument list of a function? then its the rest operator

// Destructuring
// Specify inside the function argument the object property you're interested in, in this case, name:
// Other object properties will be dropped
const printName = ({name}) => {
    console.log(name);
}
printName(person);

//You can also use destructuring outside of a function
// But oh-oh: Looks like this... needs `name` and `age` to be global variables. These are already defined so I'll comment this out
// const {name, age} = person;
// console.log(name5, age5);

const [hobby1, hobby2] = hobbies //These are retrieved by position it looks like
console.log(hobby1, hobby2)

// Async code and promises!
//Lets define a function that should execute after a certain timer expired
// 2000 means 2000ms. `setTimeout` is part of the nodejs
// However the rest of the program will continue executing!
const fetchData = callback => {
    setTimeout(() => {
        callback('Done!');
        }, 1500)
};

setTimeout(() => {
    console.log('Timer is done!');
    fetchData(text => {
        console.log(text);
    });
}, 2000)

console.log('Hello')
