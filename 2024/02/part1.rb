#!/usr/bin/env ruby

#
# comunity solutions:
#
# def safe?(row) = [row, row.reverse].any? { |row| row.each_cons(2).all? { |a, b| (a+1..a+3).include? b } }
# def safeish?(row) = row.each_index.any? { |i| safe? row.dup.tap { _1.delete_at(i) } }

# rows = File.read('input.txt').lines.map { _1.split.map(&:to_i) }

# puts rows.count { safe? _1 }
# puts rows.count { safeish? _1 }
# ----------------------------------------------------------


file_path = File.expand_path('../sample', __FILE__)
input     = File.readlines(file_path)

reports = input.map{ |line| line.split.map(&:to_i) }

safe_reports = 0

reports.each do |report|
  order = ''
  safe = true
  (0..report.length - 2).each do |i|
    currl = report[i]
    nextl = report [i + 1]

    if i.zero? && currl > nextl
      order = 'desc'
    elsif i.zero? && currl < nextl
      order = 'asc'
    end

    if order == 'desc'
      diff = currl - nextl
    elsif order == 'asc'
      diff = nextl - currl
    else
      diff = -1
    end

    if diff.negative? || diff > 3 || diff.zero?
      safe = false
    end
  end

  if safe
    safe_reports += 1
  end
end

puts safe_reports
