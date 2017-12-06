$(document).ready(function(){
	
	/*
	var numberOfPlayers = 6;
	$('.player-1').find('.image').addClass('yourTurn');
	$('.player-1').find('.timer').addClass('startTimer');
	function nextTurn(){
		// if it's your turn, draw blue outline and animate it
		if($('.player-1').find('.image').hasClass('yourTurn')){
			$('.player-1').find('.image').removeClass('yourTurn');
			$('.player-1').find('.timer').removeClass('startTimer');
			$('.player-2').find('.image').addClass('yourTurn');
			$('.player-2').find('.timer').addClass('startTimer');
		}
		else if($('.player-2').find('.image').hasClass('yourTurn')){
			$('.player-2').find('.image').removeClass('yourTurn');
			$('.player-2').find('.timer').removeClass('startTimer');
			if(numberOfPlayers >= 3){
				$('.player-3').find('.image').addClass('yourTurn');
				$('.player-3').find('.timer').addClass('startTimer');
			}
			else{
				$('.player-1').find('.image').addClass('yourTurn');
				$('.player-1').find('.timer').addClass('startTimer');
			}
		}
		else if($('.player-3').find('.image').hasClass('yourTurn')){
			$('.player-3').find('.image').removeClass('yourTurn');
			$('.player-3').find('.timer').removeClass('startTimer');
			if(numberOfPlayers >= 4){
				$('.player-4').find('.image').addClass('yourTurn');
				$('.player-4').find('.timer').addClass('startTimer');
			}
			else{
				$('.player-1').find('.image').addClass('yourTurn');
				$('.player-1').find('.timer').addClass('startTimer');
			}
		}
		else if($('.player-4').find('.image').hasClass('yourTurn')){
			$('.player-4').find('.image').removeClass('yourTurn');
			$('.player-4').find('.timer').removeClass('startTimer');
			if(numberOfPlayers >= 5){
				$('.player-5').find('.image').addClass('yourTurn');
				$('.player-5').find('.timer').addClass('startTimer');
			}
			else{
				$('.player-1').find('.image').addClass('yourTurn');
				$('.player-1').find('.timer').addClass('startTimer');
			}
		}
		else if($('.player-5').find('.image').hasClass('yourTurn')){
			$('.player-5').find('.image').removeClass('yourTurn');
			$('.player-5').find('.timer').removeClass('startTimer');
			if(numberOfPlayers == 6){
				$('.player-6').find('.image').addClass('yourTurn');
				$('.player-6').find('.timer').addClass('startTimer');
			}
			else{
				$('.player-1').find('.image').addClass('yourTurn');
				$('.player-1').find('.timer').addClass('startTimer');
			}
		}
		else{
			$('.player-6').find('.image').removeClass('yourTurn');
			$('.player-6').find('.timer').removeClass('startTimer');
			$('.player-1').find('.image').addClass('yourTurn');
			$('.player-1').find('.timer').addClass('startTimer');
		}
	}
	$('#game-menu').on('click', '#call', function(){
		$('.startTimer').stop(true, true);
		nextTurn();
	});
	$('#game-menu').on('click', '#fold', function(){
		$('.startTimer').stop(true, true);
		$('.yourTurn').parent().addClass('folded');
		$('.yourTurn').parent().parent().find('.hand').addClass('folded');
		nextTurn();
	});
	$('#game-menu').on('click', '#raiseButton', function(){
		$('.startTimer').stop(true, true);
		var bet = $('#raise').val();
		var stack = $('.yourTurn').parent().find('.stack').text();
		var amount = stack - bet;
		$('.yourTurn').parent().find('.stack').text(amount);
		nextTurn();
	});
	*/
	// Leave in until comment bug is fixed
	function yourTurn(player){
		$('#game').find(player).find('.image').addClass('yourTurn');
	}
	function folded(player){
		$('#game').find(player).addClass('folded');
	}
	
	function updatePot(amount){
		$('#game').find('#pot').text(amount);
	}
	function updateHand(player, card1, card2, card3, card4, card5){
		
	}
	function updatePot(amount){
		$('#game').find('#pot').text(amount);
	}
	$('#game-menu').on('click', '#check', function(){
		
	});
	$('form').submit(function(e){
		e.preventDefault();
	});
});