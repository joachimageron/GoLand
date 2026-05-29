package main

import "fmt"

// 1. Personne
type Personne struct {
	Prenom string
	Nom    string
	Age    int
	Email  string
}

func (p Personne) NomComplet() string {
	return p.Prenom + " " + p.Nom
}

func (p Personne) Presentation() string {
	return fmt.Sprintf("%s, %d ans (%s)", p.NomComplet(), p.Age, p.Email)
}

// 2. Adresse
type Adresse struct {
	Rue        string
	Ville      string
	CodePostal string
}

func (a Adresse) String() string {
	return fmt.Sprintf("%s, %s %s", a.Rue, a.CodePostal, a.Ville)
}

// 3. Employe (embedding Personne + Adresse)
type Employe struct {
	Personne
	Adresse
	Poste   string
	Salaire float64
}

func (e Employe) FicheEmploye() string {
	return fmt.Sprintf(
		"--- Fiche employé ---\n%s\nPoste : %s\nSalaire : %.2f €\nAdresse : %s",
		e.Presentation(), e.Poste, e.Salaire, e.Adresse.String(),
	)
}

func (e *Employe) AugmenterSalaire(pct float64) {
	e.Salaire += e.Salaire * pct / 100
}

// 4. Etudiant (embedding Personne)
type Etudiant struct {
	Personne
	Promo   string
	Moyenne float64
}

func (e Etudiant) MentionObtenue() string {
	switch {
	case e.Moyenne >= 16:
		return "Très bien"
	case e.Moyenne >= 14:
		return "Bien"
	case e.Moyenne >= 12:
		return "Assez bien"
	case e.Moyenne >= 10:
		return "Passable"
	default:
		return "Échec"
	}
}

func (e Etudiant) Fiche() string {
	return fmt.Sprintf(
		"--- Fiche étudiant ---\n%s\nPromo : %s\nMoyenne : %.2f (%s)",
		e.Presentation(), e.Promo, e.Moyenne, e.MentionObtenue(),
	)
}

// 5. main
func main() {
	emp := Employe{
		Personne: Personne{Prenom: "Alice", Nom: "Martin", Age: 32, Email: "alice.martin@entreprise.fr"},
		Adresse:  Adresse{Rue: "12 rue des Lilas", Ville: "Lyon", CodePostal: "69000"},
		Poste:    "Développeuse",
		Salaire:  42000,
	}

	etu1 := Etudiant{
		Personne: Personne{Prenom: "Bob", Nom: "Dupont", Age: 20, Email: "bob.dupont@univ.fr"},
		Promo:    "L3 Informatique",
		Moyenne:  16.5,
	}

	etu2 := Etudiant{
		Personne: Personne{Prenom: "Clara", Nom: "Petit", Age: 21, Email: "clara.petit@univ.fr"},
		Promo:    "M1 Réseaux",
		Moyenne:  11.2,
	}

	fmt.Println(emp.FicheEmploye())
	emp.AugmenterSalaire(10)
	fmt.Printf("Après augmentation : %.2f €\n\n", emp.Salaire)

	fmt.Println(etu1.Fiche())
	fmt.Println()
	fmt.Println(etu2.Fiche())
}
