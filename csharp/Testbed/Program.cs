using System;

Console.WriteLine("Enter a number!");
string userInput = Console.ReadLine() ?? "";

bool isParsable = int.TryParse(userInput, out int number);

if (isParsable)
{
    Console.WriteLine($"You entered {number}");
}
else
{
    Console.WriteLine("You did not enter a number!");
}
Console.ReadKey();

