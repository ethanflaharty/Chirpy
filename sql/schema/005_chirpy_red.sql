-- +goose up
alter table users
add is_chirpy_red bool not null Default false;

-- +goose down
alter table users
drop is_chirpy_red;