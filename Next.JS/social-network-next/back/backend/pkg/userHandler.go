package pkg

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/user" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// On check dans la présence d'un param id dans l'url.

	switch r.Method {
	case "POST":

		// Retrieve user ID from session cookie
		cookie, err := r.Cookie("session")
		if err != nil {
			return
		}

		foundVal := cookie.Value
		curr, err := CurrentUser(foundVal)
		if err != nil {
			return
		}
		id := curr.Id

		fmt.Println("id current user:", id)

		if id != 0 {

			db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
			if err != nil {
				fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
			}
			// Toggle the privacy setting
			_, err = db.Exec("UPDATE USERS SET PrivateProfile = CASE WHEN PrivateProfile = 1 THEN 0 ELSE 1 END WHERE ID = ?", id)
			if err != nil {
				fmt.Println("Erreur lors de la mise à jour du profil utilisateur:", err)
				http.Error(w, "500 internal server error", http.StatusInternalServerError)
				return
			}

			// Retrieve the updated user information
			var user User
			err = db.QueryRow("SELECT ID, Email, FirstName, LastName, DateOfBirth, Avatar, Nickname, AboutMe, PrivateProfile FROM USERS WHERE ID = ?", id).Scan(&user.Id, &user.Email, &user.Firstname, &user.Lastname, &user.DateOfBirth, &user.Avatar, &user.Nickname, &user.AboutMe, &user.PrivateProfile)
			if err != nil {
				fmt.Println("Erreur lors de la récupération des données utilisateur:", err)
				http.Error(w, "500 internal server error", http.StatusInternalServerError)
				return
			}

			resp, err := json.Marshal(user)
			if err != nil {
				http.Error(w, "500 internal server error", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(resp)
		}

	case "GET":

		id := r.URL.Query().Get("id")

		if id == "" || id == "0" {
			db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
			if err != nil {
				fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
			}

			rows, err := db.Query("SELECT ID, Email, FirstName, LastName, DateOfBirth, Avatar, Nickname, AboutMe, PrivateProfile FROM USERS ORDER BY ID ASC")

			var users []User
			for rows.Next() {
				var user User
				err := rows.Scan(&user.Id, &user.Email, &user.Firstname, &user.Lastname, &user.DateOfBirth, &user.Avatar, &user.Nickname, &user.AboutMe, &user.PrivateProfile)
				if err != nil {
					break
				}
				// user.ListFollowings = ListFollow(user.Id, "followings", 0, db)
				// user.ListFollowers = ListFollow(user.Id, "followers", 0, db)
				users = append(users, user)
			}

			// On renvoit une array de structures users.
			resp, err := json.Marshal(users)
			if err != nil {
				http.Error(w, "500 internal server error", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(resp)
			// Sinon on recherche un user par son id.
		} else {

			db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
			if err != nil {
				fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
			}

			var user User
			var test string
			i, err := strconv.Atoi(id)
			err = db.QueryRow("SELECT ID, Email, FirstName, LastName, DateOfBirth, Avatar, Nickname, AboutMe, PrivateProfile FROM USERS WHERE ID = ?", i).Scan(&user.Id, &user.Email, &user.Firstname, &user.Lastname, &test, &user.Avatar, &user.Nickname, &user.AboutMe, &user.PrivateProfile)

			user.DateOfBirth = strings.Split(test, "T")[0]
			user.ListFollowings = ListFollow(user.Id, "followings", 0)
			user.ListFollowers = ListFollow(user.Id, "followers", 0)
			user.ListFollowersToValidate = ListFollow(user.Id, "followers", 1)

			// On renvoit la structure user recherchée.
			resp, err := json.Marshal(user)
			if err != nil {
				http.Error(w, "500 internal server error", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(resp)
		}

	default:
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

// Récupération de tous les users de la database.
func FindAllUsers() ([]User, error) {
	db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT * FROM USERS ORDER BY ID ASC`)
	if err != nil {
		return []User{}, errors.New("failed to find users")
	}

	users, err := ConvertRowToUser(rows)
	if err != nil {
		return []User{}, errors.New("failed to convert")
	}

	return users, nil
}

// Récupération d'un user en fonction d'un paramètre, exemple id.
func FindUserByParam(parameter string, data int) (User, error) {
	var q *sql.Rows
	db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
	}
	defer db.Close()
	// fmt.Println(reflect.TypeOf(data), data)

	switch parameter {
	case "id":
		// fmt.Println("look id")
		if err != nil {
			return User{}, errors.New("id must be an integer")
		}
		q, err = db.Query(`SELECT * FROM USERS WHERE ID = ?`, data)
		if err != nil {
			return User{}, errors.New("could not find id")
		}
	case "username":
		q, err = db.Query(`SELECT * FROM USERS WHERE Nickname = ?`, data)
		if err != nil {
			return User{}, errors.New("could not find username")
		}
	case "email":
		q, err = db.Query(`SELECT * FROM USERS WHERE Email = ?`, data)
		if err != nil {
			return User{}, errors.New("could not find email")
		}
	default:
		return User{}, errors.New("cannot search by that parameter")
	}

	user, err := ConvertRowToUser(q)
	if err != nil {
		fmt.Println("errconv")
		return User{}, errors.New("failed to convert")
	}
	fmt.Println(user)
	return user[0], nil
}

// Retourne l'user qui est loggé et dispose d'une session cookie valide.
func CurrentUser(val string) (User, error) {
	db, err := sql.Open("sqlite3", "backend/pkg/db/database.db")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
	}
	defer db.Close()

	q, err := db.Query(`SELECT UserID FROM SESSIONS WHERE SessionToken = ?`, val)
	if err != nil {
		fmt.Println("error")
		return User{}, err
	}

	users, err := ConvertRowToUser(q)
	if err != nil {
		return User{}, err
	}
	if len(users) == 0 {
		fmt.Println("no user")
		return User{}, errors.New("no user found")
	}
	return users[0], nil
}

// Mise en forme des rows en une array de structures User.
func ConvertRowToUser(rows *sql.Rows) ([]User, error) {
	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id)
		if err != nil {
			break
		}
		users = append(users, u)
	}
	return users, nil
}
