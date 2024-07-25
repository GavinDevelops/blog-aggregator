-- name: CreateFeedFollows :one
INSERT INTO feed_follows (id, feed_id, user_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: DeleteFeedFollow :one
DELETE FROM feed_follows WHERE id = $1 and user_id = $2
RETURNING *;

-- name: GetFeedFollow :one
SELECT * from feed_follows where id = $1;

-- name: GetFeedFollowsForUser :many
SELECT * from feed_follows where user_id = $1;
