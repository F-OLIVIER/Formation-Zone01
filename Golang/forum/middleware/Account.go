package utils

import (
	"database/sql"
	"net/http"
	"strings"
)

func UpdateAccount(w http.ResponseWriter, r *http.Request, user_id, linkPhoto, username string, database *sql.DB) (string, string, string) {
	var msgUser string
	r.ParseForm()
	// Si deconnection demandé par l'utilisateur
	if r.FormValue("logout") == "logout" {
		Logout(w, r, database)
	}
	// chargement de l'image
	fileImage, header, errFormImage := r.FormFile("uploadImage")
	if errFormImage == nil { // Si pas d'erreur, alors c'est qu'il y a une nouvelle image de profil
		defer fileImage.Close()
		newNameFile := RandomFileName() + header.Filename
		uploaded := UploadPicture(fileImage, header, "./assets/photoCompte/"+newNameFile)
		if uploaded { // photo uploader, mise à jour de la db
			UpdateProfilPicture(user_id, newNameFile, database)
			linkPhoto = "/assets/photoCompte/" + newNameFile
		} else { // photo non uploader
			msgUser = username + " format d'image non supporté"
		}
	} else if r.FormValue("newUsername") != "" { // changement Username
		newUsername := r.FormValue("newUsername")
		uploaded := UpdateUsername(user_id, newUsername, database)
		if uploaded {
			username = newUsername
			msgUser = username + ",<br>Username mis à jour."
		} else {
			msgUser = username + ",<br>Le Username \"" + newUsername + "\" est déjà utilisé."
		}
	} else if r.FormValue("oldPass") != "" && r.FormValue("newPass") != "" && r.FormValue("confirmNewPass") != "" { // changement du mot de passe
		oldPass := r.FormValue("oldPass")
		newPass := r.FormValue("newPass")
		confirmNewPass := r.FormValue("confirmNewPass")
		mdpUpdated := UpdatePassword(user_id, oldPass, newPass, confirmNewPass, database)
		if mdpUpdated {
			msgUser = username + ",<br>votre mot de passe a bien été mis à jour"
		} else {
			msgUser = username + ",<br>Erreur de saisie, impossible de mettre à jour votre mot de passe"
		}
	} else if strings.ToUpper(r.FormValue("deleteAccount")) == "DELETE" { // supression du compte utilisateur
		// FIXME: apres suppression du compte, il faut soit :
		// 1- attribué au post de cette utilisateur le "user-deleted"
		// 2- gérer l'erreur généré au chargement d'un post ou d'un comment pour renvoyer les information de "user-deleted"
		deleted := DeleteAccount(user_id, database)
		if deleted {
			Logout(w, r, database)
			return "Compte supprimé avec succée", "", ""
		}
	} else if r.FormValue("deleteAccount") != "" {
		msgUser = username + ",<br>Erreur dans la suppression du compte &nbsp;: &nbsp;pour supprimer votre compte vous devez confirmer la demande de suppression en saisissant \"DELETE\""
	}
	return msgUser, linkPhoto, username
}
