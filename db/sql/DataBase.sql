------------------------- Tables ---------------------------------------------------
CREATE TABLE users(
	id SERIAL PRIMARY KEY NOT NULL,
	name TEXT
);

CREATE TABLE users_credentials(
	id SERIAL PRIMARY KEY NOT NULL,
	user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
	login TEXT UNIQUE NOT NULL,
	password_hash TEXT NOT NULL
);

CREATE TABLE authors(
	id SERIAL PRIMARY KEY NOT NULL,
	rating NUMERIC(4, 3),
	user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL  
); 

CREATE TABLE themes(
	id SERIAL PRIMARY KEY NOT NULL,
	name TEXT UNIQUE NOT NULL
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
------------------------- Functions ------------------------------------------------
CREATE FUNCTION sign_up(_name TEXT, _login TEXT, _password_hash TEXT, is_author bool)
RETURNS JSON
LANGUAGE PLPGSQL
AS $$
DECLARE
	new_user_id INT;
	result JSON;
BEGIN
	IF NOT EXISTS (SELECT 1 FROM users_credentials WHERE login = _login)
	THEN
		INSERT INTO users(name)
		VALUES (_name)
		RETURNING id INTO new_user_id;
	
		INSERT INTO users_credentials(user_id, login, password_hash)
		VALUES (new_user_id, _login, _password_hash);
	
		IF is_author
		THEN
			INSERT INTO authors(user_id)
			VALUES (new_user_id);
		END IF;
		
		result := json_build_object('success', true,
									'message', 'User registered successfully');
	ELSE
		result := json_build_object('success', true,
									'message', 'User registered successfully');
	END IF;
	
	RETURN result;
END
$$; 


CREATE OR REPLACE FUNCTION get_hash(_login TEXT)
RETURNS TEXT
LANGUAGE PLPGSQL
AS $$
BEGIN
	RETURN COALESCE((SELECT password_hash FROM users_credentials WHERE login = _login), '');
END
$$;


CREATE OR REPLACE FUNCTION add_article(_name TEXT, authors INT[], themes TEXT[], _link TEXT, _file_path TEXT)
RETURNS BOOL
LANGUAGE PLPGSQL
AS $$
DECLARE
	new_article_id INT;
	theme TEXT;
	author_id INT;
BEGIN
	IF NOT EXISTS   (
					SELECT 1
					FROM articles
					WHERE link = _link OR _file_path = _file_path
					)
	THEN
		INSERT INTO articles(name, link, file_path)
		VALUES (_name, _link, _file_path)
		RETURNING id into new_article_id;
	
		FOR theme IN SELECT UNNEST(themes)
		LOOP
			INSERT INTO article_themes(art_id, theme_id)
			VALUES (new_article_id, (SELECT id FROM themes WHERE name = theme));
		END LOOP;
	
		FOR author_id IN SELECT UNNEST(authors)
		LOOP
			INSERT INTO article_authors(art_id, auth_id)
			VALUES (new_article_id, author_id);
		END LOOP;
		
		RETURN true;
	ELSE
		RETURN false;
	END IF;
END
$$;

		
	
------------------------- Procedures -----------------------------------------------

------------------------- Tables check ---------------------------------------------
select * from users
select * from authors
select * from users_credentials

		









