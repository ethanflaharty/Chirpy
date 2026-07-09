-- +goose up 
alter TABLE users 
add hashed_password TEXT NOT NULL Default 'unset';


-- +goose down
alter table users
drop hashed_passwords;