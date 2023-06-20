def my_function
  # Add your function here
end

def ai_function
  # Add AI function here
end

num_runs = 3  # Number of times to run each function

my_times = []
ai_times = []

num_runs.times do
  my_start_time = Time.now
  my_function
  my_end_time = Time.now
  my_times << my_end_time - my_start_time

  ai_start_time = Time.now
  ai_function
  ai_end_time = Time.now
  ai_times << ai_end_time - ai_start_time
end

my_total_time = my_times.inject(:+)
ai_total_time = ai_times.inject(:+)

puts "My average time: #{'%.10f' % (my_total_time / num_runs)} seconds"
puts "AI average time: #{'%.10f' % (ai_total_time / num_runs)} seconds"
puts "\n"

puts "My execution times:"
my_times.each_with_index do |my_time, i|
  puts "Run #{i+1}: #{'%.10f' % my_time} seconds"
end

puts "AI execution times:"
ai_times.each_with_index do |ai_time, i|
  puts "Run #{i+1}: #{'%.10f' % ai_time} seconds"
end