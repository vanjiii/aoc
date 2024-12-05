file_path = File.expand_path('../input', __FILE__)
input = File.read(file_path)

pairs = input.scan(/mul\(\d+,\d+\)/)

result = 0
pairs.each do |pair|
  one, two = pair.scan(/\d+/).map(&:to_i)
  result += (one * two)
end

p result
