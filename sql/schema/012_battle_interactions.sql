-- +goose Up
CREATE TYPE special_action_type AS ENUM ('bribe', 'steal-gil', 'steal-item', 'transfer-overdrive');


CREATE TYPE critical_type AS ENUM ('crit', 'crit+ability%', 'crit+weapon%');


CREATE TYPE break_dmg_lmt_type AS ENUM ('always', 'auto-ability');


CREATE TYPE acc_source_type AS ENUM ('accuracy', 'rate');


CREATE TYPE attack_type AS ENUM ('attack', 'heal', 'absorb');


CREATE TYPE damage_type AS ENUM ('physical', 'magical', 'special');


CREATE TYPE damage_formula AS ENUM ('str-vs-def', 'str-ign-def', 'mag-vs-mdf', 'mag-ign-mdf', 'percentage-current', 'percentage-max', 'healing', 'special-no-var', 'special-var', 'special-magic', 'special-gil', 'special-kills', 'special-9999', 'fixed-9999', 'user-max-hp', 'swallowed-a', 'swallowed-b');


CREATE TYPE duration_type AS ENUM ('blocks', 'endless', 'instant', 'turns', 'user-turns', 'auto');


CREATE TYPE ctb_attack_type AS ENUM ('attack', 'heal');


CREATE TYPE delay_type AS ENUM ('ctb-based', 'tick-speed-based');


CREATE TYPE calculation_type AS ENUM ('added-percentage', 'added-value', 'multiply', 'multiply-highest', 'set-value');


CREATE TABLE damages (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    critical critical_type,
    critical_plus_val INTEGER,
    is_piercing BOOLEAN NOT NULL,
    break_dmg_limit break_dmg_lmt_type,
    element_id INTEGER REFERENCES elements(id)
);


CREATE TABLE ability_damages (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    condition TEXT,
    attack_type attack_type NOT NULL,
    stat_id INTEGER NOT NULL,
    damage_type damage_type NOT NULL,
    damage_formula damage_formula NOT NULL,
    damage_constant uint8 NOT NULL,
    CHECK(stat_id IN (1, 2))
);


CREATE TABLE ability_accuracies (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    acc_source acc_source_type NOT NULL,
    hit_chance uint8,
    acc_modifier REAL
);


CREATE TABLE battle_interactions (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    target target_type NOT NULL,
    based_on_phys_attack BOOLEAN NOT NULL,
    range distance,
    shatter_rate uint8,
    accuracy_id INTEGER NOT NULL REFERENCES ability_accuracies(id),
    hit_amount INTEGER NOT NULL,
    special_action special_action_type
);


CREATE TABLE inflicted_statusses (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    status_condition_id INTEGER NOT NULL REFERENCES status_conditions(id),
    probability uint8 NOT NULL,
    duration_type duration_type NOT NULL,
    amount INTEGER
);


CREATE TABLE inflicted_delays (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    condition TEXT,
    ctb_attack_type ctb_attack_type NOT NULL,
    delay_type delay_type NOT NULL,
    damage_constant uint8 NOT NULL
);


CREATE TABLE stat_changes (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    stat_id INTEGER NOT NULL REFERENCES stats(id),
    calculation_type calculation_type NOT NULL,
    value REAL NOT NULL
);


CREATE TABLE modifier_changes (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data_hash TEXT UNIQUE NOT NULL,
    modifier_id INTEGER NOT NULL REFERENCES modifiers(id),
    calculation_type calculation_type NOT NULL,
    value REAL NOT NULL
);


-- +goose Down
DROP TABLE IF EXISTS modifier_changes;
DROP TABLE IF EXISTS stat_changes;
DROP TABLE IF EXISTS inflicted_delays;
DROP TABLE IF EXISTS inflicted_statusses;
DROP TABLE IF EXISTS battle_interactions;
DROP TABLE IF EXISTS ability_accuracies;
DROP TABLE IF EXISTS ability_damages;
DROP TABLE IF EXISTS damages;
DROP TYPE IF EXISTS calculation_type;
DROP TYPE IF EXISTS delay_type;
DROP TYPE IF EXISTS ctb_attack_type;
DROP TYPE IF EXISTS duration_type;
DROP TYPE IF EXISTS damage_formula;
DROP TYPE IF EXISTS damage_type;
DROP TYPE IF EXISTS attack_type;
DROP TYPE IF EXISTS acc_source_type;
DROP TYPE IF EXISTS break_dmg_lmt_type;
DROP TYPE IF EXISTS critical_type;
DROP TYPE IF EXISTS special_action_type;