package dto

// Player defines model for Player.
type Player struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Rating float32 `json:"rating"`
}

// PlayerRequest defines model for PlayerRequest.
type PlayerRequest struct {
	Name   string  `json:"name"`
	Rating float32 `json:"rating"`
}
