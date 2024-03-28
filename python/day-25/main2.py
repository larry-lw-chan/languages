import csv
import pandas

data = pandas.read_csv("./csv/weather_data.csv")

print(data)

# # Data frame basically a table with rows and columns
# print(type(data))
# # Series is a single column
# print(type(data['temp']))

# # Convert Dataframe to dictionary
# dict_data = data.to_dict()
# print(dict_data)

# # Convert Series to list
# series_list = data['temp'].to_list()
# print(series_list)

# # Calculate average of a series
# average = sum(series_list) / len(series_list)
# print(average)

# # Calculate average of a series using mean function
# print(data['temp'].mean())
# print(data["temp"].max())

# # Get data in columns
# print(data["condition"])
# print(data.condition)

# # Get data in Row
# row = data[data.day == "Monday"]

# Get Highest Temperature
# row = data[data.temp == data.temp.max()]
# print(row)

# Get Monday Temperature in Fahrenheit
# monday = data[data.day == "Monday"]
# monday_temp = monday.temp[0]
# monday_temp_f = monday_temp * (9 / 5) + 32
# print(monday_temp_f)

# Create a dataframe from scratch
data_dict = {
    "students": ["Amy", "James", "Angela"],
    "scores": [76, 56, 65]
}

data = pandas.DataFrame(data_dict)
print(data)
data.to_csv("new_data.csv", index=False)


