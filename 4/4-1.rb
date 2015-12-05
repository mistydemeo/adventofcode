require "digest"

INPUT = "bgvyzdsv"
MD5 = Digest::MD5.new

i = 0
begin
  i += 1
  digest = MD5.hexdigest INPUT + i.to_s
end while digest[0..4] != "00000"

puts i
