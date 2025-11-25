package apimiddlewareauth

import pkgjwt "github.com/Rizkyprawirap/nextmed-activity-tracking/pkg/jwt"

type middlewareAuth struct {
	PkgJWT pkgjwt.IPkgJWT
}

func New(pkgJWT pkgjwt.IPkgJWT) IMiddlewareAuth {
	return &middlewareAuth{
		PkgJWT: pkgJWT,
	}
}
