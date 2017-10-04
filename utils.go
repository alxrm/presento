package main

import (
	"github.com/russross/blackfriday"
	"io/ioutil"
	"math/rand"
	"net"
	"time"
  "path"
  "github.com/kardianos/osext"
)

func generateRandomKey(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	sequence := make([]rune, length)

	rand.Seed(time.Now().UnixNano())

	for i := range sequence {
		sequence[i] = letters[rand.Intn(len(letters))]
	}

	return string(sequence)
}

func readMdFileToHtml(filename string) []byte {
  normalizedRelativePath := normalizeRelativePath(filename)

	if bytes, err := ioutil.ReadFile(normalizedRelativePath); err == nil {
		return blackfriday.MarkdownBasic(bytes)
	}

	return []byte{}
}

func normalizeRelativePath(relativePath string) string {
  filename, _ := osext.Executable()

  return path.Join(path.Dir(filename), relativePath)
}

func findLocalIpAddress() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		panic(err)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	panic("Could not find any suitable IP address")
}
