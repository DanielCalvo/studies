try {
    console.log("Just doing my thing!")
    throw new Error("Oh no, something went wrong!")
    console.log("We never get here")
} catch (e) {
    console.log("Caught error:")
} finally {
    console.log("Did the thing!")
}

console.log("Helloo")
