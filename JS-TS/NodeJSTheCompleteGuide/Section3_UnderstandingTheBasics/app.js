const http = require('http');
const fs = require('fs')
//A path to your files would have to start with `./`

function myListener(req, res){

}
// http.createServer(myListener);
//or
// http.createServer(function(req, res){})
//or

const server = http.createServer((req, res) => {
    console.log(req.url, req.method, req.headers)
    if (req.url === '/'){
        res.write('<html>');
        res.write('<head><title>Enter Message</title><head>');
        res.write('<body><form action="/message" method="POST"><input type="text" name="message"><button type="submit">Send</button></form></body>');
        res.write('</html>');
    }
    if (req.url === '/message' && req.method === 'POST'){ //it works if I change the method to GET, but POST not really :{
        //An event listener! The data event will be fired whenever a new chunk is ready to be read
        // We also specify a function to be executed for every data event
        const body = [];
        req.on('data', (chunk) => {
            console.log(chunk)
            body.push(chunk);
        });
        req.on('end', () => { //runs on the end event, not immediately, interesting
            const parsedBody = Buffer.concat(body).toString();
            const message  = parsedBody.split('')[1];
            // fs.writeFileSync('message.txt', message); //synchronous mode, its better to use it asynchronous
            fs.writeFile('message.txt', message, (err) => {
                res.statusCode = 302
                res.setHeader('Location', '/')
                return res.end
            });
        })
    }

    res.end();
});

server.listen(3000)
