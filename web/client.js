var ws;

function connect() {
    ws = new WebSocket("ws://" + window.location.host + "/ws");

    ws.onopen = function () {
        console.log("Connected to server");
        addMessage("System: Connected to server");
    };

    ws.onmessage = function (event) {
        console.log("Message received:", event.data);
        addMessage("Server: " + event.data);
    };

    ws.onclose = function () {
        console.log("Disconnected from server");
        addMessage("System: Disconnected from server");
        // 尝试5秒后重连
        setTimeout(connect, 5000);
    };

    ws.onerror = function (err) {
        console.log("WebSocket error:", err);
    };
}

function addMessage(msg) {
    var messages = document.getElementById("messages");
    messages.innerHTML += "<p>" + msg + "</p>";
    messages.scrollTop = messages.scrollHeight;
}

function sendMessage() {
    var input = document.getElementById("messageInput");
    var message = input.value;
    if (message && ws.readyState === WebSocket.OPEN) {
        ws.send(message);
        addMessage("You: " + message);
        input.value = "";
    }
}

// 页面加载时连接
window.onload = connect;