package utils

import (
	"database/sql"
	"fmt"
	data "forum/internal"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
)

// ------------------------------------------------------------
// ------------ Fonction de gestion des pages html ------------
// ------------------------------------------------------------

// Fonction qui permet de charger une page (gestion de données user simple en front) "tmplName = nom de la page"
func RenderTemplate(w http.ResponseWriter, tmplName string, td *data.DataTemplate) {
	tmpl, ok := template.ParseFiles("./templates/" + tmplName + ".page.tmpl")
	// fmt.Println("\n", ok)
	if ok != nil {
		http.Error(w, "No such template in here", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, td)
}

// Fonction qui permet de charger une page (gestion de données user complexe en front) "tmplName = nom de la page"
func UserRenderTemplate(w http.ResponseWriter, tmplName string, Check *data.GlobalData) {
	t, err := template.ParseFiles("./templates/" + tmplName + ".page.tmpl")
	CheckErr(err, "UserRenderTemplate dans Handlers")
	t.Execute(w, Check)
}

func HomeNotConnected(w http.ResponseWriter, r *http.Request, catFilter []string, database *sql.DB) {
	postData := DisplayForum(w, r, catFilter, database, "", "", "", "")
	userData := &data.User{
		StructCategories: AllColumOfCategory(database),
		Logged:           false,
	}
	data := &data.GlobalData{
		UserData:     *userData,
		PostListData: postData,
	}
	UserRenderTemplate(w, "Home", data)
}

// Il gère les messages d'erreur sur le site Web (temporaire)
func ErrorMessage(w http.ResponseWriter, r *http.Request, errType string) {
	switch errType {
	case "email":
		var data = &data.DataTemplate{
			Msgerr: "Cet email est déjà utilisé.<br>Déja membre ? <a href=\"/Login\">Se connecter</a>",
		}
		RenderTemplate(w, "Register", data)
	case "nomatch":
		var data = &data.DataTemplate{
			Msgerr: "Les mots de passe ne correspondent pas.",
		}
		RenderTemplate(w, "Register", data)
	case "username":
		var data = &data.DataTemplate{
			Msgerr: "Cet Username est déjà utilisé.",
		}
		RenderTemplate(w, "Register", data)
	case "form":
		var data = &data.DataTemplate{
			Msgerr: "Veuillez rentrer tous les champs",
		}
		RenderTemplate(w, "Register", data)
	case "cookieUsed":
		var data = &data.DataTemplate{
			Msgerr: "Vous ne pouvez vous connecter qu'un seul appareil à la fois.",
		}
		RenderTemplate(w, "Login", data)
	}
}

// ------------------------------------------------------------
// ------------------- Fonction de sécurité -------------------
// ------------------------------------------------------------

// Fonction de sécurité (vérification que les champs ne soient pas vides)(en cas d'utilisateur malveillant)
func CheckFormFull(r *http.Request) bool {
	for _, v := range r.Form {
		for _, a := range v {
			if len(a) == 0 {
				return false
			}
		}
	}
	return true
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
		image, err := gif.Decode(fileUpload)

		// Création du fichier image
		fileLocal, err := os.Create(filePath)
		CheckErr(err, "OS.Create gif (UploadFile annexe)")

		// Spécifiez la qualité d'image, entre 0 et 100
		opt := gif.Options{
			NumColors: 256,
		}

		// Sauvegarde de l'image en local
		err = gif.Encode(fileLocal, image, &opt)
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

// Fonction qui permet d'envoyer un email
// https://dev.to/go/sending-e-mail-from-gmail-using-golang-20bi
func SendMail(mailDestinataire, Subject, msgMail string) bool {
	// Sender data. boite mail sur https://my.mail.fr
	from := "couiierreburgond@mail.fr" // expéditeur
	password := "ArthourCouilliere"
	// smtp server configuration.
	smtpHost := "smtp.mail.fr" // serveur d'émission
	smtpPort := "587"          // 465"(SSL/TLS) ou "587"(STARTTLS/TLS) ou 25 (STARTTLS/TLS)

	// Création du contenu du mail (assemblage expéditeur, sujet, message, etc.).
	contenuMail := "From: " + from + "\n" +
		"To: " + mailDestinataire + "\n" +
		"Subject: " + Subject + "\r\n" +
		msgMail + "\r\n"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Connexion au serveur, authentification et expédition au destinataire,
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{mailDestinataire}, []byte(contenuMail))
	CheckErr(err, "send Mail")
	if err == nil {
		// fmt.Println("Email Sent Successfully!")
		return true
	} else {
		return false
	}
}

// Fonction qui formate la date en bon François
func GetTime() string {
	now := time.Now()
	y, m, d := now.Date()
	day := strconv.Itoa(d)
	year := strconv.Itoa(y)
	month := strconv.Itoa(int(m))
	date := day + "-" + month + "-" + year
	return date
}
