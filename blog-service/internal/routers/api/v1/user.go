package v1

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
)

type User struct {
	Id int32 `json:"id"`
	Name sql.NullString `json:"name"`
	Create_time sql.NullString `json:"create_time"`
}

func NewUser() User {
	return User{}
}


func (t User) Get(){


}
func (t User) List(c *gin.Context,db *sql.DB)[]User{
	var userList []User
	mSql := "select * from user"
	rows, err := db.Query(mSql)
	if err !=nil{
		log.Fatalf("error is %v/n",err)
		return userList
	}
	for rows.Next(){
		var user User
		err := rows.Scan(&user.Id,&user.Name,&user.Create_time)
		if err !=nil{
			log.Fatalf("error is %v/n",err)
			return userList
		}
		userList = append(userList, user)
	}

	return userList
}
func (t User) Create(c *gin.Context){}
func (t User) Update(c *gin.Context){}
func (t User) Delete(c *gin.Context){}
