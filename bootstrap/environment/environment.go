package environment

import (
	"os"
	"strconv"

	"github.com/subosito/gotenv"
)

var (
	// DATABASEHOST : Database Host
	DATABASEHOST string

	// DATABASEPORT : Database Port
	DATABASEPORT int

	// DATABASENAME : Database Name
	DATABASENAME string

	// DATABASEUSERNAME : Database Username
	DATABASEUSERNAME string

	// DATABASEPASSWORD : Database Password
	DATABASEPASSWORD string

	// DATABASEADAPTER : Database Adapter
	DATABASEADAPTER string

	// DATABASETIMEOUT : Database Timeout
	DATABASETIMEOUT int

	// ENVIRONMENT : Environment
	ENVIRONMENT string
)

// Init environment
func initialize() {
	gotenv.Load(os.Getenv("GOPATH") + "\\src\\github.com\\agungcandra\\dao\\.env")
	DATABASEHOST = databaseHost()
	DATABASEPORT = databasePort()
	DATABASENAME = databaseName()
	DATABASEUSERNAME = databaseUsername()
	DATABASEPASSWORD = databasePassword()
	DATABASEADAPTER = databaseAdapter()
	DATABASETIMEOUT = databaseTimeout()
	ENVIRONMENT = environment()
}

func Load() {
	initialize()
}

// IsDevelopment : Check whether application run in development mode
func IsDevelopment() bool {
	return ENVIRONMENT == "development"
}

// IsProduction : Check whether application run in production mode
func IsProduction() bool {
	return ENVIRONMENT == "production"
}

func databaseHost() string {
	host := env("DATABASE_HOST")

	if host != "" {
		return host
	}
	return env("DATABASE_HOST")
}

func databasePort() int {
	value, _ := strconv.Atoi(env("DATABASE_PORT"))

	return value
}

func databaseName() string {
	return env("DATABASE_NAME")
}

func databaseUsername() string {
	return env("DATABASE_USERNAME")
}

func databasePassword() string {
	return env("DATABASE_PASSWORD")
}

func databaseAdapter() string {
	return env("DATABASE_ADAPTER")
}

func databaseTimeout() int {
	value, _ := strconv.Atoi(env("DATABASE_TIMEOUT"))
	return value
}

func environment() string {
	return env("ENV")
}

func env(key string) string {
	return os.Getenv(key)
}

// GetEnv : Get Environment Variable
func GetEnv(key string) string {
	return env(key)
}
