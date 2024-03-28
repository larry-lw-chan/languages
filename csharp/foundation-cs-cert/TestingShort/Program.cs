using System.Runtime.CompilerServices;
using static System.Console;
using Evil;

class Program
{    
    static void Main() 
    {
        // WriteLine("Hello");
        // Dog dog = new ("Fido", 3, "Golden Retriever", "Golden", "Woof");
        // dog.Bark();
        // Write(dog.Greetings());

        // Cat kitty = new ("Kitty", "Meow", 5);
        // WriteLine(kitty.Name);
        // WriteLine(kitty.Sound);
        // WriteLine(kitty.Age);

        // kitty.Name = "Bad Cat";
        // WriteLine(kitty.Name);
        
        // kitty.Age = 300;
        // kitty.MakeSicker();
        // kitty.MakeSicker();
        // kitty.MakeSicker();
        // kitty.ShowSickness();

        Rat rat = new ("Rat the brat", 1, "Brown");
        Console.WriteLine(rat.GetName());
        Console.WriteLine(rat.GetAge());
        Console.WriteLine(rat.GetColor());

        // Lion lion = new ("Lion", 5, "Roar");
        // WriteLine(lion.Name);
        // WriteLine(lion.Age);
        // WriteLine(lion.Sound);
        // WriteLine($"Mr Lion ate {lion.HumanEaten} so far.");
        // lion.HumanEaten += 5;
        // WriteLine($"Mr Lion ate {lion.HumanEaten} so far.");

    }
}
