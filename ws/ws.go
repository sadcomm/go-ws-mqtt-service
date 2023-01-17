package ws

import (
    "log"
    "net/http"

    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { 
        return true 
    },
}


func wsEndpoint(w http.ResponseWriter, r *http.Request) {
    // Upgrade upgrades the HTTP server connection   to the WebSocket protocol.
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Print("upgrade failed: ", err)
        return
    }
    defer conn.Close()

    // Continuosly read and write message
    for {
        mt, message, err := conn.ReadMessage()
        if err != nil {
            log.Println("read failed:", err)
            break
        }
        err = conn.WriteMessage(mt, message)
        if err != nil {
            log.Println("write failed:", err)
            break
        }
    }
}

func testingPage(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./ui/test/websockets.html")
}

func SetupRoutes() {
    http.HandleFunc("/", testingPage);
    http.HandleFunc("/todo", wsEndpoint);

    http.ListenAndServe(":8080", nil)
}