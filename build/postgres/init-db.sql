create table Faculties(
    id serial not null primary key,
    name varchar
);

create table Groups(
    id serial not null primary key,
    name varchar,
    faculty_id integer not null references Faculties(id) on delete cascade
);

create table Students(
    id serial not null primary key,
    lastname varchar,
    firstname varchar,
    middlename varchar,
    birthdate date,
    group_id integer not null references Groups(id) on delete cascade,
    phone varchar,
    sex int
);

create table Users(
    id serial not null primary key,
    login varchar,
    pwd varchar,
    user_type int
);

insert into users (login, pwd, user_type) values ('admin', 'admin', 1), ('user', 'user', 0);

-- insert into users (id) select * from generate_series(1, 1000000)