func longestPalindrome(s string) string {
    for windowLen := len(s); windowLen >= 0; windowLen-- {
        for startIdx := 0; startIdx <= len(s) - windowLen; startIdx++ {
            if isPalindrom(s[startIdx:startIdx + windowLen]) {
                return s[startIdx:startIdx + windowLen]
            }
        }
    }
    
    return string(s[0])
}

func isPalindrom(str string) bool {
    strLen := len(str)
    
    for i := 0; i < strLen; i++ {
        if str[i] != str[strLen - i - 1] {
            return false
        }
    }
    
    return true
}