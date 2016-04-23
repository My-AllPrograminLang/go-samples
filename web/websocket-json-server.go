// Websocket server that receives Person structs encoded in JSON from the
// client. A JS client that sends messages is included.
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
							var msg = JSON.parse(e.data);
							console.log("JSON parsed: ", msg);
            }
        };

        function send() {
					var msg = {
						Name: document.getElementById('name').value,
						Emails: ["1@foo.org", "2@bar.com"],
					};
					sock.send(JSON.stringify(msg));
        };
    </script>
    <h1>WebSocket JSON sender</h1>
    <form>
        <p>
            Name: <input id="name" type="text" value="bruh">
        </p>
    </form>
    <button onclick="send();">Send Message</button>
</body>
</html>
`

type Person struct {
	Name   string
	Emails []string
}

func receivePerson(ws *websocket.Conn) {
	var person Person
	err := websocket.JSON.Receive(ws, &person)
	if err != nil {
		fmt.Println("Can't receive")
	} else {

		fmt.Println("Name: " + person.Name)
		for _, e := range person.Emails {
			fmt.Println("An email: " + e)
		}

		// Tweak the name a bit and pong it back to the client.
		person.Name = "!!" + person.Name
		err := websocket.JSON.Send(ws, person)
		if err != nil {
			fmt.Println("Couldn't send msg " + err.Error())
		}
	}
}

func serveHtml(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, page)
}

func main() {
	http.HandleFunc("/", serveHtml)
	http.Handle("/ws", websocket.Handler(receivePerson))
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
