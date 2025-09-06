-- +goose Up
CREATE TYPE weapon_type AS ENUM ('sword', 'staff', 'blitzball', 'doll', 'spear', 'blade', 'claw', 'seymour-staff');
CREATE TYPE armor_type AS ENUM ('shield', 'ring', 'armguard', 'bangle', 'armlet', 'bracer', 'targe', 'seymour-armor');

CREATE TABLE characters (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    weapon_type weapon_type NOT NULL,
    armor_type armor_type NOT NULL,
    physical_attack_range INTEGER NOT NULL,
    can_fight_underwater BOOLEAN NOT NULL
);


CREATE TYPE aeon_category AS ENUM ('standard-aeons', 'magus-sisters');
CREATE TYPE accuracy_source AS ENUM ('accuracy', 'rate');

CREATE TABLE aeons (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    category aeon_category,
    is_optional BOOLEAN NOT NULL,
    battles_to_regenerate INTEGER NOT NULL,
    phys_atk_damage_constant INTEGER,
    phys_atk_range INTEGER,
    phys_atk_shatter_rate INTEGER,
    phys_atk_acc_source accuracy_source,
    phys_atk_hit_chance INTEGER,
    phys_atk_acc_modifier REAL
);


CREATE TABLE default_abilities (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL
);


-- +goose Down
DROP TABLE IF EXISTS default_abilities;
DROP TABLE IF EXISTS aeons;
DROP TYPE IF EXISTS aeon_category;
DROP TYPE IF EXISTS accuracy_source;
DROP TABLE IF EXISTS characters;
DROP TYPE IF EXISTS weapon_type;
DROP TYPE IF EXISTS armor_type;