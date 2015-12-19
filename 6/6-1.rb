class LightArray
  def initialize
    a = []
    (0..999).each do |i|
      a[i] = [false] * 1000
    end
    @ary = a
  end

  def [] *args
    @ary[*args]
  end

  def count_lights_on
    @ary.map do |ary|
      ary.select(&:itself).length
    end.inject(&:+)
  end

  def inspect; "#<LightArray>"; end
end

class Command
  COMMAND_REGEX = /(toggle|turn off|turn on) (\d+),(\d+) through (\d+),(\d+)/

  def initialize command, startx, starty, endx, endy
    @command = command
    @startx = startx
    @starty = starty
    @endx = endx
    @endy = endy
  end

  def self.parse string
    command, startx, starty, endx, endy = string.scan(COMMAND_REGEX)[0]
    Command.new command.to_sym, startx.to_i, starty.to_i, endx.to_i, endy.to_i
  end

  def execute ary
    (@startx..@endx).each do |x|
      (@starty..@endy).each do |y|
        execute_action @command, ary, x, y
      end
    end
  end

  def inspect
    "#<Command @command=#{@command} @startx=#{@startx} @endx=#{@endx} @starty=#{@starty} @endy=#{@endy}>"
  end

  private

  def execute_action action, ary, x, y
    case action
    when :toggle
      ary[x][y] = !ary[x][y]
    when :"turn on"
      ary[x][y] = true
    when :"turn off"
      ary[x][y] = false
    end
  end
end

def main
  lights = LightArray.new
  commands = File.read("input.txt").split("\n").map {|line| Command.parse(line) }
  commands.each do |command|
    command.execute lights
  end
  puts lights.count_lights_on

  return 0
end

if to_s == 'main'
  exit(main)
end
