// See https://aka.ms/new-console-template for more information
Console.WriteLine("Hello!");
Console.WriteLine("Input the first number: ");
var input1 = Console.ReadLine();

Console.WriteLine("Input the second number: ");
var input2 = Console.ReadLine();

Console.WriteLine("What do you want to do?");
Console.WriteLine("[A]dd numbers");
Console.WriteLine("[S]ubtract numbers");
Console.WriteLine("[M]ultiply numbers");
var choice = Console.ReadLine()!;

if (choice != "A" && choice != "S" && choice != "M")
{
    Console.WriteLine("Invalid Choice");
    Console.WriteLine("Press any key to close.");
}

if (String.Equals(choice.ToLower(), "a"))
{
    int result = Convert.ToInt32(input1) + Convert.ToInt32(input2);
    Console.WriteLine($"{input1} + {input2} = {result}");
}
else if (String.Equals(choice.ToLower(), "s"))
{
    int result = Convert.ToInt32(input1) - Convert.ToInt32(input2);
    Console.WriteLine($"{input1} - {input2} = {result}");
}
else if (String.Equals(choice.ToLower(), "m"))
{
    int result = Convert.ToInt32(input1) * Convert.ToInt32(input2);
    Console.WriteLine($"{input1} * {input2} = {result}");
    Console.WriteLine($"The product of {input1} and {input2} is {result}");
}