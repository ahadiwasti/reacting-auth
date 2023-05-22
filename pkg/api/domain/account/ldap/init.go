package ldap

import (
	"./pkg/api/log"
	"github.com/spf13/viper"
)

var ldapConn LDAP_CONFIG

func Setup() {
	ldapConn = LDAP_CONFIG{
		Addr:       viper.GetString("ldap.addr"),
		BaseDn:     viper.GetString("ldap.baseDn"),
		UserDn:     viper.GetString("ldap.userDn"),
		BindDn:     viper.GetString("ldap.bindDn"),
		BindPass:   viper.GetString("ldap.bindPass"),
		AuthFilter: viper.GetString("ldap.authFilter"),
		Attributes: viper.GetStringSlice("ldap.attributes"),
		TLS:        viper.GetBool("ldap.tls"),
		StartTLS:   viper.GetBool("ldap.startTLS"),
	}
	log.Info("Successfully init ldap config")
}

func ConnectLdap() {
	e := ldapConn.Connect()
	if e != nil {
		log.Fatal("ldap connect fail!")
	}
}

func GetLdap() LDAP_CONFIG {
	ConnectLdap()
	return ldapConn
}
