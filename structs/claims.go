package structs

import (
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claim struct {
	ID						primitive.ObjectID			`bson:"_id" json:"_id,omitempty`
	jwt.StandardClaims
}