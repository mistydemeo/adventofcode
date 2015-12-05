puts File.read("input.txt").chomp.each_char.map {|c| c == "(" ? 1 : -1}.inject(&:+)
