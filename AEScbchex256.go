package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func main() {
	datarr := "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
	key, _ := base64.StdEncoding.DecodeString(datarr)
	//key := []byte("AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg=")
	//key := make([]byte, 32)

	//str := "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="

	//str_copy := copy(key, str)
	//fmt.Println(str_copy)
	//key := copy(byte_slice, str)
	fmt.Println(key)
	//cipherTextDecoded := make([]byte, 2048)
	cipherTextDecoded := []byte("tIaqmKs5mivkyhqeVeKRI0hljS1yDGbwJh4DPoYMTmGbp8/52kDwgM+DeFAre/p6QG2wRT1gVqrJyOhFsu6aRNhOmrCwHvy8XGBGW1GOirE1lzujvAHhmP4NZyz03+I/wY6F4/fOBbV4IH/olKNhCV7DhSsRl/iSLZZdCBMjg28zvSqKtFbkZIP6dQXDHfY1ZUIWamlaSchLZ2jmdjeMP4sh3gIbcc8mnQWrlTZoVdFFYjWScXvIDgobPSBS6x8b7cGCKChtqgT5R6BLtmuIPfftX/o8LrfTrsQujKVQEZY3n4vqe7oD0xt84gTj4SEAbLxvvuA6gHQApJeu2bgpeqcGutt2APPOfQuDbb2CSg8Ij6EZ/JojGmDJE5rMklUeaddWYjgBYJcGNxk43NXhEiXuqrKTHDkyRWiCDe/tU9adFF7nmrQvlQFgel/XMgrlbngi7Kd8njguU04fbH2npGnwwJLtzIM6TAaKIEmYBsA31AaKq7hmm0/zn5BFYvxzit2TCTVNlMKsn2NBZVugeBtsNktZubdTaNEd1oGwU8d3r0fAaSo9W7WzNONxyTts3Lz1O+Y3SnHmjR9t9qUre5g8lxEYwQHYof9f7df90RVZZpo2YuZcLXlOtgpSpD1Xz4LCwv8KQGfWgudS1vncOR17137KiKPNUN+1J039Vu6Pb55FwIAN+WfiL9tS76Ml4WgmgsfdcZJm04PMACdjLvWeraTv0yrOlVEj/FSgPLN7qyIX2Hl0iU4uV5KPpNigamhSfOhr234GAMW4NHrLvpXd3bvWxzb/ce4BeLCez8v9fPdIEmyIp3ni5wU6SNznCtTq6YHCRtQL8hjAtwuhqQoh34qrvVHB3x6Qsdfeqz21nL3MqiCDwCl/5xXawlE59wTN9CMxLme0esoR15n7IA9nT4n6rBMrxOgGuYhTRMhtITTKsb0isZ2IhnJ4Yr/02+Be0T/c+YMupMoKFV4bsqWeEHqFJ/hpkTpmgFJHCKyjAGnQ1S7ogeHnFTuN9UHAsFZtGlD7Uzt8JEAfRbw9SBM+apqVMwJvjPgU6pc9i9PzJa8f0K8BZY0Th6jhF6wiVW4sOUw9sqK+QScVJ7b1dEoBCqnbmuN2MYTXw/2QEhElU1FX2jNL+591VR8sj2qSiynWdMOnosx0AXzvnNIOQNnU1ei6gNzT1y9P2lpBByKezG3dbMxFMsZrRAy4enDInOdGA+aKbNfFk/Dq41zIjQ==")
	//copy(cipherTextDecoded, strText) // 256-bit key in hex format
	//ciphertext := []byte("tIaqmKs5mivkyhqeVeKRI0hljS1yDGbwJh4DPoYMTmGbp8/52kDwgM+DeFAre/p6QG2wRT1gVqrJyOhFsu6aRNhOmrCwHvy8XGBGW1GOirE1lzujvAHhmP4NZyz03+I/wY6F4/fOBbV4IH/olKNhCV7DhSsRl/iSLZZdCBMjg28zvSqKtFbkZIP6dQXDHfY1ZUIWamlaSchLZ2jmdjeMP4sh3gIbcc8mnQWrlTZoVdFFYjWScXvIDgobPSBS6x8b7cGCKChtqgT5R6BLtmuIPfftX/o8LrfTrsQujKVQEZY3n4vqe7oD0xt84gTj4SEAbLxvvuA6gHQApJeu2bgpeqcGutt2APPOfQuDbb2CSg8Ij6EZ/JojGmDJE5rMklUeaddWYjgBYJcGNxk43NXhEiXuqrKTHDkyRWiCDe/tU9adFF7nmrQvlQFgel/XMgrlbngi7Kd8njguU04fbH2npGnwwJLtzIM6TAaKIEmYBsA31AaKq7hmm0/zn5BFYvxzit2TCTVNlMKsn2NBZVugeBtsNktZubdTaNEd1oGwU8d3r0fAaSo9W7WzNONxyTts3Lz1O+Y3SnHmjR9t9qUre5g8lxEYwQHYof9f7df90RVZZpo2YuZcLXlOtgpSpD1Xz4LCwv8KQGfWgudS1vncOR17137KiKPNUN+1J039Vu6Pb55FwIAN+WfiL9tS76Ml4WgmgsfdcZJm04PMACdjLvWeraTv0yrOlVEj/FSgPLN7qyIX2Hl0iU4uV5KPpNigamhSfOhr234GAMW4NHrLvpXd3bvWxzb/ce4BeLCez8v9fPdIEmyIp3ni5wU6SNznCtTq6YHCRtQL8hjAtwuhqQoh34qrvVHB3x6Qsdfeqz21nL3MqiCDwCl/5xXawlE59wTN9CMxLme0esoR15n7IA9nT4n6rBMrxOgGuYhTRMhtITTKsb0isZ2IhnJ4Yr/02+Be0T/c+YMupMoKFV4bsqWeEHqFJ/hpkTpmgFJHCKyjAGnQ1S7ogeHnFTuN9UHAsFZtGlD7Uzt8JEAfRbw9SBM+apqVMwJvjPgU6pc9i9PzJa8f0K8BZY0Th6jhF6wiVW4sOUw9sqK+QScVJ7b1dEoBCqnbmuN2MYTXw/2QEhElU1FX2jNL+591VR8sj2qSiynWdMOnosx0AXzvnNIOQNnU1ei6gNzT1y9P2lpBByKezG3dbMxFMsZrRAy4enDInOdGA+aKbNfFk/Dq41zIjQ==") // Encrypted message in hex format
	//iv, _ := hex.DecodeString("AAAAAAAAAAAAAAAAAAAAAA==")
	//iv := make([]byte, 16) // 128-bit initialization vector in hex format
	//striv := "AAAAAAAAAAAAAAAAAAAAAA=="
	//str_copyiv := copy(iv, striv)
	//fmt.Println(str_copyiv)

	datarriv := "AAAAAAAAAAAAAAAAAAAAAA=="
	iv, _ := base64.StdEncoding.DecodeString(datarriv)
	fmt.Println(iv)
	plaintext, err := decrypt1(cipherTextDecoded, key, iv)
	if err != nil {
		panic(err)
	}
	//str := base64.StdEncoding.EncodeToString(plaintext)
	///fmt.Println(str)
	//myString := hex.EncodeToString(plaintext)
	fmt.Printf("%s", string(plaintext))
	//fmt.Printf("Plaintext: %s", myString)
}

// The decrypt function takes a ciphertext, a key, and an initialization vector (IV) as inputs, and returns the plaintext
func decrypt1(ciphertext []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	// Remove padding
	padding := plaintext[len(plaintext)-1]
	plaintext = plaintext[:len(plaintext)-int(padding)]

	return plaintext, nil
}
