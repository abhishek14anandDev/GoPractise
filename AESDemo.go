package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func main() {

	// cipher key
	//key := "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="

	//var json string = ReadJsonFile()
	//fmt.Println(json)
	// plaintext
	byteKey := []byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8}
	pt := "tIaqmKs5mivkyhqeVeKRI0hljS1yDGbwJh4DPoYMTmGbp8/52kDwgM+DeFAre/p6QG2wRT1gVqrJyOhFsu6aRNhOmrCwHvy8XGBGW1GOirE1lzujvAHhmP4NZyz03+I/wY6F4/fOBbV4IH/olKNhCV7DhSsRl/iSLZZdCBMjg28zvSqKtFbkZIP6dQXDHfY1ZUIWamlaSchLZ2jmdjeMP4sh3gIbcc8mnQWrlTZoVdFFYjWScXvIDgobPSBS6x8b7cGCKChtqgT5R6BLtmuIPfftX/o8LrfTrsQujKVQEZY3n4vqe7oD0xt84gTj4SEAbLxvvuA6gHQApJeu2bgpeqcGutt2APPOfQuDbb2CSg8Ij6EZ/JojGmDJE5rMklUeaddWYjgBYJcGNxk43NXhEiXuqrKTHDkyRWiCDe/tU9adFF7nmrQvlQFgel/XMgrlbngi7Kd8njguU04fbH2npGnwwJLtzIM6TAaKIEmYBsA31AaKq7hmm0/zn5BFYvxzit2TCTVNlMKsn2NBZVugeBtsNktZubdTaNEd1oGwU8d3r0fAaSo9W7WzNONxyTts3Lz1O+Y3SnHmjR9t9qUre5g8lxEYwQHYof9f7df90RVZZpo2YuZcLXlOtgpSpD1Xz4LCwv8KQGfWgudS1vncOR17137KiKPNUN+1J039Vu6Pb55FwIAN+WfiL9tS76Ml4WgmgsfdcZJm04PMACdjLvWeraTv0yrOlVEj/FSgPLN7qyIX2Hl0iU4uV5KPpNigamhSfOhr234GAMW4NHrLvpXd3bvWxzb/ce4BeLCez8v9fPdIEmyIp3ni5wU6SNznCtTq6YHCRtQL8hjAtwuhqQoh34qrvVHB3x6Qsdfeqz21nL3MqiCDwCl/5xXawlE59wTN9CMxLme0esoR15n7IA9nT4n6rBMrxOgGuYhTRMhtITTKsb0isZ2IhnJ4Yr/02+Be0T/c+YMupMoKFV4bsqWeEHqFJ/hpkTpmgFJHCKyjAGnQ1S7ogeHnFTuN9UHAsFZtGlD7Uzt8JEAfRbw9SBM+apqVMwJvjPgU6pc9i9PzJa8f0K8BZY0Th6jhF6wiVW4sOUw9sqK+QScVJ7b1dEoBCqnbmuN2MYTXw/2QEhElU1FX2jNL+591VR8sj2qSiynWdMOnosx0AXzvnNIOQNnU1ei6gNzT1y9P2lpBByKezG3dbMxFMsZrRAy4enDInOdGA+aKbNfFk/Dq41zIjQ=="
	//Key [ ]byte  = []byte(pt)
	//c := EncryptAES([]byte(key), pt)

	// plaintext
	//fmt.Println(pt)

	// ciphertext
	//fmt.Println(c)

	// decrypt
	DecryptAES(byteKey, pt)
}

func ReadJsonFile() string {
	jsonFile, err := os.Open("input.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	fmt.Println("Successfully Opened input.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	return "sdvfsd"
}

func EncryptAES(key []byte, plaintext string) string {

	c, err := aes.NewCipher(key)
	CheckError(err)

	out := make([]byte, len(plaintext))

	c.Encrypt(out, []byte(plaintext))

	return hex.EncodeToString(out)
}

func DecryptAES(key []byte, ct string) {
	ciphertext, _ := hex.DecodeString(ct)

	c, err := aes.NewCipher(key)
	CheckError(err)

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)

	s := string(pt[:])
	fmt.Println("DECRYPTED:", s)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
