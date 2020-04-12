create schema myhood;

create table cities
(
	city_id bigint auto_increment,
	name nvarchar(100) not null,
	constraint cities_pk
		primary key (city_id)
) ENGINE = INNODB;

insert into cities(name)
values('Москва'),('Санкт-Петербург'),('Казань'),('Нижний Новгород');



create table users
(
	user_id int auto_increment,
	email nvarchar(255) not null,
    hash nvarchar(255) not null,
	name nvarchar(50) not null,
	surname nvarchar(50) not null,
	date_of_birth datetime not null,
	gender nvarchar(1) not null,
	interests json null,
	city_id int not null,
	page_slug nvarchar(50) null,
	page_is_private bool not null,
	avatar nvarchar(400) null,
	constraint users_pk
		primary key (user_id)
) ENGINE = INNODB;



create table sessions
(
	session_id nvarchar(100) not null,
	user_id int not null,
	created datetime not null,
    constraint sessions_pk
        primary key (session_id)
) ENGINE = INNODB;


create table friends
(
	user_id int not null,
	friend_id int not null
) ENGINE = INNODB;
