package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/go-sql-driver/mysql"
)

// DB adalah koneksi global yang dipakai di seluruh aplikasi.
var DB *sql.DB

// InitDB membuka koneksi ke MySQL dan menjalankan migrasi.
// Dipanggil sekali saat aplikasi start di main.go.
func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("[DB] gagal membuka koneksi: %v", err)
	}

	// Connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(2 * time.Minute)

	// Pastikan koneksi benar-benar berhasil
	if err := db.Ping(); err != nil {
		log.Fatalf("[DB] gagal ping database: %v", err)
	}

	log.Println("[DB] koneksi berhasil")

	DB = db

	RunMigrations(db)
}

// RunMigrations menjalankan semua file migrasi yang belum diapply.
// File migrasi dibaca dari folder ./migrations.
func RunMigrations(db *sql.DB) {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("[MIGRATE] gagal membuat driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatalf("[MIGRATE] gagal inisialisasi migrate: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("[MIGRATE] gagal menjalankan migrasi: %v", err)
	}

	log.Println("[MIGRATE] migrasi selesai")
}