function resolveAfterTime (time) { //oof horrendous naming, but lets continue
    return new Promise((resolve) => {
        setTimeout(() => { //set timeout will exeute a function after a certain timer expired
            // Oooooooooooo so you could do this without set timeout!
            resolve('resolved')
        }, time)
    })
}

//interesting, the program does not block and continues, so we need to use something else to block here
//or we need another function to run to wait on it

// console.log('started!');
// resolveAfterTime(1000)
// console.log("ended!")

// so lets introduce the asyncCall function:
async function asyncCall() {
    console.log("Calling in asyncCall!")
    const result = await resolveAfterTime(1000) //so looks like "await" just means "block" in golang terminology
    console.log(result)
}

asyncCall()