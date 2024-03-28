defmodule Sayings.Greetings do
  def basic(name), do: "Hi, #{name}"
end

defmodule Example do
  alias Sayings.Greetings
  def greeting(name), do: Greetings.basic(name)
end

# Without alias
defmodule Example do
  def greeting(name), do: Sayings.Greetings.basic(name)
end


# With the use macro we can enable another module
# to modify our current module’s definition
defmodule Hello do
  defmacro __using__(_opts) do
    greeting = Keyword.get(opts, :greeting, "Hi")
    quote do
      def hello(name), do: "#{greeting}, #{name}"
    end
  end
end

defmodule Example do
  use Hello, greeting: "Hola"
end
