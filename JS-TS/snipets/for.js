// Lets start at the basic: How do you loop over an array again

strArray = ["banana", "apple", "orange", "melon"]

// Interesting, an array has a forEach method
strArray.forEach(function (element) {
    process.stdout.write(element + " ")
})
console.log()

//There's also this other method. I like this one.
for (let element of strArray){
    process.stdout.write(element + " ")
}
console.log()

//You can also use the map function of the array object, which in this case is identical to the forEach one:
strArray.map(function(element) {
    process.stdout.write(element + " ")
});

console.log()

