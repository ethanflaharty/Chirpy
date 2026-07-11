-- +goose up
create TABLE refresh_tokens(
    token TEXT primary KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP not NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    expires_at TIMESTAMP not NULL,
    revoked_at TIMESTAMP
);

-- +goose down
drop table refresh_tokens;