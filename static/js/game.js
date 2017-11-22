$(document).ready(function(){
	// initialize first turn (random)
	$('.position1').find('img').addClass('yourTurn');
	$('.position1').find('img').after('<div class=\"timer\"><div class=\"startTimer\"></div></div>');
	$('.startTimer').on('animationend', nextTurn);
	function nextTurn(){
		// if it's your turn, draw blue outline and animate it
		if($('.position1').find('img').hasClass('yourTurn')){
			$('.position1').find('img').removeClass('yourTurn');
			$('.timer').remove();
			$('.startTimer').remove();
			$('.position2').find('img').addClass('yourTurn');
			$('.position2').find('img').after('<div class=\"timer\"><div class=\"startTimer\"></div></div>');
			$('.startTimer').on('animationend', nextTurn);
		}
		else if($('.position2').find('img').hasClass('yourTurn')){
			$('.position2').find('img').removeClass('yourTurn');
			$('.timer').remove();
			$('.startTimer').remove();
			$('.position3').find('img').addClass('yourTurn');
			$('.position3').find('img').after('<div class=\"timer\"><div class=\"startTimer\"></div></div>');
			$('.startTimer').on('animationend', nextTurn);
		}
		else if($('.position3').find('img').hasClass('yourTurn')){
			$('.position3').find('img').removeClass('yourTurn');
			$('.timer').remove();
			$('.startTimer').remove();
			$('.position4').find('img').addClass('yourTurn');
			$('.position4').find('img').after('<div class=\"timer\"><div class=\"startTimer\"></div></div>');
			$('.startTimer').on('animationend', nextTurn);
		}
		else if($('.position4').find('img').hasClass('yourTurn')){
			$('.position4').find('img').removeClass('yourTurn');
			$('.timer').remove();
			$('.startTimer').remove();
			$('.position5').find('img').addClass('yourTurn');
			$('.position5').find('img').after('<div class=\"timer\"><div class=\"startTimer\"></div></div>');
			$('.startTimer').on('animationend', nextTurn);
		}
		else if($('.position5').find('img').hasClass('yourTurn')){
			$('.position5').find('img').removeClass('yourTurn');
			$('.timer').remove();
			$('.startTimer').remove();
			$('.position6').find('img').addClass('yourTurn');
			$('.position6').find('img').after('<div class=\"timer\"><div class=\"startTimer\"></div></div>');
			$('.startTimer').on('animationend', nextTurn);
		}
		else{
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
		$('.yourTurn').closest('img').addClass('folded');
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
