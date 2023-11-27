package peda

import "go.mongodb.org/mongo-driver/bson/primitive"

type GeometryPolygon struct {
	Coordinates [][][]float64 `json:"coordinates" bson:"coordinates"`
	Type        string        `json:"type" bson:"type"`
}

type GeometryLineString struct {
	Coordinates [][]float64 `json:"coordinates" bson:"coordinates"`
	Type        string      `json:"type" bson:"type"`
}

type GeometryPoint struct {
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
	Type        string    `json:"type" bson:"type"`
}
type GeoJsonPoint struct {
	Type       string        `json:"type" bson:"type"`
	Properties Properties    `json:"properties" bson:"properties"`
	Geometry   GeometryPoint `json:"geometry" bson:"geometry"`
}
type GeoJsonLineString struct {
	Type       string             `json:"type" bson:"type"`
	Properties Properties         `json:"properties" bson:"properties"`
	Geometry   GeometryLineString `json:"geometry" bson:"geometry"`
}

type GeoJsonPolygon struct {
	Type       string          `json:"type" bson:"type"`
	Properties Properties      `json:"properties" bson:"properties"`
	Geometry   GeometryPolygon `json:"geometry" bson:"geometry"`
}

type Geometry struct {
	Coordinates interface{} `json:"coordinates" bson:"coordinates"`
	Type        string      `json:"type" bson:"type"`
}
type GeoJson struct {
	Type       string     `json:"type" bson:"type"`
	Properties Properties `json:"properties" bson:"properties"`
	Geometry   Geometry   `json:"geometry" bson:"geometry"`
}

type Properties struct {
	Name string `json:"name" bson:"name"`
}

type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role,omitempty" bson:"role,omitempty"`
	Token    string `json:"token,omitempty" bson:"token,omitempty"`
	Private  string `json:"private,omitempty" bson:"private,omitempty"`
	Publick  string `json:"publick,omitempty" bson:"publick,omitempty"`
}

type Credential struct {
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type Product struct {
	Nomorid     int    `json:"nomorid" bson:"nomorid"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Price       int    `json:"price" bson:"price"`
	Stock       int    `json:"stock" bson:"stock"`
	Size        string `json:"size" bson:"size"`
	Image       string `json:"image" bson:"image"`
}

type Response struct {
	Status  bool        `json:"status" bson:"status"`
	Message string      `json:"message" bson:"message"`
	Data    interface{} `json:"data" bson:"data"`
}

type Coordinate struct {
	Type        string    `json:"type" bson:"type"`
	Name        string    `json:"name" bson:"name"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

type Token struct {
	Key    string
	Values string
}

type PostToken struct {
	Response string `json:"response"`
}

// --------------------------------------------------------------------- START GIS 9 ---------------------------------------------------------------------

type Lokasi struct { //lokasi yang bisa melakukan presensi
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Properties Name               `bson:"properties,omitempty"`
	Geometry   Geometry           `bson:"geometry,omitempty"`
	Kategori   string             `bson:"kategori,omitempty"`
}

type Name struct { //lokasi yang bisa melakukan presensi
	Name string `bson:"name,omitempty"`
}

type LongLat struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Pesan struct {
	Status  bool   `json:"status" bson:"status"`
	Message string `json:"message" bson:"message"`
}
