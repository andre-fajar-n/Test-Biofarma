package runtime

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func (r *Runtime) db() *Runtime {
	r.Logger.Info().Msg("Initiate connection to DB...")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		r.Cfg.DBHost,
		r.Cfg.DBUser,
		r.Cfg.DBPassword,
		r.Cfg.DBName,
		r.Cfg.DBPort,
	)

	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	gormConfig := &gorm.Config{
		// enhance performance config
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		Logger:                 dbLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		r.Logger.Error().Err(err).Msg("Error connect to DB")
		log.Panicf("Error connect to DB : %v", err)
	}

	r.Db = db

	r.Logger.Info().Msg("DB successfully connected")

	return r
}

func (r *Runtime) MigrateUp(path string) {
	r.prepareMigration(path, "up")
}

func (r *Runtime) MigrateDown(path string) {
	r.prepareMigration(path, "down")
}

func (r *Runtime) prepareMigration(path, migrationType string) {
	r.Logger.Info().Msgf("Initiate db migration %s", migrationType)

	m := r.prepareMigrator(path)
	defer m.Close()

	var err error
	switch migrationType {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	}
	if err != nil && err != migrate.ErrNoChange {
		r.Logger.Error().Err(err).Msgf("Error migration %s", migrationType)
		log.Panicf("Error migration %s: %v", migrationType, err)
	}

	if err == migrate.ErrNoChange {
		r.Logger.Info().Msg("There are no change migration")
		return
	}

	r.Logger.Info().Msgf("Migrating %s db has been done", migrationType)
}

func (r *Runtime) ForceLastestVersion(path string) {
	m := r.prepareMigrator(path)
	defer m.Close()

	version, _, err := m.Version()
	if err != nil {
		r.Logger.Error().Err(err).Msg("error get version")
		log.Fatalf("error get version: %v", err)
	}

	err = m.Force(int(version))
	if err != nil {
		r.Logger.Error().Err(err).Msgf("error force version %d", version)
		log.Fatalf("error force version %d: %v", version, err)
	}
}

func (r *Runtime) prepareMigrator(path string) *migrate.Migrate {
	sqlDB, err := r.Db.DB()
	if err != nil {
		r.Logger.Error().Err(err).Msg("Error return sql.DB")
		log.Panicf("Error return sql.DB: %v", err)
	}

	driver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
	if err != nil {
		r.Logger.Error().Err(err).Msg("Error create instance")
		log.Panicf("Error create instance: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", path), "postgres", driver)
	if err != nil {
		r.Logger.Error().Err(err).Msg("Error create new migrator")
		log.Panicf("Error create new migrator: %v", err)
	}

	return m
}

func (r *Runtime) CreateFileMigration(path, filename string) error {
	version := time.Now().UTC().Format("20060102150405")
	nameWithVersion := version + "_" + filename
	formatFilename := "%s/%s.%s.sql"
	filenameUp := fmt.Sprintf(formatFilename, path, nameWithVersion, "up")
	filenameDown := fmt.Sprintf(formatFilename, path, nameWithVersion, "down")

	if err := r.createFile(filenameUp); err != nil {
		r.Logger.Error().Err(err).Msg("error create file migration up")
		return err
	}

	if err := r.createFile(filenameDown); err != nil {
		r.Logger.Error().Err(err).Msg("error create file migration down")
		return err
	}

	return nil
}

func (r *Runtime) createFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		r.Logger.Error().Err(err).Msg("Error create file")
		return err
	}
	defer f.Close()

	fmt.Println("Generated new migration files...", f.Name())
	return nil
}
