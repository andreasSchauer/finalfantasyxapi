-- name: RefreshMonsterItemDropsView :exec
REFRESH MATERIALIZED VIEW mv_monster_item_drops;


-- name: RefreshMonsterEncountersView :exec
REFRESH MATERIALIZED VIEW mv_monster_encounters;


-- name: RefreshGeographyView :exec
REFRESH MATERIALIZED VIEW mv_geography;


-- name: RefreshItemSourcesView :exec
REFRESH MATERIALIZED VIEW mv_item_sources;