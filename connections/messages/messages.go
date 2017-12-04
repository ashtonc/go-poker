package messages

const (
	// KindConnected is sent when a player connects
	KindWatching = iota + 1

	// KindPlayerJoined is sent when a player joins
	KindWatcherJoins

	//KindLeftTable when player leave table
	KindLeftTable

	KindSitting

	KindSitterJoins

	KindLeftSeat

	KindPlaying

	KindPlayerJoins

	KindLeftGame

	// KindPlayerQuit is sent when a player leaves
	KindPlayerQuit

	// KindPlayerCheck is sent when a player checks
	KindPlayerCheck

	// KindPlayerBet is sent when a player bets
	KindPlayerBet

	// KindPlayerCall is sent when a player calls
	KindPlayerCall

	// KindPlayerFold is sent when a player folds
	KindPlayerFold

	// KindPlayerTrade is sent when a player trades cards
	KindPlayerTrade

	// Kind for starting round
	// KindStartRound

	// Kind for finishing round
	// KindFinishRound
)
