package utility

import (
	"fmt"
	"github.com/Vladimir77715/fibonacciservice/redis/fibonaccicashed"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GinWrapper struct {
	RouterGroup *gin.RouterGroup
	Cs          *fibonaccicashed.CashedService
}
type GinWrapperContext struct {
	Context *gin.Context
	Cs      *fibonaccicashed.CashedService
}

type HandlerFunction func(wrapper *GinWrapperContext)

func (gw *GinWrapper) GET(relativePath string, hf HandlerFunction) {
	gw.RouterGroup.GET(relativePath, func(context *gin.Context) {
		hf(&GinWrapperContext{Context: context, Cs: gw.Cs})
	})
}

type Variant struct {
	value interface{}
}

func (v *Variant) ToInt() int {
	if v.value == nil {
		return 0
	}
	var iValue, err = strconv.Atoi(fmt.Sprintf("%v", v.value))
	if err != nil {
		iValue = 0
	}
	return iValue
}

func (gwc *GinWrapperContext) GetVariantQuery(key string) *Variant {
	if q, ok := gwc.Context.GetQuery(key); ok && len(q) > 0 {
		return &Variant{value: q}
	} else {
		return &Variant{value: nil}
	}
}

func (gwc *GinWrapperContext) SuccessResponse(val interface{}) {
	gwc.Context.JSON(http.StatusOK, val)
}
