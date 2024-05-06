package Struct

// Appel de l'API contenant les adresses de 4 db et récupération des informations contenu dans le json
type Api struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

// Appel de l'API Artist et récupération des informations contenu dans le json
type Artists struct {
	Art_id           int      `json:"id"`
	Art_image        string   `json:"image"`
	Art_name         string   `json:"name"`
	Art_members      []string `json:"members"`
	Art_creationDate int      `json:"creationDate"`
	Art_firstAlbum   string   `json:"firstAlbum"`

	Art_locations        []string
	Art_nbMembers        int
	Art_AnneeFirst_Album int
}

// Appel de l'API Relation et récupération des informations contenu dans le json
type IndexRelation struct {
	Index []Relation `json:"index"`
}
type Relation struct {
	Rel_id             int                 `json:"id"`
	Rel_datesLocations map[string][]string `json:"datesLocations"`
}

// Appel de l'API Locations et récupération des informations contenu dans le json
type IndexLocations struct {
	Index []Locations `json:"index"`
}
type Locations struct {
	Loc_id        int      `json:"id"`
	Loc_locations []string `json:"locations"`
	Loc_dates     string   `json:"dates"`
}

// Appel de l'API Date et récupération des informations contenu dans le json
type IndexDates struct {
	Index []Dates `json:"index"`
}
type Dates struct {
	Dat_id    int      `json:"id"`
	Dat_dates []string `json:"dates"`
}

// Structure Contenant le retour de geocode
type Geocode struct {
	GPSresults []Geometry_geocode `json:"results"`
	GPSstatus  string             `json:"status"`
}

type Geometry_geocode struct {
	GPSgeometry Location_geocode `json:"geometry"`
}
type Location_geocode struct {
	GPSlocation Gps_geocode `json:"location"`
}
type Gps_geocode struct {
	GPSlat float64 `json:"lat"`
	GPSlng float64 `json:"lng"`
}

// structure globale qui rassemble les données de toutes les db
type Global struct {
	GArt_id           int
	GArt_image        string
	GArt_name         string
	GArt_members      []string
	GArt_creationDate int
	GArt_firstAlbum   string

	GRel_datesLocations map[string][]string

	GLoc_locations []string

	GDat_dates []string

	// Point GPS
	GPS map[string]string
}
