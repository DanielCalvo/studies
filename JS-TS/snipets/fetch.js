//Huh, there's a fetch API: https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API/Using_Fetch

const apiUrl = 'https://jsonplaceholder.typicode.com/posts/1';

const thing = await fetch(apiUrl)

console.log(thing)