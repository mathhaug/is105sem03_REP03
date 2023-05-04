package mycrypt

import (
	"unicode"
)

var ALF_SEM03 = []rune("abcdefghijklmnopqrstuvwxyzæøåABCDEFGHIJKLMNOPQRSTUVWXYZÆØÅ0123456789.,:; ")

func Krypter(melding []rune, alphabet []rune, chiffer int) []rune {
	kryptertMelding := make([]rune, len(melding))
	for i := 0; i < len(melding); i++ {
		indeks := sokIAlfabetet(melding[i], alphabet)
		isUppercase := unicode.IsUpper(melding[i])
		if isUppercase {
			indeks += 44 // convert uppercase to lowercase index
		}
		indeks = (indeks + chiffer) % len(alphabet)
		if indeks < 0 {
			indeks += len(alphabet)
		}
		if isUppercase {
			kryptertMelding[i] = unicode.ToUpper(alphabet[indeks])
		} else {
			kryptertMelding[i] = alphabet[indeks]
		}
	}
	return kryptertMelding
}

func sokIAlfabetet(symbol rune, alphabet []rune) int {
	for i := 0; i < len(alphabet); i++ {
		if symbol == alphabet[i] {
			return i
		}
	}
	return -1
}
