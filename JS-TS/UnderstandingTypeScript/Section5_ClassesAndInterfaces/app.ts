// An interface for functions, uh-oh. An alternative to function types
interface AddFn {
    (a: number, b: number): number;
}

let add: AddFn;

//You can also implement inheritance in interfaces
interface Named {
    readonly name: string;
    outputName?: string; //? tells typescript that this atribute is optional
}

interface Greetable extends Named {
    // you can also have read only here, but no public or private
    readonly name: string;
    greet(phrase: string): void;
}
// implements the greetable interface. You can also implement multiple interfaces
class Person implements Greetable, Named {
    name: string;

    constructor(n: string) {
        this.name = n
    }
    greet(phrase: string) {
        console.log("Hello! ", phrase)
    }
}

let user1: Greetable;

user1 = new Person('aaa');
user1.greet('Howdy');

// typecasting!
// The exclamation mark tells typescript that the expression in front of it will never yield null
// const userInputElement = <HTMLInputElement>document.getElementById('user-input')!;
// If you remove the exclamation mark you can do things a bit differently:
// const userInputElement = <HTMLInputElement>document.getElementById('user-input')! as HTMLInputElement;
const userInputElement = <HTMLInputElement>document.getElementById('user-input') as HTMLInputElement;
if (userInputElement) {
    (userInputElement as HTMLInputElement).value = 'Hi there!';
}

interface ErrorContainer { //Whatever object you're constructing based on this error container interface, must have properties which are strings
    id: string;
    [prop: string]: string; //index types!
}

const errorBag: ErrorContainer = {
    id: 'FFF',
    email: 'Not a valid email'
}