//but dont use alert as it blocks everything, you want to use the console!
//very obvious to say, but this shows up in the console.
console.log("Hello World!");

//an error, neat!
console.error("This is an error");

console.warn("you can also warn!");

//setting variables:
//you can use var, let and const. You dont want to us var anymore as its globally scoped
//let and const are scoped
const age = 31
//fails:
//age = 30

//widely adopted practice: always use const unless you know youre going to reasing the value