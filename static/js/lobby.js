$(document).ready(function(){
	$('.menu').on('click', '#createTable', function(){
		var name = $('#name').val();
		var players = $('#players').find(":selected").text();
		var stakes = $('#stakes').find(":selected").text();
		$('.tables').append("<tr><td>" + name + "</td><td>" + stakes + "</td><td>1/" + players + "</td><td><a href = \"game.html\"><button type=\"button\">Join Table</button></a></td></tr>");
	});
});
