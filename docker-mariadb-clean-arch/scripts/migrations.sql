-- In this script, 'dmca' stands for 'Docker MariaDB Clean Arch'.
DROP DATABASE IF EXISTS fiber_dmca;
CREATE DATABASE IF NOT EXISTS fiber_dmca;
USE fiber_dmca;

-- Create a sample table.
CREATE TABLE users (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  address VARCHAR(255) NOT NULL,
  created VARCHAR(255) NOT NULL,
  modified VARCHAR(255) NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB CHARACTER SET utf8;

-- Populate table with 10 users.
INSERT INTO users VALUES
(1, 'Sayu Ogiwara', 'Hokkaido, Japan', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(2, 'Chizuru Ichinose', 'Tokyo, Japan', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(3, 'Asagi Aiba', 'Kyoto, Japan', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(4, 'Rin Tohsaka', 'Kobe, Japan', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(5, 'Mai Sakurajima', 'Fujisawa, Japan', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(6, 'Aki Adagaki', 'Fukuoka, Japan', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(7, 'Asuna Yuuki', 'Shinagawa, Japan', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(8, 'Ruka Sarashina', 'Gotenba, Japan', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(9, 'Miyuki Shiba', 'Nagano, Japan', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(10, 'Fumino Furuhashi', 'Niigata, Japan', UNIX_TIMESTAMP(), UNIX_TIMESTAMP());

-- Create another sample table.
CREATE TABLE cities (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  created VARCHAR(255) NOT NULL,
  modified VARCHAR(255) NOT NULL,
  user INT NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (user) REFERENCES users (id)
) ENGINE=InnoDB CHARACTER SET utf8;

-- Create five data samples.
INSERT INTO cities VALUES
(1, 'Hokkaido', UNIX_TIMESTAMP(), UNIX_TIMESTAMP(), 1),
(2, 'Tokyo', UNIX_TIMESTAMP(), UNIX_TIMESTAMP(), 1),
(3, 'Kyoto', UNIX_TIMESTAMP(), UNIX_TIMESTAMP(), 1),
(4, 'Osaka', UNIX_TIMESTAMP(), UNIX_TIMESTAMP(), 2),
(5, 'Fukuoka', UNIX_TIMESTAMP(), UNIX_TIMESTAMP(), 3);
