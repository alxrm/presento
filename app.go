package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

const port = ":5000"
const prompterFile = "./prompter.md"

//go:generate file2const -package main assets/index.html:indexHtml index_html.go

func main() {

	keys := NewKeyHolder()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	socket := melody.New()

	address := findLocalIpAddress()
	randomKey := generateRandomKey(4)
	resultLink := "http://" + address + port + "/" + randomKey

	router.GET("/"+randomKey, func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/html")
		c.String(200, indexHtml)
		//c.File("./assets/index.html")
	})

	router.GET("/presento", func(c *gin.Context) {
		socket.HandleRequest(c.Writer, c.Request)
	})

	socket.HandleConnect(func(s *melody.Session) {
		s.Write(readMdFileToHtml(prompterFile))
	})

	socket.HandleMessage(func(s *melody.Session, msg []byte) {
		command := string(msg)

		switch command {
		case "left":
			keys.PressLeft()
		case "right":
			keys.PressRight()
		}
	})

	printQrCode(resultLink, true)
	println("Go to " + resultLink + " to control")
	println("Or scan this QR code to open the link on your device")

	errServer := router.Run(port)

	if errServer != nil {
		killWithMessage("Couldn't start the server, releasing the port, try again")
	}
}
