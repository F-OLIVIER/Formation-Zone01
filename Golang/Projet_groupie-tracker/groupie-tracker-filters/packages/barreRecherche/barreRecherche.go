package annexe

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	GT "groupie-tracker-filters/packages/api"
	Struct "groupie-tracker-filters/packages/struct"
)

var Globalgroupe Struct.Global

func RequestGroupe(id int) error {
	Globalgroupe.GArt_id = id
	// récupération des données dans la bd artist
	for a := range GT.ArtistAPI {
		if GT.ArtistAPI[a].Art_id == id {
			Globalgroupe.GArt_image = GT.ArtistAPI[a].Art_image
			Globalgroupe.GArt_name = GT.ArtistAPI[a].Art_name
			Globalgroupe.GArt_members = GT.ArtistAPI[a].Art_members
			Globalgroupe.GArt_creationDate = GT.ArtistAPI[a].Art_creationDate
			Globalgroupe.GArt_firstAlbum = GT.ArtistAPI[a].Art_firstAlbum
		}
	}
	// récupération des données dans la bd relation
	for a := range GT.RelationAPI.Index {
		if GT.RelationAPI.Index[a].Rel_id == id {
			Globalgroupe.GRel_datesLocations = GT.RelationAPI.Index[a].Rel_datesLocations

		}
	}
	// récupération des données dans la bd location
	for a := range GT.LocationsAPI.Index {
		if GT.LocationsAPI.Index[a].Loc_id == id {
			Globalgroupe.GLoc_locations = GT.LocationsAPI.Index[a].Loc_locations
		}
	}
	// récupération des données dans la bd date
	for a := range GT.DatesAPI.Index {
		if GT.DatesAPI.Index[a].Dat_id == id {
			Globalgroupe.GDat_dates = GT.DatesAPI.Index[a].Dat_dates
		}
	}
	// Transformation en coordonnée GPS
	var map_temp = make(map[string]string)
	for _, adress := range Globalgroupe.GLoc_locations {
		// Gestion d'erreur du retour google geocode
		err := GT.RequestAdress(adress)
		if GT.Geocodeadress.GPSstatus != "OK" {
			fmt.Println("GPSstatus : ", GT.Geocodeadress.GPSstatus)
			return errors.New("Erreur 500: Geocode API Google")
		}
		if err != nil {
			fmt.Println("Err API Google : ", err)
			return err
		}
		lat := fmt.Sprintf("%f", GT.Geocodeadress.GPSresults[0].GPSgeometry.GPSlocation.GPSlat)
		lng := fmt.Sprintf("%f", GT.Geocodeadress.GPSresults[0].GPSgeometry.GPSlocation.GPSlng)
		CoordonnéGPS := lat + "," + lng
		map_temp[adress] = CoordonnéGPS
	}
	Globalgroupe.GPS = map_temp
	return nil
}

// Gestion des recherches
var SearchArtists []Struct.Artists

func removesearch(slice []Struct.Artists) {
	var tempSearchArtists []Struct.Artists
	for a := range slice {
		if slice[a].Art_name != "sup" {
			tempSearchArtists = append(tempSearchArtists, slice[a])
		}
	}
	SearchArtists = tempSearchArtists
}

func RequestAnnees(annee_min int, annee_max int, rech_creation string, rech_album string) {
	SearchArtists = GT.ArtistAPI
	// Gestion de l'inversion années min et max
	var min int
	var max int
	if annee_min >= annee_max {
		min = annee_max
		max = annee_min
	} else {
		min = annee_min
		max = annee_max
	}

	if rech_creation == "on" {
		// Boucle de selection par années de création
		for i := range SearchArtists {
			if SearchArtists[i].Art_creationDate >= min && SearchArtists[i].Art_creationDate <= max {
				// ne rien faire
			} else {
				SearchArtists[i].Art_name = "sup"
			}
		}
	}

	if rech_album == "on" {
		// Boucle de selection par 1er album
		for i := range SearchArtists {
			if SearchArtists[i].Art_AnneeFirst_Album >= min && SearchArtists[i].Art_AnneeFirst_Album <= max {
				// ne rien faire
			} else {
				SearchArtists[i].Art_name = "sup"
			}
		}
	}
	removesearch(SearchArtists)
}

func Requestnbmembers(rech_nbmember1 string, rech_nbmember2 string, rech_nbmember3 string, rech_nbmember4 string, rech_nbmember5 string, rech_nbmember6 string, rech_nbmember7 string, rech_nbmember8 string) {
	if rech_nbmember1 == "" && rech_nbmember2 == "" && rech_nbmember3 == "" && rech_nbmember4 == "" && rech_nbmember5 == "" && rech_nbmember6 == "" && rech_nbmember7 == "" && rech_nbmember8 == "" {
		return
	}

	// Boucle de selection si nb members = ...
	for i := range SearchArtists {
		if rech_nbmember1 != "on" {
			if SearchArtists[i].Art_nbMembers == 1 {
				SearchArtists[i].Art_name = "sup"
			}
		}
		if rech_nbmember2 != "on" {
			if SearchArtists[i].Art_nbMembers == 2 {
				SearchArtists[i].Art_name = "sup"
			}
		}
		if rech_nbmember3 != "on" {
			if SearchArtists[i].Art_nbMembers == 3 {
				SearchArtists[i].Art_name = "sup"
			}
		}
		if rech_nbmember4 != "on" {
			if SearchArtists[i].Art_nbMembers == 4 {
				SearchArtists[i].Art_name = "sup"
			}
		}
		if rech_nbmember5 != "on" {
			if SearchArtists[i].Art_nbMembers == 5 {
				SearchArtists[i].Art_name = "sup"
			}
		}
		if rech_nbmember6 != "on" {
			if SearchArtists[i].Art_nbMembers == 6 {
				SearchArtists[i].Art_name = "sup"
			}
		}
		if rech_nbmember7 != "on" {
			if SearchArtists[i].Art_nbMembers == 7 {
				SearchArtists[i].Art_name = "sup"
			}
		}
		if rech_nbmember8 != "on" {
			if SearchArtists[i].Art_nbMembers == 8 {
				SearchArtists[i].Art_name = "sup"
			}
		}
	}
	removesearch(SearchArtists)
}

func BarreRecherche(recherche string) {
	// mise en minuscule de tout les caractéres
	tmp := ""
	for _, a := range recherche {
		if a >= 'A' && a <= 'Z' {
			tmp = tmp + string(a+32)
		} else if a == '-' || a == '_' {
			tmp = tmp + " "
		} else {
			tmp = tmp + string(a)
		}
	}
	recherche = tmp
	arrayRecherche := SplitWhiteSpaces(recherche)

	for i := range SearchArtists {
		// Les caractéres en string
		stringachercher := ""
		stringachercher = stringachercher + SearchArtists[i].Art_name + " "

		for _, j := range SearchArtists[i].Art_members {
			stringachercher = stringachercher + j + " "
		}

		var temp string
		// Boucle de gestion des localisations (ex : north_carolina-usa)
		for _, j := range SearchArtists[i].Art_locations {
			for i := 0; i < len(j); i++ {
				if string(j[i]) == "-" {
					temp = temp + " "
				} else if string(j[i]) == "_" {
					temp = temp + " "
				} else {
					temp = temp + string(j[i])
				}
			}
			temp = temp + " "
		}
		stringachercher = stringachercher + temp

		// Boucle de gestion des date de first album (ex : 05-08-1967)
		temp = ""
		for _, j := range SearchArtists[i].Art_firstAlbum {
			if string(j) == "-" {
				temp = temp + " "
			} else {
				temp = temp + string(j)
			}
		}
		stringachercher = stringachercher + temp + " "

		// Les caractéres en int
		stringachercher = stringachercher + strconv.Itoa(SearchArtists[i].Art_creationDate)

		// mise en minuscule de tout les caractéres
		tmp := ""
		for _, a := range stringachercher {
			if a >= 'A' && a <= 'Z' {
				tmp = tmp + string(a+32)
			} else if a == '-' || a == '_' {
				tmp = tmp + " "
			} else {
				tmp = tmp + string(a)
			}
		}
		stringachercher = tmp

		// StringaChercher = [Name, Members, Locations, string(creation), string(Album)]
		arrayDeStringaChercher := SplitWhiteSpaces(stringachercher)
		count := 0
		for _, k := range arrayRecherche {
			for _, l := range arrayDeStringaChercher {
				if strings.Contains(l, k) == true {
					count++
					break
				}
			}

			if k == "groupe" || k == "membre" || k == "creation" || k == "premier" || k == "album" || k == "concert" {
				count++
			}
		}
		// if count > 0 {
		// 	fmt.Println()
		// 	fmt.Println("recherche : ", recherche)
		// 	fmt.Println("count : ", count)
		// 	fmt.Println("stringachercher : ", stringachercher)
		// 	fmt.Println("arrayRecherche : ", arrayRecherche)
		// 	fmt.Println("arrayDeStringaChercher : ", arrayDeStringaChercher)
		// }

		if count != len(arrayRecherche) {
			SearchArtists[i].Art_name = "sup"
		}
	}
	removesearch(SearchArtists)
}

func SplitWhiteSpaces(s string) []string {
	var result []string
	var temp string
	for i := 0; i < len(s); i++ {
		if string(s[i]) != " " {
			if string(s[i]) != "," {
				if string(s[i]) != "-" {
					if string(s[i]) != "_" {
						if string(s[i]) != "." {
							if string(s[i]) != "(" {
								if string(s[i]) != ")" {
									temp = temp + string(s[i])
								}
							}
						}
					}
				}
			}
		} else {
			result = append(result, temp)
			temp = ""
		}
	}
	result = append(result, temp)
	return result
}
