/***
 * Excerpted from "Programming Groovy",
 * published by The Pragmatic Bookshelf.
 * Copyrights apply to this code. It may not be used to create training material,
 * courses, books, articles, and the like. Contact us if you are in doubt.
 * We make no guarantees that this code is fit for any purpose.
 * Visit http://www.pragmaticprogrammer.com/titles/vslg for more book information.
 ***/
//Java code
public class Car_java
{
    private int miles;
    private int year;

    public Car_java(int theYear) {
        year = theYear;
    }

    public int getMiles() {
        return miles;
    }

    public void setMiles(int theMiles) {
        miles = theMiles;
    }

    public int getYear() {
        return year;
    }

    public static void main(String[] args)
    {
        Car_java carJava = new Car_java(2008);

        System.out.println("Year: " + carJava.getYear());
        System.out.println("Miles: " + carJava.getMiles());
        System.out.println("Setting miles");
        carJava.setMiles(25);
        System.out.println("Miles: " + carJava.getMiles());
    }
}
