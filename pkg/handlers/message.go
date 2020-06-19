package handlers

import (
	"context"
	"errors"
	v1 "myproject/pkg/apis/v1"
)


var userMap = make(map[string]v1.User)

//Add a user
func AddUser(ctx context.Context, Identifier string,Name string,Age string)(v1.Operate, error){
	var newUser v1.User
	newUser.Identifier=Identifier
	newUser.Name=Name
	newUser.Age=Age
	userMap[Identifier] = newUser

	return v1.Operate{
		Message: "Add a user successfully",
	},nil
}


// AllUsers returns all users' information.
func AllUsers(ctx context.Context) ([]v1.User, error) {
	users := make([]v1.User, len(userMap))
	var i =0
	for key := range userMap {
		thisUser,_:=userMap[key]
		users[i].Identifier=thisUser.Identifier
		users[i].Name=thisUser.Name
		users[i].Age=thisUser.Age
		i++
	}
	return users, nil
}

// OneUser return a user by id.
func OneUser(ctx context.Context, id string) (v1.User, error) {
	checkUser, ok := userMap[id]
	if ok==false{
		noneUser:=v1.User{

		}
		return noneUser,errors.New("the user does not exist")
	} else{
		return checkUser,nil
	}
}

// OneUser return a user by id.
func DeleteUser(ctx context.Context, id string) (v1.Operate, error) {
	_, ok := userMap[id]
	if ok==false{
		noneOperate:=v1.Operate{
			Message: "Failed",
		}
		return noneOperate,errors.New("the user does not exist")
	} else{
		delete(userMap, id)
		return v1.Operate{
			Message: "Delete a user successfully",
		},nil
	}
}

func ChangeInfo(ctx context.Context, Identifier string,Name string,Age string)(v1.Operate, error){
	var tempUser v1.User
	_, ok := userMap[Identifier]
	OperateMess:=v1.Operate{}
	if ok==false{
		OperateMess.Message="Filed."
		return OperateMess,errors.New("the user does not exist")
	}else{
		tempUser.Identifier = userMap[Identifier].Identifier
		tempUser.Name = userMap[Identifier].Name
		tempUser.Age = userMap[Identifier].Age
		if len(Name)==0 && len(Age)==0{
			OperateMess.Message= "There is no change for the user."

		}else {
			if len(Age)!=0{
				tempUser.Age=Age
				OperateMess.Message="User's age is changed."
			}
			if len(Name)!=0{
				tempUser.Name=Name
				OperateMess.Message="User's name is changed."
			}
			if len(Age)!=0&&len(Name)!=0{
				OperateMess.Message="User's name and age is changed."
			}

		}
		userMap[Identifier]=tempUser
	}
	return OperateMess,nil

}
