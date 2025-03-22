-- name: CreateFeed :one
INSERT INTO feeds (id, user_id, name, url, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT
    f.id,
    f.name,
    f.url,
    u.name AS username
FROM feeds f
JOIN users u ON u.id = f.user_id;

-- name: GetFeed :one
SELECT
    f.id,
    f.name,
    f.url,
    u.name AS username
FROM feeds f
JOIN users u ON u.id = f.user_id
WHERE f.url = $1;

-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
        VALUES (
                   $1,
                   $2,
                   $3,
                   $4,
                   $5
               )
        RETURNING *
) SELECT
      inserted_feed_follow.*,
      feeds.name AS feed_name,
      users.name AS user_name
FROM inserted_feed_follow
    INNER JOIN users ON users.id = inserted_feed_follow.user_id
    INNER JOIN feeds ON feeds.id = inserted_feed_follow.feed_id;

-- name: GetFeedFollowsForUser :many
SELECT
    f.id feed_id,
    f.name AS feed_name,
    f.url AS feed_url,
    u.name AS user_name
FROM feed_follows ff
JOIN feeds f ON f.id = ff.feed_id
JOIN users u ON u.id = ff.user_id
WHERE ff.user_id = $1;


-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows ff
USING feeds f
WHERE ff.feed_id = f.id 
  AND ff.user_id = $1
  AND f.url = $2;
