using System.Numerics;

class Basic
{
    static T Add<T>(T left, T right) where T : INumber<T>
    {
        return left + right;
    }
}