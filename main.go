package main

//Cliente API que consulta los últimos terremotos recientes
//En este caso, de magnitud 5 del 15 al 20 de noviembre del 2024

//Nota: Se puede cambiar las fechas únicamente cambiandolo directamente de la URL que fue proporcionada en la página

// Adrian Manuel Escogido Antonio
// Alberto Brenes Fernandez
// Topicos para el despliegue de aplicaciones

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Terremoto struct {
	Properties struct {
		Mag   float64 `json:"mag"`
		Place string  `json:"place"`
		Time  int64   `json:"time"`
	} `json:"properties"`
}

type Respuesta struct {
	Features []Terremoto `json:"features"`
}

func main() {
	// URL para consultar terremotos del día 15 a 20 de noviembre del 2024
	//url := "https://earthquake.usgs.gov/fdsnws/event/1/query?format=geojson&starttime=2024-10-15&endtime=2024-10-20"
	// Misma URL pero ahora se establece para que únicamente sea con una magnitud minima de 5
	url := "https://earthquake.usgs.gov/fdsnws/event/1/query?format=geojson&starttime=2024-10-15&endtime=2024-10-20&minmagnitude=5"

	// Hacemos la solicitud HTTP GET
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error al hacer la solicitud: %v", err)
	}
	defer resp.Body.Close()

	// Leemos y procesamos la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error al leer la respuesta: %v", err)
	}

	var datosTerremoto Respuesta
	err = json.Unmarshal(body, &datosTerremoto)
	if err != nil {
		log.Fatalf("Error al decodificar JSON: %v", err)
	}

	//Impresión de la información de cada terremoto que se saca de la página
	fmt.Println("Terremotos recientes:")
	for _, quake := range datosTerremoto.Features {
		fmt.Printf("Magnitud: %.1f, Lugar: %s, Fecha: %s\n",
			quake.Properties.Mag,
			quake.Properties.Place,
			time.UnixMilli(quake.Properties.Time).Format("2006-01-02 15:04:05"), //Dicha fecha es un Standard del formato de GO
		)
	}
}
