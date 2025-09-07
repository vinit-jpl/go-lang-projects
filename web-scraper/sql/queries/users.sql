-- name: CreateUser :one
-- Creates a new user and returns the created user's details.

INSERT INTO users (id, created_at, updated_at, name)
VALUES($1, $2, $3, $4) 
RETURNING *; 