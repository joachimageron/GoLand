package main

import "fmt"
import "math"

func main() {
	var poids, taille = 70.5, 1.75
	imc := poids / math.Pow(taille, 2)

	const (
		IMCMaigre = 18.5
		IMCNormal = 25
		IMCSurpoids = 30
		Nom = "John Doe"
	)

	fmt.Printf("Bonjour %s, votre IMC est de %.2f\n", Nom, imc)

	if imc < IMCMaigre {
		fmt.Println("Vous êtes maigre.")
	} else if imc < IMCNormal {
		fmt.Println("Vous avez un poids normal.")
	} else if imc < IMCSurpoids {
		fmt.Println("Vous êtes en surpoids.")
	} else {
		fmt.Println("Vous êtes obèse.")
	}
}