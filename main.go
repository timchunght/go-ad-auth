package main

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/ldap.v2"
	"log"
	"os"
)

func main() {

	// The username and password we want to check
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	// bindusername := os.Getenv("READONLY_USERNAME")
	// bindpassword := os.Getenv("READONLY_PASSWORD")
	ldapUrl := "ldap://xnmh.nhs.uk" //
	// ldapPort := 80 // ldapUrl := fmt.Sprintf("%s:%s", ldapUrl, ldapPort)
	// baseDN := "dc=example,dc=com"
	objectClass := "organizationalUnit" // organizationalPerson
	baseDN := "dc=xnmh,dc=nhs,dc=uk"
	log.Println("username: ", username)
	log.Println("password: ", password)
	// log.Println("bindusername: ", bindusername)
	// log.Println("bindpassword: ", bindpassword)

	l, err := ldap.Dial("tcp", ldapUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	// Reconnect with TLS
	err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Fatal(err)
	}

	// // First bind with a read only user
	// err = l.Bind(bindusername, bindpassword)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Search for the given username
	searchRequest := ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=%s)&(uid=%s))", objectClass, username),
		[]string{"dn"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	if len(sr.Entries) != 1 {
		log.Fatal("User does not exist or too many entries returned")
	}

	userdn := sr.Entries[0].DN

	// Bind as the user to verify their password
	err = l.Bind(userdn, password)
	if err != nil {
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
