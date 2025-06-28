-- Создаем базу данных для приложения
CREATE DATABASE frodo;

-- Создаем пользователя для приложения (работает с public-схемой)
CREATE USER ektov WITH PASSWORD '12345';

-- Даем права на дефолтную схему public
\c frodo
GRANT ALL PRIVILEGES ON SCHEMA public TO ektov;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO ektov;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON TABLES TO ektov;