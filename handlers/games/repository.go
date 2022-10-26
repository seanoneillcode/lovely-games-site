package games

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

var games = []Game{
	{
		Id:          "0",
		Name:        "Wolfenstein",
		Description: "A raycast FPS - shoot nazis and their dogs in a tiled maze",
		Screenshot:  "wolfenstein.png",
	},
	{
		Id:          "1",
		Name:        "Doom",
		Description: "FPS - shoot demons and their demons. ",
		Screenshot:  "doom.png",
	},
	{
		Id:          "2",
		Name:        "Super Mario Bros",
		Description: "stomp on creatures",
		Screenshot:  "super-mario.png",
	},
}

func (r *Repository) GetGames() ([]Game, error) {
	return games, nil
}

func (r *Repository) AddGame(game Game) error {
	games = append(games, game)
	return nil
}
