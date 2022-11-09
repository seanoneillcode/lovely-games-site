package games

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetGames() ([]*Game, error) {
	const limit = 10
	var allGames []*Game

	q := `SELECT id, title, description, screenshot, game_file, frame_width, frame_height FROM games limit $1`
	rows, err := r.db.Query(q, limit)
	if err != nil {
		return allGames, err
	}
	defer rows.Close()

	for rows.Next() {
		game := &Game{}
		err = rows.Scan(
			&game.Id,
			&game.Name,
			&game.Description,
			&game.Screenshot,
			&game.GameFile,
			&game.FrameWidth,
			&game.FrameHeight,
		)
		if err != nil {
			return allGames, err
		}
		allGames = append(allGames, game)
	}

	return allGames, nil
}

func (r *Repository) GetGame(id string) (*Game, error) {
	q := `SELECT id, title, description, screenshot, game_file, frame_width, frame_height FROM games WHERE id = $1`
	row := r.db.QueryRow(q, id)
	game := &Game{}
	err := row.Scan(
		&game.Id,
		&game.Name,
		&game.Description,
		&game.Screenshot,
		&game.GameFile,
		&game.FrameWidth,
		&game.FrameHeight,
	)
	if err != nil {
		return game, err
	}

	return game, nil
}

func (r *Repository) AddGame(game *Game) error {
	v := `INSERT INTO games (id, title, description, screenshot, game_file, frame_width, frame_height, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())`
	_, err := r.db.Exec(v, game.Id, game.Name, game.Description, game.Screenshot, game.GameFile, game.FrameWidth, game.FrameHeight)
	return err
}

func (r *Repository) EditGame(id string, game *Game) error {
	v := `UPDATE games SET title = $1, description = $2, screenshot = $3, game_file = $4, frame_width = $5,frame_height = $6,updated_at = NOW() WHERE id = $7`
	_, err := r.db.Exec(v, game.Name, game.Description, game.Screenshot, game.GameFile, game.FrameWidth, game.FrameHeight, id)
	return err
}
