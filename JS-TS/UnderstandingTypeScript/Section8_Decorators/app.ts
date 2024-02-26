// A decorator is just a function that you apply to something (like a class!)
//Lets create a decorator factory
//Returns a decorator function but allows you to configure it when you assign it as a decorator to something
function Logger(logString: string){
    return function(constructor: Function) {
        console.log(logString)
    }
}

function WithTemplate(template: string, hookId: string) {
    return function (constructor: Function) {
        const hookEl = document.getElementById(hookId);
        if ()
    }
}

// Decorators in general are all about classes
// @Logger points at a function (Logger) that is our decorator
// Decorators are executed when your class is defined, not instantiated!
// @Logger('LOGGING PERSON')
@WithTemplate('', 'app')
class Person {
    name = 'Aaa'
    constructor() {
        console.log('Creating person object')
    }
}

const pers = new Person();
console.log(pers)