package main

import (
	"fmt"
	"math"
)

// 1. L'interface Payeur — le contrat commun à tous les modes de paiement

// Aucun type ne déclare "implements Payeur" : il suffit d'avoir une
// méthode Payer avec EXACTEMENT cette signature pour la satisfaire.
type Payeur interface {
	Payer(montant float64) (string, error)
}

// 2. CarteCredit
type CarteCredit struct {
	Numero    string
	Titulaire string
	Solde     float64
}

// Récepteur POINTEUR (*CarteCredit) car Payer modifie le Solde.
// Conséquence : c'est *CarteCredit (et non CarteCredit) qui est un Payeur.
func (cc *CarteCredit) Payer(montant float64) (string, error) {
	if montant > cc.Solde {
		// fmt.Errorf construit une valeur qui satisfait l'interface error.
		return "", fmt.Errorf("solde insuffisant : %.2f€ demandés, %.2f€ disponibles", montant, cc.Solde)
	}
	cc.Solde -= montant

	// On n'affiche que les 4 derniers chiffres (bonne pratique bancaire).
	quatreDerniers := cc.Numero
	if len(cc.Numero) >= 4 {
		quatreDerniers = cc.Numero[len(cc.Numero)-4:]
	}
	return fmt.Sprintf("Transaction CB #%s confirmée", quatreDerniers), nil
}

// 3. PayPal
type PayPal struct {
	Email string
	Solde float64
}

func (pp *PayPal) Payer(montant float64) (string, error) {
	if montant > pp.Solde {
		return "", fmt.Errorf("solde PayPal insuffisant : %.2f€ demandés, %.2f€ disponibles", montant, pp.Solde)
	}
	pp.Solde -= montant
	return fmt.Sprintf("Paiement PayPal de %.2f€ vers %s", montant, pp.Email), nil
}

// 4. Crypto
const tauxBTC = 50000.0 // 1 BTC = 50 000 €

type Crypto struct {
	Adresse string
	Solde   float64 // solde exprimé en euros
	Monnaie string
}

func (c *Crypto) Payer(montant float64) (string, error) {
	if montant > c.Solde {
		return "", fmt.Errorf("solde crypto insuffisant : %.2f€ demandés, %.2f€ disponibles", montant, c.Solde)
	}
	c.Solde -= montant

	// Conversion €→crypto, arrondie à 3 décimales (indice du sujet).
	montantCrypto := math.Round(montant/tauxBTC*1000) / 1000
	return fmt.Sprintf("Paiement de %.3f %s (%.2f€) vers %s", montantCrypto, c.Monnaie, montant, c.Adresse), nil
}

// Vérifications statiques à la compilation (bonne pratique du cours)
// Le compilateur refuse de compiler si un type ne satisfait plus Payeur.
var (
	_ Payeur = &CarteCredit{}
	_ Payeur = &PayPal{}
	_ Payeur = &Crypto{}
)

// 5. ProcesserPanier — fonction polymorphique
// Elle accepte N'IMPORTE QUEL Payeur. Le type switch sert uniquement à
// afficher le mode : la logique de paiement, elle, reste générique.
func ProcesserPanier(payeur Payeur, articles []float64) {
	total := 0.0
	for _, prix := range articles {
		total += prix
	}

	// Type switch : on inspecte le type concret rangé dans l'interface.
	switch p := payeur.(type) {
	case *CarteCredit:
		fmt.Printf("Mode : Carte crédit (%s)\n", p.Titulaire)
	case *PayPal:
		fmt.Printf("Mode : PayPal (%s)\n", p.Email)
	case *Crypto:
		fmt.Printf("Mode : Crypto %s\n", p.Monnaie)
	default:
		fmt.Println("Mode : inconnu")
	}

	fmt.Printf("Total du panier : %.2f€\n", total)

	recu, err := payeur.Payer(total)
	if err != nil {
		fmt.Println("❌ Échec :", err)
		return
	}
	fmt.Println("✅", recu)
}

// main -> démonstration
func main() {
	articles := []float64{29.99, 15.50, 4.90} // total = 50.39 €

	cb := &CarteCredit{Numero: "4539123456789012", Titulaire: "Alice Martin", Solde: 100}
	pp := &PayPal{Email: "bob@example.com", Solde: 40}
	btc := &Crypto{Adresse: "bc1qxy2k...", Solde: 200000, Monnaie: "BTC"}

	// Le même appel fonctionne pour les 3 types : c'est le polymorphisme.
	ProcesserPanier(cb, articles)
	fmt.Println()
	ProcesserPanier(pp, articles) // ← échouera : solde 40 < 50.39
	fmt.Println()
	ProcesserPanier(btc, articles)
}
