file_path = File.expand_path('../input', __FILE__)
input = File.readlines(file_path).map(&:strip)

# community solutions
#input = File.read('input.txt')
# grid = input.lines.flat_map.with_index { |line, y| line.chomp.chars.map.with_index { |char, x| [x + y.i, char] } }.to_h

# def text(grid, start, indices) = grid.values_at(*indices.map { start + _1 }).join
# def take(grid, start, dir) = text grid, start, 4.times.map { dir * _1 }

# p grid.keys.sum { |i| [-1, 0, 1].product([-1, 0, 1]).count { take(grid, i, _1 + _2.i) == 'XMAS' } }
# p grid.keys.map { text grid, _1, [0, -1-1i, -1+1i, 1-1i, 1+1i] }.count { _1[0] == 'A' && _1.count(?M) == 2 && _1.count(?S) == 2 && _1[1] != _1[4] }
#
# text взема offset-и и връща стринг .text(grid, 0+0i, [0, 1, 2, 3]) връща първата дума на първия ред. ако бяха [0, 1i, 2i, 3i] щеше да е първата колона
# take(grid, 0+0i, 1) взема четири символа от (0+0i), демек горе-ляво, и посока 1 (демек надясно). всичките посоки са 1 (надясно), 1i (надолу), -1 (наляво), -1-1i (диагонал наляво и нагоре) и т.н. [-1, 0, 1].product([-1, 0, 1]) ги генерира впрочем
# предпоследния ред взема всички точки, на всяка взема посоките около нея и вади по един стринг от четири символа, след това брои колко от тях са XMAS
# последния взема X-а около всеки символ, проверява че по средата имаме А, че имаме два пъти M, два пъти S, и по единия диагонал някъде пише MAS
# структурата е кеш, където ключа е координат (0+0i), а стойността е буквата

row_len = input.length - 1
col_len = input[0].length - 1

occ = 0

def check(str)
  str == 'XMAS'
end

(0..row_len).each do |row|
  (0..col_len).each do |col|
    letter = input[row][col]

    if letter == 'X'
      # rows
      # forward
      if col + 3 <= col_len && check(input[row][col..col + 3])
        occ += 1
      end
      # backward
      if col - 3 >= 0 && check(input[row][col - 3..col].reverse)
        occ += 1
      end

      # columns
      # forward
      if row + 3 <= row_len
        tmp_word = input[row][col] << input[row+1][col] << input[row+2][col] << input[row+3][col]
        if check(tmp_word)
          occ += 1
        end
      end
      # backwards
      if row - 3 >= 0
        tmp_word = input[row][col] << input[row - 1][col] << input[row - 2][col] << input[row - 3][col]
        if check(tmp_word)
          occ += 1
        end
      end

      # diagonal
      if row+3 <= row_len && col+3 <= col_len
        tmp_word = input[row][col] << input[row+1][col+1] << input[row+2][col+2] << input[row+3][col+3]
        if check(tmp_word)
          occ += 1
        end
      end

      if row-3 >= 0 && col+3 <= col_len
        tmp_word = input[row][col] << input[row-1][col+1] << input[row-2][col+2] << input[row-3][col+3]
        if check(tmp_word)
          occ += 1
        end
      end

      if row+3 <= row_len && col-3 >= 0
        tmp_word = input[row][col] << input[row+1][col-1] << input[row+2][col-2] << input[row+3][col-3]
        if check(tmp_word)
          occ += 1
        end
      end

      if row-3 >= 0 && col-3 >= 0
        tmp_word = input[row][col] << input[row-1][col-1] << input[row-2][col-2] << input[row-3][col-3]
        if check(tmp_word)
          occ += 1
        end
      end
    end
  end
end

puts occ
