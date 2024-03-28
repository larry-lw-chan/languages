class Dog
{
    public string Name { get; set; }
    public int Age { get; set; }
    public string Breed { get; set; }
    public string Color { get; set; }
    public string BarkSound { get; set; }

    public Dog(string name, int age, string breed, string color, string barkSound)
    {
        Name = name;
        Age = age;
        Breed = breed;
        Color = color;
        BarkSound = barkSound;
    }

    public void Bark()
    {
        Console.WriteLine(BarkSound);
    }

    public string Greetings()
    {
        Name = Name.ToUpper();
        return $"Hello, my name is {Name} and I am a {Breed}!";
    }
}