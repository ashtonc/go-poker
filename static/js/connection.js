function WebSocketTest() {
	if ("WebSocket" in window) {
		alert("WebSocket is supported by your Browser!");
		// Let us open a web socket
		var ws = new WebSocket("ws://localhost:8000/poker/game/limit-10/ws");
		// don't hardcode this
		// WebSocket('ws://' + location.hostname + ':8000');

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