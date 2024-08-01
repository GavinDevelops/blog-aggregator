-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT * from feeds;

-- name: GetFeed :one
SELECT * from feeds where id = $1;

-- name: GetNextFeedsToFetch :many
SELECT * from feeds ORDER BY last_fetched_at LIMIT $1;

-- name: UpdateLastFetchedAt :one
UPDATE feeds SET last_fetched_at = NOW() WHERE id = $1
RETURNING *;
