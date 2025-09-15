package strand

func ToRNA(dna string) string {
	rna := ""
    for _, c := range dna {
        switch c {
        case 'G':
        	rna += string('C')
        case 'C':
            rna += string('G')
        case 'T':
            rna += string('A')
        case 'A':
            rna += string('U')
        default :
            return ""
        }
    }
    return rna
}
