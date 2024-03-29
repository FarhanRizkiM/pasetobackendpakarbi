package pasetobackendpakarbi

import (
	"fmt"
	"testing"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCreateNewUserRole(t *testing.T) {
	var userdata User
	userdata.Username = "pakarbi"
	userdata.Password = "lodons"
	userdata.Role = "user"
	mconn := SetConnection("MONGOSTRING", "dbpakarbi")
	CreateNewUserRole(mconn, "user", userdata)
}

// func TestDeleteUser(t *testing.T) {
// 	mconn := SetConnection("MONGOSTRING", "pasabar13")
// 	var userdata User
// 	userdata.Username = "lolz"
// 	DeleteUser(mconn, "user", userdata)
// }

func CreateNewUserToken(t *testing.T) {
	var userdata User
	userdata.Username = "pakarbi"
	userdata.Password = "pakarbipass"
	userdata.Role = "user"

	// Create a MongoDB connection
	mconn := SetConnection("MONGOSTRING", "dbpakarbi")

	// Call the function to create a user and generate a token
	err := CreateUserAndAddToken("your_private_key_env", mconn, "user", userdata)

	if err != nil {
		t.Errorf("Error creating user and token: %v", err)
	}
}

func TestGFCPostHandlerUser(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "dbpakarbi")
	var userdata User
	userdata.Username = "pakarbi"
	userdata.Password = "pakarbipass"
	userdata.Role = "user"
	CreateNewUserRole(mconn, "user", userdata)
}

func TestGeneratePasswordHash(t *testing.T) {
	password := "testpass"
	hash, _ := HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)
	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
}
func TestGeneratePrivateKeyPaseto(t *testing.T) {
	privateKey, publicKey := watoken.GenerateKey()
	fmt.Println(privateKey)
	fmt.Println(publicKey)
	hasil, err := watoken.Encode("testpakarbi", privateKey)
	fmt.Println(hasil, err)
}

func TestHashFunction(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "dbpakarbi")
	var userdata User
	userdata.Username = "rizki"
	userdata.Password = "testpass"

	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[User](mconn, "user", filter)
	fmt.Println("Mongo User Result: ", res)
	hash, _ := HashPassword(userdata.Password)
	fmt.Println("Hash Password : ", hash)
	match := CheckPasswordHash(userdata.Password, res.Password)
	fmt.Println("Match:   ", match)

}

func TestIsPasswordValid(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "dbpakarbi")
	var userdata User
	userdata.Username = "rizki"
	userdata.Password = "testpass"

	anu := IsPasswordValid(mconn, "user", userdata)
	fmt.Println(anu)
}

func TestUserFix(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "dbpakarbi")
	var userdata User
	userdata.Username = "pakarbi"
	userdata.Password = "pakarbipass"
	userdata.Role = "user"
	CreateUser(mconn, "user", userdata)
}
