create table genres (
    id_genre serial primary key,
    title varchar(20) not null
);

create table films (
    id_film serial primary key,
    name varchar(20) not null
);

create table list (
    film_id int,
    genre_id int,
    FOREIGN KEY (film_id) REFERENCES films (id_film),
    FOREIGN KEY (genre_id) REFERENCES genres (id_genre)
)

