$(document).ready(function(){
	WebSocketTest();
});

function WebSocketTest() {
	if ("WebSocket" in window) {
		alert("WebSocket is supported by your Browser!");
		// Let us open a web socket
		var ws = new WebSocket("ws://localhost:8000/poker/ws");
		// don't hardcode this: WebSocket('ws://' + location.hostname + ':8000' + '/poker/ws'); or something

		ws.onopen = function () {
			alert("WS is opened...");
		};

		ws.onmessage = function (evt) {
			var received_msg = evt.data;
			// uncomment next line if you want to get alerts on each message
			//alert("Message is received..." + received_msg);
			//document.getElementById("temp").innerHTML = "<pre>" + received_msg + "</pre>"
			alert(received_msg);
			ws.send("Message received")
		};

		ws.onclose = function () {
			// websocket is closed.
			alert("Connection is closed...");
		};

	} else {
		// The browser doesn't support WebSocket
		alert("WebSocket NOT supported by your Browser!");
	}
}

MESSAGE_WATCHING = 1;
MESSAGE_WATCHER_JOINS = 2;
MESSAGE_LEFT_TABLE = 3;


function onMessage(message) {
switch (message.kind) {
  case MESSAGE_WATCHING:
    for (var i = 0; i < message.users.length; i++) {
      var user = message.users[i];
      otherWatchers[user.id] = user.name;
    }
    break;
  case MESSAGE_WATCHER_JOINS:
    otherWatchers[message.user.id] = message.user.name;
    break;
  case MESSAGE_LEFT_TABLE:
    delete otherNames[message.userId];
    update();
    break;
}
}

socket.onmessage = function (event) {
  var messages = event.data.split('\n');
  for (var i = 0; i < messages.length; i++) {
    var message = JSON.parse(messages[i]);
    onMessage(message);
  }
};