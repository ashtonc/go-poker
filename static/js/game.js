$(document).ready(function(){
	// initialize first turn (random)
	var numberOfPlayers = 6;
	addPlayers();
	deal();
	$('.position1').find('img').addClass('yourTurn');
	$('.position1').find('img').after('<div class=\"timer\"><div class=\"startTimer\"></div></div>');
	$('.startTimer').on('animationend', nextTurn);
	function addPlayers(){
		if (numberOfPlayers == 2){
			$('.bottomRow').append('<div class=\"position2\"><img src=\"profile.jpg\" title =\"picture\"><div class=\"name\">Player 2</div><div class=\"stack\">1000000</div><ul class=\"table\"></ul></div>');
		}
		if (numberOfPlayers == 3 || numberOfPlayers == 4){
			$('.topRow').append('<div class=\"position2\"><img src=\"profile.jpg\" title =\"picture\"><div class=\"name\">Player 2</div><div class=\"stack\">1000000</div><ul class=\"table\"></ul></div>');
			if (numberOfPlayers == 4){
				$('.bottomRow').append('<div class=\"position4\"><img src=\"profile.jpg\" title =\"picture\"><div class=\"name\">Player 4</div><div class=\"stack\">1000000</div><ul class=\"table\"></ul></div>');
			}
			$('.bottomRow').append('<div class=\"position3\"><img src=\"profile.jpg\" title =\"picture\"><div class=\"name\">Player 3</div><div class=\"stack\">1000000</div><ul class=\"table\"></ul></div>');
		}
		if (numberOfPlayers == 5 || numberOfPlayers == 6){
			$('.topRow').append('<div class=\"position2\"><img src=\"profile.jpg\" title =\"picture\"><div class=\"name\">Player 2</div><div class=\"stack\">1000000</div><ul class=\"table\"></ul></div>');
			$('.topRow').append('<div class=\"position3\"><img src=\"profile.jpg\" title =\"picture\"><div class=\"name\">Player 3</div><div class=\"stack\">1000000</div><ul class=\"table\"></ul></div>');
			if (numberOfPlayers == 6){
				$('.bottomRow').append('<div class=\"position6\"><img src=\"profile.jpg\" title =\"picture\"><div class=\"name\">Player 6</div><div class=\"stack\">1000000</div><ul class=\"table\"></ul></div>');
			}
				$('.bottomRow').append('<div class=\"position5\"><img src=\"profile.jpg\" title =\"picture\"><div class=\"name\">Player 5</div><div class=\"stack\">1000000</div><ul class=\"table\"></ul></div>');
				$('.bottomRow').append('<div class=\"position4\"><img src=\"profile.jpg\" title =\"picture\"><div class=\"name\">Player 4</div><div class=\"stack\">1000000</div><ul class=\"table\"></ul></div>');
		}
	}
	function deal(){
		for(i = 0; i < numberOfPlayers * 5; i++){
			currentPlayer = '.position' + ((i%numberOfPlayers) + 1);
			var rank = (i%numberOfPlayers + 2);
			var suit = 'spades';
			$(currentPlayer).find('.table').append('<li><div class=\"card rank-' + rank + ' ' + suit + '\"><span class=\"rank\">' + rank + '</span><span class=\"suit\">&' + suit + ';</span></div></li>');
		}
	}
	function nextTurn(){
		// if it's your turn, draw blue outline and animate it
		if($('.position1').find('img').hasClass('yourTurn')){
			$('.yourTurn').parent().addClass('folded');
			$('.position1').find('img').removeClass('yourTurn');
			$('.timer').remove();
			$('.startTimer').remove();
			$('.position2').find('img').addClass('yourTurn');
			$('.position2').find('img').after('<div class=\"timer\"><div class=\"startTimer\"></div></div>');
			$('.startTimer').on('animationend', nextTurn);
		}
		else if($('.position2').find('img').hasClass('yourTurn')){
			$('.yourTurn').parent().addClass('folded');
			$('.position2').find('img').removeClass('yourTurn');
			$('.timer').remove();
			$('.startTimer').remove();
			if(numberOfPlayers >= 3){
				$('.position3').find('img').addClass('yourTurn');
				$('.position3').find('img').after('<div class=\"timer\"><div class=\"startTimer\"></div></div>');
				$('.startTimer').on('animationend', nextTurn);
			}
			else{
				$('.position1').find('img').addClass('yourTurn');
				$('.position1').find('img').after('<div class=\"timer\"><div class=\"startTimer\"></div></div>');
				$('.startTimer').on('animationend', nextTurn);
			}
		}
		else if($('.position3').find('img').hasClass('yourTurn')){
			$('.yourTurn').parent().addClass('folded');
			$('.position3').find('img').removeClass('yourTurn');
			$('.timer').remove();
			$('.startTimer').remove();
			if(numberOfPlayers >= 4){
				$('.position4').find('img').addClass('yourTurn');
				$('.position4').find('img').after('<div class=\"timer\"><div class=\"startTimer\"></div></div>');
				$('.startTimer').on('animationend', nextTurn);
			}
			else{
				$('.position1').find('img').addClass('yourTurn');
				$('.position1').find('img').after('<div class=\"timer\"><div class=\"startTimer\"></div></div>');
				$('.startTimer').on('animationend', nextTurn);
			}
		}
		else if($('.position4').find('img').hasClass('yourTurn')){
			$('.yourTurn').parent().addClass('folded');
			$('.position4').find('img').removeClass('yourTurn');
			$('.timer').remove();
			$('.startTimer').remove();
			if(numberOfPlayers >= 5){
				$('.position5').find('img').addClass('yourTurn');
				$('.position5').find('img').after('<div class=\"timer\"><div class=\"startTimer\"></div></div>');
				$('.startTimer').on('animationend', nextTurn);
			}
			else{
				$('.position1').find('img').addClass('yourTurn');
				$('.position1').find('img').after('<div class=\"timer\"><div class=\"startTimer\"></div></div>');
				$('.startTimer').on('animationend', nextTurn);
			}
		}
		else if($('.position5').find('img').hasClass('yourTurn')){
			$('.yourTurn').parent().addClass('folded');
			$('.position5').find('img').removeClass('yourTurn');
			$('.timer').remove();
			$('.startTimer').remove();
			if(numberOfPlayers == 6){
				$('.position6').find('img').addClass('yourTurn');
				$('.position6').find('img').after('<div class=\"timer\"><div class=\"startTimer\"></div></div>');
				$('.startTimer').on('animationend', nextTurn);
			}
			else{
				$('.position1').find('img').addClass('yourTurn');
				$('.position1').find('img').after('<div class=\"timer\"><div class=\"startTimer\"></div></div>');
				$('.startTimer').on('animationend', nextTurn);
			}
		}
		else{
			$('.yourTurn').parent().addClass('folded');
			$('.position6').find('img').removeClass('yourTurn');
			$('.timer').remove();
			$('.startTimer').remove();
			$('.position1').find('img').addClass('yourTurn');
			$('.position1').find('img').after('<div class=\"timer\"><div class=\"startTimer\"></div></div>');
			$('.startTimer').on('animationend', nextTurn);
		}
	}
	$('.menu').on('click', '#check', function(){
		$('.startTimer').stop(true, true);
		nextTurn();
	});
	$('.menu').on('click', '#call', function(){
		$('.startTimer').stop(true, true);
		nextTurn();
	});
	$('.menu').on('click', '#fold', function(){
		$('.startTimer').stop(true, true);
		$('.yourTurn').parent().addClass('folded');
		nextTurn();
	});
	$('.menu').on('click', '#raiseButton', function(){
		$('.startTimer').stop(true, true);
		var bet = $('#raise').val();
		var stack = $('.yourTurn').parent().find('.stack').text();
		var amount = stack - bet;
		$('.yourTurn').parent().find('.stack').text(amount);
		nextTurn();
	});
	$('form').submit(function(e){
		e.preventDefault();
	});
});
