file_path = File.expand_path('../input', __FILE__)
input = File.readlines(file_path).map(&:strip)

row_len = input.length - 1
col_len = input[0].length - 1

occ = 0

def check(str)
  str == 'MAS' || str == 'SAM'
end

(1..row_len - 1).each do |row|
  (1..col_len - 1).each do |col|
    letter = input[row][col]

    if letter == 'A'
      first = input[row-1][col-1] << input[row][col] << input[row+1][col+1]
      second = input[row-1][col+1] << input[row][col] << input[row+1][col-1]

      if check(first) && check(second)
        occ += 1
      end
    end
  end
end

puts occ
