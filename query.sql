-- name: CreateUserCommand :one
INSERT INTO users (
    email, hashed_password
) VALUES ( 
    $1, $2 
) RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;
