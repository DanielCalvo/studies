class LogLineParser
  def initialize(line)
    @line = line
  end

  def message
    line_array = @line.split(' ')
    line_array = line_array[1..-1] #Remove the first element, aka the log_level. Also trims whitespace
    return line_array.join(' ')
  end

  def log_level
    line_array = @line.split(' ')
    return line_array[0].delete(':').delete('[').delete(']').downcase #incredibly ugly
  end

  def reformat #This all turned out ugly -- there oughta be a prettier way of doing this!
    line_array = @line.split(' ')
    log_level = line_array[0]
    line_array = line_array[1..] #Removes the log level from the array
    log_level = log_level.delete(':').gsub('[', '(').gsub(']', ')').downcase
    line_array.append(log_level)
    return line_array.join(" ")
  end
end

puts LogLineParser.new('[ERROR]: Stack overflow').message
