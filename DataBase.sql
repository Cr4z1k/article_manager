CREATE TABLE users(
	id SERIAL PRIMARY KEY NOT NULL,
	name TEXT
);

CREATE TABLE authors(
	id SERIAL PRIMARY KEY NOT NULL,
	rating NUMERIC(4, 3),
	user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL  
); 

CREATE TABLE themes(
	id SERIAL PRIMARY KEY NOT NULL,
	name TEXT
);

CREATE TABLE articles(
	id SERIAL PRIMARY KEY NOT NULL,
	name TEXT,
	rating NUMERIC(4, 3),
	link TEXT,
	file_path TEXT NOT NULL
);

CREATE TABLE article_themes(
	id SERIAL PRIMARY KEY NOT NULL,
	art_id INT REFERENCES articles(id) ON DELETE CASCADE NOT NULL,
	theme_id INT REFERENCES themes(id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE article_authors(
	id SERIAL PRIMARY KEY NOT NULL,
	art_id INT REFERENCES articles(id) ON DELETE CASCADE NOT NULL,
	auth_id INT REFERENCES authors(id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE authors_rating(
	user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
	auth_id INT REFERENCES authors(id) ON DELETE CASCADE NOT NULL,
	rating INT CHECK (rating >= 1 AND rating <= 5)
);

CREATE TABLE articles_rating(
	user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
	art_id INT REFERENCES articles(id) ON DELETE CASCADE NOT NULL,
	rating INT CHECK (rating >= 1 AND rating <= 5)
);

