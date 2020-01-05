USE processor;

CREATE TABLE IF NOT EXISTS acquirers (
                                     id INTEGER PRIMARY KEY AUTO_INCREMENT,
                                     url VARCHAR(255) NOT NULL
) ENGINE = INNODB;

INSERT INTO processor.acquirers
(url) VALUES
("http://api.fake.cielo.com/buy"),
("http://fake.api.stone.com/sell");