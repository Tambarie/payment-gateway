package domain

type User struct {
	Reference    string `json:"reference" bson:"reference"`
	FirstName    string `json:"first_name" bson:"first_name"`
	LastName     string `json:"last_name" bson:"last_name"`
	Email        string `json:"email" bson:"email"`
	Password     string `json:"-" bson:"password"`
	HashPassword []byte `json:"-" bson:"hash_password"`
	TimeCreated  string `json:"time-created" bson:"time_created"`
}
