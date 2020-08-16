package config

import (
	"crypto/rand"
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/pprof"
	"github.com/gofiber/template/html"
	. "github.com/itsursujit/fiber-boilerplate/app"
	"github.com/itsursujit/flash"
	"github.com/markbates/pkger"
	"os"
	"path/filepath"
)

type AppConfiguration struct {
	App_Name        string
	App_Upload_Path string
	App_Upload_Size int
	App_Env         string
	App_Key         string
	App_Url         string
	App_Port        string
}

var AppConfig *AppConfiguration //nolint:gochecknoglobals

var BlastlistedDomains = []string{}

func LoadAppConfig() {
	loadDefaultConfig()
	ViperConfig.Unmarshal(&AppConfig)
	if AppConfig.App_Url == "" {
		AppConfig.App_Url = fmt.Sprintf("http://localhost:%s", AppConfig.App_Port)
	}
	AppConfig.App_Upload_Path = filepath.Join(".", AppConfig.App_Upload_Path)
	AppConfig.App_Upload_Size = AppConfig.App_Upload_Size * 1024 * 1024
	if _, err := os.Stat(AppConfig.App_Upload_Path); os.IsNotExist(err) {
		os.MkdirAll(AppConfig.App_Upload_Path, os.ModePerm)
	}
}

func loadDefaultConfig() {
	ViperConfig.SetDefault("APP_NAME", "fiber-boilerplate")
	ViperConfig.SetDefault("APP_ENV", "dev")
	ViperConfig.SetDefault("APP_UPLOAD_PATH", "uploads")
	ViperConfig.SetDefault("APP_UPLOAD_SIZE", 4)
	ViperConfig.SetDefault("APP_KEY", "1894cde6c936a294a478cff0a9227fd276d86df6573b51af5dc59c9064edf426")
	ViperConfig.SetDefault("APP_PORT", "8080")
}

func GenerateAppKey(length int) {
	key := make([]byte, length)

	_, err := rand.Read(key)
	if err != nil {
		// handle error here
	}
}

func BootApp() {
	LoadAppConfig()
	TemplateEngine = html.NewFileSystem(pkger.Dir("/resources/views"), ".html")
	App = fiber.New(&fiber.Settings{
		ErrorHandler:          CustomErrorHandler,
		ServerHeader:          "fiber-boilerplate",
		Prefork:               true,
		DisableStartupMessage: false,
		Views:                 TemplateEngine,
		BodyLimit:             AppConfig.App_Upload_Size,
	})

	App.Use(pprof.New())
	App.Use(LoadHeaders)
	App.Use(middleware.Favicon())
	App.Use(middleware.Recover())
	App.Use(middleware.Compress(middleware.CompressLevelBestSpeed))
	/*App.Use(csrf.New(csrf.Config{
		CookieSecure:   true,
	}))*/

	App.Static("/assets", "resources/assets", fiber.Static{
		Compress: true,
	})

	App.Use(LoadCacheHeaders)
	Hash = NewHashDriver()
	_, err := SetupDB()
	if err != nil {
		panic(err)
	}
	SetupPermission()
	LoadSession()
	Flash = &flash.Flash{
		CookiePrefix: "fiber-boilerplate",
	}
}

func CustomErrorHandler(c *fiber.Ctx, err error) {
	// StatusCode defaults to 500
	code := fiber.StatusInternalServerError
	//nolint:misspell    // Retrieve the custom statuscode if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	} //nolint:gofmt,wsl
	if c.Is("json") {
		c.SendStatus(code)
		_ = c.JSON(err)
	} else {
		c.SendStatus(code)
		_ = c.Render(fmt.Sprintf("errors/%d", code), fiber.Map{ //nolint:nolintlint,errcheck
			"error": err,
		})
	}
}
