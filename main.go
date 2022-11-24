package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"
	"reflect"

	"ecpoint2/ecpoint"
)

func main() {
	aliceBytes, _, _, _ := elliptic.GenerateKey(ecpoint.Curve, rand.Reader)
	bobBytes, _, _, _ := elliptic.GenerateKey(ecpoint.Curve, rand.Reader)
	// a = Alice's private key
	a := *big.NewInt(0).SetBytes(aliceBytes)
	fmt.Printf("Alice's private = %x\n", a.Bytes())
	// b = Bob's private key
	b := *big.NewInt(0).SetBytes(bobBytes)
	fmt.Printf("Bob's private = %x\n", b.Bytes())

	// Ha = a * G
	Ha := ecpoint.ECPointGen(ecpoint.Curve.ScalarBaseMult(aliceBytes))
	fmt.Printf("Ha = %s\n", ecpoint.ECPointToString(Ha))

	// Hb = b * G
	Hb := ecpoint.ECPointGen(ecpoint.Curve.ScalarBaseMult(bobBytes))
	fmt.Printf("Hb = %s\n", ecpoint.ECPointToString(Hb))

	// a*Hb
	aHb := ecpoint.ScalarMult(Hb, a)
	fmt.Printf("aHb = %s\n", ecpoint.ECPointToString(aHb))
	// b*Ha
	bHa := ecpoint.ScalarMult(Ha, b)
	fmt.Printf("bHa = %s\n", ecpoint.ECPointToString(bHa))

	fmt.Printf("a*Hb(%s) == b*Ha(%s)\n", aHb, bHa)

	// Must be true (a * Hb == b * Ha)
	fmt.Println(reflect.DeepEqual(aHb, bHa))
}
