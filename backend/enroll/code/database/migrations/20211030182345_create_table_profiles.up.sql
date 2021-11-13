CREATE TABLE IF NOT EXISTS profiles (id SERIAL PRIMARY KEY, profile_name VARCHAR(50) NOT NULL, created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, deleted_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP);

INSERT INTO profiles (id, profile_name) VALUES (1, 'ADMIN');
INSERT INTO profiles (id, profile_name) values (2, 'USER');
