using System;

string input;
List<string> todos = new();

void showTodos()
{
    for (int i = 1; i <= todos.Count; i++)
    {
        Console.WriteLine($"{i}. {todos[i - 1]}");
    }
    Console.WriteLine("");
}

void addTodo()
{
    bool success;
    do
    {
        success = true;
        Console.WriteLine("Enter the TODO description: ");
        string todo = Console.ReadLine() ?? "";
        foreach (string t in todos)
        {
            if (t == todo)
            {
                Console.WriteLine("The description must be unique.");
                Console.WriteLine("");
                success = false;
            }
        }
        if (success)
        {
            todos.Add(todo);
        }
    } while (!success);
}

void removeTodo()
{
    if (todos.Count == 0)
    {
        Console.WriteLine("There are no TODOs to remove.");
        return;
    }

    while (true)
    {
        Console.WriteLine("Select the index of the TODO you want to remove");
        showTodos();

        string input = Console.ReadLine() ?? "";
        bool success = int.TryParse(input, out int index);
        if (success && index > 0 && index <= todos.Count)
        {
            todos.RemoveAt(index - 1);
            return;
        }
        else
        {
            Console.WriteLine("Invalid input");
            Console.ReadKey();
            Console.Clear();
        }
    }
}

do
{
    Console.WriteLine("What do you want to do?");
    Console.WriteLine("[S]ee all TODOs");
    Console.WriteLine("[A]dd a TODO");
    Console.WriteLine("[R]emove a TODO");
    Console.WriteLine("[E]xit");

    input = Console.ReadLine() ?? "";
    input = input.ToUpper();

    switch (input)
    {
        case "S":
            showTodos();
            break;
        case "A":
            addTodo();
            break;
        case "R":
            removeTodo();
            break;
        case "E":
            Console.WriteLine("Exit");
            break;
        default:
            Console.WriteLine("Invalid input");
            Console.ReadKey();
            Console.Clear();
            break;
    }
} while (input != "E");
