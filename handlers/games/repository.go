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
			Id:               "pluto",
			Name:             "Plutos Revenge",
			ShortDescription: "Blast away invaders from Pluto.",
			DescriptionFile:  "description.md",
			Screenshot:       "screenshot.png",
			HeaderImage:      "header.png",
		},
		{
			Id:               "superb",
			Name:             "Super B Saves The City",
			ShortDescription: "Aliens are attacking the city, blast them away!",
			DescriptionFile:  "description.md",
			Screenshot:       "screenshot.png",
			HeaderImage:      "header.png",
		},
		{
			Id:               "chess",
			Name:             "Super Fun Action Chess",
			ShortDescription: "Realtime action chess",
			DescriptionFile:  "description.md",
			Screenshot:       "screenshot.png",
			HeaderImage:      "header.png",
		},
		{
			Id:               "tactics",
			Name:             "Fantasy Tactics",
			ShortDescription: "Tough Tactics in a fantasy world.",
			DescriptionFile:  "description.md",
			Screenshot:       "screenshot.png",
			HeaderImage:      "header.png",
		},
		{
			Id:               "beautiful",
			Name:             "Beautiful Castle Death Machine",
			ShortDescription: "Smash blocks to build a matching row.",
			DescriptionFile:  "description.md",
			Screenshot:       "screenshot.png",
			HeaderImage:      "header.png",
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
