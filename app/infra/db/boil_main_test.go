// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package db

import (
	"database/sql"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/kat-co/vala"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/volatiletech/sqlboiler/boil"
)

var flagDebugMode = flag.Bool("test.sqldebug", false, "Turns on debug mode for SQL statements")

var (
	dbMain tester
)

type tester interface {
	setup() error
	conn() (*sql.DB, error)
	teardown() error
}

func TestMain(m *testing.M) {
	if dbMain == nil {
		fmt.Println("no dbMain tester interface was ready")
		os.Exit(-1)
	}

	rand.Seed(time.Now().UnixNano())
	var err error

	// Load configuration
	err = initViper()
	if err != nil {
		fmt.Println("unable to load config file")
		os.Exit(-2)
	}

	setConfigDefaults()
	if err := validateConfig("mysql"); err != nil {
		fmt.Println("failed to validate config", err)
		os.Exit(-3)
	}

	// Set DebugMode so we can see generated sql statements
	flag.Parse()
	boil.DebugMode = *flagDebugMode

	if err = dbMain.setup(); err != nil {
		fmt.Println("Unable to execute setup:", err)
		os.Exit(-4)
	}

	conn, err := dbMain.conn()
	if err != nil {
		fmt.Println("failed to get connection:", err)
	}

	var code int
	boil.SetDB(conn)
	code = m.Run()

	if err = dbMain.teardown(); err != nil {
		fmt.Println("Unable to execute teardown:", err)
		os.Exit(-5)
	}

	os.Exit(code)
}

func initViper() error {
	var err error

	viper.SetConfigName("sqlboiler")

	configHome := os.Getenv("XDG_CONFIG_HOME")
	homePath := os.Getenv("HOME")
	wd, err := os.Getwd()
	if err != nil {
		wd = "../"
	} else {
		wd = wd + "/.."
	}

	configPaths := []string{wd}
	if len(configHome) > 0 {
		configPaths = append(configPaths, filepath.Join(configHome, "sqlboiler"))
	} else {
		configPaths = append(configPaths, filepath.Join(homePath, ".config/sqlboiler"))
	}

	for _, p := range configPaths {
		viper.AddConfigPath(p)
	}

	// Ignore errors here, fall back to defaults and validation to provide errs
	_ = viper.ReadInConfig()
	viper.AutomaticEnv()

	return nil
}

// setConfigDefaults is only necessary because of bugs in viper, noted in main
func setConfigDefaults() {
	if viper.GetString("postgres.sslmode") == "" {
		viper.Set("postgres.sslmode", "require")
	}
	if viper.GetInt("postgres.port") == 0 {
		viper.Set("postgres.port", 5432)
	}
	if viper.GetString("mysql.sslmode") == "" {
		viper.Set("mysql.sslmode", "true")
	}
	if viper.GetInt("mysql.port") == 0 {
		viper.Set("mysql.port", 3306)
	}
	if viper.GetString("mssql.sslmode") == "" {
		viper.Set("mssql.sslmode", "true")
	}
	if viper.GetInt("mssql.port") == 0 {
		viper.Set("mssql.port", 1433)
	}
}

func validateConfig(driverName string) error {
	if driverName == "postgres" {
		return vala.BeginValidation().Validate(
			vala.StringNotEmpty(viper.GetString("postgres.user"), "postgres.user"),
			vala.StringNotEmpty(viper.GetString("postgres.host"), "postgres.host"),
			vala.Not(vala.Equals(viper.GetInt("postgres.port"), 0, "postgres.port")),
			vala.StringNotEmpty(viper.GetString("postgres.dbname"), "postgres.dbname"),
			vala.StringNotEmpty(viper.GetString("postgres.sslmode"), "postgres.sslmode"),
		).Check()
	}

	if driverName == "mysql" {
		return vala.BeginValidation().Validate(
			vala.StringNotEmpty(viper.GetString("mysql.user"), "mysql.user"),
			vala.StringNotEmpty(viper.GetString("mysql.host"), "mysql.host"),
			vala.Not(vala.Equals(viper.GetInt("mysql.port"), 0, "mysql.port")),
			vala.StringNotEmpty(viper.GetString("mysql.dbname"), "mysql.dbname"),
			vala.StringNotEmpty(viper.GetString("mysql.sslmode"), "mysql.sslmode"),
		).Check()
	}

	if driverName == "mssql" {
		return vala.BeginValidation().Validate(
			vala.StringNotEmpty(viper.GetString("mssql.user"), "mssql.user"),
			vala.StringNotEmpty(viper.GetString("mssql.host"), "mssql.host"),
			vala.Not(vala.Equals(viper.GetInt("mssql.port"), 0, "mssql.port")),
			vala.StringNotEmpty(viper.GetString("mssql.dbname"), "mssql.dbname"),
			vala.StringNotEmpty(viper.GetString("mssql.sslmode"), "mssql.sslmode"),
		).Check()
	}

	return errors.New("not a valid driver name")
}
