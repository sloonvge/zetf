package db

import "syscall"

/** 
* Created by wanjx in 2019/2/21 22:47
**/

func Initdb() {
	q := `
		CREATE TABLE users (
		user_id INT(8) NOT NULL AUTO_INCREMENT,
		user_name VARCHAR(30) NOT NULL,
		user_passwd VARCHAR(30) NOT NULL,
		user_date DATETIME NOT NULL,
		user_role INT(8) NOT NULL,
		UNIQUE INDEX user_name_unique (user_name),
		PRIMARY KEY (user_id)
		) TYPE=INNODB;`
	q1 := `
		CREATE TABLE articles (
		art_id INT(8) NOT NULL AUTO_INCREMENT,
		art_title VARCHAR(255) NOTNULL,
		art_content TEXT NOT NULL,
		art_date DATETIME NOT NULL,
		art_by INT(8) NOT NULL,
		PRIMARY KEY (art_id)
		) TYPE INNODB;`
	q2 := `
		CREATE TABLE comments (
		comment_id INT(8) NOT NULL AUTO_INCREMENT,
		comment_content VARCHAR(255) NOT NULL,
		comment_date DATETIME NOT NULL,
		comment_by INT(8) NOT NULL,
		comment_art INT(8) NOT NULL
		) TYPE=INNODB;`
	FOREIGN KEY
	ALTER TABLE comments ADD FOREIFN KEY(comment_art) REFERENCE articles(art_id) ON DELETE CASCADE ON UPDATE CASCADE
	
}
