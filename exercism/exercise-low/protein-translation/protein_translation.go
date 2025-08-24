// Package protein has functionality to translate RNA sequences into proteins.
package protein

import (
	"errors"
)

var (
	// ErrStop indicates the translation was stopped.
	ErrStop = errors.New("stop")

	// ErrInvalidBase indicates the codon was invalid.
	ErrInvalidBase = errors.New("invalid codon")

	noProteins = []string{}
)

const codonLength = 3

// func FromRNA accept rna and return list of codon
func FromRNA(rna string) ([]string, error) {
	proteins := []string{}
	for i := 0; i < len(rna); i = i + codonLength {
		rna_string := rna[i : i+codonLength]
		codon, err := FromCodon(rna_string)
		switch err {
		case ErrStop:
			return proteins, nil
		case ErrInvalidBase:
			return noProteins, err
		default:
			proteins = append(proteins, codon)
		}
	}

	return proteins, nil
}

// FromCodon either translates a codon to a protein or returns an error for a stop codon or invalid codon.
func FromCodon(codon string) (protein string, err error) {
	switch codon {
	case "AUG":
		protein = "Methionine"
	case "UUU", "UUC":
		protein = "Phenylalanine"
	case "UUA", "UUG":
		protein = "Leucine"
	case "UCU", "UCC", "UCA", "UCG":
		protein = "Serine"
	case "UAU", "UAC":
		protein = "Tyrosine"
	case "UGU", "UGC":
		protein = "Cysteine"
	case "UGG":
		protein = "Tryptophan"
	case "UAA", "UAG", "UGA":
		err = ErrStop
	default:
		err = ErrInvalidBase
	}
	return
}
