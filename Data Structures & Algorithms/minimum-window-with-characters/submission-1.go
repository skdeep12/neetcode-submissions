func minWindow(s string, t string) string {
    if len(t) > len(s) {
        return ""
    }
    smap := make(map[byte]int)
    tmap := make(map[byte]int)
    for i:=0;i<len(t);i+=1{
        if _, ok := tmap[t[i]]; !ok{
            tmap[t[i]] = 0
        }
        tmap[t[i]]+=1
    }
    l := 0
    r := 0
    ans := ""
    // fmt.Println(tmap)
    for r < len(s) {
        if _, ok := smap[s[r]]; !ok {
            smap[s[r]]=0
        }
        smap[s[r]]+=1

        for isValid(smap, tmap) {
            // fmt.Println(smap, l, r)
            if r-l+1 < len(ans) || ans == "" {
                ans = s[l:r+1]
            }
            smap[s[l]]-=1
            l+=1
        }
        r+=1
    }
    return ans
}

func isValid(smap, tmap map[byte]int) bool {

    for k,v := range tmap {
        if sv, ok := smap[k]; !ok {
            return false
        } else if v>sv{
            return false
        }
    }
    return true
}
