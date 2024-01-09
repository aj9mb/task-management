create table user(user_id bigint not null AUTO_INCREMENT,
user_name varchar(80) not null,
password varchar(80) not null,
name varchar(100) not null,
active tinyint(1) default 1 not null,
create_dt datetime not null default current_timestamp,
last_update datetime default null on update current_timestamp,
PRIMARY KEY(user_id));

create table board(board_id bigint AUTO_INCREMENT,
name varchar(100) not null,
created_by bigint not null,
create_dt datetime DEFAULT CURRENT_TIMESTAMP,
last_update datetime DEFAULT NULL 
ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY(board_id));

create table board_user(board_user_id bigint auto_increment,
board_id bigint not null,
user_id bigint not null,
added_by bigint not null,
active tinyint(1) default 1 not null,
create_dt datetime not null default current_timestamp,
last_update datetime default null 
on update current_timestamp,
primary key(board_user_id),
unique(board_id, user_id));

create table task(task_id bigint not null AUTO_INCREMENT,
board_id bigint not null,
added_by bigint not null,
assigned_to bigint not null,
task_desc varchar(200) not null,
active tinyint(1) default 1 not null,
create_dt datetime not null default current_timestamp,
last_update datetime default null on update current_timestamp,
PRIMARY KEY(task_id));