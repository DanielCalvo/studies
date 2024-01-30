function add(n1: number, n2: number): number {
    return n1 + n2;
}

function printResult(num: number){
    console.log('Result: ' + num)
}
printResult(add(2,3))

let combineValues: (a: number, b: number) => number;
combineValues = add;
console.log(combineValues(9,9))