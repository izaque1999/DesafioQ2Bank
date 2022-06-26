drop database if exists q2bank;

create database q2bank;

use q2bank;

create table users(
    id      int primary key not null auto_increment,
    name_user    varchar(50) not null,
    type_user    enum('common', 'storekeeper'),
    CPF_CNPJ       varchar(50) unique,
    email   varchar(50) unique not null,
    password_user   varchar(50) not null,
    balance   FLOAT(10.2) not null
);

create table transactions(
    id int primary key auto_increment,
    values_transaction FLOAT (10.2) not null,
    id_payer   int not null,
    id_payee int not null,
    date_transaction datetime not null
);

alter table transactions
add foreign key (id_payer) references users(id);

alter table transactions
add foreign key (id_payee) references users(id);

insert into users (name_user, type_user, CPF_CNPJ, email, password_user, balance) values ('Izaque Magalhaes Silva ', 'storekeeper', '321.123.321-80','izaquemagalhaes@outlook.com', '12300596', 2300.30);
insert into users (name_user, type_user, CPF_CNPJ, email, password_user, balance) values ('Eliana Inacio Magalhaes', 'common', '450.155.711-55', 'elianainacio@gmail.com', '22111976e', 950.15);
insert into users (name_user, type_user, CPF_CNPJ, email, password_user, balance) values ('Jose Robson da Silva', 'common', '888.584.211-44', 'robson123@gmail.com', '08101970r', 1350.00);
insert into users (name_user, type_user, CPF_CNPJ, email, password_user, balance) values ("Matheus Magalhaes Oliveira", "storekeeper", "785.585.154-20", "matheusm@hotmail.com", "12300596m", 35000.00);
insert into users (name_user, type_user, CPF_CNPJ, email, password_user, balance) values ("Joao Vitor Leite", "storekeeper", "450.252.147-20", "jvitor@gmail.com", "geleia123", 12350.00);
insert into users (name_user, type_user, CPF_CNPJ, email, password_user, balance) values ("Helena Maria", "common", "225.351.121-21", "helena1997@hotmail.com", "helena1997", 450.00);