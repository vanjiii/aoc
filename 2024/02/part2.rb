#!/usr/bin/env ruby

file_path = File.expand_path('../input', __FILE__)
input     = File.readlines(file_path)

reports = input.map { |line| line.split.map(&:to_i) }

def calc(reports)
  safe_reports = 0

  reports.each do |report|
    safe = false

    safe = calc_single(report)
    if safe
      safe_reports += 1
      next
    end

    (0..report.length - 1).each do |i|
      safe = calc_single(delete_report(report, i))
      if safe
        safe_reports += 1
        next
      end
    end
  end

  safe_reports
end

def delete_report(report, idx)
  if idx + 2 > report.length
    return report[0..idx]
  end

  report[0..idx] + report[idx + 2..]
end

def calc_single(report)
  order = ''
  (0..report.length - 2).each do |i|
    currl = report[i]
    nextl = report [i + 1]

    if i.zero? && currl > nextl
      order = 'desc'
    elsif i.zero? && currl < nextl
      order = 'asc'
    end

    diff = -1
    if order == 'desc'
      diff = currl - nextl
    elsif order == 'asc'
      diff = nextl - currl
    end

    if diff.negative? || diff > 3 || diff.zero?
      return false
    end
  end

  true
end

puts calc(reports)
