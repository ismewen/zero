package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"zero/mxshop-api/user-web/validators"
)

func InitValidator(){
	if v, ok := binding.Validator.Engine().(*validator.Validate);ok {
		err := v.RegisterValidation("mxmobile", validators.ValidateMobile)
		fmt.Println("register")
		fmt.Println("register")
		if err!= nil {
			fmt.Println(err.Error())
		}
	}else{
		fmt.Println("register failed")
	}
}