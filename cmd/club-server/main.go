package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/rbicker/club/internal/mailer"
	"github.com/rbicker/club/internal/server"
	"github.com/rbicker/club/internal/store"
	"github.com/rbicker/club/internal/utils"
)

func main() {
	infoLogger := log.New(os.Stdout, "INFO: ", log.Lmsgprefix+log.LstdFlags)
	errorLogger := log.New(os.Stderr, "ERROR: ", log.Lmsgprefix+log.LstdFlags)
	var dbOpts []func(*store.MGO) error
	var srvOpts []func(*server.Server) error
	// get or create secret
	var secret string
	if s, ok := os.LookupEnv("CLUB_SECRET"); ok {
		secret = s
	} else {
		infoLogger.Printf("because no secret was given, a random string will be used")
		infoLogger.Printf("this means things like pagination tokens won't survive a restart of the application")
		infoLogger.Printf("make sure to set the CLUB_SECRET environment variable in production")
		secret = utils.RandomString(20)
	}
	// init db connection
	mongoUrl := utils.LookupEnv("CLUB_MONGO_URL", "mongodb://localhost:27017")
	dbOpts = append(dbOpts, store.WithUrl(mongoUrl))
	dbName := utils.LookupEnv("CLUB_MONGO_DB", "db")
	dbOpts = append(dbOpts, store.WithDBName(dbName))
	dbOpts = append(dbOpts, store.WithInfoLogger(infoLogger))
	dbOpts = append(dbOpts, store.WithErrorLogger(errorLogger))
	db, err := store.NewMongoConnection(secret, dbOpts...)
	if err != nil {
		errorLogger.Fatalf("unable to create mongodb connection: %s", err)
	}
	err = db.Connect()
	if err != nil {
		errorLogger.Fatalf("unable to connect to mongodb: %s", err)
	}
	infoLogger.Println("connected to mongodb")
	// mailer
	var mailClient mailer.MailClient
	smtpHost, ok := os.LookupEnv("CLUB_SMTP_HOST")
	if ok {
		smtpPort := utils.LookupEnv("CLUB_SMTP_PORT", "587")
		smtpUsername, _ := os.LookupEnv("CLUB_SMTP_USERNAME")
		smtpPassword, _ := os.LookupEnv("CLUB_SMTP_PASSWORD")
		mailClient, err = mailer.NewTLSMailer(smtpHost, smtpPort, smtpUsername, smtpPassword)
		if err != nil {
			errorLogger.Fatalf("error while creating mail client: %s", err)
		}
	} else {
		infoLogger.Println("no SMTP settings given, sending mails by logging them to stdout")
		infoLogger.Println("to send real mails, have a look at the CLUB_SMTP_* environment variables")
		mailClient = mailer.NewLogMailer(infoLogger)
	}
	mailFrom, _ := os.LookupEnv("CLUB_MAIL_FROM")
	siteName := utils.LookupEnv("CLUB_SITE_NAME", "gooser")
	gooserTarget := utils.LookupEnv("CLUB_GOOSER_TARGET", "gooser:50051")
	mailer, err := mailer.NewMailer(mailClient, mailFrom, siteName)
	if err != nil {
		errorLogger.Fatalf("error while creating mailer: %s", err)
	}
	// init server
	srvOpts = append(srvOpts, server.EnableReflection())
	srvOpts = append(srvOpts, server.WithInfoLogger(infoLogger))
	srvOpts = append(srvOpts, server.WithErrorLogger(errorLogger))
	srvOpts = append(srvOpts, server.WithGooserTarget(gooserTarget))
	p := utils.LookupEnv("CLUB_PORT", "50051")
	srvOpts = append(srvOpts, server.WithPort(p))
	srv, err := server.NewServer(db, mailer, srvOpts...)
	if err != nil {
		errorLogger.Fatalf("unable to create new gooser server: %s", err)
	}
	// channels
	errChan := make(chan error)
	stopChan := make(chan os.Signal)
	// bind OS events to the signal channel
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	// serve in a go routine
	go func() {
		infoLogger.Println("starting server")
		if err := srv.Serve(); err != nil {
			errChan <- err
		}
	}()
	// terminate gracefully before leaving the main function
	defer func() {
		infoLogger.Println("stopping server")
		srv.Stop()
		infoLogger.Println("disconnecting from mongodb")
		err := db.Disconnect(context.TODO())
		if err != nil {
			errorLogger.Fatalf("error while disconnecting from mongodb: %s", err)
		}
	}()
	// block until either OS signal, or server fatal error
	select {
	case err := <-errChan:
		errorLogger.Printf("Fatal error: %v\n", err)
	case <-stopChan:
	}
}
