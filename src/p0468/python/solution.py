class Solution:
    def validIPAddress(self, queryIP: str) -> str:
        if self.isIPv4(queryIP):
            return "IPv4"
        elif self.isIPv6(queryIP):
            return "IPv6"
        
        return "Neither"
    
    def isIPv4(self, queryIP: str) -> bool:
        items = queryIP.split('.')
        if len(items) != 4:
            return False
        for item in items:
            if len(item) == 0 or len(item) > 3:
                return False
            temp = [c for c in item if c.isdigit()]
            if len(temp) != len(item):
                return False
            if len(temp) > 1 and temp[0] == '0':
                return False
            if int(item) > 255:
                return False 
        
        return True

    def isIPv6(self, queryIP: str) -> bool:
        items = queryIP.split(':')
        if len(items) != 8:
            return False
        for item in items:
            if len(item) == 0 or len(item) > 4:
                return False
            for c in item:
                if not c.isalnum():
                    return False
                if c.isalpha() and c.lower() not in 'abcdef':
                    return False
                    
        return True                        