package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	AppName     string `env:"APP_NAME" env-default:"mPharma-icd10"`
	AppAddress  string `env:"APP_ADDRESS" env-default:"0.0.0.0:3000"`
	AppLogLevel string `env:"APP_LOG_LEVEL" env-default:"debug"`

	SMTP struct {
		Email     string `env:"SMTP_EMAIL" env-default:"admin@admin.com"`
		Server    string `env:"SMTP_SERVER" env-default:"smtp.google.com"`
		Port      int    `env:"SMTP_PORT" env-default:"587"`
		Password  string `env:"SMTP_PASSWORD" env-default:"PP"`
		EnableTLS bool   `env:"SMTP_TLS_ENABLED" env-default:"false"`
	}

	File struct {
		UploadDirectory string `env:"FILE_UPLOAD_DIR" env-default:"uploads"`
	}

	Database struct {
		PoolMax              int    `env:"DB_POOL_MAX"  env-default:"1"`
		URL                  string `env:"DB_URL" env-default:"./icd_codes_db.sqlite"`
		Driver               string `env:"DB_DRIVER" env-default:"sqliteshim"`
		PrintQueriesToStdout bool   `env:"DB_PRINT_TO_STDOUT"  env-default:"true"`
	}
}

// Loads config file, overwrite/parses with env variables
// & returns filled struct.
//
// Take note: if configFile is nil, it skips to env
func Load(configFile string, config any) {
	var err error

	if configFile != "" {
		err = cleanenv.ReadConfig(configFile, config)
		if err != nil {
			log.Fatalf("unable read config from file: '%s' | %s", configFile, err.Error())
		}
	}

	// env overwriting config
	err = cleanenv.ReadEnv(config)
	if err != nil {
		log.Fatalf("unable read config from env: %s", err.Error())
	}

	// updating env variables change via runtime
	err = cleanenv.UpdateEnv(config)
	if err != nil {
		log.Fatalf("unable to update config from env: %s", err.Error())
	}
}
