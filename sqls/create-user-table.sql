CREATE TABLE IF NOT EXISTS `api-go`.`users` (
	id INTEGER UNSIGNED auto_increment NOT NULL,
	name varchar(150) NOT NULL,
	CONSTRAINT user_PK PRIMARY KEY (id)
);