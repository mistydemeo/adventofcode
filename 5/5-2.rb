REPEATED = /(..).*\1/
SEPARATED = /(.).\1/

def nice? string
  !!(string =~ REPEATED && string =~ SEPARATED)
end

puts File.read("input.txt").chomp.split("\n").select(&method(:nice?)).length
