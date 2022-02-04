create database if not exists app;
create table if not exists app.users (
    id serial primary key,
    first_name text not null,
    last_name text not null,
    email text not null,
    created_at timestamp default now()
);
