package api

import (
	"net/http"

	goose "github.com/pressly/goose/v3"
)

func (cfg *Config) HandlerResetDatabase(w http.ResponseWriter, r *http.Request) {
	providedKey := r.Header.Get("X-Admin-Key")
	if providedKey != cfg.adminApiKey {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if cfg.platform != "dev" {
		http.Error(w, "Reset is only allowed in dev environment.", http.StatusForbidden)
		return
	}

	goose.SetDialect("postgres")

	err := goose.DownTo(cfg.dbConn, "./sql/schema", 0)
	if err != nil {
		http.Error(w, "Failed to reset database", http.StatusInternalServerError)
		return
	}

	err = goose.Up(cfg.dbConn, "./sql/schema")
	if err != nil {
		http.Error(w, "Failed to rebuild database", http.StatusInternalServerError)
		return
	}

	// might simply put the seeding of the database here

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Database reset successfully"))
}
