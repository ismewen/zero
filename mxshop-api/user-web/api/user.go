package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
	"time"
	"zero/mxshop-api/user-web/forms"
	"zero/mxshop-api/user-web/proto"
	"zero/mxshop-api/user-web/response"
	"zero/mxshop-api/user-web/srv"
)

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			default:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "Bad Request",
				})

			}
		}
	}
}

func QueryValueConvertToInt(value string, defaultValue int) int {
	v, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return v
}

func GetUserList(ctx *gin.Context) {
	svc := srv.NewServiceContext()

	ps := ctx.DefaultQuery("PageSize", "10")
	pn := ctx.DefaultQuery("PageNum", "1")

	paginator := proto.Paginator{
		PageSize: uint32(QueryValueConvertToInt(ps, 10)),
		PageNum:  uint32(QueryValueConvertToInt(pn, 1)),
	}

	userList, err := svc.UserRpc.List(context.Background(), &paginator)
	if err != nil {
		zap.S().Errorw("query failed", "error", err.Error())
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	datas := make([]response.UserResponse, 0)
	for _, value := range userList.Results {
		data := response.UserResponse{
			Id:       value.Id,
			NickName: value.NickName,
			BirthDay: response.JsonTime(time.Unix(int64(value.BirthDay), 0)),
			Gender:   value.Gender,
		}
		datas = append(datas, data)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"RichMan":  "IsMeWen",
		"pageSize": paginator.PageSize,
		"pageNum":  paginator.PageNum,
		"results":  datas,
	})
}

func Login(ctx *gin.Context) {
	form := forms.PasswordLoginForm{}
	err := ctx.ShouldBindJSON(&form)
	if err != nil {
		// 校验错误
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "todo",
	})
}
