package cex

type Api struct {
	Cex        Cex
	ApiKey     string `json:"apiKey,omitempty" bson:"apiKey"`
	SecretKey  string `json:"secretKey,omitempty" bson:"secretKey"`
	Passphrase string `json:"passphrase,omitempty" bson:"passphrase"`
}
