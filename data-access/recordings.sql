drop database if exists recordings;
create database recordings;

use recordings;

drop table if exists album;
create table album (
	id int primary key auto_increment,
    title varchar(128) not null,
    artist varchar(255) not null,
    price decimal(5,2) not null
);

insert into album
  (title, artist, price)
values
  ('Blue Train', 'John Coltrane', 56.99),
  ('Giant Steps', 'John Coltrane', 63.99),
  ('Jeru', 'Gerry Mulligan', 17.99),
  ('Sarah Vaughan', 'Sarah Vaughan', 34.98);
  
select * from album;