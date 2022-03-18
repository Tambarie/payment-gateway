package domain

type User struct {
	Reference    string `json:"id,omitempty" bson:"id"`
	FirstName    string `json:"first_name,omitempty" bson:"first_name"`
	LastName     string `json:"last_name,omitempty" bson:"last_name"`
	Email        string `json:"email,omitempty" bson:"email"`
	Password     string `json:"-" bson:"password"`
	HashPassword []byte `json:"-" bson:"hash_password"`
	TimeCreated  string `json:"-" bson:"time_created"`
}
