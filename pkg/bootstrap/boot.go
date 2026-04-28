package boot

import (
	"context"
	"fmt"
	"log"
	"net"
	"tipe-handling/config"
	"tipe-handling/internal/handler"
	"tipe-handling/internal/repository/am"
	"tipe-handling/internal/repository/cms"
	"tipe-handling/internal/repository/outbox"
	"tipe-handling/internal/service"
	"tipe-handling/internal/worker"
	"tipe-handling/pkg/router"
)

type App struct {
	Engine *router.EngineWrapper
	DB     config.DB
}

func NewApp() *App {

	// =========================
	// DATABASE LAYER
	// =========================
	dbs := config.Init()
	amDB := dbs.MUFAM()
	cmsDB := dbs.MUFCMS()

	// =========================
	// REPOSITORY LAYER
	// =========================
	amRepo := am.NewHandlingSettingRepository(amDB)
	cmsRepo := cms.NewHandlingSettingRepository(cmsDB)
	amRepoSpSpt := am.NewHandlingSpSptRepository(amDB)
	cmsRepoSpSpt := cms.NewHandlingSpSptRepository(cmsDB)
	outboxRepo := outbox.NewRepository(amDB)

	// =========================
	// SERVICE LAYER
	// =========================
	handlingService := service.NewHandlingSettingService(
		amDB,
		amRepo,
		cmsRepo,
		outboxRepo,
	)
	handlingSpSptService := service.NewHandlingSpSptService(
		amDB,
		amRepoSpSpt,
		cmsRepoSpSpt,
		outboxRepo,
	)

	// =========================
	// HANDLER LAYER
	// =========================
	handlingHandler := handler.NewHandlingSettingHandler(handlingService, handlingSpSptService)

	// =========================
	// ROUTER (GIN)
	// =========================
	engine := router.SetupRoutes(handlingHandler)

	// =========================
	// WORKER START (CMS SYNC)
	// =========================
	ctx := context.Background()

	cmsWorker := worker.NewCmsSyncWorker(cmsDB, outboxRepo, cmsRepo)

	if config.GetBool("ENABLE_CMS_WORKER", false) {
		go cmsWorker.Start(ctx)
	} else {
		log.Println("CMS Worker disabled")
	}

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
