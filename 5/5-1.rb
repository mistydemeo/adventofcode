FORBIDDEN = /(ab|cd|pq|xy)/
VOWELS = /[aeiou]/
REPEATED = /(.)\1/

nice = []

def nice? string
  !!((string =~ FORBIDDEN).nil? && string.scan(VOWELS).length >= 3 && string =~ REPEATED)
end

File.read("input.txt").chomp.split("\n").each do |line|
  nice << line if nice? line
end

puts nice.length
