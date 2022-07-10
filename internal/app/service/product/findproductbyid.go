package service

import (
	"ms-entity/internal/infra/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FindProductByIdService struct {
	FindProductByIdPort repository.FindProductByIdPort
}

func NewFindProductByIdService(port repository.FindProductByIdPort) *FindProductByIdService {
	return &FindProductByIdService{
		FindProductByIdPort: port,
	}
}

func (pr FindProductByIdService) Execute(ctx echo.Context) error {
	id := ctx.Param("id")
	product, err := pr.FindProductByIdPort.FindById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, product)
}

func (pr FindProductByIdService) RegisterServer(server *echo.Echo) {
	server.GET("/products/:id", pr.Execute)
}
