package GT

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	Struct "groupie-tracker-filters/packages/struct"
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

	// récupération des années de 1er Album (transformation de "02-02-2023" en "2023")
	for x := range ArtistAPI {
		date := ArtistAPI[x].Art_firstAlbum
		annee := 0
		count := 0
		for _, a := range date {
			if a == '-' {
				count++
			} else if count == 2 {
				annee = annee*10 + int(a-'0')
			}
		}
		ArtistAPI[x].Art_AnneeFirst_Album = annee
	}

	// récupération du nombre de membres en int dans chaque groupe
	for x := range ArtistAPI {
		count := 0
		for range ArtistAPI[x].Art_members {
			count++
		}
		ArtistAPI[x].Art_nbMembers = count
	}
	return erreur
}

// Ouverture de l'API contenant les liens
var AdresseAPI Struct.Api

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

// Ouverture de l'API Artist
var ArtistAPI []Struct.Artists

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

// Ouverture de l'API Relation
var RelationAPI Struct.IndexRelation

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

// Ouverture de l'API Locations
var LocationsAPI Struct.IndexLocations

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

// Ouverture de l'API dates
var DatesAPI Struct.IndexDates

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

// Conversion des coordonnées GPS
var Geocodeadress Struct.Geocode

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
