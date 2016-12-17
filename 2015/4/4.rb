require "digest"

INPUT = "bgvyzdsv"
MD5 = Digest::MD5.new

begin
  prefix_size = ARGV[0].nil? ? 5 : Integer(ARGV[0], 10)
rescue ArgumentError
  $stderr.puts "Argument was not an integer: #{ARGV[0]}"
  exit 1
end

cmp_string = (["0"] * prefix_size).join

i = 0
begin
  i += 1
  digest = MD5.hexdigest INPUT + i.to_s
end while digest[0...prefix_size] != cmp_string

puts i
