defmodule Greeter do
  def hello(name) do
    "Hello, " <> name
  end

  def shorter(name), do: "Hello, #{name}"
end

result = Greeter.hello("John")
IO.puts(result)

result = Greeter.shorter("Barry")
IO.puts(result)

defmodule Length do
  def of([]), do: 0
  def of([_ | tail]), do: 1 + of(tail)
end

result = Length.of([1,2,3,4,5])
IO.puts(result)

defmodule Greeter2 do
  def hello(), do: "Hello, anoymous person"
  def hello(name), do: "Hello, " <> name
  def hello(name1, name2), do: "Hello, #{name1} and #{name2}"
end

IO.puts(Greeter2.hello())
IO.puts(Greeter2.hello("John"))
IO.puts(Greeter2.hello("John", "Barry"))
