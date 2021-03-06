CREATE TABLE IF NOT EXISTS users (id uuid PRIMARY KEY, name VARCHAR(200) NOT NULL, email VARCHAR(150) NOT NULL, username VARCHAR(150) NOT NULL, password VARCHAR(300) NOT NULL, profile_id INT NOT NULL, created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, deleted_at TIMESTAMP);

ALTER TABLE users ADD CONSTRAINT fk_profile FOREIGN KEY (profile_id) REFERENCES profiles(id);
