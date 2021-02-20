package leaderboard

import (
	"context"
	"encore.dev/storage/sqldb"
	"fmt"
	"time"
)

type Leaderboard struct {
	// ID is the unique id for the leaderboard.
	ID int64

	// Name is an alphanumeric string.
	Name string
}

type Score struct {
	// ID is the unique id for the score.
	ID int64

	// LeaderboardID is the unique id of the leaderboard the score is part of.
	LeaderboardID int64

	// Score is the numeric value representing the user score.
	Score float32

	// Created is the time the board was created.
	Created time.Time
}

type CreateLeaderboardParams struct {
	Name string
}

type PublishScoreParams struct {
	Score float32
}

// CreateLeaderboard creates a new leaderboard.
// encore:api public
func CreateLeaderboard(ctx context.Context, params *CreateLeaderboardParams) (*Leaderboard, error) {
	l := &Leaderboard{Name: params.Name}

	err := sqldb.QueryRow(ctx, `
		INSERT INTO leaderboard (name)
		VALUES ($1)
		RETURNING id
	`, l.Name).Scan(&l.ID)

	if err != nil {
		return nil, fmt.Errorf("could not create leaderboard: %v", err)
	}

	return l, nil
}

// PublishScore publishes a new score to an existing leaderboard.
// encore:api public
func PublishScore(ctx context.Context, params *PublishScoreParams) (*Score, error) {
	s := &Score{LeaderboardID: 1, Score: params.Score, Created: time.Now()}

	err := sqldb.QueryRow(ctx, `
		INSERT INTO score (leaderboard_id, score,  created)
		VALUES ($1, $2, $3)
		RETURNING id
	`, s.LeaderboardID, params.Score, s.Created).Scan(&s.ID)

	if err != nil {
		return nil, fmt.Errorf("could not add score: %v", err)
	}

	return s, nil
}
