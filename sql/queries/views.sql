-- name: RefreshMonsterItemDropsView :exec
REFRESH MATERIALIZED VIEW mv_monster_item_drops;


-- name: RefreshMonsterEquipmentDropsView :exec
REFRESH MATERIALIZED VIEW mv_monster_equipment_drops;


-- name: RefreshMonsterEncountersView :exec
REFRESH MATERIALIZED VIEW mv_monster_encounters;


-- name: RefreshGeographyView :exec
REFRESH MATERIALIZED VIEW mv_geography;


-- name: RefreshGeographyGraphView :exec
REFRESH MATERIALIZED VIEW mv_geography_graph;


-- name: RefreshItemSourcesView :exec
REFRESH MATERIALIZED VIEW mv_item_sources;


-- name: RefreshEquipmentSourcesView :exec
REFRESH MATERIALIZED VIEW mv_equipment_sources;


-- name: RefreshAutoAbilitySourcesView :exec
REFRESH MATERIALIZED VIEW mv_auto_ability_sources;


-- name: RefreshAvailabilityView :exec
REFRESH MATERIALIZED VIEW mv_availabilities;


-- name: RefreshAbilitiesView :exec
REFRESH MATERIALIZED VIEW mv_abilities;