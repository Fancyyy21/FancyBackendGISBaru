package peda

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// }
// func TestGCFCreateHandler(t *testing.T) {
// 	// Simulate input parameters
// 	MONGOCONNSTRINGENV := "mongodb://raulgantengbanget:0nGCVlPPoCsXNhqG@ac-oilbpwk-shard-00-00.9ofhjs3.mongodb.net:27017,ac-oilbpwk-shard-00-01.9ofhjs3.mongodb.net:27017,ac-oilbpwk-shard-00-02.9ofhjs3.mongodb.net:27017/test?replicaSet=atlas-13x7kp-shard-0&ssl=true&authSource=admin"
// 	dbname := "petapedia"
// 	collectionname := "user"

// 	// Create a test User
// 	datauser := User{
// 		Username: "testuser",
// 		Password: "testpassword",
// 		Role:     "user",
// 	}

// 	// Call the handler function
// 	result := GCFCreateHandler(MONGOCONNSTRINGENV, dbname, collectionname, datauser)
// 	fmt.Println(result)
// 	// You can add assertions here to validate the result, or check the database for the created user.
// }

func TestCreateNewUserRole(t *testing.T) {
	var userdata User
	userdata.Username = "raulmahya"
	userdata.Password = "banget"
	userdata.Role = "admin"
	mconn := SetConnection("MONGOULBI", "petapedia")
	CreateNewUserRole(mconn, "user", userdata)
}

func TestCreateNewUserToken(t *testing.T) {
	// Create a User struct
	var userdata User
	userdata.Username = "asd"
	userdata.Password = "banget"
	userdata.Role = "admin"

	// Generate private and public keys using watoken.GenerateKey
	privateKey, publicKey := watoken.GenerateKey()

	// Store the private and public keys in the userdata
	userdata.Private = privateKey
	userdata.Publick = publicKey // Corrected the field name from Publick to Public

	// Encode a token using the privateKey
	hasil, err := watoken.Encode("banget", privateKey)
	fmt.Println(hasil, err)
	if err != nil {
		t.Errorf("Failed to create user and token: %v", err)
	} else {
		t.Logf("User and token created successfully")

		// Assuming you have a MongoDB client and a database connection, use the client and connection to insert the userdata
		// Replace "yourDatabaseName" with your actual database name
		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://raulgantengbanget:0nGCVlPPoCsXNhqG@ac-oilbpwk-shard-00-00.9ofhjs3.mongodb.net:27017,ac-oilbpwk-shard-00-01.9ofhjs3.mongodb.net:27017,ac-oilbpwk-shard-00-02.9ofhjs3.mongodb.net:27017/test?replicaSet=atlas-13x7kp-shard-0&ssl=true&authSource=admin"))
		if err != nil {
			t.Errorf("Failed to create MongoDB client: %v", err)
		} else {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			err = client.Connect(ctx)
			if err != nil {
				t.Errorf("Failed to connect to MongoDB: %v", err)
			} else {
				// Use the database name and collection name where you want to insert the user data
				db := client.Database("petapedia")
				collection := db.Collection("user")

				_, err = collection.InsertOne(ctx, userdata)
				if err != nil {
					t.Errorf("Failed to insert user data into MongoDB: %v", err)
				} else {
					t.Logf("User data inserted into MongoDB successfully")
				}
			}
		}
	}
}

func TestDeleteUser(t *testing.T) {

	mconn := SetConnection("mongodb://raulgantengbanget:0nGCVlPPoCsXNhqG@ac-oilbpwk-shard-00-00.9ofhjs3.mongodb.net:27017,ac-oilbpwk-shard-00-01.9ofhjs3.mongodb.net:27017,ac-oilbpwk-shard-00-02.9ofhjs3.mongodb.net:27017/test?replicaSet=atlas-13x7kp-shard-0&ssl=true&authSource=admin", "petapedia")
	var userdata User
	userdata.Username = "maulana"
	DeleteUser(mconn, "user", userdata)
}

func TestGFCPostHandlerUser(t *testing.T) {
	mconn := SetConnection("MONGOULBI", "petapedia")
	var userdata User
	userdata.Username = "raulmahya"
	userdata.Password = "banget"
	userdata.Role = "admin"
	CreateNewUserRole(mconn, "user", userdata)
}

func TestFunciionUser(t *testing.T) {
	mconn := SetConnection("MONGOULBI", "petapedia")
	var userdata User
	userdata.Username = "raulmahya"
	userdata.Password = "banget"
	userdata.Role = "admin"
	CreateNewUserRole(mconn, "user", userdata)
}

func TestProduct(t *testing.T) {
	mconn := SetConnection("MONGOULBI", "petapedia")
	var productdata Product
	productdata.Nomorid = 1
	productdata.Name = "raul"
	productdata.Description = "mahya"
	productdata.Price = 1000
	productdata.Size = "XL"
	productdata.Stock = 100
	productdata.Image = "https://images3.alphacoders.com/165/thumb-1920-165265.jpg"
	CreateNewProduct(mconn, "product", productdata)
}

func TestAllProduct(t *testing.T) {
	mconn := SetConnection("MONGOULBI", "petapedia")
	product := GetAllProduct(mconn, "product")
	fmt.Println(product)
}

func TestGeneratePasswordHashh(t *testing.T) {
	password := "secret"
	hash, _ := HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
}
func TestHashFunctionn(t *testing.T) {
	mconn := SetConnection("MONGOULBI", "petapedia")
	var userdata User
	userdata.Username = "zz"
	userdata.Password = "mahya"

	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[User](mconn, "user", filter)
	fmt.Println("Mongo User Result: ", res)
	hash, _ := HashPassword(userdata.Password)
	fmt.Println("Hash Password : ", hash)
	match := CheckPasswordHash(userdata.Password, res.Password)
	fmt.Println("Match:   ", match)

}
func TestFindUser(t *testing.T) {
	var userdata User
	userdata.Username = "petped"
	mconn := SetConnection("mongodb://raulgantengbanget:0nGCVlPPoCsXNhqG@ac-oilbpwk-shard-00-00.9ofhjs3.mongodb.net:27017,ac-oilbpwk-shard-00-01.9ofhjs3.mongodb.net:27017,ac-oilbpwk-shard-00-02.9ofhjs3.mongodb.net:27017/test?replicaSet=atlas-13x7kp-shard-0&ssl=true&authSource=admin", "petapedia")
	res := FindUser(mconn, "user", userdata)
	fmt.Println(res)
}

func TestGeneratePasswordHash(t *testing.T) {
	password := "ganteng"
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
	hasil, err := watoken.Encode("bangsat", privateKey)
	fmt.Println(hasil, err)
}

func TestHashFunction(t *testing.T) {
	mconn := SetConnection("mongodb://raulgantengbanget:0nGCVlPPoCsXNhqG@ac-oilbpwk-shard-00-00.9ofhjs3.mongodb.net:27017,ac-oilbpwk-shard-00-01.9ofhjs3.mongodb.net:27017,ac-oilbpwk-shard-00-02.9ofhjs3.mongodb.net:27017/test?replicaSet=atlas-13x7kp-shard-0&ssl=true&authSource=admin", "petapedia")
	var userdata User
	userdata.Username = "bangsat"
	userdata.Password = "ganteng"

	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[User](mconn, "user", filter)
	fmt.Println("Mongo User Result: ", res)
	hash, _ := HashPassword(userdata.Password)
	fmt.Println("Hash Password : ", hash)
	match := CheckPasswordHash(userdata.Password, res.Password)
	fmt.Println("Match:   ", match)

}

func TestIsPasswordValid(t *testing.T) {
	mconn := SetConnection("mongodb://raulgantengbanget:0nGCVlPPoCsXNhqG@ac-oilbpwk-shard-00-00.9ofhjs3.mongodb.net:27017,ac-oilbpwk-shard-00-01.9ofhjs3.mongodb.net:27017,ac-oilbpwk-shard-00-02.9ofhjs3.mongodb.net:27017/test?replicaSet=atlas-13x7kp-shard-0&ssl=true&authSource=admin", "petapedia")
	var userdata User
	userdata.Username = "bangsat"
	userdata.Password = "ganteng"

	anu := IsPasswordValid(mconn, "user", userdata)
	fmt.Println(anu)
}

// --------------------------------------------------------------------- START GIS 9 ---------------------------------------------------------------------

func TestUpdateGetData(t *testing.T) {
	mconn := SetConnectionTest("mongodb+srv://Fancy21:Acermaulana21@maulana.uiq9nmy.mongodb.net/?retryWrites=true&w=majority", "gis")
	datagedung := GeoIntersects(mconn, 108.22501803948785, -6.835460702759789)
	fmt.Println(datagedung)
}
