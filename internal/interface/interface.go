package inf

import (
	"https://github.com/gmmoshopping/GoFramework/router"
	"https://github.com/gmmoshopping/GoFrameworkutils"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Handler struct {
	Store          *gorm.DB
	Version        string
	CircuitBreaker appcore_utils.CircuitBreaker
}

type Module interface {
	ModuleAPI(r *appcore_router.Router)
	GrpcServer() *grpc.Server
}
