package protein

import "errors"

const STOP = "STOP"
var(
    ErrStop = errors.New("Stop Iterating")
    ErrInvalidBase = errors.New("Invalid Base")
)
var mep = map[string]string{
    "AUG": "Methionine",
    "UUU": "Phenylalanine",
    "UUC": "Phenylalanine",
    "UUA": "Leucine",
    "UUG": "Leucine",
    "UCU": "Serine",
    "UCC": "Serine",
    "UCA": "Serine",
    "UCG": "Serine",
    "UAU": "Tyrosine",
    "UAC": "Tyrosine",
    "UGU": "Cysteine",
    "UGC": "Cysteine",
    "UGG": "Tryptophan",
    "UAA": STOP,
    "UAG": STOP,
    "UGA": STOP,
}

func FromRNA(rna string) ([]string, error) {
    if len(rna) % 3 != 0 {
        return []string{}, ErrInvalidBase
    }
	var ans []string
    for i := 0;i < len(rna) - 2; i += 3 {
        codon := rna[i:i+3]
        acid, err := FromCodon(codon)
        if err != nil && err == ErrStop {
            return ans, nil
        } 
        if err != nil && err == ErrInvalidBase {
            return []string{}, err
        }
        ans = append(ans, acid)
    }

    return ans, nil
}

func FromCodon(codon string) (string, error) {
	acid, ok := mep[codon]
    if !ok {
        return "", ErrInvalidBase
    }
    if acid == STOP {
        return acid, ErrStop
    }
    return acid, nil
}
