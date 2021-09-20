import Vue from "vue"

let socket = undefined;

const emitter = new Vue({
    methods: {
        send(message) {
            if (1 === socket.readyState)
                socket.send(message)
        },
        connect() {
            socket = new WebSocket('ws://' + document.location.host + '/ws');

            socket.onmessage = function (msg) {
                let payload = JSON.parse(msg.data)
                console.debug(`Received ${payload.Class} message: ${payload.Message}`)
                emitter.$emit(payload.Class, payload.Message)
            }

            socket.onclose = function () {
                console.info("Websocket closed")
                emitter.$emit("LocalEvent", { name: "WS.OnClose" })
                setTimeout(emitter.connect, 3000);
            }

            socket.onopen = function () {
                console.info("Websocket connected")
                emitter.$emit("LocalEvent", { name: "WS.OnOpen" })
            }

            socket.onerror = function (err) {
                console.warn(`Websocket connection error:`);
                console.warn(err);
                emitter.$emit("LocalEvent", { name: "WS.OnError" });
            }
        }
    }
})

emitter.connect()

export default emitter