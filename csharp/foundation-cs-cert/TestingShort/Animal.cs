class Animal
{
    public string Name { get; set; }
    public int Age { get; set; }
    public string Sound { get; set; }

    public Animal(string name, int age, string sound)
    {
        Name = name;
        Age = age;
        Sound = sound;
    }
}