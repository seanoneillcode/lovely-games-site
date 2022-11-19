package games

type Repository struct {
	games []*Game
}

func NewRepository() *Repository {
	return &Repository{
		games: []*Game{
			{
				Id:           "587sdnre86dh",
				Name:         "Plutos Revenge",
				Description:  "Blast away invaders from Pluto.",
				Screenshot:   "VbhV3vC5AWX39IVU.png",
				ReleaseState: "Released",
				GameFile:     "WSP2NcHciWvqZTa2.wasm",
				FrameWidth:   720,
				FrameHeight:  960,
			},
		},
	}
}

func (r *Repository) GetGames() []*Game {
	return r.games
}

func (r *Repository) GetGame(id string) *Game {
	for _, g := range r.games {
		if g.Id == id {
			return g
		}
	}
	return nil
}
