defmodule Example.Greetings do
  def morning(name) do
    "Good morning, #{name}"
  end

  def evening(name) do
    "Good evening, #{name}"
  end
end

Example.Greetings.morning("Sean") |> IO.puts;
Example.Greetings.evening("Sarah") |> IO.puts;

defmodule Example do
  @greeting "Hello"

  def greet(name) do
    "#{@greeting}, #{name}"
  end
end

Example.greet("Sean") |> IO.puts;
