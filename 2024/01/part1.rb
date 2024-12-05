#!/usr/bin/env ruby

# comunity solutions
#
# a, b = File.readlines('input').map { |l|
#   l.split.map(&:to_i)
# }.transpose

# # a
# p a.sort.zip(b.sort).map { (_1 - _2).abs }.sum

# # b
# p a.map { b.count(_1) * _1 }.sum
# -----------------------------------------------
# p File.read('input.txt').lines.map { _1.split.map(&:to_i) }.transpose.map(&:sort).transpose.sum { |a, b| (a - b).abs }

file_path = File.expand_path('../input', __FILE__)
input     = File.read(file_path)

left = []
right = []

lines = input.split("\n")

lines.each do |line|
  row = line.split

  left << row.first.to_i
  right << row.last.to_i
end

left.sort!
right.sort!

distance = 0

(0..lines.length - 1).each do |i|
  distance += (right[i] - left[i]).abs
end

puts distance
