class Cat 
{
    public string Name { get; set; }
    public string Sound { get; set; }
    public int Age { get; set; }

    private int _sickness;

    public Cat(string name, string sound, int age)
    {
        Name = name;
        Sound = sound;
        Age = age;
    }

    public void MakeSicker()
    {
        _sickness += 1;
    }

    public void ShowSickness()
    {
        Console.WriteLine(_sickness);
    }

    
}