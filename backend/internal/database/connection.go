package database

import (
	"fmt"
	"log"
	"os"
	"github.com/ilyes-rhdi/buildit-Gql/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
func InitDB() {
  var err error

	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("DSN non défini dans .env")
	}

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données :", err)
	}

	modelsToMigrate := []interface{}{ 
		&models.User{},
		&models.Workspace{},
		&models.Page{},
		&models.WorkspaceMember{},
		&models.Block{},
	
	}
	for _, model := range modelsToMigrate {
        if err := db.AutoMigrate(model); err != nil {
            log.Printf(" Erreur migration pour %T : %v", model, err)
        } else {
            log.Printf(" Table migrée : %T", model)
        }
}
	if err != nil {
		log.Fatal("Erreur lors de l'exécution de AutoMigrate :", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Erreur lors de l'accès à la connexion SQL :", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Impossible de se connecter à la base de données :", err)
	}

	fmt.Println("Connexion à la base de données PostgreSQL réussie.")
}

// GetDB retourne la connexion à la base de données
func GetDB() *gorm.DB {
	return db
}

// CloseDB ferme la connexion à la base de données
func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Erreur lors de la fermeture de la base de données:", err)
	}
	sqlDB.Close()
}