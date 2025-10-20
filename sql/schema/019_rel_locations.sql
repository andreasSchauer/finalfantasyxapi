-- +goose Up
CREATE TYPE area_connection_type AS ENUM('both-directions', 'one-direction', 'warp');

CREATE TABLE area_connections (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    area_id INTEGER NOT NULL REFERENCES areas(id),
    connection_type area_connection_type NOT NULL,
    story_only BOOLEAN NOT NULL,
    notes TEXT
);


CREATE TABLE j_area_connection (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    area_id INTEGER NOT NULL REFERENCES areas(id),
    connection_id INTEGER NOT NULL REFERENCES area_connections(id)
);



-- +goose Down
DROP TABLE IF EXISTS j_area_connection;
DROP TABLE IF EXISTS area_connections;
DROP TYPE IF EXISTS area_connection_type;