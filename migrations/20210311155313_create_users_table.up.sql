CREATE TABLE users (
  id int PRIMARY KEY AUTO_INCREMENT NOT NULL,
  email VARCHAR(60) NOT NULL UNIQUE, 
  password VARCHAR(60) NOT NULL,
  first_name VARCHAR(60) NOT NULL, 
  last_name VARCHAR(60) NOT NULL,
  visible BOOLEAN
)