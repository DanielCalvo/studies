// Lets make this add two numbers or concatenate two strings
// So you can have different logic in your function based on what type you're getting, interesting

//create type alias
type Combinable = number | string;
type ConversionDescriptor = 'as-number' | 'as-text';

function combine(
    // input1: number | string,
    // input2: number | string,
    input1: Combinable,
    input2: Combinable,
    // resultConversion: 'as-number' | 'as-text') {
    resultConversion: ConversionDescriptor) {
    let result;
    if (typeof input1 === "number" && typeof input2 === "number" || resultConversion === 'as-number') {
        result = +input1 + +input2;
    }
    if (typeof input1 === "string" && typeof input2 === "string" || resultConversion === 'as-text') {
        result = input1.toString() + input2.toString();
    }
    return result;
}

const combinedAges = combine(33,33, 'as-number');
console.log(combinedAges)

const combinedNames = combine("Eeeo", "EEEEEEOOO", 'as-text');
console.log(combinedNames)

//oh sweet lordy
const doingItWrong = combine(22, 22, "as-text")
console.log(doingItWrong)