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
	amExcelRepo := am.NewHandlingSettingExcelRepository(amDB)
	amExcludeContractRepo := am.NewExcludeContractRepository(amDB)
	// cmsExcludeContractRepo := cms.NewExcludeContractRepository(cmsDB)

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
	handlingDataService := service.NewHandlingSettingExcelService(
		amDB,
		amExcelRepo,
	)
	excludeContractService := service.NewExcludeContractService(
		amDB,
		amExcludeContractRepo,
		outboxRepo,
	)

	// =========================
	// HANDLER LAYER
	// =========================
	handlingHandler := handler.NewHandlingSettingHandler(handlingService, handlingSpSptService)
	handlingDataHandler := handler.NewHandlingDataHandler(handlingDataService)
	excludeContractHandler := handler.NewExcludeContractHandler(excludeContractService)
	handlingSpHandler := handler.NewHandlingSpHandler(handlingSpSptService)

	// =========================
	// ROUTER (GIN)
	// =========================
	engine := router.SetupRoutes(handlingHandler, handlingDataHandler, excludeContractHandler, handlingSpHandler)

	// =========================
	// WORKER START (CMS SYNC)
	// =========================
	ctx := context.Background()

	cmsWorker := worker.NewCmsSyncWorker(cmsDB, outboxRepo, cmsRepo)
	cmsSpWorker := worker.NewCmsSyncSpWorker(cmsDB, outboxRepo, cmsRepoSpSpt)
	// cmsExcContWorker := worker.NewCmsSyncExcContWorker(cmsDB, outboxRepo, cmsExcludeContractRepo)

	if config.GetBool("ENABLE_CMS_WORKER", false) {
		go cmsWorker.Start(ctx)
		go cmsSpWorker.Start(ctx)
		// go cmsExcContWorker.Start(ctx)
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
