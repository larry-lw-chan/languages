add = fn (a,b) -> a + b end

# Anonymous functions are invoked with a dot and parentheses
x = add.(1,2)
IO.puts(x)

double = fn a -> add.(a, a) end

x = double.(2)
IO.puts(x)
IO.puts(ceil(5.6))
