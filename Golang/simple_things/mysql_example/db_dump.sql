CREATE TABLE IF NOT EXISTS person (
    person_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL
)  ENGINE=INNODB;

INSERT INTO person (name) VALUES ('Joe');
INSERT INTO person (name) VALUES ('Bob');
INSERT INTO person (name) VALUES ('Alice');