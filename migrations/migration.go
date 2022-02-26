package main

import (
	"cimble/utilities"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const migrationQueriesPath = "./migrations/migration_queries"

func main() {
	fileName := flag.String("c", "", "provide migration filename")
	operation := flag.String("o", "up", "provide number of up migrations to run")
	numDownMigrations := flag.Int("n", -1, "provide number of up migrations to run")

	flag.Parse()

	var err error

	if *fileName != "" {
		err = createMigrationFiles(*fileName)
		if err != nil {
			return
		}

		fmt.Printf("Migration files created for filename: %s", *fileName)

		return
	}

	err = utilities.LoadEnvironmentVariables()
	if err != nil {
		return
	}

	databaseUrl := utilities.ConstructDatabaseUrl()

	migrationQueriesAbsPath, err := filepath.Abs(migrationQueriesPath)
	if err != nil {
		log.Fatalf("Error getting absolute migration queries file path. %v", err)
	}

	m, err := migrate.New(fmt.Sprintf("file://%s", migrationQueriesAbsPath), databaseUrl)
	if err != nil {
		log.Fatalf("Error creating db connection: %v", err)
		return
	}

	if *operation == "up" {
		err = runUpMigrations(m)
		if err != nil {
			return
		}
	}

	if *operation == "down" {
		err = runDownMigrations(m, *numDownMigrations)
		if err != nil {
			return
		}
	}

	fmt.Println("Migrations successfully completed.")
}

func createMigrationFile(migrationType string, fileName string) (err error) {
	currentUnixTime := time.Now().Unix()
	migrationQueriesAbsPath, err := filepath.Abs(migrationQueriesPath)
	if err != nil {
		log.Fatalf("Error getting absolute migration queries file path. %v", err)
	}

	fileName = fmt.Sprintf("%s/%d_%s.%s.sql", migrationQueriesAbsPath, currentUnixTime, fileName, migrationType)

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating migration %s file: %s: Error: %v\n", migrationType, fileName, err)
		return
	}
	file.Close()

	return
}

func runUpMigrations(m *migrate.Migrate) (err error) {
	err = m.Up()
	if err != nil {
		fmt.Printf("Error running up migrations: %v\n", err)
		return
	}

	fmt.Printf("Up migrations successfully completed.")
	return
}

func runDownMigrations(m *migrate.Migrate, numDownMigrations int) (err error) {
	if numDownMigrations == -1 {
		log.Fatalf("Provide number of down migrations to run.\n")
		return
	}

	err = m.Steps(numDownMigrations * -1)
	if err != nil {
		fmt.Printf("Error running down migrations: %v\n", err)
		return
	}

	return
}

func createMigrationFiles(fileName string) (err error) {
	err = createMigrationFile("up", fileName)
	if err != nil {
		fmt.Printf("Error creating up migration file.")
		return
	}
	err = createMigrationFile("down", fileName)
	if err != nil {
		fmt.Printf("Error creating down migration file.")
		return
	}

	fmt.Printf("Migration files created for filename: %s", fileName)

	return
}
