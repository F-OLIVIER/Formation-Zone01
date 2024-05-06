package lem_in

type Link struct {
	Current string
	Target  string
}

// --------- LISTS STRUCTS ------------- //

type Room struct {
	r_type string
	Name   string
	X      int
	Y      int
}

type Matrix struct {
	Name_X string
	Name_Y string
	Linked bool
}

// Struct pour la distribution des fourmis sur les routes
type Ant struct {
	Name        string // nom de la fourmis
	Room        string // Salle de depart de la fourmis
	Road        int    // Route emprunter par la fourmis
	Road_lenght int    // taille de la route
	Position    int    // position sur la route
	Arret       bool   // condition d'arrivé
}
