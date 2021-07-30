CREATE TABLE users
(
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE events
(
    id          serial       not null unique,
    title       varchar(255) not null,
    time_utc    varchar(255) not null,
    time        varchar(255) not null,
    description varchar(255)
);