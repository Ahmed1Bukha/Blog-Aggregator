-- name: CreatePost :one
INSERT INTO posts(id,created_at,updated_at,title,url,description,published_at,feed_id,user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9
)
RETURNING *;
-- name: GetPostsUser :many
SELECT * FROM posts
WHERE posts.user_id = $1
LIMIT $2;