package sbt

import (
	"testing"

	"github.com/labbloom/go-merkletree"

	"github.com/willf/bitset"
)

func TestVerifyPresence(t *testing.T) {
	b := bitset.New(37)
	b.Set(0)
	b.Set(1)
	b.Set(8)
	b.Set(13)
	b.Set(18)
	b.Set(19)
	b.Set(20)
	b.Set(26)
	b.Set(28)
	b.Set(32)
	b.Set(33)
	b.Set(37)

	bT := NewBloomTree(b)
	proof1, data, err := bT.GenerateMultiProof([]int{8, 13, 33})
	if err != nil {
		t.Fatal(err)
	}

	proven, err := merkletree.VerifyMultiProof(data, false, proof1, bT.MT.Root())
	if err != nil {
		t.Fatal(err)
	}

	if !proven {
		t.Fatal("multiproof should be true")
	}

	proof2, data, err := bT.GenerateMultiProof([]int{0, 26, 28})
	if err != nil {
		t.Fatal(err)
	}

	proven, err = merkletree.VerifyMultiProof(data, false, proof2, bT.MT.Root())
	if err != nil {
		t.Fatal(err)
	}

	if !proven {
		t.Fatal("multiproof should be true")
	}

	_, _, err = bT.GenerateMultiProof([]int{0, 7, 28})
	if err == nil {
		t.Fatal("element is not present")
	}

	falseProof, data, err := bT.GenerateMultiProof([]int{0, 26, 28})
	if err != nil {
		t.Fatal(err)
	}

	falseProof.Indices[0] = 1

	proven, err = merkletree.VerifyMultiProof(data, false, falseProof, bT.MT.Root())
	if err != nil {
		t.Fatal(err)
	}

	if proven {
		t.Fatal("multiproof should not be true")
	}
}

func TestVerifyAbsence(t *testing.T) {
	b := bitset.New(37)
	b.Set(0)
	b.Set(1)
	b.Set(8)
	b.Set(13)
	b.Set(18)
	b.Set(19)
	b.Set(20)
	b.Set(26)
	b.Set(28)
	b.Set(32)
	b.Set(33)
	b.Set(37)

	bT := NewBloomTree(b)
	proof1, data, err := bT.GenerateAbsenceProof(16)
	if err != nil {
		t.Fatal(err)
	}

	proven, err := merkletree.VerifyMultiProof(data, false, proof1, bT.MT.Root())
	if err != nil {
		t.Fatal(err)
	}

	if !proven {
		t.Fatal("multiproof should be true")
	}

	proof2, data, err := bT.GenerateAbsenceProof(36)
	if err != nil {
		t.Fatal(err)
	}

	proven, err = merkletree.VerifyMultiProof(data, false, proof2, bT.MT.Root())
	if err != nil {
		t.Fatal(err)
	}

	if !proven {
		t.Fatal("multiproof should be true")
	}

	_, _, err = bT.GenerateAbsenceProof(1)
	if err == nil {
		t.Fatal("element is present")
	}

	falseProof, data, err := bT.GenerateAbsenceProof(22)
	if err != nil {
		t.Fatal(err)
	}

	falseProof.Indices[0] = 1

	proven, err = merkletree.VerifyMultiProof(data, false, falseProof, bT.MT.Root())
	if err != nil {
		t.Fatal(err)
	}

	if proven {
		t.Fatal("multiproof should not be true")
	}
}
