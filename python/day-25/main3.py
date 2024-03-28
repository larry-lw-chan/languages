import pandas

# X,Y,Unique Squirrel ID,Hectare,Shift,Date,Hectare Squirrel Number,
# Age,Primary Fur Color,Highlight Fur Color,Combination of Primary and Highlight Color,
# Color notes,Location,Above Ground Sighter Measurement,Specific Location,Running,Chasing,
# Climbing,Eating,Foraging,Other Activities,Kuks,Quaas,Moans,Tail flags,Tail twitches,
# Approaches,Indifferent,Runs from,Other Interactions,Lat/Long
data = pandas.read_csv("./csv/2018_Central_Park_Squirrel_Census_-_Squirrel_Data.csv")


fur_color = data["Primary Fur Color"]

grey_count = fur_color[fur_color == "Gray"].count()
black_count = fur_color[fur_color == "Black"].count()
cinnamon_count = fur_color[fur_color == "Cinnamon"].count()

data_dict = {
    "Fur Color": ["grey", "black", "cinnamon"],
    "Count": [grey_count, black_count, cinnamon_count]
}

df = pandas.DataFrame(data_dict)
df.to_csv("./csv/squirrel_count.csv")