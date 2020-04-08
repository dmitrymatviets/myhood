create schema myhood;

create table users
(
	user_id int auto_increment,
	email nvarchar(255) not null,
	name nvarchar(50) not null,
	surname nvarchar(50) not null,
	date_of_birth date not null,
	gender nvarchar(1) not null,
	interests json null,
	city_id int not null,
	page_slug nvarchar(50) null,
	page_is_private bool not null,
	avatar nvarchar(400) null,
	constraint users_pk
		primary key (user_id)
);

alter table users
	add hash nvarchar(255) not null after name;

create table sessions
(
	session_id nvarchar(100) not null,
	user_id int not null,
	created datetime not null
);

alter table sessions
	add constraint sessions_pk
		primary key (session_id);


create table friends
(
	user_id int not null,
	friend_id int not null
);

