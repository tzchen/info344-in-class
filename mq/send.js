#!/usr/bin/env node
// ^ shebang lets you run the node file like a script
// the shebang tells the system what to use to interpret
// the file, so theoretically could name it just "send"
// > chmod +x send.js
// > ./send.js
"use strict";

const amqp = require("amqplib");
const qName = "testQ";
const mqAddr = process.env.MQADDR || "localhost:5672";
const mqURL = `amqp://${mqAddr}`;

(async function() {
    console.log("connecting to %s", mqURL);
    let connection = await amqp.connect(mqURL);
    let channel = await connection.createChannel();
    let qConf = await channel.assertQueue(qName, { durable: false });

    setInterval(() => {
        let msg = "hi there: " + new Date().toLocaleTimeString();
        channel.sendToQueue(qName, Buffer.from(msg));
        console.log("sending message");
    }, 1000);
})();
