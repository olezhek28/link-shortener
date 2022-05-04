-- +goose Up
create table links
(
    id         bigserial primary key,
    long_link  text      not null,
    short_link text      not null,
    created_at timestamp not null default now()
);

-- +goose Down
drop table links;
