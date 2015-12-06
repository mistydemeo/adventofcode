FORBIDDEN = /(ab|cd|pq|xy)/
VOWELS = /[aeiou]/
REPEATED = /(.)\1/

nice = []
naughty = []

def nice? string
  !!((string =~ FORBIDDEN).nil? && string.scan(VOWELS).length >= 3 && string =~ REPEATED)
end

File.read("input.txt").chomp.split("\n").each do |line|
  nice?(line) ? nice << line : naughty << line
end

puts nice.length
