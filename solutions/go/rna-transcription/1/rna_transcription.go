package rnatranscription

func ToRNA(dna string) string {
	res := []rune{}

    for _, nucleotides := range dna {
		switch nucleotides {
            case 'G':
            res = append(res, 'C')
            case 'C':
            res = append(res, 'G')
            case 'T':
            res = append(res, 'A')
            case 'A':
            res = append(res, 'U')
        }
        
    }
    return string(res)
}
