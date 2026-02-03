-- +goose Up
CREATE TYPE area_connection_type AS ENUM('both-directions', 'one-direction', 'warp');
CREATE TYPE shop_type AS ENUM ('pre-airship', 'post-airship');


CREATE TABLE area_connections (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    area_id INTEGER NOT NULL REFERENCES areas(id),
    connection_type area_connection_type NOT NULL,
    story_only BOOLEAN NOT NULL,
    notes TEXT
);


CREATE TABLE j_area_connected_areas (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    area_id INTEGER NOT NULL REFERENCES areas(id),
    connection_id INTEGER NOT NULL REFERENCES area_connections(id)
);


CREATE TABLE found_equipment_pieces (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    equipment_name_id INTEGER NOT NULL REFERENCES equipment_names(id),
    empty_slots_amount equipment_slots NOT NULL
);


CREATE TABLE j_found_equipment_abilities (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    found_equipment_id INTEGER NOT NULL REFERENCES found_equipment_pieces(id),
    auto_ability_id INTEGER NOT NULL REFERENCES auto_abilities(id)
);


CREATE TABLE j_treasures_items (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    treasure_id INTEGER NOT NULL REFERENCES treasures(id),
    item_amount_id INTEGER NOT NULL REFERENCES item_amounts(id)
);


CREATE TABLE shop_items (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    item_id INTEGER NOT NULL REFERENCES items(id),
    price INTEGER NOT NULL
);


CREATE TABLE shop_equipment_pieces (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    found_equipment_id INTEGER NOT NULL REFERENCES found_equipment_pieces(id),
    price INTEGER NOT NULL
);


CREATE TABLE j_shops_items (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    shop_id INTEGER NOT NULL REFERENCES shops(id),
    shop_item_id INTEGER NOT NULL REFERENCES shop_items(id),
    shop_type shop_type NOT NULL
);


CREATE TABLE j_shops_equipment (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    shop_id INTEGER NOT NULL REFERENCES shops(id),
    shop_equipment_id INTEGER NOT NULL REFERENCES shop_equipment_pieces(id),
    shop_type shop_type NOT NULL
);


ALTER TABLE treasures
ADD COLUMN found_equipment_id INTEGER REFERENCES found_equipment_pieces(id);



-- +goose Down
ALTER TABLE treasures
DROP COLUMN IF EXISTS found_equipment_id;

DROP TABLE IF EXISTS j_shops_equipment;
DROP TABLE IF EXISTS j_shops_items;
DROP TABLE IF EXISTS shop_equipment_pieces;
DROP TABLE IF EXISTS shop_items;
DROP TABLE IF EXISTS j_treasures_items;
DROP TABLE IF EXISTS j_found_equipment_abilities;
DROP TABLE IF EXISTS found_equipment_pieces;
DROP TABLE IF EXISTS j_area_connected_areas;
DROP TABLE IF EXISTS area_connections;
DROP TYPE IF EXISTS shop_type;
DROP TYPE IF EXISTS area_connection_type;