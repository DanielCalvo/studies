
// Strings! This is what I came here for

println 'He said, "That is Groovy"'

mystring = 'A string'
println(toString().getClass().name)

value = 25

println("Printing value: ${value}")
println('Printing value: ${value}') // Single quotes treat things as literals apparently

// Strings in groovy are immutable

mystring2 = "banana"

println(mystring2[2])

try {
    mystring2[2] = 'a'
} catch (Exception ex) {
    println(ex)
}

// All valid
println("We paid \$$value")
println("We paid \$${value}")
println(/We paid $$value/)

println()
what = new StringBuffer('fence')
text = "The cow jumped over the $what"
println(text)

what.replace(0,5,"moon")
println(text)

//Strings created using single quotes are Java strings.
//Strings created using double quotes are groovy strings (aka gstrings).

def printClassInfo(obj){
    println "The object passed was ${obj}"
    println "class: ${obj.getClass().name}"
    println "superclass: ${obj.getClass().superclass.name}"
    println()
}

val = 125

printClassInfo("The stock closet at ${val}")
printClassInfo("This is a sample string")
printClassInfo('Single quoted string')
printClassInfo(val)


price = 500
company = "google"
quote = "Today $company stock closed at $price"
println quote

stocks = [Apple : 100, Microsoft : 200]

stocks.each { key, value ->
    company = key
    price = value
    println quote
}

//The above doesn't work as GString is not reevaluating strings. We must ask it to do it for us:

companyClosure = { it.write(company) }
priceClosure = { it.write("$price")}
quote = "Today ${companyClosure} stock closed at ${priceClosure}"

stocks.each { key, value ->
    company = key
    price = value
    println quote
}

memo = '''Several of you raised concerns about long meetings.
To discuss this, we will be holding a 3 hour meeting starting
at 9AM tomorrow. All getting this memo are required to attend.
If you can't make it, please have a meeting with your manager to explain.'''

println(memo)

message = """We're very pleased to announce
that our stock price hit a high of \$${price} per share
on December 24th. Great news in time for..."""

println(message)

mystring3 = "Its a rainy day in seattle"
println(mystring3)

mystring3 -= "rainy "

println(mystring3)


obj = ~"hello"
println(obj.getClass().name) //java.util.regex.Pattern


pattern = ~"(G|g)roovy"
text = 'Groovy is Hip'
if (text =~ pattern)
    println "match"
else
    println "no match"

// =~ performs regex to a partial match
// ==~ performs regex to an exact match


