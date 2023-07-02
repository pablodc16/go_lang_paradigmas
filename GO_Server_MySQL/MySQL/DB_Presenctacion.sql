drop database books;
CREATE DATABASE books;

use books;

CREATE TABLE authors(
id int not null primary key auto_increment,
name varchar(50)
);

CREATE TABLE books(
id int unsigned not null primary key auto_increment,
title varchar(50),
author_id int,
publish_date datetime,
foreign key (author_id) references authors(id)
);


insert into authors (name) values 
("Frank Herbert"),
("Patrick Süskind"),
("antoine de Saint-Exupéry");


insert into books (title, author_id, publish_date) values 
("El perfume", 2, current_date()),
("El principio", 3, current_date()),
("Dune", 1, current_date()),
("new book", 1, current_date());

select * from authors;
select * from books;


-- este es para consultar la primera demostración 
SELECT 
    b.id, 
    b.title, 
    b.author_id
FROM
    books b
WHERE
    b.id = 3;
    
-- este es para la segunda demostracion mas compleja
SELECT 
    b.id, 
    b.title, 
    b.author_id,
    a.name as 'author',
    b.publish_date
FROM
    books b
inner join authors a
on a.id = b.author_id
