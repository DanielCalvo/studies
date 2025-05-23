//How can I iterate over an array while removing one element at a time?
numArray = [1,2,3,4,5,6]

while (numArray.length > 0){
    console.log(numArray.shift())
}

console.log("Array:",numArray)