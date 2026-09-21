func evalRPN(tokens []string) int {
    s := make([]int, 0)
    for _, t := range tokens{
        n := len(s)
        switch t {
            case "+":
                a := s[n-1]+s[n-2]
                s = s[:n-2]
                s = append(s, a)
            case "*":
                a := s[n-1]*s[n-2]
                s = s[:n-2]
                s = append(s, a)
            case "-":
                a := s[n-2]-s[n-1]
                s = s[:n-2]
                s = append(s, a)
            case "/":
                a := int(s[n-2]/s[n-1])
                s = s[:n-2]
                s = append(s, a)
            default:
                a, _ := strconv.Atoi(t)
                s = append(s,a)
        }
    }
    return s[0]
}
