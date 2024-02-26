function add(n1: number, n2: number): number {
    return n1 + n2;
}

function printResult(num: number){
    console.log('Result: ' + num)
}
printResult(add(2,3))

function addAndHandle(n1: number, n2: number, cb: (num: number) => void) {
    const result = n1 + n2
    cb(result)
}

let combineValues: (a: number, b: number) => number;
combineValues = add;
console.log(combineValues(9,9))

//With an anonymous function!
addAndHandle(10, 22,  (result) => {
    console.log(result)
})

//From the quiz, this is bamzoozling me:
function sendRequest(data: string, cb: (response: any) => void) {
    // ... sending a request with "data"
    return cb({data: 'Hi there!'});
}

function cbFun (message: string) {
    console.log("You called cbFun with message:", message)
}

sendRequest('Send this!', (response) => {
    console.log(response);
    return "aaa";
});