// Bare block statement
{
    console.log("Hello")
    console.log("World")
}

let x = 0;
while (x < 10) {
    x++;
}

let y = 1;
{
    let y = 2;
}
console.log(y) //prints 1

var yy = 1;
{
    var yy = 2;
}
console.log(yy) //prints 2, I guess this is why authors tell me to avoid using var, makes sense

