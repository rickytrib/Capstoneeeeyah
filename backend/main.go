package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"

	"my-app/backend/database"
	"my-app/backend/models"
	"my-app/backend/routes"
)

type App struct {
	DB *gorm.DB
}


func main() {
	 database.ConnectDB()


	 database.DB.Migrator().DropTable(
		&models.Mengelola{},
		&models.Dikelola{},
		&models.Pendaftaran{},
		&models.Masyarakat{},
		&models.Admin{},
		&models.SuperAdmin{},
	)
	log.Println("Dropping old tables...")
	log.Println("Migrating new tables...")


	// Migrasi semua tabel
	err := database.DB.AutoMigrate(
    	&models.Admin{},
    	&models.Masyarakat{},
		&models.Mendaftar{},
    	&models.Pendaftaran{},
   		&models.Mengelola{},
    	&models.Dikelola{},
    	&models.SuperAdmin{},
    )
    if err != nil {
        log.Fatal("Migrasi gagal:", err)
    }

    log.Println("Migrasi berhasil!")

	// Fiber setup
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	// Routes
	app.Get("/api/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Halo dari Fiber!"})
	})

	// Register routes
	routes.SetupLoginRoutes(app)
	routes.SetupDaftarRoutes(app)

app.Use(func(c *fiber.Ctx) error {
	// Lewati kalau ini permintaan API
	if len(c.Path()) >= 4 && c.Path()[:4] == "/api" {
		return c.Next()
	}
	return c.SendFile("../frontend/dist/index.html")
})

	// Start server
	if err := app.Listen(":8070"); err != nil {
		log.Fatal("Gagal menjalankan server:", err)
	}

}