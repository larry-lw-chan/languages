case {1, 2, 3} do
  {4, 5, 6} ->
    IO.puts("This clause won't match")
  {1, x, 3} ->
    IO.puts("This clause will match and bind #{x} to 2 in this clause")
  _ ->
    IO.puts("This clause would match any value")
end

x = 1
case 10 do
  ^x -> "Won't match"
  _ -> "Will match"
end

# Guards Activated
case {1, 2, 3} do
  {1, x, 3} when x > 0 ->
    IO.puts("Will match")
  _ ->
    IO.puts("Would match, if guard condition were not satisfied")
end
