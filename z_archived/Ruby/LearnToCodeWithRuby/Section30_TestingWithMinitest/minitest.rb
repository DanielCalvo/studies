#Unit tests with minitest! Miinitest is a unit test framework that comes with Ruby.

require "minitest/autorun"
class InvalidAttackError < StandardError
end
#Test class -- any tests need to inherit from this class!

def sum(a, b)
  a + b
end

class TestMathematics < Minitest::Test
    #method needs to start with test!
    def test_sum
        #asset equal is available because you inherited from Minitest::Test
        #first argument is the value you expect, second argument is whatever you're running that is going to produce that value
        assert_equal 8, sum(2, 6)
    end

    def test_sum_again #two dots means two tests passed, neat!
        assert_equal 4, sum(2, 2)
    end
end

#if the tests fail, you get an F instead of a dot!

#Testing a class! Author says: Typically you'll write your application code and your test code in separate files!

class Pokemon
    attr_reader :name, :type, :attacks
    def initialize(name, type)
        @name = name
        @type = type
        @attacks = []
    end
    def add_attack(attacks)
        raise InvalidAttackError unless attacks.is_a?(String)
        @attacks << attacks
    end
end

#set up and teardown methods
class TestPokemon < Minitest::Test
    def setup #Runs before each test! Allows you to set up common code that all of your tests can use. Useful for cuttind down on code duplication.
        #But by default pikachu is a local variable, so you need to make it an instance variable so that it can be accessed by all the tests.
        @pikachu = Pokemon.new("Pikachu", :electric)
    end

    def teardown #Runs after each test!
        #puts "Test is done, deleting pokemon from database (in theory)"
    end

    def test_name #You can optionally pass a third argument, a string describing the test in case the assertion is not met
        assert_equal "Pikachu", @pikachu.name, "The pokemon object did not assign its name correctly!"
    end
    #You want the tests to be isolated. Assume that your tests are running in random order and are independent of each other.
    def test_type
        assert_equal :electric, @pikachu.type
    end
    #Assert include! Checks that something is included in something else
    def test_adding_new_power
        @pikachu.add_attack("Thunderbolt")
        assert_includes @pikachu.attacks, "Thunderbolt", "The pokemon object did not add the attack correctly!"
    end
    def test_adding_fake_power
        #assert raises looks for this exception
        assert_raises InvalidAttackError do
            @pikachu.add_attack(123) #in the block you add the code that should trigger this error!
        end
    end
end