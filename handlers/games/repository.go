package games

import "errors"

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

var games = []*Game{
	{
		Id:           "3489fhjbef",
		Name:         "2048",
		Description:  "A puzzle game where the tiles must add up.",
		Screenshot:   "86hrusuwnv84kshg.png",
		GameFile:     "game-2048.wasm",
		ReleaseState: "released",
	},
	{
		Id:           "hd74hsndkghd",
		Name:         "Flappy",
		Description:  "Fly past the pipes using a bird.",
		Screenshot:   "BpLnfgDsc2WD8F2q.png",
		GameFile:     "NfHK5a84jjJkwzDk.wasm",
		ReleaseState: "released",
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
