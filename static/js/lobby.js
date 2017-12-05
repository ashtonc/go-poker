$(document).ready(function(){
	$('.menu').on('click', '#createTable', function(){
		var name = $('#name').val();
		var players = $('#players').find(":selected").text();
		var stakes = $('#stakes').find(":selected").text();
		$('.tables').append("<tr><td>" + name + "</td><td>" + stakes + "</td><td>1/" + players + "</td><td><a href = \"game.html\"><button type=\"button\">Join Table</button></a></td></tr>");
	});
});


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

function onMessage(message) {
  switch (message.kind) {
    case MESSAGE_WATCHING:
      break;
    case MESSAGE_WATCHER_JOINS:
      break;
    case MESSAGE_LEFT_TABLE:
      break;

  }
}