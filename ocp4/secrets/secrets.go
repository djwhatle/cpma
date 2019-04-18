package secrets

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// TODO: Comment exported functions and structures.
// We may want to unexport some...

type FileSecret struct {
	HTPasswd string `yaml:"htpasswd"`
}

type LiteralSecret struct {
	ClientSecret string `yaml:"clientSecret"`
}

type Secret struct {
	APIVersion string      `yaml:"apiVersion"`
	Kind       string      `yaml:"kind"`
	Type       string      `yaml:"type"`
	MetaData   MetaData    `yaml:"metaData"`
	Data       interface{} `yaml:"data"`
}

type MetaData struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
}

var APIVersion = "v1"

func GenSecretFile(name string, encodedSecret string, namespace string) *Secret {
	var secret = Secret{
		APIVersion: APIVersion,
		Data:       FileSecret{HTPasswd: encodedSecret},
		Kind:       "Secret",
		Type:       "Opaque",
		MetaData: MetaData{
			Name:      name,
			Namespace: namespace,
		},
	}
	return &secret
}

func GenSecretLiteral(name string, clientSecret string, namespace string) *Secret {
	var secret = Secret{
		APIVersion: APIVersion,
		Data:       LiteralSecret{ClientSecret: clientSecret},
		Kind:       "Secret",
		Type:       "Opaque",
		MetaData: MetaData{
			Name:      name,
			Namespace: namespace,
		},
	}
	return &secret
}

// FIXME: The name of the function is misleading
func (secret *Secret) PrintCRD() string {
	yamlBytes, err := yaml.Marshal(&secret)
	if err != nil {
		logrus.WithError(err).Fatal("Cannot generate CRD")
		logrus.Debugf("%+v", secret)
	}
	return string(yamlBytes)
}
