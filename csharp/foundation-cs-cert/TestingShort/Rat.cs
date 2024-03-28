
namespace Evil 
{
    class Rat
    {
        private string name;
        private int age;
        private string color;

        public Rat(string name, int age, string color)
        {
            this.name = name;
            this.age = age;
            this.color = color;
        }

        public string GetName()
        {
            return this.name;
        }

        public int GetAge()
        {
            return this.age;
        }

        public string GetColor()
        {
            return this.color;
        }
    }
}

