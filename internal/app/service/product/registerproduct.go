package service

import (
	"errors"
	"ms-entity/internal/app/model"
	"ms-entity/internal/app/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RegisterProductService struct {
	SaveProductPort repository.SaveProductPort
}

func NewRegisterProductService(port repository.SaveProductPort) *RegisterProductService {
	return &RegisterProductService{
		SaveProductPort: port,
	}
}

func (pr RegisterProductService) Execute(ctx echo.Context) error {
	request := &model.NewProductRequest{}
	if err := ctx.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if err := validateRequest(*request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	product := model.NewProduct(request)
	result, err := pr.SaveProductPort.Save(*product)
	if err != nil {
		echo.NewHTTPError(http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusOK, result)
}

func validateRequest(request model.NewProductRequest) error {
	if request.Description == "" || request.Coast == 0 || request.Price == 0 {
		return errors.New("Invalid data")
	}

	return nil
}

func (pr RegisterProductService) RegisterServer(server *echo.Echo) {
	server.POST("/products", pr.Execute)
}
