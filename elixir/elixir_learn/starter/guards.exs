f = fn
  x, y when x > 0 -> x + y
  x, y -> x * y
end

y = f.(1, 3)
IO.puts(y)

z = f.(-1, 3)
IO.puts(z)
