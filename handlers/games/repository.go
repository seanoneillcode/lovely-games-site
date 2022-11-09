package games

import "errors"

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

var games = []*Game{
	{
		Id:           "587sdnre86dh",
		Name:         "Pluto's Revenge",
		Description:  "Blast away invaders from Pluto.",
		Screenshot:   "VbhV3vC5AWX39IVU.png",
		GameFile:     "WSP2NcHciWvqZTa2.wasm",
		ReleaseState: "released",
		FrameWidth:   720,
		FrameHeight:  960,
	},
	{
		Id:           "3489fhjbef",
		Name:         "2048",
		Description:  "A puzzle game where the tiles must add up.",
		Screenshot:   "86hrusuwnv84kshg.png",
		GameFile:     "86hrusuwnv84kshg.wasm",
		ReleaseState: "released",
		FrameWidth:   420,
		FrameHeight:  600,
	},
}

func (r *Repository) GetGames() ([]*Game, error) {
	return games, nil
}

func (r *Repository) GetGame(id string) (*Game, error) {
	for _, v := range games {
		if v.Id == id {
			return v, nil
		}
	}
	return nil, errors.New("game with id not found: " + id)
}

func (r *Repository) AddGame(game *Game) error {
	games = append(games, game)
	return nil
}

func (r *Repository) EditGame(id string, game *Game) error {
	for _, g := range games {
		if g.Id == id {
			g.Name = game.Name
			g.Description = game.Description
			g.ReleaseState = game.ReleaseState
			//g.Screenshot = game.Screenshot
			//g.GameFile = game.GameFile
			return nil
		}
	}
	return nil
}
