CREATE TABLE urls
(
    id            serial       not null unique,
    fullURL       varchar(255) not null unique,
    tinyURL      varchar(255) not null  unique
);