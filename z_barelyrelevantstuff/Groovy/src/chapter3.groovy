
//Java code, it works!
//public class Greetings
//{
//    public static void main(String[] args)
//    {
//        for(int i = 0; i < 3; i++)
//        {
//            System.out.print("ho " );
//        }
//        System.out.println("Merry Groovy!" );
//    }
//}

//Also works if you remove the class stuff
for(int i = 0; i < 3; i++)
{
    System.out.print("ho " )
}
System.out.println("Merry Groovy!" )

//More groovy-like:
for(i in 0..2) { print 'ho ' }
println 'Merry Groovy!'

0.upto(1){
    println("hey")
}

3.times {
    println ("Look, printing 3 times! $it")
}

0.step(10, 2){
    print ("$it ")
}
println()
3.times {
    print 'ho '
}

robot = new Robot(type: 'arm' , width: 10, height: 40)
println "$robot.type, $robot.height, $robot.width"
robot.access(50, x: 30, y: 20, z: 10, true)
println 'Merry Groovy!'

//Very cool, much less verbose than what Java would look like to do it.
println "git help".execute().text
println "ls".execute().text

/*
def foo(str){
    str?.reverse()
}

foo.('hello')
foo.(null)

*/

// Exception handling:

try {
    myfile = new File("asd")
    println(myfile.text)
} catch (FileNotFoundException ex) { // Catches only file not found exception
    println("Looks like that file was not found")
}


try {
    myfile = new File("asd")
    println(myfile.text)
} catch (ex) { // catches any exception
    println("There was an exception!")
    println(ex)
}


robot = new Robot(type: 'arm' , width: 10, height: 40)
println "$robot.type, $robot.height, $robot.width"
robot.access(50, x: 30, y: 20, z: 10, true)

str = "hello"

if (str) {
    println("str is not empty")
}

// Operator overloading

for(i = 'a' ; i < 'd' ; i++)
{
    println i
}

for (i in 'a'..'e'){
    println(i)
}


lst = ['hello']
lst << 'there'
println lst //You can also print with no parenthesis!

String[] greetings = ["Hi", "Hello", "Howdy"]

for(String greet : greetings) {
    println(greet)
}
for (greet in greetings){
    println(greet)
}




