func titleToNumber(columnTitle string) int {
    result := 0

    for _, char := range columnTitle {
        result = result * 26

        letterValue := int(char - 'A' + 1)

        result = result + letterValue
    }

    return result
}
