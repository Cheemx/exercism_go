package triangle

type Kind int

const (
    NaT = iota
    Equ  
    Iso 
    Sca 
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
    if !IsATriangle(a, b, c) {
        return NaT
    }
    if a == b && b == c {
        return Equ
    }
    if (a == b && a != c) || (b == c && b != a) || (a == c && a != b) {
        return Iso
    }
    if a != b && a != c && b != c {
        return Sca
    }
    var k Kind
	return k
}

func IsATriangle(a, b, c float64) bool {
    return a > 0 && b > 0 && c > 0 && a + b >= c && b + c >= a && a + c >= b
}