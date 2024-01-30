console.log('Your code goes hereeee!')

//if you don't pass a type here, typescript will interpret this as you not caring about type
function add(n1: number, n2: number, showResult: boolean, phrase: string) {
    const result1 = n1 + n2 //you need this here otherwise typescript will try to concatenate 2 numbers and a string as a string, oof
    if (showResult){
        console.log(phrase + result1)
    }
    return n1 + n2
}

const number1 = 5
const number2 = 2.8
const printResult = true

const result = add(number1, number2, printResult, "Bazinga")
