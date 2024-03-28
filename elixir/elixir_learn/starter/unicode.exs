IO.puts(?a)
IO.puts(?Z)
IO.puts(?😃)

string = "héllo"
String.length(string)

cp = String.codepoints("👩‍🚒")
IO.puts(cp)
# ["👩", "‍", "🚒"]

gp = String.graphemes("👩‍🚒")
IO.puts(gp)
# ["👩‍🚒"]

IO.puts(String.length("👩‍🚒"))
