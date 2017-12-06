$(document).ready(function () {
	WebSocketTest();
});

function WebSocketTest() {
	if ("WebSocket" in window) {
		alert("WebSocket is supported by your Browser!");

		var wsPath = "ws://" + location.hostname + ":" + location.port;
		var pathArray = location.pathname.split('/');

		for (i = 0; i < pathArray.length - 1; i++) {
			if (i != 0) {
				wsPath += "/";
			}
			wsPath += pathArray[i];
		}

		var ws = new WebSocket(wsPath + "/ws");

		// don't hardcode this: WebSocket('ws://' + location.hostname + location.port + '/poker/ws'); or something

		ws.onopen = function () {
			alert("WS opened.");
		};

		ws.onmessage = function (evt) {
			var received_msg = evt.data;
			alert(received_msg);
			ws.send("Message received")
		};

		ws.onclose = function () {
			// websocket is closed.
			alert("Connection closed.");
		};

	} else {
		// The browser doesn't support WebSocket
		alert("WebSocket NOT supported by your Browser!");
	}
}

KindGameState = 1;

KindPlayerSits = 2;
KindPlayerLeaves = 3;
KindTimedOut = 4;

KindTakeSeat = 5;
KindLeaveSeat = 6;

KindCheck = 7;
KindBet = 8;
KindCall = 9;
KindFold = 10;
KindDiscard = 11;

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