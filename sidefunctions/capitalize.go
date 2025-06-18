package sidefunctions

func Capitalize(s string) string {
    result := ""
    for i, c := range s {
        if i == 0 || s[i-1] == ' ' || s[i-1] == '+' {
            if c >= 'a' && c <= 'z' {
                result += string(c - 32)
            } else {
                result += string(c)
            }
        } else {
            if c >= 'A' && c <= 'Z' {
                result += string(c + 32)
            } else {
                result += string(c)
            }
        }
    }
    return result
}