package main

import "golang.org/x/crypto/bcrypt"

func HashAndSort(passWd string) (pwdHashed []byte, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(passWd), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	return hash, nil
}

func ComparePassword(inputPassWd string, hashedPasswd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPasswd), []byte(inputPassWd))
	if err != nil {
		return false
	} else {
		return true
	}
}
