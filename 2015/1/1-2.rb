floor = 0

File.read("input.txt").each_char.with_index do |c, i|
  floor += c == "(" ? 1 : -1
  if floor == -1
    puts i + 1
    exit
  end
end
