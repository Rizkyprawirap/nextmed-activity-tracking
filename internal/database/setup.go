package internaldatabase

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Rizkyprawirap/nextmed-activity-tracking/pkg"
	pkgpostgre "github.com/Rizkyprawirap/nextmed-activity-tracking/pkg/postgre"
	pkgredis "github.com/Rizkyprawirap/nextmed-activity-tracking/pkg/redis"
)

func New() pkg.PkgDatabase {

	//setup redis database
	pkgRedis := pkgredis.New(
		fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		"",
		0,
	)
	pkgPostgrePort, err := strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil {
		panic(fmt.Errorf("invalid DB_PORT: %w", err))
	}
	pkgPostgreMaxCon, err := strconv.Atoi(os.Getenv("DB_MAX_OPENCONN"))
	if err != nil {
		panic(fmt.Errorf("invalid DB_PORT: %w", err))
	}
	pkgPostgreMaxIdleConn, err := strconv.Atoi(os.Getenv("DB_MAX_IDLECONN"))
	if err != nil {
		panic(fmt.Errorf("invalid DB_PORT: %w", err))
	}
	pkgPostgreMaxLifeTime, err := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME"))
	if err != nil {
		panic(fmt.Errorf("invalid DB_PORT: %w", err))
	}
	pkgPostgreIdleTime, err := strconv.Atoi(os.Getenv("DB_MAX_IDLETIME"))
	if err != nil {
		panic(fmt.Errorf("invalid DB_PORT: %w", err))
	}
	pkgPostgre, err := pkgpostgre.New(pkgpostgre.Config{
		Host:            os.Getenv("DB_HOST"),
		User:            os.Getenv("DB_USER"),
		Password:        os.Getenv("DB_PASSWORD"),
		DBName:          os.Getenv("DB_NAME"),
		Port:            pkgPostgrePort,
		SSLMode:         "disable",
		MaxOpenConns:    pkgPostgreMaxCon,
		MaxIdleConns:    pkgPostgreMaxIdleConn,
		ConnMaxLifetime: time.Duration(pkgPostgreMaxLifeTime) * time.Second,
		ConnMaxIdleTime: time.Duration(pkgPostgreIdleTime) * time.Second,
	})

	if err != nil {
		panic(err)
	}

	return pkg.PkgDatabase{
		PkgRedis:   pkgRedis,
		PkgPostgre: pkgPostgre,
	}

}
