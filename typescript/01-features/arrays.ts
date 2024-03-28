"use strict";

const carMakers = ["ford", "toyota", "chevy"];
const dates = [new Date(), new Date()];
const carsByMake: string[][] = [];

// Help with inference when extracting values
const car = carMakers[0];
const myCar = carMakers.pop();

// Prevent incompatible values
carMakers.push(100);

// Help with 'map'
carMakers.map((car: string): string => car.toUpperCase());

// Flexible types
const importantDates: (Date | string)[] = [];
importantDates.push(new Date());
importantDates.push("2030-10-10");

// Where to use typed arrays
// Any time we need to represent a collection of records
// with some arbitrary sort order
