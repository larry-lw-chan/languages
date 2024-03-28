"use strict";

// Objects
const drink = {
  color: "brown",
  carbonated: true,
  sugar: 40,
};

// Type alias
type Drink = [string, boolean, number];
const pepsi: Drink = ["brown", true, 40]; // Tuple

// Tuple hard to represent meaningful data
const carSpecs: [number, number] = [400, 3354]; // Tuple

// Object more readable
const carStats = {
  horsepower: 400,
  weight: 3354,
};
