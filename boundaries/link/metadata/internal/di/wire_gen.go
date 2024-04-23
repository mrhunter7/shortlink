// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package metadata_di

import (
	"context"
	"github.com/google/wire"
	v1_2 "github.com/shortlink-org/shortlink/boundaries/link/metadata/internal/domain/metadata/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/metadata/internal/infrastructure/mq"
	"github.com/shortlink-org/shortlink/boundaries/link/metadata/internal/infrastructure/repository/media"
	"github.com/shortlink-org/shortlink/boundaries/link/metadata/internal/infrastructure/repository/store"
	"github.com/shortlink-org/shortlink/boundaries/link/metadata/internal/infrastructure/rpc/metadata/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/metadata/internal/usecases/metadata"
	"github.com/shortlink-org/shortlink/boundaries/link/metadata/internal/usecases/parsers"
	"github.com/shortlink-org/shortlink/boundaries/link/metadata/internal/usecases/screenshot"
	"github.com/shortlink-org/shortlink/internal/di"
	"github.com/shortlink-org/shortlink/internal/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/internal/di/pkg/config"
	"github.com/shortlink-org/shortlink/internal/di/pkg/context"
	"github.com/shortlink-org/shortlink/internal/di/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/di/pkg/mq"
	"github.com/shortlink-org/shortlink/internal/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/internal/di/pkg/store"
	"github.com/shortlink-org/shortlink/internal/di/pkg/traicing"
	"github.com/shortlink-org/shortlink/pkg/db"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/mq"
	"github.com/shortlink-org/shortlink/pkg/notify"
	"github.com/shortlink-org/shortlink/pkg/observability/monitoring"
	"github.com/shortlink-org/shortlink/pkg/rpc"
	"github.com/shortlink-org/shortlink/pkg/s3"
	"go.opentelemetry.io/otel/trace"
)

// Injectors from wire.go:

func InitializeMetaDataService() (*MetaDataService, func(), error) {
	context, cleanup, err := ctx.New()
	if err != nil {
		return nil, nil, err
	}
	logger, cleanup2, err := logger_di.New(context)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	configConfig, err := config.New()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tracerProvider, cleanup3, err := traicing_di.New(context, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	monitoringMonitoring, cleanup4, err := monitoring.New(context, logger, tracerProvider)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	pprofEndpoint, err := profiling.New(context, logger)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	autoMaxProAutoMaxPro, cleanup5, err := autoMaxPro.New(logger)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	db, err := store.New(context, logger, tracerProvider, monitoringMonitoring)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	metaStore, err := NewMetaDataStore(context, logger, db)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	uc, err := NewParserUC(metaStore)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	mq, err := mq_di.New(context, logger)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	event, err := InitMetadataMQ(context, mq)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	server, err := rpc.InitServer(context, logger, tracerProvider, monitoringMonitoring)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	client, err := s3.New(context, logger)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	service, err := NewMetaDataMediaStore(context, client)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	screenshotUC, err := NewScreenshotUC(context, service)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	metadataUC, err := NewMetadataUC(logger, uc, screenshotUC)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	metadata, err := NewMetaDataRPCServer(logger, server, uc, screenshotUC, metadataUC)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	metaDataService, err := NewMetaDataService(logger, configConfig, monitoringMonitoring, tracerProvider, pprofEndpoint, autoMaxProAutoMaxPro, uc, event, metadata, metaStore)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return metaDataService, func() {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

type MetaDataService struct {
	// Common
	Log    logger.Logger
	Config *config.Config

	// Observability
	Tracer        trace.TracerProvider
	Monitoring    *monitoring.Monitoring
	PprofEndpoint profiling.PprofEndpoint
	AutoMaxPro    autoMaxPro.AutoMaxPro

	// Delivery
	metadataMQ        *metadata_mq.Event
	metadataRPCServer *v1.Metadata

	// Application
	service *parsers.UC

	// Repository
	metadataStore *storeRepository.MetaStore
}

// MetaDataService =====================================================================================================
var MetaDataSet = wire.NewSet(di.DefaultSet, mq_di.New, store.New, rpc.InitServer, s3.New, InitMetadataMQ,
	NewMetaDataRPCServer,

	NewParserUC,
	NewScreenshotUC,
	NewMetadataUC,

	NewMetaDataStore,
	NewMetaDataMediaStore,

	NewMetaDataService,
)

func InitMetadataMQ(ctx2 context.Context, dataBus mq.MQ) (*metadata_mq.Event, error) {
	metadataMQ, err := metadata_mq.New(dataBus)
	if err != nil {
		return nil, err
	}
	notify.Subscribe(v1_2.METHOD_ADD, metadataMQ)

	return metadataMQ, nil
}

func NewMetaDataStore(ctx2 context.Context, log logger.Logger, db2 db.DB) (*storeRepository.MetaStore, error) {
	store2 := &storeRepository.MetaStore{}
	metadataStore, err := store2.Use(ctx2, log, db2)
	if err != nil {
		return nil, err
	}

	return metadataStore, nil
}

func NewMetaDataMediaStore(ctx2 context.Context, s3_2 *s3.Client) (*s3Repository.Service, error) {
	client, err := s3Repository.New(ctx2, s3_2)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func NewParserUC(store2 *storeRepository.MetaStore) (*parsers.UC, error) {
	metadataService, err := parsers.New(store2)
	if err != nil {
		return nil, err
	}

	return metadataService, nil
}

func NewScreenshotUC(ctx2 context.Context, media *s3Repository.Service) (*screenshot.UC, error) {
	metadataService, err := screenshot.New(ctx2, media)
	if err != nil {
		return nil, err
	}

	return metadataService, nil
}

func NewMetadataUC(log logger.Logger, parsersUC *parsers.UC, screenshotUC *screenshot.UC) (*metadata.UC, error) {
	metadataService, err := metadata.New(log, parsersUC, screenshotUC)
	if err != nil {
		return nil, err
	}

	return metadataService, nil
}

func NewMetaDataRPCServer(log logger.Logger, runRPCServer *rpc.Server, parsersUC *parsers.UC, screenshotUC *screenshot.UC, metadataUC *metadata.UC) (*v1.Metadata, error) {
	metadataRPCServer, err := v1.New(log, runRPCServer, parsersUC, screenshotUC, metadataUC)
	if err != nil {
		return nil, err
	}

	return metadataRPCServer, nil
}

func NewMetaDataService(

	log logger.Logger, config2 *config.Config, monitoring2 *monitoring.Monitoring,
	tracer trace.TracerProvider,
	pprofHTTP profiling.PprofEndpoint,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,

	service *parsers.UC,

	metadataMQ *metadata_mq.Event,
	metadataRPCServer *v1.Metadata,

	metadataStore *storeRepository.MetaStore,
) (*MetaDataService, error) {
	return &MetaDataService{

		Log:    log,
		Config: config2,

		Tracer:        tracer,
		Monitoring:    monitoring2,
		PprofEndpoint: pprofHTTP,
		AutoMaxPro:    autoMaxProcsOption,

		service: service,

		metadataMQ:        metadataMQ,
		metadataRPCServer: metadataRPCServer,

		metadataStore: metadataStore,
	}, nil
}