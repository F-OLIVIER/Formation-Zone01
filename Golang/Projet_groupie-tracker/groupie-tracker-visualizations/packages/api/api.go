package GT

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Fonction qui ouvre toutes les API pour mettre a jours les données
// Si l'ouverture echoue, retour d'une erreur
func Appeldb() (erreur bool) {
	erreur = false
	err := AppelAPI()
	if err != nil {
		fmt.Println("err_appel_api : ", err)
		erreur = true
	}
	// fmt.Println("Api : ", AdresseAPI)

	err0 := RequestDates(AdresseAPI.Dates)
	if err0 != nil {
		fmt.Println("err_dates : ", err0)
		erreur = true
	}
	// fmt.Println("DatesAPI : ", GT.DatesAPI)

	err1 := RequestLocations(AdresseAPI.Locations)
	if err1 != nil {
		fmt.Println("err_locations : ", err1)
		erreur = true
	}
	// fmt.Println("DatesAPI : ", GT.LocationsAPI)

	err2 := RequestRelation(AdresseAPI.Relation)
	if err2 != nil {
		fmt.Println("err_relation : ", err2)
		erreur = true
	}
	// fmt.Println("RelationAPI : ", GT.RelationAPI)

	err3 := RequestArtist(AdresseAPI.Artists)
	if err3 != nil {
		fmt.Println("err_Artist : ", err3)
		erreur = true
	}

	// récupération des location dans la bd relation
	for x := range ArtistAPI {
		for a := range LocationsAPI.Index {
			if ArtistAPI[x].Art_id == LocationsAPI.Index[a].Loc_id {
				ArtistAPI[x].Art_locations = LocationsAPI.Index[a].Loc_locations
			}
		}
	}

	// fmt.Println("ArtistAPI : ", GT.ArtistAPI[1])
	return erreur
}

// Appel de l'API contenant les adresses de 4 db et récupération des informations contenu dans le json
type Api struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

var AdresseAPI Api

func AppelAPI() error {
	data, err1 := http.Get("https://groupietrackers.herokuapp.com/api")
	if err1 != nil {
		return err1
	}
	defer data.Body.Close()
	All_data, err2 := io.ReadAll(data.Body)
	if err2 != nil {
		return err2
	}
	err3 := json.Unmarshal(All_data, &AdresseAPI)
	if err3 != nil {
		return err3
	}
	return nil
}

// Appel de l'API Artist et récupération des informations contenu dans le json
type Artists struct {
	Art_id           int      `json:"id"`
	Art_image        string   `json:"image"`
	Art_name         string   `json:"name"`
	Art_members      []string `json:"members"`
	Art_creationDate int      `json:"creationDate"`
	Art_firstAlbum   string   `json:"firstAlbum"`
	Art_locations    []string
}

var ArtistAPI []Artists

func RequestArtist(adresse string) error {
	data, err1 := http.Get(adresse)
	if err1 != nil {
		return err1
	}
	defer data.Body.Close()
	data_artist, err2 := io.ReadAll(data.Body)
	if err2 != nil {
		return err2
	}
	err3 := json.Unmarshal(data_artist, &ArtistAPI)
	if err3 != nil {
		return err3
	}
	return nil
}

// Appel de l'API Relation et récupération des informations contenu dans le json
type IndexRelation struct {
	Index []Relation `json:"index"`
}
type Relation struct {
	Rel_id             int                 `json:"id"`
	Rel_datesLocations map[string][]string `json:"datesLocations"`
}

var RelationAPI IndexRelation

func RequestRelation(adresse string) error {
	data, err1 := http.Get(adresse)
	if err1 != nil {
		return err1
	}
	defer data.Body.Close()
	data_relation, err2 := io.ReadAll(data.Body)
	if err2 != nil {
		return err2
	}
	err3 := json.Unmarshal(data_relation, &RelationAPI)
	if err3 != nil {
		return err3
	}
	return nil
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

var LocationsAPI IndexLocations

func RequestLocations(adresse string) error {
	data, err1 := http.Get(adresse)
	if err1 != nil {
		return err1
	}
	defer data.Body.Close()
	data_locations, err2 := io.ReadAll(data.Body)
	if err2 != nil {
		return err2
	}
	err3 := json.Unmarshal(data_locations, &LocationsAPI)
	if err3 != nil {
		return err3
	}
	return nil
}

// Appel de l'API Date et récupération des informations contenu dans le json
type IndexDates struct {
	Index []Dates `json:"index"`
}
type Dates struct {
	Dat_id    int      `json:"id"`
	Dat_dates []string `json:"dates"`
}

var DatesAPI IndexDates

func RequestDates(adresse string) error {
	data, err1 := http.Get(adresse)
	if err1 != nil {
		return err1
	}
	defer data.Body.Close()
	data_date, err2 := io.ReadAll(data.Body)
	if err2 != nil {
		return err2
	}
	err3 := json.Unmarshal(data_date, &DatesAPI)
	if err3 != nil {
		return err3
	}
	return nil
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

var Geocodeadress Geocode

func RequestAdress(adresse string) error {
	localisation, err1 := http.Get("https://maps.googleapis.com/maps/api/geocode/json?new_forward_geocoder=true&address=" + adresse + "&key=AIzaSyAlN9UWYxSRnlX-eBh5umoXGD1V0DGzCHc")
	if err1 != nil {
		fmt.Println("Err geocode : ", err1)
		return err1
	}
	defer localisation.Body.Close()
	All_data, err2 := io.ReadAll(localisation.Body)
	if err2 != nil {
		fmt.Println("Err geocode : ", err2)
		return err2
	}
	err3 := json.Unmarshal(All_data, &Geocodeadress)
	if err3 != nil {
		fmt.Println("Err geocode : ", err3)
		return err3
	}
	return nil
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

var Globalgroupe Global

func RequestGroupe(id int) error {
	Globalgroupe.GArt_id = id
	// récupération des données dans la bd artist
	for a := range ArtistAPI {
		if ArtistAPI[a].Art_id == id {
			Globalgroupe.GArt_image = ArtistAPI[a].Art_image
			Globalgroupe.GArt_name = ArtistAPI[a].Art_name
			Globalgroupe.GArt_members = ArtistAPI[a].Art_members
			Globalgroupe.GArt_creationDate = ArtistAPI[a].Art_creationDate
			Globalgroupe.GArt_firstAlbum = ArtistAPI[a].Art_firstAlbum
		}
	}
	// récupération des données dans la bd relation
	for a := range RelationAPI.Index {
		if RelationAPI.Index[a].Rel_id == id {
			Globalgroupe.GRel_datesLocations = RelationAPI.Index[a].Rel_datesLocations

		}
	}
	// récupération des données dans la bd location
	for a := range LocationsAPI.Index {
		if LocationsAPI.Index[a].Loc_id == id {
			Globalgroupe.GLoc_locations = LocationsAPI.Index[a].Loc_locations
		}
	}
	// récupération des données dans la bd date
	for a := range DatesAPI.Index {
		if DatesAPI.Index[a].Dat_id == id {
			Globalgroupe.GDat_dates = DatesAPI.Index[a].Dat_dates
		}
	}
	// Transformation en coordonnée GPS
	// var map_temp = make(map[string]string)
	// for _, adress := range Globalgroupe.GLoc_locations {
	// 	// Gestion d'erreur du retour google geocode
	// 	err := RequestAdress(adress)
	// 	if Geocodeadress.GPSstatus != "OK" {
	// 		fmt.Println("GPSstatus : ", Geocodeadress.GPSstatus)
	// 		return errors.New("Erreur 500: Geocode API Google")
	// 	}
	// 	if err != nil {
	// 		fmt.Println("Err API Google : ", err)
	// 		return err
	// 	}
	// 	lat := fmt.Sprintf("%f", Geocodeadress.GPSresults[0].GPSgeometry.GPSlocation.GPSlat)
	// 	lng := fmt.Sprintf("%f", Geocodeadress.GPSresults[0].GPSgeometry.GPSlocation.GPSlng)
	// 	CoordonnéGPS := lat + "," + lng
	// 	map_temp[adress] = CoordonnéGPS
	// }
	// Globalgroupe.GPS = map_temp
	return nil
}

// Fonction de recherche de l'id a partir du nom du groupe
func RechercheID(name string) (id int) {
	for a := range ArtistAPI {
		if ArtistAPI[a].Art_name == name {
			id = ArtistAPI[a].Art_id
		}
	}
	return id
}
