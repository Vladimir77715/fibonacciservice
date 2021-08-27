package fibonacciservice

import (
	"github.com/Vladimir77715/fibonacciservice/fibonacci"
	"github.com/Vladimir77715/fibonacciservice/redis/fibonaccicashed"
	"github.com/Vladimir77715/fibonacciservice/rest/server/utility"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FibonacciResponse struct {
	FibSlice []uint64 `json:"fib_slice"`
}

func RegisterFibonacciService(router *gin.RouterGroup, cs *fibonaccicashed.CashedService) {
	fibonacciRoute := utility.GinWrapper{Cs: cs, RouterGroup: router.Group("/fibonacci")}
	fibonacciRoute.GET("/fibonacciSlice", fibonacciSlice)
}

func fibonacciSlice(context *utility.GinWrapperContext) {
	var x, y = context.GetVariantQuery("x").ToInt(), context.GetVariantQuery("y").ToInt()
	var fibSlice []uint64
	var err error
	if context.Cs != nil {
		fibSlice, err = context.Cs.FibonacciSlice(x, y)
	} else {
		fibSlice, err = fibonacci.FibonacciSlice(x, y, nil)
	}
	if err != nil {
		context.Context.String(http.StatusBadRequest, err.Error())
		return
	}
	context.SuccessResponse(FibonacciResponse{FibSlice: fibSlice})
}
