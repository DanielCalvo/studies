const http = require('http');

const express = require('express');
const app = express(); //Neat

//This function that you pass to app.use() will run at every single request

// app.use((req, res, next) => {
//     console.log("In a middleware!"); //I don't see this :{
//     next();
// });

app.use('/', (req, res, next) => {
    console.log("this always runs");
    next()
});


app.use('/add-product', (req, res, next) => {
    console.log("In add product middleware!"); //I don't see this :{
    res.send('<h1>aDD PROduct page</h1>')
});

app.use('/', (req, res, next) => {
    console.log("In / middleware!"); //I don't see this :{
    res.send('<h1>Hello from express!</h1>')
});

app.listen(3000)