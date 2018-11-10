package main

import (
	"github.com/fumiyas/go-tty"
	"github.com/kardianos/osext"
	"github.com/mattn/go-colorable"
	"github.com/qpliu/qrencode-go/qrencode"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"path"
	"time"
)

func printQrCode(src string, inverse bool) {
	grid, errQr := qrencode.Encode(src, qrencode.ECLevelL)

	if errQr != nil {
		panic(errQr)
	}

	da1, errQrText := tty.GetDeviceAttributes1(os.Stdout)

	if errQrText == nil && da1[tty.DA1_SIXEL] {
		printSixelQr(os.Stdout, grid, inverse)
	} else {
		stdout := colorable.NewColorableStdout()
		printAnsiArtQr(stdout, grid, inverse)
	}
}

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
