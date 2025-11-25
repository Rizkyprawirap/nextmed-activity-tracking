package main

import (
	"fmt"
	"os"

	controllerAdmin "github.com/Rizkyprawirap/nextmed-activity-tracking/api/controllers/admin"
	controllerClients "github.com/Rizkyprawirap/nextmed-activity-tracking/api/controllers/client"
	controllerUsages "github.com/Rizkyprawirap/nextmed-activity-tracking/api/controllers/usage"

	apimodelsadmin "github.com/Rizkyprawirap/nextmed-activity-tracking/api/models/admin"
	apimodelsclient "github.com/Rizkyprawirap/nextmed-activity-tracking/api/models/client"
	apimodelsusage "github.com/Rizkyprawirap/nextmed-activity-tracking/api/models/usage"

	routerAdmin "github.com/Rizkyprawirap/nextmed-activity-tracking/api/routes/admin"
	routerClient "github.com/Rizkyprawirap/nextmed-activity-tracking/api/routes/client"
	routerUsage "github.com/Rizkyprawirap/nextmed-activity-tracking/api/routes/usage"

	internaldatabase "github.com/Rizkyprawirap/nextmed-activity-tracking/internal/database"

	pkgenv "github.com/Rizkyprawirap/nextmed-activity-tracking/pkg/env"
	pkgjwt "github.com/Rizkyprawirap/nextmed-activity-tracking/pkg/jwt"
	pkgratelimit "github.com/Rizkyprawirap/nextmed-activity-tracking/pkg/rate_limit"

	auth "github.com/Rizkyprawirap/nextmed-activity-tracking/api/middleware/auth"

	corsMiddleware "github.com/Rizkyprawirap/nextmed-activity-tracking/api/middleware"
	rateLimitMiddleware "github.com/Rizkyprawirap/nextmed-activity-tracking/api/middleware/rate_limit"

	"github.com/gin-gonic/gin"
)

func main() {
	pkgenv.New()
	pkgDatabaseContract := internaldatabase.New()

	// jwt
	pkgJWTAdmin := pkgjwt.New(
		os.Getenv("JWT_SECRET_ADMIN"),
		pkgDatabaseContract.PkgRedis,
	)
	authenticationAdmin := auth.New(pkgJWTAdmin)

	rateLimitService := pkgratelimit.New(pkgDatabaseContract.PkgRedis)
	rateLimiter := rateLimitMiddleware.New(
		pkgJWTAdmin,
		rateLimitService,
	)

	r := gin.Default()
	r.UseRawPath = true

	apimodelsadmin := apimodelsadmin.New(pkgDatabaseContract.PkgPostgre)
	apimodelsclients := apimodelsclient.New(
		pkgDatabaseContract.PkgPostgre,
		pkgDatabaseContract.PkgRedis,
	)
	apimodelsusage := apimodelsusage.New(
		pkgDatabaseContract.PkgPostgre,
		pkgDatabaseContract.PkgRedis,
	)

	r.Use(corsMiddleware.CorsMiddleware)

	// Client
	controllerClients := controllerClients.New(
		apimodelsclients,
	)
	routerClient.New(
		r.Group("/api"),
		rateLimiter,
		controllerClients,
	)

	// Usage
	controllerUsages := controllerUsages.New(
		apimodelsusage,
		pkgJWTAdmin,
	)
	routerUsage.New(
		r.Group("/api"),
		authenticationAdmin,
		rateLimiter,
		controllerUsages,
	)

	// Admin
	controllerAdmin := controllerAdmin.New(
		apimodelsadmin,
		pkgJWTAdmin,
	)
	routerAdmin.New(
		r.Group("/api"),
		authenticationAdmin,
		rateLimiter,
		controllerAdmin,
	)

	r.Run(fmt.Sprintf("%s%s", ":", os.Getenv("PORT")))
}
