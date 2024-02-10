package cex

type Api struct {
	Cex        Cex    `json:"cex" bson:"cex" yaml:"cex"`
	ApiKey     string `json:"apiKey,omitempty" bson:"apiKey" yaml:"apiKey"`
	SecretKey  string `json:"secretKey,omitempty" bson:"secretKey" yaml:"secretKey"`
	Passphrase string `json:"passphrase,omitempty" bson:"passphrase" yaml:"passphrase"`
}
