class BrowserHistory:

    def __init__(self, homepage: str):
        self.history = [homepage]
        self.top = 1
        self.pos = 0


    def visit(self, url: str) -> None:
        if self.top == len(self.history) and self.top - 1 == self.pos:
            self.history.append(url)
            self.top = len(self.history)
            self.pos = self.top - 1
        else:
            self.pos += 1
            self.history[self.pos] = url
            self.top = self.pos + 1
  
            
    def back(self, steps: int) -> str:
        if self.pos < steps:
            self.pos = steps
        self.pos -= steps
        return self.history[self.pos]

    def forward(self, steps: int) -> str:
        if self.pos + steps >= self.top:
            steps = self.top - self.pos - 1
        self.pos += steps
        return self.history[self.pos]


# Your BrowserHistory object will be instantiated and called as such:
# obj = BrowserHistory(homepage)
# obj.visit(url)
# param_2 = obj.back(steps)
# param_3 = obj.forward(steps)