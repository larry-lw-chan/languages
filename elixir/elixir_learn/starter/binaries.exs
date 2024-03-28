test = is_bitstring(<<3::4>>)
IO.puts(test)

test = is_binary(<<3::4>>)
IO.puts(test)

test = is_bitstring(<<0, 255, 42>>)
IO.puts(test)

test = is_binary(<<0, 255, 42>>)
IO.puts(test)

test = is_binary(<<42::16>>)
IO.puts(test)
