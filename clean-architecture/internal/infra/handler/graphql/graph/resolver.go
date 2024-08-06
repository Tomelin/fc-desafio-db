package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"github.com/Tomelin/fc-desafio-db/clean-architecture/internal/core/service"
)

type Resolver struct{
	Service service.ServiceOrderInterface
}
