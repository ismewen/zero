package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"zero/mxshop-api/user-web/proto"
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

func GetUserList(ctx *gin.Context) {
	svc := srv.NewServiceContext()
	paginator := proto.Paginator{
		PageSize: 2,
		PageNum:  1,
	}
	userList, err := svc.UserRpc.List(context.Background(), &paginator)
	if err != nil {
		zap.S().Errorw("query failed", "error", err.Error())
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	datas := make([]interface{}, 0)
	for _, value := range userList.Results {
		data := make(map[string]interface{})
		data["id"] = value.Id
		data["name"] = value.NickName
		data["birthday"] = value.BirthDay
		data["gender"] = value.Gender
		data["mobile"] = value.Mobile
		datas = append(datas, data)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"pageSize": paginator.PageSize,
		"pageNum":  paginator.PageNum,
		"results":  datas,
	})
}
