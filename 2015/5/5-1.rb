FORBIDDEN = /(ab|cd|pq|xy)/
VOWELS = /[aeiou]/
REPEATED = /(.)\1/

def nice? string
  !!((string =~ FORBIDDEN).nil? && string.scan(VOWELS).length >= 3 && string =~ REPEATED)
end

puts File.read("input.txt").chomp.split("\n").select(&method(:nice?)).length
