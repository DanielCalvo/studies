const fs = require('fs');

console.log('Hello world!')
fs.writeFileSync('/tmp/hello.txt', "Writing to hello file!")