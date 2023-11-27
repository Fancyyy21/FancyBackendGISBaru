package peda

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
)

func ReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

func GCFHandler(MONGOCONNSTRINGENV, dbname, collectionname string) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	datagedung := GetAllUser(mconn, collectionname)
	return GCFReturnStruct(datagedung)
}

func GCFFindUserByID(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		return err.Error()
	}
	user := FindUser(mconn, collectionname, datauser)
	return GCFReturnStruct(user)
}

func GCFFindUserByName(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		return err.Error()
	}

	// Jika username kosong, maka respon "false" dan data tidak ada
	if datauser.Username == "" {
		return "false"
	}

	// Jika ada username, mencari data pengguna
	user := FindUserUser(mconn, collectionname, datauser)

	// Jika data pengguna ditemukan, mengembalikan data pengguna dalam format yang sesuai
	if user != (User{}) {
		return GCFReturnStruct(user)
	}

	// Jika tidak ada data pengguna yang ditemukan, mengembalikan "false" dan data tidak ada
	return "false"
}

func GCFDeleteHandler(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		return err.Error()
	}
	DeleteUser(mconn, collectionname, datauser)
	return GCFReturnStruct(datauser)
}

func GCFUpdateHandler(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		return err.Error()
	}
	ReplaceOneDoc(mconn, collectionname, bson.M{"username": datauser.Username}, datauser)
	return GCFReturnStruct(datauser)
}

// add encrypt password to database and tokenstring
// func GCFCreateHandler(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {

// 	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
// 	var datauser User
// 	err := json.NewDecoder(r.Body).Decode(&datauser)
// 	if err != nil {
// 		return err.Error()
// 	}
// 	CreateNewUserRole(mconn, collectionname, datauser)
// 	return GCFReturnStruct(datauser)
// }

func GCFCreateHandlerTokenPaseto(PASETOPRIVATEKEYENV, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		return err.Error()
	}
	hashedPassword, hashErr := HashPassword(datauser.Password)
	if hashErr != nil {
		return hashErr.Error()
	}
	datauser.Password = hashedPassword
	CreateNewUserRole(mconn, collectionname, datauser)
	tokenstring, err := watoken.Encode(datauser.Username, os.Getenv(PASETOPRIVATEKEYENV))
	if err != nil {
		return err.Error()
	}
	datauser.Token = tokenstring
	return GCFReturnStruct(datauser)
}

func GCFCreateHandler(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		return err.Error()
	}

	// Hash the password before storing it
	hashedPassword, hashErr := HashPassword(datauser.Password)
	if hashErr != nil {
		return hashErr.Error()
	}
	datauser.Password = hashedPassword

	createErr := CreateNewUserRole(mconn, collectionname, datauser)
	fmt.Println(createErr)

	return GCFReturnStruct(datauser)
}
func GFCPostHandlerUser(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	var Response Credential
	Response.Status = false

	// Mendapatkan data yang diterima dari permintaan HTTP POST
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		Response.Message = "error parsing application/json: " + err.Error()
	} else {
		// Menggunakan variabel MONGOCONNSTRINGENV untuk string koneksi MongoDB
		mongoConnStringEnv := MONGOCONNSTRINGENV

		mconn := SetConnection(mongoConnStringEnv, dbname)

		// Lakukan pemeriksaan kata sandi menggunakan bcrypt
		if IsPasswordValid(mconn, collectionname, datauser) {
			Response.Status = true
			Response.Message = "Selamat Datang"
		} else {
			Response.Message = "Password Salah"
		}
	}

	// Mengirimkan respons sebagai JSON
	responseJSON, _ := json.Marshal(Response)
	return string(responseJSON)
}

func GCFPostHandler(PASETOPRIVATEKEYENV, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	var Response Credential
	Response.Status = false
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		Response.Message = "error parsing application/json: " + err.Error()
	} else {
		if IsPasswordValid(mconn, collectionname, datauser) {
			Response.Status = true
			tokenstring, err := watoken.Encode(datauser.Username, os.Getenv(PASETOPRIVATEKEYENV))
			if err != nil {
				Response.Message = "Gagal Encode Token : " + err.Error()
			} else {
				Response.Message = "Selamat Datang"
				Response.Token = tokenstring
			}
		} else {
			Response.Message = "Password Salah"
		}
	}

	return GCFReturnStruct(Response)
}

func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

// product
func GCFGetAllProduct(MONGOCONNSTRINGENV, dbname, collectionname string) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	datagedung := GetAllProduct(mconn, collectionname)
	return GCFReturnStruct(datagedung)
}

func GCFCreateProduct(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var dataproduct Product
	err := json.NewDecoder(r.Body).Decode(&dataproduct)
	if err != nil {
		return err.Error()
	}
	CreateNewProduct(mconn, collectionname, dataproduct)
	return GCFReturnStruct(dataproduct)
}

func GCFLoginTest(username, password, MONGOCONNSTRINGENV, dbname, collectionname string) bool {
	// Membuat koneksi ke MongoDB
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)

	// Mencari data pengguna berdasarkan username
	filter := bson.M{"username": username}
	collection := collectionname
	res := atdb.GetOneDoc[User](mconn, collection, filter)

	// Memeriksa apakah pengguna ditemukan dalam database
	if res == (User{}) {
		return false
	}

	// Memeriksa apakah kata sandi cocok
	return CheckPasswordHash(password, res.Password)
}

func InsertDataUserGCF(Mongoenv, dbname string, r *http.Request) string {
	resp := new(Credential)
	userdata := new(User)
	resp.Status = false
	conn := SetConnection(Mongoenv, dbname)
	err := json.NewDecoder(r.Body).Decode(&userdata)
	if err != nil {
		resp.Message = "error parsing application/json: " + err.Error()
	} else {
		resp.Status = true
		hash, err := HashPassword(userdata.Password)
		if err != nil {
			resp.Message = "Gagal Hash Password" + err.Error()
		}
		InsertUserdata(conn, userdata.Username, userdata.Role, hash)
		resp.Message = "Berhasil Input data"
	}
	return GCFReturnStruct(resp)
}

func GCFCreatePostLineStringg(MONGOCONNSTRINGENV, dbname, collection string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var geojsonline GeoJsonLineString
	err := json.NewDecoder(r.Body).Decode(&geojsonline)
	if err != nil {
		return err.Error()
	}

	// Mengambil nilai header PASETO dari permintaan HTTP
	pasetoValue := r.Header.Get("PASETOPRIVATEKEYENV")

	// Disini Anda dapat menggunakan nilai pasetoValue sesuai kebutuhan Anda
	// Misalnya, menggunakannya untuk otentikasi atau enkripsi.
	// Contoh sederhana menambahkan nilainya ke dalam pesan respons:
	response := GCFReturnStruct(geojsonline)
	response += " PASETO value: " + pasetoValue

	PostLinestring(mconn, collection, geojsonline)
	return response
}

func GCFCreatePostLineString(MONGOCONNSTRINGENV, dbname, collection string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var geojsonline GeoJsonLineString
	err := json.NewDecoder(r.Body).Decode(&geojsonline)
	if err != nil {
		return err.Error()
	}
	PostLinestring(mconn, collection, geojsonline)
	return GCFReturnStruct(geojsonline)
}

func AmbilDataGeojsonToken(mongoenv, dbname, collname string, r *http.Request) string {
	var atmessage PostToken
	if r.Header.Get("token") == os.Getenv("TOKEN") {
		mconn := SetConnection(mongoenv, dbname)
		datagedung := GetAllBangunanLineString(mconn, collname)
		var geojsonpoint GeoJsonPoint
		err := json.NewDecoder(r.Body).Decode(&geojsonpoint)
		if err != nil {
			atmessage.Response = "error parsing application/json: " + err.Error()
		} else {
			PostPoint(mconn, collname, geojsonpoint)
			atmessage, _ = PostStructWithToken[PostToken]("token", os.Getenv("TOKEN"), datagedung, "https://asia-southeast2-befous.cloudfunctions.net/Befous-AmbilDataGeojson")
		}
	} else {
		atmessage.Response = "Token Salah"
	}
	return GCFReturnStruct(atmessage)
}

func AmbilDataGeojson(mongoenv, dbname, collname string) string {
	mconn := SetConnection(mongoenv, dbname)
	datagedung := GetAllBangunanLineString(mconn, collname)
	return GCFReturnStruct(datagedung)
}

func MembuatGeojsonPointToken(mongoenv, dbname, collname string, r *http.Request) string {
	var atmessage PostToken
	if r.Header.Get("token") == os.Getenv("TOKEN") {
		mconn := SetConnection(mongoenv, dbname)
		var geojsonpoint GeoJsonPoint
		err := json.NewDecoder(r.Body).Decode(&geojsonpoint)
		if err != nil {
			atmessage.Response = "error parsing application/json: " + err.Error()
		} else {
			PostPoint(mconn, collname, geojsonpoint)
			atmessage, _ = PostStructWithToken[PostToken]("token", os.Getenv("TOKEN"), geojsonpoint, "https://asia-southeast2-befous.cloudfunctions.net/Befous-MembuatGeojsonPoint")
		}
	} else {
		atmessage.Response = "Token Salah"
	}
	return GCFReturnStruct(atmessage)
}

func MembuatGeojsonPolylineToken(mongoenv, dbname, collname string, r *http.Request) string {
	var atmessage PostToken
	if r.Header.Get("token") == os.Getenv("TOKEN") {
		mconn := SetConnection(mongoenv, dbname)
		var geojsonline GeoJsonLineString
		err := json.NewDecoder(r.Body).Decode(&geojsonline)
		if err != nil {
			atmessage.Response = "error parsing application/json: " + err.Error()
		} else {
			PostLinestring(mconn, collname, geojsonline)
			atmessage, _ = PostStructWithToken[PostToken]("token", os.Getenv("TOKEN"), geojsonline, "https://asia-southeast2-befous.cloudfunctions.net/Befous-MembuatGeojsonPolyline")
		}
	} else {
		atmessage.Response = "Token Salah"
	}
	return GCFReturnStruct(atmessage)
}

func MembuatGeojsonPolygonToken(mongoenv, dbname, collname string, r *http.Request) string {
	var atmessage PostToken
	if r.Header.Get("token") == os.Getenv("TOKEN") {
		mconn := SetConnection(mongoenv, dbname)
		var geojsonpolygon GeoJsonPolygon
		err := json.NewDecoder(r.Body).Decode(&geojsonpolygon)
		if err != nil {
			atmessage.Response = "error parsing application/json: " + err.Error()
		} else {
			PostPolygon(mconn, collname, geojsonpolygon)
			atmessage, _ = PostStructWithToken[PostToken]("token", os.Getenv("TOKEN"), geojsonpolygon, "https://asia-southeast2-befous.cloudfunctions.net/Befous-MembuatGeojsonPolygon")
		}
	} else {
		atmessage.Response = "Token Salah"
	}
	return GCFReturnStruct(atmessage)
}

func MembuatGeojsonPoint(mongoenv, dbname, collname string, r *http.Request) string {
	mconn := SetConnection(mongoenv, dbname)
	var geojsonpoint GeoJsonPoint
	err := json.NewDecoder(r.Body).Decode(&geojsonpoint)
	if err != nil {
		return err.Error()
	}
	PostPoint(mconn, collname, geojsonpoint)
	return GCFReturnStruct(geojsonpoint)
}

func MembuatGeojsonPolyline(mongoenv, dbname, collname string, r *http.Request) string {
	mconn := SetConnection(mongoenv, dbname)
	var geojsonline GeoJsonLineString
	err := json.NewDecoder(r.Body).Decode(&geojsonline)
	if err != nil {
		return err.Error()
	}
	PostLinestring(mconn, collname, geojsonline)
	return GCFReturnStruct(geojsonline)
}

func MembuatGeojsonPolygon(mongoenv, dbname, collname string, r *http.Request) string {
	mconn := SetConnection(mongoenv, dbname)
	var geojsonpolygon GeoJsonPolygon
	err := json.NewDecoder(r.Body).Decode(&geojsonpolygon)
	if err != nil {
		return err.Error()
	}
	PostPolygon(mconn, collname, geojsonpolygon)
	return GCFReturnStruct(geojsonpolygon)
}

// --------------------------------------------------------------------- START GIS 9 ---------------------------------------------------------------------

func PostGeoIntersects(mongoenv, dbname string, r *http.Request) string {
	var longlat LongLat
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)

	err := json.NewDecoder(r.Body).Decode(&longlat)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
	} else {
		response.Status = true
		response.Message = GeoIntersects(mconn, longlat.Longitude, longlat.Latitude)
	}
	return ReturnStruct(response)
}
