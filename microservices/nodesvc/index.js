// @ts-check

"use strict";

const express = require("express");
const morgan = require("morgan");

const addr = process.env.ADDR || ":80";
const [host, port] = addr.split(":");
const portNum = parseInt(port);

const app = express();
// sets morgan to use either environment
// variable LOG_FORMAT, or if undefined, 
// uses default value "dev"
app.use(morgan(process.env.LOG_FORMAT || "dev"));
app.use(express);

app.get("/", (req, res) => {
    res.set("Content-Type", "text/plain");
    res.send("Hello, Node.js!");
});

app.get("/v1/users/me/hello", (req, res) => {
    let userJSON = req.get("X-User");
    let user = JSON.parse(userJSON);
    res.json({
        message: `Hello, ${user.firstName} ${user.lastName}`
    });
});

const handlers = require("./handlers.js")
app.use(handlers({}))

app.use((err, req, res, next) => {
    console.error(err.stack);
    res.set("Content-Type", "text/plain");
    res.send(err.message);
})

app.listen(portNum, host, () => {
    console.log(`server is listening at http://${addr}...`)
});

