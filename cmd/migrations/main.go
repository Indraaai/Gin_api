package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	// 1. Load file .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Gagal membaca file .env")
	}

	// 2. Ambil DB_URL dari environment
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL tidak ditemukan di dalam .env")
	}

	// 3. Tangkap argumen dari terminal (misal: "up" atau "down")
	// Jika tidak ada argumen yang diberikan, default ke "up"
	action := "up"
	if len(os.Args) > 1 {
		action = os.Args[1]
	}

	log.Printf("Menjalankan migrasi database: %s...", action)

	// 4. Siapkan perintah CLI migrate
	cmd := exec.Command("migrate", "-path", "migrations", "-database", dbURL, action)

	// Arahkan output command agar tampil di terminal kita
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 5. Eksekusi perintah
	if err := cmd.Run(); err != nil {
		log.Fatalf("Migrasi gagal dijalankan: %v", err)
	}

	log.Println("Migrasi berhasil dieksekusi!")
}
