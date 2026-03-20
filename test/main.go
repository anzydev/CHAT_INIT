package main

import (
	"crypto/rand"
	"encoding/binary"
	"time"
	"fmt"
)

func generateOTP() string {

	b := make([]byte, 2)

	rand.Read(b)

	number := binary.BigEndian.Uint16(b)

	return fmt.Sprintf("%04d", number%10000)

}

type record struct {
	Code       string
	Createdat time.Time
}

var otpstorage = make(map[string]record)



func main() {

	fmt.Printf("hllo \n")

	otp := generateOTP()

	newentry := record{
		Code: otp,
		Createdat: time.Now(),
	}

	otpstorage["mikey"] = newentry

	fmt.Printf(" %s \n", otp)

	fmt.Printf(" %s is the otp and its created on %v " , otpstorage["mikey"].Code , otpstorage["mikey"].Createdat)







}
