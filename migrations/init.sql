CREATE TABLE users(
    id int AUTO_INCREMENT not null,
    nick_name varchar(80),
    name varchar(80),
    email varchar(80),
    password varchar(80),
	created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    updated_at datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
 )
