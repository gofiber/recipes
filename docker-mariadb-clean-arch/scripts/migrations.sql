-- In this script, 'dmca' stands for 'Docker MariaDB Clean Arch'.
DROP DATABASE IF EXISTS fiber_dmca;
CREATE DATABASE IF NOT EXISTS fiber_dmca;
USE fiber_dmca;

-- Create a sample table.
CREATE TABLE users (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  address VARCHAR(255) NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB CHARACTER SET utf8;

-- Populate table with 10 users.
INSERT INTO users VALUES
(1, 'Sayu Ogiwara', 'Hokkaido, Japan'),
(2, 'Chizuru Ichinose', 'Tokyo, Japan'),
(3, 'Asagi Aiba', 'Kyoto, Japan'),
(4, 'Rin Tohsaka', 'Kobe, Japan'),
(5, 'Mai Sakurajima', 'Fujisawa, Japan'),
(6, 'Aki Adagaki', 'Fukuoka, Japan'),
(7, 'Asuna Yuuki', 'Shinagawa, Japan'),
(8, 'Ruka Sarashina', 'Gotenba, Japan'),
(9, 'Miyuki Shiba', 'Nagano, Japan'),
(10, 'Fumino Furuhashi', 'Niigata, Japan');