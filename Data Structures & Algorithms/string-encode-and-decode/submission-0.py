class Solution:

    def encode(self, strs: List[str]) -> str:
        ans = ""
        for s in strs:
            ans+=str(len(s))
            ans+="#"
            ans+=s
        return ans

    def decode(self, s: str) -> List[str]:
        ans =[]
        start=0
        i = 0
        while i < len(s):
            if s[i] == "#":
                l = int(s[start:i])
                ans.append(s[i+1: i+l+1])
                start = i+l+1
                i = i+l+1
            else:
                i+=1
        return ans
