package app

import (
	"dishes-admin-mod/internal/app/config"
	"dishes-admin-mod/pkg/logger"

	"github.com/tus/tusd/pkg/filestore"
	"github.com/tus/tusd/pkg/handler"
)

func InitFileServer() (*handler.UnroutedHandler, func(), error) {
	cfg := config.C.FileServer

	store := filestore.FileStore{
		Path: cfg.Directory,
	}

	composer := handler.NewStoreComposer()
	store.UseIn(composer)

	handler, err := handler.NewUnroutedHandler(handler.Config{
		StoreComposer:         composer,
		NotifyCompleteUploads: true,
	})
	if err != nil {
		return handler, nil, err
	}

	go func() {
		for {
			event := <-handler.CompleteUploads
			logger.Infof("Upload %s finished\n", event.Upload.ID)
		}
	}()

	cleanFunc := func() {
		logger.Infof("Stop File Server")
	}

	return handler, cleanFunc, nil
}
