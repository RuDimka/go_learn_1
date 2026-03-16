DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	age INTEGER,
	profession TEXT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (name, age, profession) VALUES
('Зинаида', 22, 'Бухгалтер'),
('Борис', 23, 'Слесарь'),
('Никита', 25, 'Водитель');