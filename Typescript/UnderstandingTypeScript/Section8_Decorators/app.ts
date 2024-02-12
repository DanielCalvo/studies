function Logger(logString: string){
    return function(constructor: Function) {
        console.log(logString)
    }
}

@Logger('LOGGING PERSON')
class Person {
    name = 'Aaa'
    constructor() {
        console.log('Creating person object')
    }
}

const pers = new Person();
console.log(pers)