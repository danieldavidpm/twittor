package main

import (
	"log"

	"github.com/danieldavidpm/twittor/bd"
	"github.com/danieldavidpm/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()

}
