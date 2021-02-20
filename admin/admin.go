package admin

import (
	"context"
	"encore.dev/storage/sqldb"
	"fmt"
	"time"
)

type GameStudio struct {
	ID int64

	Name string

	Created time.Time
}

type Game struct {
	ID int64

	GameStudio *GameStudio

	Name string

	Created time.Time
}

type CreateGameStudioParams struct {
	Name string
}

type CreateGameParams struct {
	GameStudioID int64

	Name string
}

// CreateGameStudio creates a new organisation entity.
// encore:api public
func CreateGameStudio(ctx context.Context, params *CreateGameStudioParams) (*GameStudio, error) {
	gameStudio := &GameStudio{Name: params.Name, Created: time.Now()}

	err := sqldb.QueryRow(ctx, `
		INSERT INTO game_studio (name, created)
		VALUES ($1, $2)
		RETURNING id
	`, gameStudio.Name, gameStudio.Created).Scan(&gameStudio.ID)

	if err != nil {
		return nil, fmt.Errorf("could not create Game Studio: %v", err)
	}

	return gameStudio, nil
}

// CreateGame creates a new Game title.
// encore:api public
func CreateGame(ctx context.Context, params *CreateGameParams) (*Game, error) {
	game := &Game{
		GameStudio: &GameStudio{ID: params.GameStudioID},
		Name:       params.Name,
		Created:    time.Now(),
	}

	err := sqldb.QueryRow(ctx, `
		WITH g AS (
			INSERT INTO game (game_studio_id, name, created)
			VALUES ($1, $2, $3)
			RETURNING id
		)
		SELECT (SELECT id FROM g) AS game_id, game_studio.name, game_studio.created
		FROM game_studio
	`, game.GameStudio.ID, game.Name, game.Created).Scan(
		&game.ID,
		&game.GameStudio.Name,
		&game.GameStudio.Created)

	if err != nil {
		return nil, fmt.Errorf("could not create Game: %v", err)
	}

	return game, nil
}
