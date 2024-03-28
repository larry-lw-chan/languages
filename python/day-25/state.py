from turtle import Turtle

class State(Turtle):
    def __init__(self, state, x, y):
        super().__init__()
        self.state = state
        self.x = x
        self.y = y
        self.write_state()

    def write_state(self):
        self.hideturtle()
        self.penup()
        self.goto(self.x, self.y)
        self.write(self.state, align="center", font=("Arial", 8, "normal"))
