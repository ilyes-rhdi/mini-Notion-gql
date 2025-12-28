package services

import (
	"github.com/ilyes-rhdi/buildit-Gql/internal/database"
	"gorm.io/gorm"
)

// getDB évite le piège des variables globales initialisées avant database.InitDB()
func getDB() *gorm.DB {
	return database.GetDB()
}
