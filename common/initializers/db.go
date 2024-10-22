package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB(schema string) {
	var err error
	// Mengambil URL dasar dari environment
	baseDSN := os.Getenv("DB_URL")
	// Menambahkan skema ke DSN
	dsn := fmt.Sprintf("%s search_path=%s", baseDSN, schema)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		log.Fatalf("Failed to connect to database with schema %s: %v", schema, err)
	}
}
