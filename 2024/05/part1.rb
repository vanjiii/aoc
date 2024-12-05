file_path = File.expand_path('../input', __FILE__)
input = File.readlines(file_path).map(&:strip)

first_section = {}
second_section =[]
is_ordered = true

input.each do |line|
  if line.include?('|')
    smaller, bigger = line.split('|').map(&:to_i)
    if first_section.key?(smaller)
      first_section[smaller] << bigger
    else
      first_section[smaller] = [bigger]
    end
  else
    is_ordered = true
    nums = line.split(',').map(&:to_i)

    (0..nums.length-2).each do |i|
      num = nums[i]
      nextn = nums[i+1]

      unless first_section[num]&.include?(nextn)
        is_ordered = false
      end
    end

    second_section << nums if is_ordered
  end
end

sum = 0
second_section.reject!(&:empty?).each do |arr|
  sum += arr[(arr.length - 1) / 2]
end

puts sum
