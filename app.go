package main

import (
  "github.com/gin-gonic/gin"
  "gopkg.in/olahol/melody.v1"
  "os"
)

const port = ":5000"

//go:generate file2const -package main assets/index.html:indexHtml index_html.go

func main() {

  keys := NewKeyHolder()

  gin.SetMode(gin.ReleaseMode)

  router := gin.Default()
  socket := melody.New()

  address := findLocalIpAdress()
  randomKey := generateRandomKey(4)

  router.GET("/" + randomKey, func(c *gin.Context) {
    c.Writer.Header().Set("Content-Type", "text/html")
    c.String(200, indexHtml)
  })

  router.GET("/presento", func(c *gin.Context) {
    socket.HandleRequest(c.Writer, c.Request)
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

  os.Stderr.WriteString("Go to http://" + address + port + "/" + randomKey + " to control\n")

  router.Run(port)
}
