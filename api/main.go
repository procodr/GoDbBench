package main

import (
	"api/core"
	"api/repo"
	"context"
	"database/sql"
	"github.com/cockroachdb/pebble"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func openPgDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgresql://postgres:example_test@127.0.0.1:54320/postgres?sslmode=disable")
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	q := `
		CREATE TABLE IF NOT EXISTS data (c1 varchar, c2 varchar, c3 varchar);
		TRUNCATE TABLE data;`

	_, err = db.Exec(q)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func openSqLite() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "data/sqliteDb")
	if err != nil {
		return nil, err
	}

	q := `
		CREATE TABLE IF NOT EXISTS data (c1 varchar, c2 varchar, c3 varchar);
		DELETE FROM data;`

	_, err = db.Exec(q)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func openPebble() (*pebble.DB, error) {
	db, err := pebble.Open("data/pebble", nil)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func openMongo() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
}

func main() {
	pgDb, err := openPgDB()
	if err != nil {
		log.Fatal(err)
	}

	sqDb, err := openSqLite()
	if err != nil {
		log.Fatal(err)
	}

	pbDb, err := openPebble()
	if err != nil {
		log.Fatal(err)
	}

	mnDb, err := openMongo()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello 👋!")
	})

	pgTest := repo.NewDbTest(pgDb)
	sqTest := repo.NewDbTest(sqDb)
	pbTest := repo.NewPbTest(pbDb)
	mnTest := repo.NewMongoTest(mnDb)

	app.Post("/pgTest", func(c *fiber.Ctx) error {
		d := new(core.Data)

		if err := c.BodyParser(d); err != nil {
			return c.Status(503).SendString(err.Error())
		}

		return pgTest.Create(d)
	})

	app.Post("/sqTest", func(c *fiber.Ctx) error {
		d := new(core.Data)

		if err := c.BodyParser(d); err != nil {
			return c.Status(503).SendString(err.Error())
		}

		return sqTest.Create(d)
	})

	app.Post("/pbTest", func(c *fiber.Ctx) error {
		d := new(core.Data)
		if err := c.BodyParser(d); err != nil {
			return c.Status(503).SendString(err.Error())
		}

		return pbTest.Create(d)
	})

	app.Post("/mnTest", func(c *fiber.Ctx) error {
		d := new(core.Data)
		if err := c.BodyParser(d); err != nil {
			return c.Status(503).SendString(err.Error())
		}

		return mnTest.Create(d)
	})

	err = app.Listen(":3000")
	if err != nil {
		return
	}
}
