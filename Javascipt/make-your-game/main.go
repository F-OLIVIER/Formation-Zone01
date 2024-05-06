package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"text/template"
	"time"

	"github.com/gorilla/websocket"
)

type PlayerData struct {
	PlayerRank  int    `json:"playerRank"`
	PlayerName  string `json:"playerName"`
	PlayerScore int    `json:"playerScore"`
	PlayerTimer string `json:"playerTimer"`
}

var addr = flag.String("addr", ":5501", "http service address")

func main() {
	flag.Parse()
	http.HandleFunc("/", serveHome)
	http.Handle("/data/", http.StripPrefix("/data/", http.FileServer(http.Dir("./js/data"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css/"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js/"))))
	http.Handle("/sprites/", http.StripPrefix("/sprites/", http.FileServer(http.Dir("./sprites/"))))
	server := &http.Server{
		Addr:              *addr,
		ReadHeaderTimeout: 3 * time.Second,
	}
	fmt.Println("Server started on port 5501 : http://localhost:5501")

	// mise en écoute du WebSocket
	http.HandleFunc("/ws-endpoint", handleWebSocket)
	http.ListenAndServe(":5501", nil)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./index.html")
	if err != nil {
		log.Fatal(err)
	}
	ts.Execute(w, nil)
}

func handleScores(w http.ResponseWriter, r *http.Request, playerData PlayerData) {
	// fmt.Printf("\tlayerName  : %v\n\tPlayerLevel : %d\n\tScore : %d points\n\tTimer : %s s\n", playerData.PlayerName, playerData.PlayerLevel, playerData.PlayerScore, playerData.PlayerTimer)
	file, err := os.Open("./js/data/scores.json")
	if err != nil {
		http.Error(w, "Erreur lors de l'ouverture du fichier scores.json", http.StatusInternalServerError)
		return
	}

	var scores []PlayerData
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&scores)
	if err != nil {
		http.Error(w, "Erreur lors du décodage du fichier JSON", http.StatusInternalServerError)
		return
	}
	file.Close()

	if scores[len(scores)-1].PlayerScore < playerData.PlayerScore {
		scores = append(scores, playerData)
		sort.Slice(scores, func(i, j int) bool {
			return scores[i].PlayerScore > scores[j].PlayerScore
		})
		// plus de 10 dans le fichier ?
		if len(scores) > 10 {
			scores = scores[:10]
		}
		for i := range scores {
			scores[i].PlayerRank = i + 1
		}
	}

	file, err = os.OpenFile("./js/data/scores.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o600)
	if err != nil {
		http.Error(w, "Erreur lors de l'ouverture du fichier scores.json", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(scores)
	if err != nil {
		http.Error(w, "Erreur lors de l'encodage du fichier JSON", http.StatusInternalServerError)
		return
	}

	// http.Redirect(w, r, "/", http.StatusSeeOther)
}

// ---------------------------------------------
// ----------------- websocket -----------------
// ---------------------------------------------

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("err Upgrade handleWebSocket : ", err)
		return
	}
	defer conn.Close()

	// Lecture des données envoyées par le client
	_, message, err := conn.ReadMessage()
	if err != nil {
		fmt.Println("err message handleWebSocket : ", err)
		return
	}

	// Décoder les données JSON et les mettre dans PlayerData
	var playerData PlayerData
	err = json.Unmarshal(message, &playerData)
	if err != nil {
		fmt.Println("err Unmarshal handleWebSocket : ", err)
		return
	}

	handleScores(w, r, playerData)
}
