package boot

import (
	"fmt"
	"log"
	"net"
	"tipe-handling/config"
	"tipe-handling/internal/handler"
	"tipe-handling/internal/repository"
	"tipe-handling/internal/service"
	"tipe-handling/pkg/router"
)

type App struct {
	Engine *router.EngineWrapper
	DB     config.DB
}

func NewApp() *App {

	log.Println(">>> BOOT STARTED")

	// =========================
	// DATABASE LAYER
	// =========================
	dbs := config.Init()

	log.Printf("TYPE OF DB = %T", dbs)

	// optional debug tambahan
	log.Printf("DB INSTANCE = %+v", dbs)

	// =========================
	// REPOSITORY LAYER
	// =========================
	handlingRepo := repository.NewHandlingSettingRepository(dbs.MUFAM(), dbs.MUFCMS())

	// =========================
	// SERVICE LAYER
	// =========================
	handlingService := service.NewHandlingSettingService(handlingRepo)

	// =========================
	// HANDLER LAYER
	// =========================
	handlingHandler := handler.NewHandlingSettingHandler(handlingService)

	// =========================
	// ROUTER (GIN)
	// =========================
	engine := router.SetupRoutes(handlingHandler)

	return &App{
		Engine: engine,
	}
}

func (a *App) Run() {
	ln, err := net.Listen("tcp4", ":0")
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

	addr := ln.Addr().(*net.TCPAddr)

	fmt.Printf("\n🚀 Server running at: http://localhost:%d\n\n", addr.Port)

	if err := a.Engine.RunListener(ln); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
