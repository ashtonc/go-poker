
//Slice of connected users
type Watching struct {
  Kind  int    `json:"kind"`
  Name string `json: "string"`
  Watchers []Player `json:"watchers"`
}

//Connects a new user
func NewWatcher(name string, watchers []Player) *Watching {
  return &Watching{
    Kind:  KindWatching,
    Name: name
    Watchers: watchers,
  }
}

type WatcherJoins struct {
  Kind int  `json:"kind"`
  Watcher Player `json:"watcher"`
}

func NewWatcherJoins(userID string) *WatcherJoins {
  return &WatcherJoins{
    Kind: KindWatcherJoins,
    User: User{ID: userID},
  }
}

type LeftTable struct {
  Kind   int    `json:"kind"`
  UserID string `json:"userId"`
}

func NewLeftTable(userID string) *LeftTable {
  return &LeftTable{
    Kind:   KindLeftTable,
    UserID: userID,
  }
}

//
type Sitting struct{
  Kind  int    `json:"kind"`
  Name string `json: "string"`
  Sitters []Player `json:"sitters"`
}

//Sits a user at a table
func NewSitter(name String, sitters []Player) *Sitting{
  return &Siting{
    Kind:  KindSitting,
    Name: name,
    Sitters: sitters,
  }
}

type SitterJoins struct {
  Kind int  `json:"kind"`
  Sitter Player `json:"sitter"`
}

func NewSitterJoins(userID string) *SitterJoins {
  return &SitterJoins{
    Kind: KindSitterJoins,
    User: User{ID: userID},
  }
}


type LeftSeat struct {
  Kind   int    `json:"kind"`
  UserID string `json:"userId"`
}

func NewLeftSeat(userID string) *LeftSeat{
  return &LeftSeat{
    Kind:   KindLeftSeat,
    UserID: userID,
  }
}

type Playing struct {
	Kind int `json: "king"`
	Name int `json: "string"`
	Players []Player `json: "Players"`
}

func NewPlayer(name String, players []Player) *Playing{
	return &Playing{
		Kind: KindPlaying,
		Name: name,
		Players: players,
	}
}


type PlayerJoins struct {
  Kind int  `json:"kind"`
  Sitter Player `json:"player"`
}

func NewPlayerJoins(userID string) *PlayerJoins {
  return &PlayerJoins{
    Kind: KindPlayerJoins,
    User: User{ID: userID},
  }
}

type LeftGame struct {
  Kind   int    `json:"kind"`
  UserID string `json:"userId"`
}

func NewLeftSeat(userID string) *LeftGame{
  return &LeftGame{
    Kind:   KindLeftGame,
    UserID: userID,
  }
}

