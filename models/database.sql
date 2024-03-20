-- Buat database myGram
CREATE DATABASE myGram;


-- Tabel User
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    age INTEGER NOT NULL CHECK (age > 8),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Photo
CREATE TABLE IF NOT EXISTS photos (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    caption TEXT,
    photo_url TEXT NOT NULL,
    user_id INTEGER REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Comment
CREATE TABLE IF NOT EXISTS comments (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    photo_id INTEGER REFERENCES photos(id),
    message TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel SocialMedia
CREATE TABLE IF NOT EXISTS social_media (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    social_media_url TEXT NOT NULL,
    user_id INTEGER REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Buat trigger untuk menghapus foto saat pengguna dihapus
CREATE OR REPLACE FUNCTION delete_user_photos()
RETURNS TRIGGER AS $$
BEGIN
    DELETE FROM photos WHERE user_id = OLD.id;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_user_photos_trigger
AFTER DELETE ON users
FOR EACH ROW
EXECUTE FUNCTION delete_user_photos();

-- Buat trigger untuk menghapus komentar saat foto dihapus
CREATE OR REPLACE FUNCTION delete_photo_comments()
RETURNS TRIGGER AS $$
BEGIN
    DELETE FROM comments WHERE photo_id = OLD.id;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_photo_comments_trigger
AFTER DELETE ON photos
FOR EACH ROW
EXECUTE FUNCTION delete_photo_comments();
