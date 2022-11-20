package games

import (
	"github.com/gomarkdown/markdown"
	"os"
)

type Repository struct {
	games []*Game
}

func NewRepository() *Repository {
	r := &Repository{games: []*Game{
		{
			Id:               "587sdnre86dh",
			Name:             "Plutos Revenge",
			ShortDescription: "Blast away invaders from Pluto.",
			DescriptionFile:  "description.md",
			Screenshot:       "screenshot.png",
			HeaderImage:      "header.png",
			ReleaseState:     "Released",
			GameFile:         "game.wasm",
			FrameWidth:       720,
			FrameHeight:      960,
		},
		{
			Id:               "dh58sgt38dd4",
			Name:             "Super B Saves The City",
			ShortDescription: "Aliens are attacking the city, blast them away!",
			DescriptionFile:  "description.md",
			Screenshot:       "screenshot.png",
			HeaderImage:      "header.png",
			ReleaseState:     "Released",
			GameFile:         "game.wasm",
			FrameWidth:       720,
			FrameHeight:      960,
		},
	}}
	for _, game := range r.games {
		data, err := os.ReadFile("static/" + game.Id + "/" + game.DescriptionFile)
		if err != nil {
			panic(err)
		}
		game.Description = string(markdown.ToHTML(data, nil, nil))
	}

	return r
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
