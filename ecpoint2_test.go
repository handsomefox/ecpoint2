package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"
	"reflect"
	"testing"

	"ecpoint2/ecpoint"
)

func TestTrue(t *testing.T) {
	aliceBytes, _, _, _ := elliptic.GenerateKey(ecpoint.Curve, rand.Reader)
	bobBytes, _, _, _ := elliptic.GenerateKey(ecpoint.Curve, rand.Reader)
	// a = Alice's private key
	a := *big.NewInt(0).SetBytes(aliceBytes)
	// b = Bob's private key
	b := *big.NewInt(0).SetBytes(bobBytes)

	// Ha = a * G
	Ha := ecpoint.ECPointGen(ecpoint.Curve.ScalarBaseMult(aliceBytes))

	// Hb = b * G
	Hb := ecpoint.ECPointGen(ecpoint.Curve.ScalarBaseMult(bobBytes))

	// a*Hb
	aHb := ecpoint.ScalarMult(Hb, a)
	// b*Ha
	bHa := ecpoint.ScalarMult(Ha, b)
	fmt.Printf("a*Hb(%s) == b*Ha(%s)\n", aHb, bHa)

	// Must be true (a * Hb == b * Ha)

	if !reflect.DeepEqual(aHb, bHa) {
		t.Fatal("a * Hb is not equal to b * Ha")
	}
}
