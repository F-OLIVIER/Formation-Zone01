package backend

import (
	"database/sql"
	"encoding/json"
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
	r.HandleFunc("/notif", pkg.NotifHandler)
	r.HandleFunc("/post", pkg.PostHandler)
	r.HandleFunc("/comment", pkg.CommentHandler)
	r.HandleFunc("/commentgroup", pkg.CommentGroupHandler)
	r.HandleFunc("/like", pkg.LikeHandler)
	r.HandleFunc("/user", pkg.UserHandler)
	r.HandleFunc("/chat", pkg.ChatHandler)
	r.HandleFunc("/message", pkg.MessageHandler)
	r.HandleFunc("/messagegroup", pkg.GroupMessageHandler)
	r.HandleFunc("/event", pkg.EventHandler)
	r.HandleFunc("/comingevent", pkg.ComingEventHandler)
	r.HandleFunc("/notcomingevent", pkg.NotComingEventHandler)
	r.HandleFunc("/follow", pkg.FollowHandler)
	r.HandleFunc("/logout", pkg.LogoutHandler)
	r.HandleFunc("/session", pkg.SessionHandler)
	r.HandleFunc("/creategroup", pkg.CreateGroupHandler)
	r.HandleFunc("/getallgroups", pkg.GetAllGroups)
	r.HandleFunc("/postgroup", pkg.PostGroupHandler)
	r.HandleFunc("/inviteinmygroup", pkg.Inviteinmygroup)
	r.HandleFunc("/getonegroup", pkg.GetOneGroup)
	r.HandleFunc("/accept-group-notification", pkg.AcceptGroupNotification)
	r.HandleFunc("/askforjoingroup", pkg.AskForJoinGroup)

	r.HandleFunc("/delete-notification", pkg.DeleteNotificationHandler)

	r.HandleFunc("/target", pkg.TargetHandler)

	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		pkg.ServeWs(hub, w, r)
	})
	
	endpoints := []string{
		"/crossplatform/login", "/crossplatform/getuser", "/crossplatform/getallmessage", "/crossplatform/listuser",
	}
	for _, endpoint := range endpoints {
		r.HandleFunc(endpoint, pkg.CrossPlatformHandler)
	}

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

func DeleteNotification(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data struct {
		NotificationID int `json:"notification_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ici, vous supprimez la notification avec l'ID fourni de votre base de données

	jsonResponse := map[string]interface{}{
		"success": true,
		"message": "Notification deleted successfully",
	}
	json.NewEncoder(w).Encode(jsonResponse)
}
