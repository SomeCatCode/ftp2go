# FTP2Go

FTP2Go ist ein einfacher FTP-Server, der in Go geschrieben ist. Er verwendet die `goftp/server` Bibliothek und liest Konfigurationsdaten aus einer JSON-Datei. Wenn die Konfigurationsdatei nicht vorhanden ist, wird sie automatisch mit Standardwerten erstellt.

## Abhängigkeiten

- [goftp/server](https://github.com/goftp/server)
- [goftp/file-driver](https://github.com/goftp/file-driver)

Diese können mit `go get` installiert werden:

```bash
go get github.com/goftp/server
go get github.com/goftp/file-driver
```

## Verwendung
Nach dem Klonen des Repositories können Sie das Programm mit go run . ausführen oder eine ausführbare Datei mit go build erstellen.

## Konfiguration
Die Konfigurationsdatei config.json sollte im selben Verzeichnis wie die ausführbare Datei liegen. Sie sollte folgende Felder enthalten:

root: Das Wurzelverzeichnis für den FTP-Server
user: Benutzername für den FTP-Server
pass: Passwort für den FTP-Server
port: Port für den FTP-Server
host: Hostname für den FTP-Server
Beispiel:

```json
{
"root": "./",
"user": "admin",
"pass": "123456",
"port": 2121,
"host": "localhost"
}
```

## Lizenz
Dieses Projekt steht unter der MIT-Lizenz - siehe die LICENSE.md Datei für weitere Informationen.

