// Function that takes a callback
function fetchData(callback) {
    // Simulate an asynchronous operation (e.g., fetching data from a server)
    setTimeout(() => {
        const data = { id: 1, name: 'John Doe', age: 30 };
        // Pass the fetched data to the callback function
        callback(data);
    }, 1000); // Simulate a delay of 1 second
}

// Callback function
function processData(data) {
    console.log('Received data:', data);
}

// Calling the function with the callback
fetchData(processData);
