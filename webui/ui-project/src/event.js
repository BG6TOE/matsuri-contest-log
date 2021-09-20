import Vue from "vue"

const socket = new WebSocket('ws://' + document.location.host + '/ws');

const emitter = new Vue({
    methods: {
        send(message) {
            if (1 === socket.readyState)
                socket.send(message)
        }
    }
})

socket.onmessage = function (msg) {
    let payload = JSON.parse(msg.data)
    console.debug(`Received ${payload.Class} message: ${payload.Message}`)
    emitter.$emit(payload.Class, payload.Message)
}

socket.onerror = function (err) {
    console.warn(`Websocket connection error:`)
    console.warn(err)
}

export default emitter