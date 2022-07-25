BEGIN;
CREATE table IF NOT EXISTS category (
    id serial primary key,
    genre text not NULL
);

CREATE TABLE IF NOT EXISTS book(
    id serial PRIMARY KEY,
    title varchar(255) NOT NULL,
    author TEXT NOT NULL,
	category_id integer REFERENCES category (id)
  
);

insert into category(genre)
values
('Fantasy'),
('Drama'),
('Comedy'),
('Triller');

insert into book (author,title,category_id)
values
('Stephene Meyer','Twilight',1),
('Cristopher Nolan','Interstellar',2),
('Nurlan Saburov','StandUp',3),
('Timur Bekmambetov','Find',4);

Select * from book

join category on book.category_id = category.id;

COMMIT;