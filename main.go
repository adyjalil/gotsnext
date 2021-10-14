package main

import (
	"bytes"
	"fmt"
	"gotsnext/internal/helpers"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	file, err := os.Open("a.txt")
	if err != nil {
		log.Fatalln(err)
	}

	io.Copy(os.Stdout, file)
	conf := fiber.Config{
		ServerHeader: "go fiber",
	}
	app := fiber.New(conf)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./html/index.html", true)
	})

	app.Get("/n/:number", func(c *fiber.Ctx) error {
		number := c.Params("number", "0")
		numberInt, err := strconv.Atoi(number)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).SendString("cannot convert int")
		}

		n := helpers.IntToString(numberInt)
		return c.SendString(n)
	})

	app.Get("/s/:number", func(c *fiber.Ctx) error {
		number := c.Params("number", "0")
		n, err := helpers.StringToInt(number)
		if err != nil {
			return fiber.ErrBadRequest
		}

		return c.SendString(fmt.Sprintf("number value: %d", n))
	})

	app.Post("/file", func(c *fiber.Ctx) error {

		fh, err := c.FormFile("file")
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err.Error())
		}
		file, err := fh.Open()
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON("error open")
		}
		defer file.Close()

		buf := new(bytes.Buffer)
		buf.ReadFrom(file)

		tp := http.DetectContentType(buf.Bytes())
		name := fh.Filename
		size := fh.Size

		res := struct {
			Type         string `json:"type"`
			Name         string `json:"name"`
			Size         int64  `json:"size"`
			Message      string `json:"message"`
			UploadedSize int    `json:"uploaded_size"`
			Filename     string `json:"filename"`
		}{
			Type:    tp,
			Name:    name,
			Size:    size,
			Message: "upload successful",
		}
		suffix := ".png"
		if strings.Contains(tp, "jpeg") {
			suffix = ".jpg"
		}
		fname := fmt.Sprintf("./assets/img_%d%s", time.Now().Unix(), suffix)
		dst, err := os.Create(fname)
		if err != nil {
			return c.Status(fiber.ErrInternalServerError.Code).JSON(err.Error())
		}
		defer dst.Close()
		res.UploadedSize, err = dst.Write(buf.Bytes())
		res.Filename = fname
		if err != nil {
			return c.Status(fiber.ErrInternalServerError.Code).JSON(err.Error())
		}

		return c.JSON(res)
	})

	app.Static("/", "./assets")
	app.Listen(":3000")
}
