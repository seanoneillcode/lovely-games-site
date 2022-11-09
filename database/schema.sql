CREATE TABLE IF NOT EXISTS public.games (
    id CHAR(24) NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    screenshot VARCHAR(255) NOT NULL,
    game_file VARCHAR(255) NOT NULL,
    frame_width integer NOT NULL,
    frame_height integer NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY (id)
);
