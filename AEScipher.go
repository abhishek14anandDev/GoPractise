package main

// Working example: https://goplay.space/#Sa7qCLm6w65
import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func main() {
	key := "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
	fmt.Printf("Decoded Key %v", []byte(key))
	iv := "AAAAAAAAAAAAAAAAAAAAAA=="
	fmt.Printf("Decoded Key %v", []byte(iv))
	plaintext := "tIaqmKs5mivkyhqeVeKRI0hljS1yDGbwJh4DPoYMTmGbp8/52kDwgM+DeFAre/p6QG2wRT1gVqrJyOhFsu6aRNhOmrCwHvy8XGBGW1GOirE1lzujvAHhmP4NZyz03+I/wY6F4/fOBbV4IH/olKNhCV7DhSsRl/iSLZZdCBMjg28zvSqKtFbkZIP6dQXDHfY1ZUIWamlaSchLZ2jmdjeMP4sh3gIbcc8mnQWrlTZoVdFFYjWScXvIDgobPSBS6x8b7cGCKChtqgT5R6BLtmuIPfftX/o8LrfTrsQujKVQEZY3n4vqe7oD0xt84gTj4SEAbLxvvuA6gHQApJeu2bgpeqcGutt2APPOfQuDbb2CSg8Ij6EZ/JojGmDJE5rMklUeaddWYjgBYJcGNxk43NXhEiXuqrKTHDkyRWiCDe/tU9adFF7nmrQvlQFgel/XMgrlbngi7Kd8njguU04fbH2npGnwwJLtzIM6TAaKIEmYBsA31AaKq7hmm0/zn5BFYvxzit2TCTVNlMKsn2NBZVugeBtsNktZubdTaNEd1oGwU8d3r0fAaSo9W7WzNONxyTts3Lz1O+Y3SnHmjR9t9qUre5g8lxEYwQHYof9f7df90RVZZpo2YuZcLXlOtgpSpD1Xz4LCwv8KQGfWgudS1vncOR17137KiKPNUN+1J039Vu6Pb55FwIAN+WfiL9tS76Ml4WgmgsfdcZJm04PMACdjLvWeraTv0yrOlVEj/FSgPLN7qyIX2Hl0iU4uV5KPpNigamhSfOhr234GAMW4NHrLvpXd3bvWxzb/ce4BeLCez8v9fPdIEmyIp3ni5wU6SNznCtTq6YHCRtQL8hjAtwuhqQoh34qrvVHB3x6Qsdfeqz21nL3MqiCDwCl/5xXawlE59wTN9CMxLme0esoR15n7IA9nT4n6rBMrxOgGuYhTRMhtITTKsb0isZ2IhnJ4Yr/02+Be0T/c+YMupMoKFV4bsqWeEHqFJ/hpkTpmgFJHCKyjAGnQ1S7ogeHnFTuN9UHAsFZtGlD7Uzt8JEAfRbw9SBM+apqVMwJvjPgU6pc9i9PzJa8f0K8BZY0Th6jhF6wiVW4sOUw9sqK+QScVJ7b1dEoBCqnbmuN2MYTXw/2QEhElU1FX2jNL+591VR8sj2qSiynWdMOnosx0AXzvnNIOQNnU1ei6gNzT1y9P2lpBByKezG3dbMxFMsZrRAy4enDInOdGA+aKbNfFk/Dq41zIjQ=="
	fmt.Println("Data to encode: ", plaintext)

	//cipherText := fmt.Sprintf("%v", Ase256Encode(plaintext, key, iv, aes.BlockSize))

	//fmt.Println("Encode Result:\t", cipherText)
	fmt.Println("Decode Result:\t", Ase256Decode(plaintext, key, iv))

}

func Ase256Encode(plaintext string, key string, iv string, blockSize int) string {
	bKey := []byte(key)
	bIV := []byte(iv)
	bPlaintext := PKCS5Padding([]byte(plaintext), blockSize, len(plaintext))
	block, err := aes.NewCipher(bKey)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, len(bPlaintext))
	fmt.Println(ciphertext)
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return hex.EncodeToString(ciphertext)
}

func Ase256Decode(cipherText string, encKey string, iv string) (decryptedString string) {
	bKey := []byte(encKey)
	bIV := []byte(iv)
	cipherTextDecoded, err := hex.DecodeString(cipherText)
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher(bKey)
	if err != nil {
		panic(err)
	}

	mode := cipher.NewCBCDecrypter(block, bIV)
	mode.CryptBlocks([]byte(cipherTextDecoded), []byte(cipherTextDecoded))
	return string(cipherTextDecoded)
}

func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
