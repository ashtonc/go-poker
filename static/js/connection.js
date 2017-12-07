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

		wsPath += "/ws"

		console.log("Opening websocket connection at " + wsPath)
		var ws = new WebSocket(wsPath);

		ws.onopen = function () {
			alert("WS opened.");
		};

		ws.onmessage = function (event) {
			alert(received_msg);
			var messages = event.data.split('\n');
			for (var i = 0; i < messages.length; i++) {
				var message = JSON.parse(messages[i]);
				onMessage(message);
			}
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

KIND_GAME_STATE = 1;
KIND_PLAYER_SITS = 2;
KIND_PLAYER_LEAVES = 3;
KIND_TIMED_OUT = 4;

KIND_TAKE_SEAT = 5;
KIND_LEAVE_SEAT = 6
KIND_CHECK = 7;
KIND_BET = 8;
KIND_CALL = 9;
KIND_FOLD = 10;
KIND_DISCARD = 11;

function onMessage(message) {
	switch (message.kind) {
		case KIND_GAME_STATE:

			break;
		case KIND_PLAYER_SITS:

			break;
		case KIND_PLAYER_LEAVES:

			break;
		case KIND_TIMED_OUT:

			break;
	}
}