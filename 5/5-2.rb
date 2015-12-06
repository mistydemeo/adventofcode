REPEATED = /(..).*\1/
SEPARATED = /(.).\1/

nice = []
naughty = []

def nice? string
  !!(string =~ REPEATED && string =~ SEPARATED)
end

File.read("input.txt").chomp.split("\n").each do |line|
  nice?(line) ? nice << line : naughty << line
end

puts nice.length
