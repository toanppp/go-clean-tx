package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/toanppp/go-clean-tx/internal/infrastructure/http/presenter"
)

func responseSuccess[T any](ctx *gin.Context, data T) {
	ctx.JSON(http.StatusOK, presenter.Response[T]{
		Data:    data,
		Message: http.StatusText(http.StatusOK),
	})

}

func responseFail(ctx *gin.Context, code int, message string, err error) {
	if message == "" {
		message = http.StatusText(code)
	}

	resp := presenter.Response[any]{
		Message: message,
	}

	if err != nil {
		resp.Error = err.Error()
	}

	ctx.AbortWithStatusJSON(code, resp)
}

func responseBadRequest(ctx *gin.Context, message string) {
	responseFail(ctx, http.StatusBadRequest, message, nil)
}

func responseError(ctx *gin.Context, err error) {
	responseFail(ctx, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
}
