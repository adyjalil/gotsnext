package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"gotsnext/internal/helpers"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5444
	user     = "db_admin"
	password = "123qwe"
	dbname   = "api"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalln("db conn error", err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatalln("db ping error", err)
	}

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

	type User struct {
		Username  string       `json:"username" form:"username"`
		Email     string       `json:"email" form:"email"`
		Password  string       `json:"password" form:"password"`
		CreatedAt time.Time    `json:"created_at" form:"created_at"`
		UpdatedAt time.Time    `json:"updated_at" form:"updated_at"`
		DeletedAt sql.NullTime `json:"deleted_at" form:"deleted_at"`
	}

	app.Post("/user", func(c *fiber.Ctx) error {
		var u *User = new(User)

		if err := c.BodyParser(u); err != nil {
			return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
		}

		u.CreatedAt = time.Now()
		u.UpdatedAt = time.Now()

		fmt.Printf("user: %+v", u)

		stmt := "INSERT INTO users (username, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"

		res, err := db.Exec(stmt, u.Username, u.Email, u.Password, u.CreatedAt, u.UpdatedAt)
		if err != nil {
			return c.Status(fiber.ErrInternalServerError.Code).SendString(err.Error())
		}

		rows, _ := res.RowsAffected()
		return c.SendString(fmt.Sprintf("user success inserted: %d", rows))
	})

	app.Put("/user", func(c *fiber.Ctx) error {
		var u *User = new(User)

		if err := c.BodyParser(u); err != nil {
			return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
		}

		u.UpdatedAt = time.Now()

		fmt.Printf("user: %+v", u)

		stmt := `UPDATE users 
			SET email = $1,
				password = $2,
				updated_at = $3
			WHERE username = $4
				`

		res, err := db.Exec(stmt, u.Email, u.Password, u.UpdatedAt, u.Username)
		if err != nil {
			return c.Status(fiber.ErrInternalServerError.Code).SendString(err.Error())
		}

		rows, _ := res.RowsAffected()
		return c.SendString(fmt.Sprintf("user success updated: %d", rows))
	})

	app.Static("/", "./assets")
	app.Listen(":3000")
}
