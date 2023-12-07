/*

 */

package server

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/dgryski/dgoogauth"
)

func makeHttpServer(addr string, r http.Handler) *http.Server {
	return &http.Server{
		Addr: addr,
		//Handler: ratelimit.RateLimit(r),
		Handler: r,
	}
}

func handleGracefulShutdown(server *http.Server, closers []io.Closer) {

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	log.Println("receive interrupt signal")
	for _, c := range closers {
		err := c.Close()
		if err != nil {
			log.Println("error during closing: ", err)
		}
	}
	if err := server.Close(); err != nil {
		log.Fatal("Server Close:", err)
	}
}

func validate2FACode(code string) (int, error) {
	if code == "" {
		return 0, errors.New("2FA code not found")
	}
	codeTrimmed := strings.TrimLeft(code, "0")
	codeInt, err := strconv.Atoi(codeTrimmed)
	if err != nil {
		return 0, errors.New("2FA code must be a number")
	}
	if codeInt < 0 || codeInt > 999999 {
		return 0, errors.New("2FA code must be six digits")
	}
	return codeInt, nil
}

func compute2FACode(sub string) int {
	// Refer to: totp.danhersam.com/
	inputNoSpaces := strings.Replace(sub, "-", "", -1)
	inputBase32 := strings.Replace(strings.Replace(strings.Replace(strings.Replace(
		inputNoSpaces, "0", "O", -1), "1", "I", -1), "8", "B", -1), "9", "6", -1)
	inputNoSpacesUpper := strings.ToUpper(inputBase32)
	code := dgoogauth.ComputeCode(inputNoSpacesUpper, time.Now().Unix()/600)
	return code
}
