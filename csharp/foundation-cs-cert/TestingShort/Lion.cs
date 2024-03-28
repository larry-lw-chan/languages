class Lion : Animal
{
    public int HumanEaten { get; set; }
    public Lion(string name, int age, string sound, int humanEaten=0) : base(name, age, sound)
    {
        HumanEaten = humanEaten;
    }
}