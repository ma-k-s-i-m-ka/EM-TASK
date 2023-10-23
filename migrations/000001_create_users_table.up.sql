CREATE TABLE IF NOT EXISTS users (
 id             bigserial   primary key,
 name           text        not null,
 surname        text        not null,
 patronymic     text,
 age            int         not null,
 gender         text        not null,
 nation         text        not null
);