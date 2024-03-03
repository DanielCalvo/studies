//use the throw statement to throw an exception
// throw 42

//Example from the docs, interesting:
function getPerson(person) {
    const people = ["Joe", "Max", "Linus", "Bob"]
    if (people[person]) {
        return people[person]
    } else {
        throw new Error("Could not find person")
    }
}

try {
    person = getPerson("Joe")
} catch (e) {
    person = "unknown"
    console.log("Could not get person")
}

console.log(person)

//Lets try some try-catch with the following: I'll try to write to /etc/passwd as my regular user, which will fail
//when it fails, lets write it to a file in /tmp/ then

const fs = require('fs')

//Huh, doesn't seem to be getting into the "catch" part here, there's obviously something I'm doing wrong
try {
    fs.writeFile("/etc/passwd", "bazinga", (err) => {
        if (err){
            throw err;
        }
        console.log("Data has been written to file")
    });
} catch (error)  {
    console.log("Error writing to file:", error)
}

