package utils

import (
	"encoding/json"
	"fmt"
	data "forum/internal"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// ------------------------------------------------------------
// ------------ Fonction de gestion des pages html ------------
// ------------------------------------------------------------

// Il gère les messages d'erreur sur le site Web
func ErrorMessage(w http.ResponseWriter, r *http.Request, errType string) {
	UserData := &data.User{
		Logged: false,
		Msgerr: "",
	}
	switch errType {
	case "email":
		UserData.Msgerr = "Cet email est déjà utilisé."
	case "nomatch":
		UserData.Msgerr = "Les mots de passe ne correspondent pas."
	case "username":
		UserData.Msgerr = "Cet Username est déjà utilisé."
	case "form":
		UserData.Msgerr = "Veuillez rentrer tous les champs"
	case "cookieUsed":
		UserData.Msgerr = "Vous ne pouvez vous connecter qu'un seul appareil à la fois."
	case "badusername":
		UserData.Msgerr = "Caractére @ non autorisé dans le username"
	}

	data := &data.GlobalData{
		UserData: *UserData,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
	fmt.Println("ErrorMessage : ", errType)
}

// ------------------------------------------------------------
// --------------------- Fonction d'upload --------------------
// ------------------------------------------------------------

// Permet de charger en local un fichiers envoyé via un formulaire
func UploadFile(fileUpload multipart.File, header *multipart.FileHeader, filePath string) {
	fileLocal, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	CheckErr(err, "OS.Create file (UploadFile annexe)")
	io.Copy(fileLocal, fileUpload)
	fileLocal.Close()
}

// Permet de charger en local une images envoyé via un formulaire (Format supporté : jpeg/jpg, png, gif)
// https://riptutorial.com/go/example/31686/loading-and-saving-image
func UploadPicture(fileUpload multipart.File, header *multipart.FileHeader, filePath string) bool {
	typeFile1 := strings.ToLower(filePath[len(filePath)-3:])
	typeFile2 := strings.ToLower(filePath[len(filePath)-4:])

	if typeFile1 == "jpg" || typeFile2 == "jpeg" { // format jpeg/jpg
		// Décodage de l'image
		image, err := jpeg.Decode(fileUpload)
		// Création du fichier image
		fileLocal, err := os.Create(filePath)
		CheckErr(err, "OS.Create jpg/jpeg (UploadFile annexe)")
		// Spécifiez la qualité d'image, entre 0 et 100
		opt := jpeg.Options{
			Quality: 90,
		}
		// Sauvegarde de l'image en local
		err = jpeg.Encode(fileLocal, image, &opt)
		CheckErr(err, "Encode picture (Annexe)")
		fileLocal.Close()
		return true
	} else if typeFile1 == "png" { // format png
		// Décodage de l'image
		image, err := png.Decode(fileUpload)
		// Création du fichier image
		fileLocal, err := os.Create(filePath)
		CheckErr(err, "OS.Create png (UploadFile annexe)")
		// Sauvegarde de l'image en local
		err = png.Encode(fileLocal, image)
		CheckErr(err, "Encode picture (Annexe)")
		fileLocal.Close()
		return true
	} else if typeFile1 == "gif" { // format gif
		// Décodage de l'image
		image, err := gif.DecodeAll(fileUpload)
		// Création du fichier image
		fileLocal, err := os.Create(filePath)
		CheckErr(err, "OS.Create gif (UploadFile annexe)")
		// Sauvegarde de l'image en local
		err = gif.EncodeAll(fileLocal, image)
		CheckErr(err, "Encode picture (Annexe)")
		fileLocal.Close()
		return true
	}
	return false
}

// ------------------------------------------------------------
// --------------------- Fonction diverse --------------------
// ------------------------------------------------------------

// Vérificateur d'erreurs
func CheckErr(err error, str string) {
	if err != nil {
		fmt.Printf("ERROR : %v\n%v\n", str, err)
		// os.Exit(1)
	}
}

// Génére une chaine de caractére aéatoire de 12 caractéres alphanumérique
func RandomFileName() string {
	base := "azertyuiopmlkjhgfdsqwxcvbn-0123456789-AZERTYUIOPMLKJHGFDSQWXCVBN"
	var randomName string
	for i := 0; i < 12; i++ {
		num := rand.Intn(len(base))
		randomName += string(base[num])
	}
	return randomName + "_"
}

// Fonction qui formate la date en bon François
func GetTime() string {
	now := time.Now()
	y, m, d := now.Date()
	day := strconv.Itoa(d)
	year := strconv.Itoa(y)
	month := strconv.Itoa(int(m))
	date := day + "/" + month + "/" + year
	return date
}
