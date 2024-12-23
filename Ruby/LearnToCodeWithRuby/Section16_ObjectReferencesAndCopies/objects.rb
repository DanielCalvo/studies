a = [1,2,3]
b = a

#we are not creating a copy! we are simply giving a new name to the same array

p a.object_id
p b.object_id #ooo they're pointing to the exact same object!
p [1,2,3].object_id

a.push(4)
p a
p b #uh-oh

#a and b are not clones or duplicates

#dup and clone objects -- create an identical object, don't use the same

a = [1,2,3]
b = a.dup
c = b.clone

p a.object_id
p b.object_id
p c.object_id

puts 

d = "weee"
e = d #some say that here you are just creating another reference to that object!

d.upcase!
puts e #e is a reference to the same object

#the freeze method -- makes an object immutable!
name = "Joe".freeze
hobbies = ["coding", "learning"]
#name << "aa" #throws FrozenError
#name.upcase! #same as above

#dup when called on a frozen object, the given object will not be frozen
#clone will give you a frozen copy!
p name.frozen?
name_dup = name.dup
p name_dup.frozen?
name_dup << "McBobson"
puts name_dup

name_clone = name.clone
puts name_clone.frozen? #this is frozen!
#name << "Joeson" #FrozenError

#if you call dup and clone on a non-frozen object they function the same way
#if you call them on a frozen object, clone gives you a frozen copy, dup gives you a non-frozen one!

def append_5(elements)
    #do note that you can modify the existing array like this, and the object outside of the function will be changed:
    elements << 5
    #however if you do something like this:
    #elements = [] 
    #it will not overwrite the object outside of the function
end

#passing objects to mutating methods
values = [1,2,3,4]
append_5(values) #mutates the underlying value array -- be careful with this
puts values

def uppercase(text)
    text.upcase!
end

name = "Joe"
uppercase(name)
puts name