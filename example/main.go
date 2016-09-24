package main

import (
	"encoding/csv"
	"github.com/ybping/metaheuristics/tsp"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input/china.csv")
	if err != nil {
		log.Panicln(err)
	}
	defer file.Close()
	var cities []tsp.City
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	for _, v := range records {
		lng, _ := strconv.ParseFloat(v[1], 64)
		lat, _ := strconv.ParseFloat(v[2], 64)
		city := tsp.City{
			Name: v[0],
			Lng:  lng,
			Lat:  lat,
		}
		cities = append(cities, city)
	}
	log.Println(cities)
}
