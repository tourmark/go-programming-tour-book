package v1

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
)

type UserPO struct {
	Id int32 `json:"id"`
	Name sql.NullString `json:"name"`
	Create_time sql.NullString `json:"create_time"`
}

type UserDTO struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	Create_time string `json:"create_time"`
}

func NewUser() UserDTO {
	return UserDTO{}
}

func (t UserDTO) List(c *gin.Context,db *sql.DB)[]UserDTO{
	var userList []UserPO
	mSql := "select * from user"
	rows, err := db.Query(mSql)
	if err !=nil{
		log.Fatalf("error is %v/n",err)
		return nil
	}
	for rows.Next(){
		var user UserPO
		err := rows.Scan(&user.Id,&user.Name,&user.Create_time)
		if err !=nil{
			log.Fatalf("error is %v/n",err)
			return nil
		}
		userList = append(userList, user)
	}
	var userListNew []UserDTO
	for _,x:=range userList{
		var userNew UserDTO
		userNew.Id=x.Id
		if x.Name.Valid {
			userNew.Name=x.Name.String
		}else {
			userNew.Name="NULL"
		}
		if x.Create_time.Valid {
			userNew.Create_time=x.Create_time.String
		}else {
			userNew.Create_time="NULL"
		}
		userListNew = append(userListNew, userNew)
	}

	return userListNew
}

