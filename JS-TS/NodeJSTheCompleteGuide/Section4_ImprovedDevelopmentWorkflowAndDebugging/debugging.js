const fs = require('fs');

var myname = "Joe McJoeson"

var sometext = `here is
some multiline
text!
Also, my name is ${myname}`

console.log(sometext)

let x = 0;
while (x<10) {
    x++
}

const filePath = "/tmp/myfile.txt"

fs.writeFile(filePath, sometext, (err) => {
    if (err) {
        console.error('Error writing to file:', err);
        return;
    }
    console.log('Data has been written to', filePath);
});

let aaa = "aaa";
console.log("Program finished!") //Huh, interesting, prints before things are written to file