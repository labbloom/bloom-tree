package bloomtree

import (
	"testing"
)

func TestPresenceProofPresentElement(t *testing.T) {
	var tests = []struct {
		element  []byte
		elements [][]byte
	}{
		{
			element:  []byte{1},
			elements: [][]byte{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}},
		},
		{
			element: []byte{1},
			elements: [][]byte{{0}, {1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}, {10}, {11}, {12}, {13},
				{14}, {15}, {16}},
		},
		{
			element:  []byte{1},
			elements: [][]byte{{0}, {1}},
		},
	}

	for _, test := range tests {
		seed := "secret seed"
		dbf := generateDBF(seed, test.elements...)
		tree, err := NewBloomTree(dbf)
		if err != nil {
			t.Fatal(err)
		}

		multiproof, err := tree.GenerateCompactMultiProof(test.element)
		if err != nil {
			t.Fatal(err)
		}

		if CheckProofType(multiproof.proofType) != true {
			t.Fatal("proof type is not presence")
		}

		present, err := tree.VerifyCompactMultiProof(test.element, []byte(seed), multiproof, tree.Root())
		if err != nil {
			t.Fatal(err)
		} else if !present {
			t.Fatal("expected element to be present, but is absent")
		}
	}
}