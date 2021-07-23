package handler

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"os"
	"testing"
	"zero/mxshop/user_srv/global"
	"zero/mxshop/user_srv/model"
	"zero/mxshop/user_srv/proto"
)

var DB *gorm.DB
var conn *grpc.ClientConn
var userClient proto.UserClient

func setup() {
	// 建立连接
	dsn := "root:ismewen@tcp(127.0.0.1:3306)/mx_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	db := global.CreateDB(dsn)
	DB = db
	// migrate 表
	db.AutoMigrate(&model.User{})
	u := model.User{}
	// 删除所有数据
	db.Where("1 = 1").Delete(&u)

	// 重建数据
	for i := 0; i < 10; i++ {
		pwd, _ := u.GetMd5Str(fmt.Sprintf("%d", i))
		user := model.User{
			NickName: fmt.Sprintf("Somebody-%d", i),
			Mobile:   fmt.Sprintf("9999%d", i),
			Password: pwd,
		}
		db.Save(&user)

	}

	// grpc 服务
	addr := "0.0.0.0:8181"
	c, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
		panic("grpc client init failed,请确保 server端 是否已经起来")
	}
	userClient = proto.NewUserClient(c)

}

func TestUserServer_List(t *testing.T) {
	paginator := proto.Paginator{
		PageSize: 2,
		PageNum:  1,
	}
	listRes, err := userClient.List(context.Background(), &paginator)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("查询成功")
	fmt.Println(listRes.Total)
	//DB.Model(&UserInfo{}).Where("name <> ?","daw").Count(&count)

	var count int64 = 0
	tx := DB.Model(&model.User{}).Where("1 = 1").Count(&count)

	if tx.Error != nil {
		t.Error(tx.Error.Error())
	}

	if count != int64(listRes.Total) {
		t.Error(fmt.Sprintf("total 结果不匹配, %d != %d", count, listRes.Total))
	}

	if paginator.PageSize != uint32(len(listRes.Results)) {
		t.Error(fmt.Sprintf("page size 结果不匹配, %d != %d", count, len(listRes.Results)))
	}

}

func TestUserServer_Retrieve(t *testing.T) {
	idRequest := &proto.RetrieveRequest{
		Mobile: "127271",
	}
	_, err := userClient.Retrieve(context.Background(), idRequest)
	if err == nil {
		t.Error("Not raise not found error")
		return
	}
	u := model.User{}
	tx := DB.First(&u)
	if tx.Error != nil {
		fmt.Println("hello world")
		t.Error("query user failed")
		return
	}
	idRequest = &proto.RetrieveRequest{
		Mobile: u.Mobile,
	}
	user, err := userClient.Retrieve(context.Background(), idRequest)
	if err != nil {
		t.Error("retrieve err" + err.Error())
		return
	}
	if user.Id != u.ID && user.NickName != u.NickName && user.Mobile != u.Mobile {
		t.Error("query failed")
		return
	}
}

func TestUserServer_Update(t *testing.T) {
	user := model.User{}
	tx := DB.First(&user)
	if tx.Error != nil {
		t.Error(tx.Error.Error())
		return
	}

	updateInfo := proto.UpdateUserInfo{
		Id:       user.ID,
		NickName: "NewNickName",
	}
	_, err := userClient.Update(context.Background(), &updateInfo)
	if err != nil {
		t.Error(err.Error())
		return
	}
	tx = DB.First(&user)
	if tx.Error != nil {
		t.Error(tx.Error.Error())
		return
	}
	if user.NickName != "NewNickName" {
		t.Error("update failed")
		return
	}

}

func TestUserServer_CheckPassWord(t *testing.T) {
	user := model.User{}
	tx := DB.First(&user)
	if tx.Error != nil {
		t.Error("query failed")
		return
	}
	pc := proto.PasswordCheckInfo{
		Password:          "test",
		EncryptedPassword: user.Password,
	}
	pcRes, err := userClient.CheckPassWord(context.Background(), &pc)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if pcRes.IsCorrect {
		t.Error("unexpected")
		return
	}
	user.Password, _ = user.GetMd5Str("0")
	tx = DB.Save(&user)
	if tx.Error != nil {
		t.Error(tx.Error.Error())
		return
	}

	pc = proto.PasswordCheckInfo{
		Password:          "0",
		EncryptedPassword: user.Password,
	}
	fmt.Printf("pass: %s, enpass: %s, user: %s\n", pc.Password, pc.EncryptedPassword, user.Password)
	pcRes, _ = userClient.CheckPassWord(context.Background(), &pc)

	if !pcRes.IsCorrect {
		t.Error("unexpected")
		return
	}

}
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}
