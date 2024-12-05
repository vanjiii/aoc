file_path = File.expand_path('../input', __FILE__)
input = File.read(file_path)

#
# Community solutions
#
# sum, enabled = 0, 1

# File.read('input.txt').scan(/do\(\)|don't\(\)|mul\((\d+),(\d+)\)/) do |a, b|
#   case $&
#   in "do()" then enabled = 1
#   in "don't()" then enabled = 0
#   else sum += a.to_i * b.to_i * enabled
#   end
# end

# puts sum

instructions = input.gsub(/(do\(\)|don't\(\))|(mul\(\d+,\d+\))/)

result = 0
is_disabled = false

instructions.each do |instruction|
  if instruction == 'do()'
    is_disabled = false
    next
  elsif instruction == "don't()"
    is_disabled = true
    next
  else

  one, two = instruction.scan(/\d+/).map(&:to_i)
  result += (one * two) unless is_disabled
  end
end

p result
