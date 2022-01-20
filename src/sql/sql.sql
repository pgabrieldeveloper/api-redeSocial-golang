create database if not exists devbook;
use devbook;

drop table if exists  usuarios;

create table usuarios (
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    password varchar(100) not null,
    email varchar(50) not null unique,
    createAt timestamp default current_timestamp()
) engine=innodb;

