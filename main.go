package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/getcalmar", getcalmar)

	port := os.Getenv("PORT")
	if port == "" {
		port = "1323" // Значение по умолчанию, если переменная окружения PORT не установлена
	}

	e.Start(":" + port)
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func getcalmar(c echo.Context) error {
	const apiKey = "6c40608948132e679640fb887ffc8b2861ddd39be695c6dca6eb92fedcfdd0c5"
	x := fmt.Sprintf("%d", time.Now().Unix())
	str := "pid=p210o9wx867h8eowp7pg&method=getRandItem&uts=" + x
	hash := GetMD5Hash(str + apiKey)
	an, err := http.Get("http://anecdotica.ru/api?" + str + "&hash=" + hash)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(an.Body)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, string(body))
}
