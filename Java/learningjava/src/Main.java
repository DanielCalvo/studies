import java.awt.*;
import java.lang.reflect.Array;
import java.util.Arrays;

public class Main {
    public static void main(String[] args){
   /*
        // String and an introduction to classes:

        System.out.println("Hello world!");
        Car myCar = new Car(25, "abcd123", Color.BLUE, true);
        Car daniCar = new Car(29.4, "1111 BBB", Color.BLACK, false);

        System.out.println("myCar license plate: " + myCar.licensePlate);
        System.out.println("daniCar license plate: " + daniCar.licensePlate);

        myCar.changePaintColor(Color.RED);
        System.out.println(myCar.paintColor.toString());

        String userInput = "entertainment";
        String upperCased = userInput.toUpperCase();
        System.out.println(upperCased);
        System.out.println(userInput);

        char firstCharacter = userInput.charAt(0);
        System.out.println(firstCharacter);

        System.out.println("Contains: " + userInput.contains("Enter".toLowerCase()));

         */

        //Introduction to arrays:


/*        int[] numbers = new int[5];
        numbers[0] = 31;
        numbers[1] = 32;
        numbers[3] = 44;

        int[] numbers2 = {1,22,33,44,55};

        Arrays.sort(numbers2);
        System.out.println(Arrays.toString(numbers2));

        String[] candyBars = {"Twix", "Hersheys", "Crunch"};
        System.out.println("Index 1: " + candyBars[1]);
        System.out.println(candyBars.length);

        System.out.println("Index 2: " + Array.get(candyBars, 2));*/

        // Challenge to create a object. Done with a bicycle

/*        Car daniCar = new Car(29.4, "1111 BBB", Color.BLACK, false);

        double myCarSpeed = 50;
        myCarSpeed = daniCar.speedingUp(myCarSpeed);

        String s = "dog";
        String replacedF = s.replace('d', 'f');
        System.out.println(replacedF);

        Bicycle daniBicycle = new Bicycle();
        daniBicycle.setFrameMaterial("Carbon");
        daniBicycle.setFrameSize(57);
        daniBicycle.setWheelSize(29);

        System.out.println(daniBicycle.getFrameMaterial());*/


        // introduction to conditional statements
        int age = 4;

        if (age >= 0 && age <= 5) {
            System.out.println("Between 0 and 5");
        } else {
            System.out.println("Other value");
        }

        // You stopped on chapter 20 of Lynda's learning Java tutorial

    }
}
