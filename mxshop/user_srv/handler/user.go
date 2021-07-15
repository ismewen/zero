package handler

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"

	"zero/mxshop/user_srv/global"
	"zero/mxshop/user_srv/model"
	"zero/mxshop/user_srv/proto"
)

func UserSerializer(user *model.User) *proto.UserInfoResponse {
	res := &proto.UserInfoResponse{
		Id:       user.ID,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
	}
	if user.Birthday != nil {
		res.BirthDay = uint64(user.Birthday.Unix())
	}
	return res
}

type UserServer struct {
	*proto.UnimplementedUserServer
}

func (s *UserServer) List(ctx context.Context, paginator *proto.Paginator) (*proto.UserListResponse, error) {
	var users []model.User
	res := global.DB.Find(&users)
	if res.Error != nil {
		// 直接返回错误
		return nil, res.Error
	}
	rsp := &proto.UserListResponse{}
	rsp.Total = int32(res.RowsAffected)
	global.DB.Scopes(global.Paginate(int(paginator.PageNum), int(paginator.PageSize))).Find(&users)
	for _, user := range users {
		userInfo := UserSerializer(&user)
		rsp.Results = append(rsp.Results, userInfo)
	}
	return rsp, nil
}

func (s *UserServer) Retrieve(ctx context.Context, request *proto.IdRequest) (*proto.UserInfoResponse, error) {
	var user *model.User
	res := global.DB.Where("id = ?", request.Id).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return UserSerializer(user), nil
}

func (s *UserServer) Create(ctx context.Context, info *proto.CreateUserInfo) (*proto.UserInfoResponse, error) {
	var user *model.User
	res := global.DB.Where("mobile = ?", info.Mobile).First(&user)
	if res.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已经存在")
	}
	user.Mobile = info.Mobile
	user.NickName = info.NickName
	pwd, _ := user.GetMd5Str(info.Password)
	user.Password = pwd
	dbRes := global.DB.Create(user)
	if dbRes.Error != nil {
		return nil, status.Errorf(codes.Internal, dbRes.Error.Error())
	}
	userInfo := UserSerializer(user)
	return userInfo, nil
}

func (s *UserServer) Update(ctx context.Context, info *proto.UpdateUserInfo) (*proto.UserInfoResponse, error) {
	var user model.User
	res := global.DB.First(&user, info.Id)
	if res.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "用户不存在")
	}
	birthDay := time.Unix(int64(info.Birthday), 0)
	user.NickName = info.NickName
	user.Birthday = &birthDay
	user.Gender = info.Gender
	dbRes := global.DB.Save(&user)
	if dbRes.Error != nil {
		return nil, status.Errorf(codes.Internal, dbRes.Error.Error())
	}
	userInfoRes := UserSerializer(&user)
	return userInfoRes, nil
}

func (s *UserServer) CheckPassWord(ctx context.Context, info *proto.PasswordCheckInfo) (*proto.CheckResponse, error) {
	user := model.User{}
	mpwd, _ := user.GetMd5Str(info.Password)
	fmt.Printf("pwd: %s, enpwd: %s, userpwd: %s",info.Password, mpwd, info.EncryptedPassword)
	response := &proto.CheckResponse{IsCorrect: mpwd == info.EncryptedPassword}
	return response, nil
}
