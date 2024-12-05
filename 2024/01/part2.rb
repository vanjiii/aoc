#!/usr/bin/env ruby

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

hash = Hash.new(0)
right.each do |num|
  hash[num] += 1
end

score = 0

left.each do |num|
  score += num * hash[num]
end

puts score
