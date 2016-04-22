// Basic websocket server, based on the code sample taken from
// https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/08.2.html
//
// Serves on port 1234
// The websocket handler listens on /ws
package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

// HTML content we're serving
var page = `
<html>
<head></head>
<body>
    <script type="text/javascript">
        var sock = null;
        var wsuri = "ws://127.0.0.1:1234/ws";

        window.onload = function() {

            console.log("onload");

            sock = new WebSocket(wsuri);

            sock.onopen = function() {
                console.log("connected to " + wsuri);
            }

            sock.onclose = function(e) {
                console.log("connection closed (" + e.code + ")");
            }

            sock.onmessage = function(e) {
                console.log("message received: " + e.data);
            }
        };

        function send() {
            var msg = document.getElementById('message').value;
            sock.send(msg);
        };
    </script>
    <h1>WebSocket Echo Test</h1>
    <form>
        <p>
            Message: <input id="message" type="text" value="Hello, world!">
        </p>
    </form>
    <button onclick="send();">Send Message</button>
</body>
</html>
`

func serveHtml(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, page)
}

func echo(ws *websocket.Conn) {
	var err error

	for {
		var msgFromClient string

		if err = websocket.Message.Receive(ws, &msgFromClient); err != nil {
			log.Printf("Receive failed: %v", err)
			if !ws.IsClientConn() {
				log.Printf("Client is no longer connected")
			}
			break
		}
		fmt.Printf("Received string from client \"%s\"\n", msgFromClient)

		if err = websocket.Message.Send(ws, "thanks"); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func main() {
	http.HandleFunc("/", serveHtml)
	http.Handle("/ws", websocket.Handler(echo))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
