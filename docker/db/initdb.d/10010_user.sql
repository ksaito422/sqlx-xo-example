SET CHARSET utf8mb4;

CREATE TABLE IF NOT EXISTS sqlx_xo.users (
	id  	INT PRIMARY KEY AUTO_INCREMENT,
	name 	VARCHAR(50) NOT NULL,
	kana 	VARCHAR(50),
	age 	TINYINT UNSIGNED
);

INSERT INTO sqlx_xo.users (name, kana, age)
VALUES ('saito', 'さいとう', 31),
		('tanaka', 'たなか', 20),
		('suzuki', 'すずき', 25);

