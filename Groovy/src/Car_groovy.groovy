/***
 * Excerpted from "Programming Groovy",
 * published by The Pragmatic Bookshelf.
 * Copyrights apply to this code. It may not be used to create training material, 
 * courses, books, articles, and the like. Contact us if you are in doubt.
 * We make no guarantees that this code is fit for any purpose. 
 * Visit http://www.pragmaticprogrammer.com/titles/vslg for more book information.
 ***/
class Car
{
    def miles = 0
    final year // final means read only

    Car(theYear)
    {
        year = theYear
    }
}

Car car = new Car(2008)

println "Year: $car.year"
println "Miles: $car.miles"
println  'Setting miles'
//car.miles = 25 // can't do this
//println "Miles: $car.miles"


//class Car2
//{
//    final miles = 0
//    def getMiles()
//    {
//        println "getMiles called"
//        miles
//    }
//    def drive(dist) { if (dist > 0) miles += dist }
//}
//
//Car2 car2 = new Car2()