package games

import (
	"github.com/gomarkdown/markdown"
	"os"
)

type Repository struct {
	games []*Game
}

func NewRepository() *Repository {
	r := &Repository{games: []*Game{}}
	g := &Game{
		Id:               "587sdnre86dh",
		Name:             "Plutos Revenge",
		ShortDescription: "Blast away invaders from Pluto.",
		DescriptionFile:  "description.md",

		Screenshot:   "screenshot.png",
		HeaderImage:  "header.png",
		ReleaseState: "Released",
		GameFile:     "game.wasm",
		FrameWidth:   720,
		FrameHeight:  960,
	}
	data, err := os.ReadFile("static/" + g.Id + "/" + g.DescriptionFile)
	if err != nil {
		panic(err)
	}
	g.Description = string(markdown.ToHTML(data, nil, nil))

	r.games = append(r.games, g)

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
