-- name: GetPosts :many
SELECT * FROM post;

-- name: GetPost :one
SELECT * FROM post WHERE id = $1;

-- name: NewPost :one
INSERT INTO post (title, content, author) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdatePost :one
UPDATE post SET title = $1, content = $2, author = $3 WHERE id = $4 RETURNING *;

-- name: DeletePost :exec
DELETE FROM post WHERE id = $1;

