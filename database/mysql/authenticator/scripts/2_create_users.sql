USE authenticator;

CREATE TABLE
    IF NOT EXISTS users
(
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    username VARCHAR
           (255) NOT NULL,
    pass VARCHAR
           (255) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE = INNODB;

INSERT INTO authenticator.users
(id, username, pass)
VALUES(1, "admin", SHA1("@#$RF@!718"));