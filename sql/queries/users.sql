-- name: CreateUser :one
Insert into users (id, created_at, updated_at, email)
values (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1
)
returning *;

-- name: DeleteUsers :exec
Delete from users;