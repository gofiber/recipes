-- name: GetAuthors :many
SELECT * FROM author;

-- name: GetAuthor :one
SELECT * FROM author WHERE id = $1;

-- name: NewAuthor :one
INSERT INTO author (email, name) VALUES ($1, $2) RETURNING *;

-- name: UpdateAuthor :one
UPDATE author SET email = $1, name = $2 WHERE id = $3 RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM author WHERE id = $1;
