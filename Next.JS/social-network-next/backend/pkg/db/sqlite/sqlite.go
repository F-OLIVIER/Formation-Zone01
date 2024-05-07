package sqlite

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func RunMigrations(db *sql.DB, migrationsDir string) error {
	// Récupère la liste des migrations
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".up.sql") {
			migrationPath := filepath.Join(migrationsDir, file.Name())
			query, err := os.ReadFile(migrationPath)
			if err != nil {
				return err
			}

			_, err = db.Exec(string(query))
			if err != nil {
				return err
			}

			log.Printf("Migration réussie: %s\n", file.Name())
		}
	}

	return nil
}
