CREATE SCHEMA IF NOT EXISTS ecounter;

--* Tabela dos usuarios
CREATE TABLE IF NOT EXISTS ecounter.users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


--* Tabela das Contagens 
CREATE TABLE IF NOT EXISTS ecounter.quotes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    quote INT,
    userid INT REFERENCES ecounter.users(id)
);