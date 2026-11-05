package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	goose "github.com/pressly/goose/v3"
)

func (srv *Server) HandlerResetDatabase(w http.ResponseWriter, r *http.Request) {
	providedKey := r.Header.Get("X-Admin-Key")
	if providedKey != config.Cfg.AdminAPIKey {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if config.Cfg.Platform != "dev" {
		http.Error(w, "Reset is only allowed in dev environment", http.StatusForbidden)
		return
	}

	goose.SetDialect("postgres")

	err := goose.DownTo(config.Cfg.DBConn, "./sql/schema", 0)
	if err != nil {
		http.Error(w, "Failed to reset database", http.StatusInternalServerError)
		return
	}

	err = goose.Up(config.Cfg.DBConn, "./sql/schema")
	if err != nil {
		http.Error(w, "Failed to rebuild database", http.StatusInternalServerError)
		return
	}

	// might simply put the seedtypes of the database here

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Database reset successfully"))
}
