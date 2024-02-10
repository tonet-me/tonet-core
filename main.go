package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/tonet-me/tonet-core/config"
	"github.com/tonet-me/tonet-core/entity"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
	mongodb "github.com/tonet-me/tonet-core/repository/mongo"
	usermongo "github.com/tonet-me/tonet-core/repository/mongo/user"
)

func main() {
	cfx := config.C()
	clientCfg := mongodb.Config{
		Host:     cfx.MongoClient.Host,
		Port:     cfx.MongoClient.Port,
		Username: cfx.MongoClient.Username,
		Password: cfx.MongoClient.Password,
	}
	db := mongodb.New(clientCfg)
	userDB := usermongo.New(usermongo.Config{
		DBName:   "test",
		CollName: "user",
	}, db)

	//isEist, user, err := userDB.IsUserExistByEmail(context.TODO(), "kswsssss@gmail.com")
	//if err != nil {
	//	fmt.Println("err", err)
	//}
	//s, err := userDB.DeActiveUser(context.TODO(), "65c728c64bb1081b4046d682")
	s, err := userDB.UpdateUser(context.TODO(), "65c728c64bb1081b4046d68x", entity.User{
		//ID:              "",
		FirstName:       "q1q",
		LastName:        "qqq",
		Email:           "q@gmail.com",
		PhoneNumber:     "0912",
		ProfilePhotoURL: "qq",
		Status:          2,
	})
	richErr := new(richerror.RichError)
	if errors.As(err, &richErr) {
		fmt.Println("rich ", richErr.Message())
		fmt.Println("kind ", richErr.Kind())
	} else {
		fmt.Println("no rich", err)
	}
	fmt.Println(s, err)
	//user, err := userDB.CreateNewUser(context.TODO(), entity.User{
	//	LastName:    "k",
	//	Email:       "kswssss@gmail.com",
	//	PhoneNumber: "091221702858",
	//	Status:      1,
	//})
	//if err != nil {
	//	fmt.Print(err)
	//}
	//fmt.Printf("%s\n", user)
	//result, err := db.GetClient().Database("test").Collection("movie").InsertOne(
	//	context.TODO(),
	//	bson.D{
	//		{"item", "canvas"},
	//		{"qty", 100},
	//		{"tags", bson.A{"cotton"}},
	//		{"size", bson.D{
	//			{"h", 28},
	//			{"w", 35.5},
	//			{"uom", "cm"},
	//		}},
	//	})
	//if err != nil {
	//	fmt.Print(err)
	//}
	//fmt.Printf("%s\n", result)
}
