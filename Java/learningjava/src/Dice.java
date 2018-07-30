import java.util.Random;

public class Dice {

    private int rolledNum;
    int previousRoll = -1;

    public Dice(){
    }

    public int roll(){
        Random rand = new Random();
        rolledNum = rand.nextInt(6) + 1;
        this.previousRoll= rolledNum;
        return rolledNum;
    }
}
