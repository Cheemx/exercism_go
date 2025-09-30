package proverb

func Proverb(rhyme []string) []string {
    var proverbs []string
    if len(rhyme) == 0 {
        return proverbs
    }
    for i := 0; i < len(rhyme) - 1; i++ {
        want := rhyme[i]
        lost := rhyme[i + 1]
        proverbs = append(proverbs, "For want of a "+want+" the "+lost+" was lost.")
    }
    proverbs = append(proverbs, "And all for the want of a "+rhyme[0]+".")
    return proverbs
}
