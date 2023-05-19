-- migrate:up
CREATE TABLE IF NOT EXISTS books (
	id INT PRIMARY KEY,
  title TEXT ,
	is_publish BOOLEAN,
	author_id INT,
	price INT, 
	CONSTRAINT author_fk FOREIGN KEY (author_id) REFERENCES authors(id)
	)

-- migrate:down
DROP TABLE IF EXISTS books

