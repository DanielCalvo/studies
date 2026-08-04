class Attendee
  def initialize(num)
    @height = num
    @pass_id = nil
  end

  def height
    @height
  end

  def pass_id
    @pass_id
  end

  def issue_pass!(num)
    @pass_id = num
  end

  def revoke_pass!
    @pass_id = nil
  end
end

att = Attendee.new(150)

puts att.height
puts att.pass_id
att.issue_pass(22)
puts att.pass_id
att.revoke_pass!
puts att.pass_id