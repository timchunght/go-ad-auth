package main

import (
	// "crypto/tls"
	// "fmt"
	"gopkg.in/ldap.v2"
	"log"
	"os"
)

func main() {

	// The username and password we want to check
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	userDN := os.Getenv("USERDN")

	// bindusername := os.Getenv("READONLY_USERNAME")
	// bindpassword := os.Getenv("READONLY_PASSWORD")
	ldapUrl := os.Getenv("LDAP_URL") // "ldap://xnmh.nhs.uk:80" //
	// ldapPort := 80 // ldapUrl := fmt.Sprintf("%s:%s", ldapUrl, ldapPort)
	// baseDN := "dc=example,dc=com"
	// objectClass := "organizationalUnit" // organizationalPerson
	// baseDN := "dc=xnmh,dc=nhs,dc=uk"
	// userDN := fmt.Sprintf("%s@xnmh.nhs.uk", username)
	log.Println("username: ", username)
	log.Println("password: ", password)
	log.Println("userdn: ", userDN)
	log.Println("ldapUrl: ", ldapUrl)
	// log.Println("bindusername: ", bindusername)
	// log.Println("bindpassword: ", bindpassword)

	l, err := ldap.Dial("tcp", ldapUrl)
	if err != nil {
		log.Println("Connection Error")
		log.Fatal(err)
	}
	defer l.Close()

	// // Reconnect with TLS
	// err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Bind as the user to verify their password
	err = l.Bind(userDN, password)
	if err != nil {
		log.Println("Authentication Error")
		log.Fatal(err)
	} else {
		log.Println("Successfully authenticated")
	}

	// // Rebind as the read only user for any futher queries
	// err = l.Bind(bindusername, bindpassword)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
