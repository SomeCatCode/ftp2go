package main

import (
	"encoding/json"
	"log"
	"os"

	filedriver "github.com/goftp/file-driver"
	"github.com/goftp/server"
)

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

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Fehler beim Parsen der Config-Datei: %v", err)
		return
	}

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

	log.Printf("Starting ftp server on %v:%v", opts.Hostname, opts.Port)
	log.Printf("Username %v, Password %v", config.User, config.Pass)

	ftpServer := server.NewServer(opts)
	err = ftpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("Fehler beim Starten des FTP-Servers: %v", err)
	}
}
