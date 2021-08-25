package fibonacciservice

import (
	"fibonacciservice/fibonacci"
	"fibonacciservice/rest/server/utility"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FibonacciResponse struct {
	FibSlice []uint64 `json:"fib_slice"`
}

func RegisterFibonacciService(router *gin.RouterGroup) {
	fibonacciRoute := utility.GinWrapper{RouterGroup: router.Group("/fibonacci")}
	fibonacciRoute.GET("/fibonacciSlice", fibonacciSlice)
}

func fibonacciSlice(context *utility.GinWrapperContext) {
	var x, y = context.GetVariantQuery("x").ToInt(), context.GetVariantQuery("y").ToInt()
	var fibSlice, err = fibonacci.FibonacciSlice(x, y)
	if err != nil {
		context.Context.String(http.StatusBadRequest, err.Error())
		return
	}
	context.SuccessResponse(FibonacciResponse{FibSlice: fibSlice})
}
