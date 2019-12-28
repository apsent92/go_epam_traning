create table genres (
    id_genre serial primary key,
    title varchar(20) not null
);

create table films (
    id_film serial primary key,
    id_genre int not null,
    name varchar(20) not null,
    FOREIGN KEY (id_genre) REFERENCES genres (id_genre)
);

