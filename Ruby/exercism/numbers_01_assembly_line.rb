class AssemblyLine
  CARS_PER_HOUR = 221.0
  def initialize(speed)
    @speed = speed
    @prod_rate_per_hour = CARS_PER_HOUR * @speed * 1.0 if @speed >= 1 && @speed <= 4
    @prod_rate_per_hour = CARS_PER_HOUR * @speed * 0.9 if @speed >= 5 && @speed <= 8
    @prod_rate_per_hour = CARS_PER_HOUR * @speed * 0.8 if @speed == 9
    @prod_rate_per_hour = CARS_PER_HOUR * @speed * 0.77 if @speed == 10
  end

  def production_rate_per_hour
    @prod_rate_per_hour
  end

  def working_items_per_minute
    (@prod_rate_per_hour / 60).floor #need to round down
  end
end

a = AssemblyLine.new(5)
puts a.production_rate_per_hour.class
puts a.working_items_per_minute
