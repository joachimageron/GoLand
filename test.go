package main

import "fmt"
import "math"
import "time"

func main() {
	// affiche "Hello, World!" dans la console
    fmt.Println("Hello, World!")

	// affiche le résultat de 64 élevé à la puissance 8
    fmt.Println(math.Pow(64, 8))

	// affiche la date et l'heure actuelles
	fmt.Println(time.Now().Format("Today is Monday, January 2, 2006 and it's 3:04 PM"))

	// age calculé à partir de la date de naissance
	birthDate := time.Date(2003, time.May, 20, 0, 0, 0, 0, time.UTC)
	age := time.Now().Year() - birthDate.Year()
	fmt.Printf("Age calculated from birth date: %d years", age)

}