let userInput: unknown

//All of this is allowed with the unkwnon type

let userName: string

userInput = 2
userInput = "asd"

//Javascript won't let you do this, but if userInput was "any" then it would!
// unknown is a bit more restrictive than any!
//userName = userInput

// this is ok!
if (typeof userInput === "string") {
    userName = userInput
}

function generateError (message: string, code: number) {
    throw {message: message, errorCode: code};
}

generateError('something went wrong', 555)