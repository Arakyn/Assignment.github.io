package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
)

type Car struct {
	Car                                string
	MPG, Cylinders                     float64
	Displacement, Weight, Acceleration float64
	Horsepower                         float64
	Model                              int
	Origin                             string
}

const (
	Name int = iota
	MPG
	Cylinders
	Displacement
	Weight
	Acceleration
	Horsepower
	Model
	Origin
)

func main() {
	var lol [406]Car
	f, err := os.Open("bettercars.csv")
	if err != nil {
		log.Fatal("Error loading in the file from csv")
	}
	reader := csv.NewReader(bufio.NewReader(f))

	// discarding the header
	_, err = reader.Read()
	if err != nil {
		log.Fatal("Error at disposing the first line")
	}
	no := 0
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break

		}
		if err != nil {
			log.Fatal(err)
		}
		model, _ := strconv.Atoi(row[Model])
		mpg, _ := strconv.ParseFloat(row[MPG], 64)
		cyl, _ := strconv.ParseFloat(row[Cylinders], 64)
		disp, _ := strconv.ParseFloat(row[Displacement], 64)
		hp, _ := strconv.ParseFloat(row[Horsepower], 64)
		weight, _ := strconv.ParseFloat(row[Weight], 64)
		accl, _ := strconv.ParseFloat(row[Acceleration], 64)
		lol[no] = Car{
			Car:          row[Name],
			MPG:          mpg,
			Cylinders:    cyl,
			Displacement: disp,
			Horsepower:   hp,
			Weight:       weight,
			Acceleration: accl,
			Model:        model,
			Origin:       row[Origin],
		}
		no++

		jsonB, _ := json.Marshal(Car{
			Car:          row[Name],
			MPG:          mpg,
			Cylinders:    cyl,
			Displacement: disp,
			Horsepower:   hp,
			Weight:       weight,
			Acceleration: accl,
			Model:        model,
			Origin:       row[Origin],
		})

		fmt.Fprintf(os.Stdout, "%s\n", jsonB)

	}
	fmt.Println("Parsed the data into jsonB")
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("Printing the Data in a table form")
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(w, "Car \t\t\t\t\t\t\t\t MPG \t\t\t\t\t\t Cylinders \t\t\t\t\t\tDisplacement \t\t\t\t\t\t\t\t Horsepower \t\t\t\t\t\t\t\t Weight \t\t\t\t\t\t\t\t Acceleration\t\t\t\t\t\t\t\t Model \t\t\t\t\t\t\t\t Origin  ")
	for i := range lol {
		fmt.Fprintln(w, lol[i].Car, "\t\t\t\t\t\t\t\t", lol[i].MPG, "\t\t\t\t\t\t", lol[i].Cylinders, "\t\t\t\t\t\t", lol[i].Displacement, "\t\t\t\t\t\t\t\t", lol[i].Horsepower, "\t\t\t\t\t\t\t\t", lol[i].Weight, "\t\t\t\t\t\t\t\t", lol[i].Acceleration, "\t\t\t\t\t\t\t\t", lol[i].Model, "\t\t\t\t\t\t\t\t", lol[i].Origin)
	}
	w.Flush()
}
