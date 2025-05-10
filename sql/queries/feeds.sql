-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeedByUrl :one
SELECT * FROM feeds
WHERE url = $1
LIMIT 1;

-- name: GetFeeds :many
SELECT feeds.name, feeds.url, users.name AS username FROM feeds
INNER JOIN users ON users.id = feeds.user_id;

-- name: MarkFeedFetched :one
UPDATE feeds
SET 
    last_fetched_at = NOW(),
    updated_at = NOW()
WHERE id = $1
RETURNING *;