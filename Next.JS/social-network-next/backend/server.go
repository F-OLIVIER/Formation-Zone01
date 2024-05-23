package backend

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"social/backend/pkg"
	"social/backend/pkg/db/sqlite"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func StartServer() {
	dbFile := filepath.Join("backend", "pkg", "db", "database.db")

	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		log.Println("La base de données n'existe pas, création en cours...")
		_, err := os.Create(dbFile)
		if err != nil {
			log.Fatal("Erreur lors de la création de la base de données:", err)
		}
	}

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal("Erreur lors de l'ouverture de la base de données:", err)
	}
	defer db.Close()

	migrationsDir := filepath.Join("backend", "pkg", "db", "migrations", "sqlite")

	err = sqlite.RunMigrations(db, migrationsDir)
	if err != nil {
		log.Fatal("Erreur lors de l'exécution des migrations:", err)
	}

	port := 8080

	r := mux.NewRouter()

	hub := pkg.NewHub()
	go hub.Run()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	r.HandleFunc("/login", pkg.LoginHandler)
	r.HandleFunc("/register", pkg.RegisterHandler)
	r.HandleFunc("/post", pkg.PostHandler)
	r.HandleFunc("/comment", pkg.CommentHandler)
	r.HandleFunc("/like", pkg.LikeHandler)
	r.HandleFunc("/user", pkg.UserHandler)
	r.HandleFunc("/chat", pkg.ChatHandler)
	r.HandleFunc("/message", pkg.MessageHandler)
	r.HandleFunc("/follow", pkg.FollowHandler)
	r.HandleFunc("/logout", pkg.LogoutHandler)
	r.HandleFunc("/session", pkg.SessionHandler)
	r.HandleFunc("/creategroup", pkg.CreateGroupHandler)
	r.HandleFunc("/getallgroups", pkg.GetAllGroups)
	r.HandleFunc("/inviteinmygroup", pkg.Inviteinmygroup)
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		pkg.ServeWs(hub, w, r)
	})

	// nécessaire pour éviter l'erreur CORS
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	credentials := handlers.AllowCredentials()

	server := &http.Server{
		Addr:    ":" + fmt.Sprint(port),
		Handler: handlers.CORS(headers, origins, methods, credentials)(r),
	}

	fmt.Println("Server listening on :8080...")
	log.Fatal(server.ListenAndServe())
}
