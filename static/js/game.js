$(document).ready(function(){
	// initialize first turn (random)
	$('.position1').addClass('.yourTurn');
	$('.position1').find('img').css({'outline': '5px solid blue'});
	$('.position1').find('img').after('<div class=\"timer\"><div class=\"startTimer\"></div></div>');
	$('.startTimer').on('animationend', nextTurn);
	function nextTurn(){
		// if it's your turn, draw blue outline and animate it
		if($('.position1').hasClass('.yourTurn')){
			$('.position1').removeClass('.yourTurn');
			$('.position1').find('img').css({'outline': 'none'});
			$('.timer').remove();
			$('.startTimer').remove();
			$('.position2').addClass('.yourTurn');
			$('.position2').find('img').css({'outline': '5px solid blue'});
			$('.position2').find('img').after('<div class=\"timer\"><div class=\"startTimer\"></div></div>');
			$('.startTimer').on('animationend', nextTurn);
		}
		else if($('.position2').hasClass('.yourTurn')){
			$('.position2').removeClass('.yourTurn');
			$('.position2').find('img').css({'outline': 'none'});
			$('.timer').remove();
			$('.startTimer').remove();
			$('.position3').addClass('.yourTurn');
			$('.position3').find('img').css({'outline': '5px solid blue'});
			$('.position3').find('img').after('<div class=\"timer\"><div class=\"startTimer\"></div></div>');
			$('.startTimer').on('animationend', nextTurn);
		}
		else if($('.position3').hasClass('.yourTurn')){
			$('.position3').removeClass('.yourTurn');
			$('.position3').find('img').css({'outline': 'none'});
			$('.timer').remove();
			$('.startTimer').remove();
			$('.position4').addClass('.yourTurn');
			$('.position4').find('img').css({'outline': '5px solid blue'});
			$('.position4').find('img').after('<div class=\"timer\"><div class=\"startTimer\"></div></div>');
			$('.startTimer').on('animationend', nextTurn);
		}
		else if($('.position4').hasClass('.yourTurn')){
			$('.position4').removeClass('.yourTurn');
			$('.position4').find('img').css({'outline': 'none'});
			$('.timer').remove();
			$('.startTimer').remove();
			$('.position5').addClass('.yourTurn');
			$('.position5').find('img').css({'outline': '5px solid blue'});
			$('.position5').find('img').after('<div class=\"timer\"><div class=\"startTimer\"></div></div>');
			$('.startTimer').on('animationend', nextTurn);
		}
		else if($('.position5').hasClass('.yourTurn')){
			$('.position5').removeClass('.yourTurn');
			$('.position5').find('img').css({'outline': 'none'});
			$('.timer').remove();
			$('.startTimer').remove();
			$('.position6').addClass('.yourTurn');
			$('.position6').find('img').css({'outline': '5px solid blue'});
			$('.position6').find('img').after('<div class=\"timer\"><div class=\"startTimer\"></div></div>');
			$('.startTimer').on('animationend', nextTurn);
		}
		else{
			$('.position6').removeClass('.yourTurn');
			$('.position6').find('img').css({'outline': 'none'});
			$('.timer').remove();
			$('.startTimer').remove();
			$('.position1').addClass('.yourTurn');
			$('.position1').find('img').css({'outline': '5px solid blue'});
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
		// $('.game').find('.yourTurn').closest('img').css({'opacity': '0.5'});
		nextTurn();
	});
	$('.menu').on('click', '#raise', function(){
		$('.startTimer').stop(true, true);
		nextTurn();
	});
	$('form').submit(function(e){
		e.preventDefault();
	});
});
