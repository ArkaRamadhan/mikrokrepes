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
	if baseDSN == "" {
		log.Println("DB_URL is not set in the environment variables")
		return
	}

	// Menambahkan skema ke DSN
	dsn := fmt.Sprintf("%s search_path=%s", baseDSN, schema)
	log.Printf("Connecting to database with DSN: %s\n", dsn)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		log.Printf("Failed to connect to database with schema %s: %v\n", schema, err)
		// Mungkin tambahkan retry atau exit berdasarkan kebutuhan
		return
	}

	log.Println("Successfully connected to the database")
	// Mungkin tambahkan logika untuk memverifikasi koneksi dengan query sederhana jika perlu
}

func main() {
	ConnectToDB("your_schema_here")
}
