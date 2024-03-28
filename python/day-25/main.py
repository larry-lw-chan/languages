import pandas
import turtle
from state import State

# Initialize the screen
IMAGE = "blank_states_img.gif"
screen = turtle.Screen()
screen.setup(width=725, height=491)
screen.title("U.S. States Game")
screen.addshape(IMAGE)
turtle.shape(IMAGE)

# Get the data
data = pandas.read_csv("./csv/50_states.csv")
all_states = data.state.values
guessed_states = []

while len(guessed_states) < 50:
    answer_state = screen.textinput(title=f"{len(guessed_states)}/{len(all_states)} States Correct", prompt="What's another state's name?")

    if answer_state == "exit":
        learn_states = [state for state in all_states if state not in guessed_states]
        pandas.DataFrame({"states_to_learn": learn_states}).to_csv("states_to_learn.csv")
        break

    # Clean user input
    answer_state = answer_state.title()
    
    # Check if answer is in the data
    if answer_state in data.state.values:
        guessed_states.append(answer_state)
        state_data = data[data["state"] == answer_state]
        # Extract the state, x, and y values from the state_data
        state = state_data.state.values[0]
        x = state_data.x.values[0]
        y = state_data.y.values[0]
        # Create new state object
        State(state, x, y)

    # Check if user clicked cancel
    if answer_state is None:
        break

# screen.exitonclick()