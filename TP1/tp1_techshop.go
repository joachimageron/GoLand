package main

import (
	"errors"
	"fmt"
	"strings"
)

// ----- DONNÉES -----

type Produit struct {
	ID        int
	Nom       string
	Marque    string
	Prix      float64
	Stock     int
	Categorie string
	Actif     bool
}

type Catalogue struct {
	produits []Produit
}

// ----- FONCTIONNALITÉS (méthodes sur Catalogue) -----

// AjouterProduit ajoute un produit, erreur si l'ID existe déjà.
func (c *Catalogue) AjouterProduit(p Produit) error {
	for i := range c.produits {
		if c.produits[i].ID == p.ID {
			return fmt.Errorf("ID %d déjà utilisé", p.ID)
		}
	}
	c.produits = append(c.produits, p)
	return nil
}

// TrouverParID retourne le produit correspondant à l'ID.
func (c Catalogue) TrouverParID(id int) (Produit, error) {
	for i := range c.produits {
		if c.produits[i].ID == id {
			return c.produits[i], nil
		}
	}
	return Produit{}, fmt.Errorf("aucun produit avec l'ID %d", id)
}

// TrouverParCategorie retourne tous les produits d'une catégorie (insensible à la casse).
func (c Catalogue) TrouverParCategorie(cat string) []Produit {
	var resultat []Produit
	for i := range c.produits {
		if strings.EqualFold(c.produits[i].Categorie, cat) {
			resultat = append(resultat, c.produits[i])
		}
	}
	return resultat
}

// AppliquerReduction applique une réduction (en %) à une catégorie,
// retourne le nombre de produits modifiés.
func (c *Catalogue) AppliquerReduction(categorie string, pct float64) int {
	n := 0
	for i := range c.produits {
		if strings.EqualFold(c.produits[i].Categorie, categorie) {
			c.produits[i].Prix -= c.produits[i].Prix * pct / 100
			n++
		}
	}
	return n
}

// Vendre réduit le stock d'un produit, erreur si stock insuffisant.
func (c *Catalogue) Vendre(id int, qte int) error {
	if qte <= 0 {
		return errors.New("la quantité doit être positive")
	}
	for i := range c.produits {
		if c.produits[i].ID == id {
			if c.produits[i].Stock < qte {
				return fmt.Errorf("stock insuffisant (%d disponible, %d demandé)",
					c.produits[i].Stock, qte)
			}
			c.produits[i].Stock -= qte
			return nil
		}
	}
	return fmt.Errorf("aucun produit avec l'ID %d", id)
}

// Rapport retourne un résumé : nombre de produits et valeur totale du stock.
func (c Catalogue) Rapport() string {
	var valeurTotale float64
	for i := range c.produits {
		valeurTotale += c.produits[i].Prix * float64(c.produits[i].Stock)
	}
	return fmt.Sprintf("Catalogue : %d produits | valeur totale du stock : %.2f €",
		len(c.produits), valeurTotale)
}

func (p Produit) String() string {
	etat := "actif"
	if !p.Actif {
		etat = "inactif"
	}
	return fmt.Sprintf("#%d %s (%s) — %.2f € — stock %d — %s [%s]",
		p.ID, p.Nom, p.Marque, p.Prix, p.Stock, p.Categorie, etat)
}

// ----- MAIN -----

func main() {
	cat := Catalogue{}
	produitsInitiaux := []Produit{
		{1, "iPhone 15 Pro", "Apple", 1229.00, 12, "Smartphone", true},
		{2, "MacBook Air M3", "Apple", 1499.00, 7, "Ordinateur", true},
		{3, "Galaxy S24", "Samsung", 899.00, 15, "Smartphone", true},
		{4, "ThinkPad X1", "Lenovo", 1799.00, 4, "Ordinateur", true},
		{5, "AirPods Pro", "Apple", 279.00, 30, "Accessoire", true},
	}
	for _, p := range produitsInitiaux {
		cat.AjouterProduit(p)
	}

	for {
		fmt.Println("\n===== TechShop =====")
		fmt.Println("[1] Ajouter  [2] Chercher  [3] Soldes  [4] Vendre  [5] Rapport  [0] Quitter")
		fmt.Print("Votre choix : ")

		var choix int
		if _, err := fmt.Scan(&choix); err != nil {
			fmt.Println("choix invalide")
			return
		}

		switch choix {
		case 1:
			ajouter(&cat)
		case 2:
			chercher(&cat)
		case 3:
			soldes(&cat)
		case 4:
			vendre(&cat)
		case 5:
			fmt.Println(cat.Rapport())
		case 0:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("option inconnue")
		}
	}
}

func ajouter(cat *Catalogue) {
	var p Produit
	fmt.Print("ID Nom Marque Prix Stock Categorie : ")
	if _, err := fmt.Scan(&p.ID, &p.Nom, &p.Marque, &p.Prix, &p.Stock, &p.Categorie); err != nil {
		fmt.Println("saisie invalide :", err)
		return
	}
	p.Actif = true
	if err := cat.AjouterProduit(p); err != nil {
		fmt.Println("erreur :", err)
		return
	}
	fmt.Println("produit ajouté :", p)
}

func chercher(cat *Catalogue) {
	fmt.Print("ID du produit : ")
	var id int
	if _, err := fmt.Scan(&id); err != nil {
		fmt.Println("saisie invalide")
		return
	}
	p, err := cat.TrouverParID(id)
	if err != nil {
		fmt.Println("erreur :", err)
		return
	}
	fmt.Println(p)
}

func soldes(cat *Catalogue) {
	fmt.Print("Catégorie et pourcentage de réduction : ")
	var categorie string
	var pct float64
	if _, err := fmt.Scan(&categorie, &pct); err != nil {
		fmt.Println("saisie invalide")
		return
	}
	n := cat.AppliquerReduction(categorie, pct)
	fmt.Printf("%d produit(s) modifié(s)\n", n)
}

func vendre(cat *Catalogue) {
	fmt.Print("ID et quantité : ")
	var id, qte int
	if _, err := fmt.Scan(&id, &qte); err != nil {
		fmt.Println("saisie invalide")
		return
	}
	if err := cat.Vendre(id, qte); err != nil {
		fmt.Println("erreur :", err)
		return
	}
	fmt.Println("vente effectuée")
}
