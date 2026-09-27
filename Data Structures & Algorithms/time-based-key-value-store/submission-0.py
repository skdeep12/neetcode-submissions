class TimeMap:

    def __init__(self):
        self.map = dict()

    def set(self, key: str, value: str, timestamp: int) -> None:
        if key not in self.map.keys():
            self.map[key]= []
        self.map[key].append((timestamp, value))


    def get(self, key: str, timestamp: int) -> str:
        values = self.map.get(key, None)
        if values is None or values[0][0] > timestamp:
            return ""
        if len(values) == 1:
            return values[0][1]
        return self.binary_search(values, timestamp)
    
    def binary_search(self, values, timestamp) -> str:
        l = 0
        r = len(values)-1
        while l < r:
            m = int((l+r)/2)
            if values[m][0] > timestamp:
                r = m-1
            else:
                l = m+1
        if values[l][0] > timestamp:
            return values[l-1][1]
        else:
            return values[l][1]
         


