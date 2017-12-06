package connection

const (
	KindGameState = iota + 1

	KindTakeSeat
	KindLeaveSeat

	KindPlayerSits
	KindPlayerLeaves

	KindCheck
	KindBet
	KindCall
	KindFold
	KindDiscard

	KindTimedOut
)

type SeatMessage struct {
	Kind int `json:"kind"`
	Seat int `json:"seattaken"`
}

type PlayerSitMessage struct {
	Kind int `json:"kind"`
}

type PlayerLeaveMessage struct {
	Kind int `json:"kind"`
}

type DiscardMessage struct {
	Kind int `json:"kind"`
}

type CheckMessage struct {
	Kind int `json:"kind"`
}
