package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// DISCLAIMER : Ce code n'est pas uniquement de moi, aide sur git, GPT et tutoriels.

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

var newline = []byte{'\n'}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		return origin == "http://localhost:3000"
	},
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub        *Hub
	conn       *websocket.Conn // The websocket connection.
	send       chan []byte     // Buffered channel of outbound messages.
	userID     int             // The user id of the client
	typing     bool            // Track the typing status
	typingLock chan bool
	isReceiver bool // Track if the client is the receiver
}

// serveWs handles websocket requests from the peer.
func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	//fmt.Println("call ws")

	cookie, err := r.Cookie("session")
	if err != nil {
		return
	}

	foundVal := cookie.Value

	curr, err := CurrentUser(foundVal)
	if err != nil {
		return
	}

	client := &Client{
		hub:        hub,
		conn:       conn,
		send:       make(chan []byte, 256),
		userID:     curr.Id,
		typing:     false,
		typingLock: make(chan bool),
		isReceiver: false, // true or false based on your logic to determine if the client is the receiver,
	}

	//log.Println("Client isReceiver:", client.isReceiver)
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()

	// Start a goroutine to listen for typing events
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		//log.Printf("Received message from client: %s", string(message))

		var msg Message
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Printf("Error unmarshaling message: %v", err)
			break
		}

		msg.Sender_id = c.userID

		if msg.Msg_type == "msg" {
			msg.Date = time.Now().Format("01-02-2006 15:04:05")
			fmt.Println("mp")
			err = NewMessage(msg)
			if err != nil {
				log.Printf("Error storing new message: %v", err)
				break
			}

		}

		sendMsg, err := json.Marshal(msg)
		if err != nil {
			log.Printf("Error marshaling message: %v", err)
			break
		}

		//log.Printf(string(sendMsg))

		c.hub.broadcast <- sendMsg
	}

	// Stop typing when the readPump exits
	c.typingLock <- false
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
// writePump pumps messages from the hub to the websocket connection.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		//log.Print((c.send))
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub  the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			//log.Printf("Sended message from client: %s", string(message))

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
