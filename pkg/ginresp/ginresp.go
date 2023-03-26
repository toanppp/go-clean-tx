package ginresp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/toanppp/go-clean-tx/pkg/response"
)

func ResponseSuccess[T any](ctx *gin.Context, data T) {
	ctx.JSON(http.StatusOK, response.Response[T]{
		Data:    data,
		Message: http.StatusText(http.StatusOK),
	})

}

func ResponseFail(ctx *gin.Context, code int, message string, err error) {
	if message == "" {
		message = http.StatusText(code)
	}

	resp := response.Response[any]{
		Message: message,
	}

	if err != nil {
		resp.Error = err.Error()
	}

	ctx.AbortWithStatusJSON(code, resp)
}

func ResponseBadRequest(ctx *gin.Context, message string) {
	ResponseFail(ctx, http.StatusBadRequest, message, nil)
}

func ResponseError(ctx *gin.Context, err error) {
	ResponseFail(ctx, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
}
