package sockets

import (
	"goiris/admin/app/middleware/jwt"
	"goiris/admin/app/middleware/sockets/gorilla"

	//"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
	"log"
)

var G_ws *neffos.Server

// values should match with the client sides as well.
const enableJWT = false
const namespace = "default"

// if namespace is empty then simply websocket.Events{...} can be used instead.
var serverEvents = websocket.Namespaces{
	namespace: websocket.Events{
		websocket.OnNamespaceConnected: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			// with `websocket.GetContext` you can retrieve the Iris' `Context`.
			ctx := websocket.GetContext(nsConn.Conn)
			log.Printf("[%s] connected to namespace [%s] with IP [%s]", nsConn, msg.Namespace, ctx.RemoteAddr())
			return nil
		},
		websocket.OnNamespaceDisconnect: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			log.Printf("[%s] disconnected from namespace [%s]", nsConn, msg.Namespace)
			return nil
		},
		"chat": func(nsConn *websocket.NSConn, msg websocket.Message) error {
			// room.String() returns -> NSConn.String() returns -> Conn.String() returns -> Conn.ID()
			log.Printf("[%s] sent: %s", nsConn, string(msg.Body))
			nsConn.Conn.Server().Broadcast(nsConn, msg)
			return nil
		},
	},
}

func InitWebsocket(app *iris.Application) {
	websocketServer := websocket.New(
		//websocket.DefaultGorillaUpgrader, /* DefaultGobwasUpgrader can be used too. */
		gorilla.DefaultUpgrader,
		serverEvents)
	G_ws = websocketServer

	//j := jwt.New(jwt.Config{
	//	// Extract by the "token" url,
	//	// so the client should dial with ws://localhost:8080/echo?token=$token
	//	Extractor: jwt.FromParameter("token"),
	//	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
	//		return []byte("My Secret"), nil
	//	},
	//	SigningMethod: jwt.SigningMethodHS256,
	//})

	//idGen := func(ctx iris.Context) string {
	//	if username := ctx.GetHeader("X-Username"); username != "" {
	//		return username
	//	}
	//	return websocket.DefaultIDGenerator(ctx)
	//}

	// serves the endpoint of ws://localhost:8080/echo
	// with optional custom ID generator.
	websocketRoute := app.Get("/echo", websocket.Handler(websocketServer))
	if enableJWT {
		// Register the jwt middleware (on handshake):
		//websocketRoute.Use(j.Serve)
		websocketRoute.Use(jwt.G_JWT.ServeWebsocket)
	}
}
