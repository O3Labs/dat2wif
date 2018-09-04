package main

import (
	"flag"
	"log"
	"os"

	"github.com/o3labs/ontology-crypto/keypair"
	"github.com/ontio/ontology/account"
)

func main() {
	fileString := flag.String("file", "", "path to .dat file from owallet")
	passwordStr := flag.String("password", "", "A password to decrypt a file")
	flag.Parse()
	if *fileString == "" || *passwordStr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	path := *fileString
	w, err := account.NewClientImpl(path)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	passwd := *passwordStr
	a, err := w.GetDefaultAccount([]byte(passwd))
	if err != nil {
		log.Printf("Invalid password: %v", err)
		return
	}
	wif, err := keypair.Key2WIF(a.PrivateKey)
	if err != nil {
		return
	}
	log.Printf("Your WIF: %v", string(wif))
}
