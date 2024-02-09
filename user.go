package cex

type Api struct {
	ApiKey     string `json:"apiKey,omitempty" bson:"apiKey"`
	SecretKey  string `json:"secretKey,omitempty" bson:"secretKey"`
	Passphrase string `json:"passphrase,omitempty" bson:"passphrase"`
}

type User interface {
	Cex() Cex
	Api() Api
}
