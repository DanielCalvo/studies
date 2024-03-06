const http = require('http');

console.log("Hello world!")

const server = http.createServer((req, res) => {
    console.log(req.url, req.method, req.headers);
    res.write("Hello from section 4!!!");
    res.end();
})

server.listen(3000)