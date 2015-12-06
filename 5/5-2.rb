REPEATED = /(..).*\1/
SEPARATED = /(.).\1/

nice = []

def nice? string
  !!(string =~ REPEATED && string =~ SEPARATED)
end

File.read("input.txt").chomp.split("\n").each do |line|
  nice << line if nice? line
end

puts nice.length
