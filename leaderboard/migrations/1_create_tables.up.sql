CREATE TABLE leaderboard
(
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE score
(
    id BIGSERIAL PRIMARY KEY,
    leaderboard_id BIGINT NOT NULL REFERENCES leaderboard (id),
    score NUMERIC DEFAULT 0,
    created TIMESTAMP WITH TIME ZONE NOT NULL
);
