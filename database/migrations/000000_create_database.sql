CREATE TABLE if not exists users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    lastname      varchar(255) not null,
    email         varchar(255) not null unique,
    password_hash varchar(255) not null
);
CREATE TABLE if not exists vkusers
(
    vkid          integer      not null unique,
    name          varchar(255) not null,
    lastname      varchar(255) not null,
    email         varchar(255) not null unique
);