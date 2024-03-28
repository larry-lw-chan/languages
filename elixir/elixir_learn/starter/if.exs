defmodule Learn do
  def wrong_math() do
    if (4 == 5) do
      "This is sooo wrong"
    else
      "This is sooo right"
    end
  end
end

IO.puts(Learn.wrong_math())


x = 5
IO.puts("Before change: #{x}")
x = if true do x + 10 end
IO.puts("After change: #{x}")
