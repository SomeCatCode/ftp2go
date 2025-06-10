package main

import (
	"encoding/json"
	"log"
	"os"

	filedriver "github.com/goftp/file-driver"
	"github.com/goftp/server"
)

// Config definiert die Struktur der Konfigurationsdaten.
type Config struct {
	Root string `json:"root"`
	User string `json:"user"`
	Pass string `json:"pass"`
	Port int    `json:"port"`
	Host string `json:"host"`
}

func main() {
	configFile := "config.json"

	// Überprüfe, ob Config-Datei existiert
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		// Erstelle Config-Datei mit Standardwerten
		defaultConfig := Config{
			Root: "C://",
			User: "admin",
			Pass: "123456",
			Port: 2121,
			Host: "0.0.0.0",
		}

		data, err := json.MarshalIndent(defaultConfig, "", "  ")
		if err != nil {
			log.Fatalf("Fehler beim Erzeugen der Standard-Config: %v", err)
			return
		}

		// Schreibe die Standard-Konfiguration in die Datei.
		err = os.WriteFile(configFile, data, 0644)
		if err != nil {
			log.Fatalf("Fehler beim Schreiben der Standard-Config: %v", err)
			return
		}
	}

	// Lese Config-Datei
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Fehler beim Lesen der Config-Datei: %v", err)
		return
	}

	// Parse die Konfigurationsdaten in die Config-Struktur.
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Fehler beim Parsen der Config-Datei: %v", err)
		return
	}

	// Überprüfe, ob der angegebene Root-Pfad existiert und ein Verzeichnis ist.
	rootInfo, err := os.Stat(config.Root)
	if os.IsNotExist(err) {
		log.Fatalf("Der angegebene Root-Pfad existiert nicht: %v", config.Root)
		return
	} else if !rootInfo.IsDir() {
		log.Fatalf("Der angegebene Root-Pfad ist kein Verzeichnis: %v", config.Root)
		return
	}

	// Überprüfe die Schreibrechte im Root-Pfad.
	testFile := config.Root + "/test.txt"
	err = os.WriteFile(testFile, []byte("test"), 0644)
	if err != nil {
		log.Fatalf("Keine Schreibrechte im Root-Pfad: %v, Fehler: %v", config.Root, err)
		return
	}

	// Lösche die Testdatei, um zu bestätigen, dass keine unnötigen Dateien hinterlassen werden.
	err = os.Remove(testFile)
	if err != nil {
		log.Fatalf("Fehler beim Entfernen der Testdatei: %v", err)
		return
	}

	// Konfiguriere den FTP-Server mit den Einstellungen aus der Konfigurationsdatei.
	factory := &filedriver.FileDriverFactory{
		RootPath: config.Root,
		Perm:     server.NewSimplePerm("user", "group"),
	}

	opts := &server.ServerOpts{
		Factory:  factory,
		Port:     config.Port,
		Hostname: config.Host,
		Auth:     &server.SimpleAuth{Name: config.User, Password: config.Pass},
	}

	// Starte den FTP-Server.
	log.Printf("Starting ftp server on %v:%v", opts.Hostname, opts.Port)
	log.Printf("Username %v, Password %v", config.User, config.Pass)
	ftpServer := server.NewServer(opts)
	err = ftpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("Fehler beim Starten des FTP-Servers: %v", err)
	}
}
