// Bare bones syntax:
const examplePromise = new Promise (() => {});

// Promise: Takes a function as argument with 2 parameters: resolve and reject
const myPromise = new Promise ((resolve, reject) => {
    setTimeout(() => {
        const randomNumber = Math.random();
        if (randomNumber > 0.5) {
            resolve(randomNumber);
        } else {
            reject(new Error('Random number is too small'));
        }
    }, 1000)
});

myPromise.then((result) => {
    console.log("Worked:", result)
}).catch((e) => {
    console.log("error:", e)
})

//Lets look at this other example that calls a placeholder API:

const apiUrl = 'https://jsonplaceholder.typicode.com/posts/1';

// fetchData is a function with no arguments that returns a promise
const fetchData = () => {
    //in the return statement, a new promise is created
    return new Promise((resolve, reject) => {
        //Huh, there's a fetch function? Lets investigate this and brb
        fetch(apiUrl)
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                return response.json();
            })
            .then(data => resolve(data))
            .catch(error => reject(error));
    });
};

fetchData()
    .then(data => console.log('Data fetched successfully:', data))
    .catch(error => console.error('Error fetching data:', error));
