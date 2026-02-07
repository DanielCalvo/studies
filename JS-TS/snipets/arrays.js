const numbers = [1,2,3,4,5]

//the classic way of iterating over an array:
for (let i = 0; i < numbers.length; i++) {
    console.log(numbers[i])
}

//the modern way of doing it, python-like!
//since the variable is redefined in a new scope everytime the loop iterates, const is valid!
// also it means you cant change that variable as you iterate!
for (const number of numbers) {
    console.log(number)
}

//arrays also have a built in forEach method!
//the for each method executes a function for each element in the array
numbers.forEach(function (number) {
    console.log(number)
})

//ah, the foreach method receives 3 arguments:
//geremias explains:
// When forEach() executes, for every element in the array, it automatically passes three pieces of data to your callback function, in this specific order:
// Argument 1 (The Element): The value of the item (e.g., 'red').
// Argument 2 (The Index): The numerical position of the item (e.g., 0, then 1, then 2).
// Argument 3 (The Array): The whole array itself.

//so you can print this:
numbers.forEach(function (number, index, array) {
    console.log(number, index, array)
})

//lets have a look at the map method
//the map function takes a callback function and applies it to every element in the array
//in other words, it transforms every element in the array
const doubleNumbers = numbers.map (function (number){
    return number * 2
});

//ah so it returns an array! .
console.log(doubleNumbers)

//lets turn every number into a string!
const stringNumbers = numbers.map (function (number){
    return `${number}`
});
console.log(stringNumbers)

//the filter element creates a new array containing the elements that pass a test
const evenNumbers = numbers.filter(function (number){
    return number % 2 === 0
})

console.log(evenNumbers)

//oh you can make it even shorter by using an arrow function
const moreEvenNumbers = numbers.filter((number => number % 2 === 0))
console.log(moreEvenNumbers)

//the reduce method executes a reducer function on each element of the array, resulting in a single output value
const sum = numbers.reduce((acc, number) => {
    return acc + number
})

console.log(sum)