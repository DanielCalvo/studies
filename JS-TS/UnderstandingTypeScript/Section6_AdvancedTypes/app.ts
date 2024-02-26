// Defining a type
type Admin = {
    name: string;
    privileges: string[];
};

type Employee = {
    name: string;
    startDate: Date;
}

// Oooo an intersection type -- must have... both attributes?
type ElevatedEmployee = Admin & Employee;

const e1: ElevatedEmployee = {
    name: 'aaa',
    privileges: ['free coffees'],
    startDate: new Date()
}

type Combinable = string | number;
type Numeric = number | boolean;

// Of type number as that is the only intersection between Combinable and Numeric, uh-oh
type Universal = Combinable & Numeric;

// typeguard!
// uh-oh, function overloads
function add(a: number, b: number): number;
function add(a: string, b: string): string;
function add(a: Combinable, b: Combinable) {
    if (typeof a === 'string' || typeof b === 'string'){
        return a.toString() + b.toString();
    }
    return a+b;
}

const result = add('aaaa','bbb');
result.split(' '); //With funtion overloading above now typescript knows my result is of type string!

// optional chaining
const fetchedUserData = {
    id: 'DDD',
    name: 'Userino',
    job: {title: 'Bossman', description: 'My own company'}
};
// we try fetching job and if that works we try fetching title
console.log(fetchedUserData.job && fetchedUserData.job.title);
console.log(fetchedUserData?.job.title); //this tells TS: Does this exist? If so, access job. This is called optional chaining

// nullish coalescing
const userInput = null;
const storedData = userInput ?? 'DEFAULT'; //nullish coalescing operator: If this is null or undefined (but not empty string, for example)

type UnknownEmployee = Employee | Admin;

function printEmployeeInformation(emp: UnknownEmployee){
    console.log('You called printEmployeeInformation')
    console.log('Name: ', emp.name) //Will work without issue as both types have a name property
    if ('privileges' in emp) {
        console.log('Privileges: '+ emp.privileges)
    }
    if ('startDate' in emp) {
        console.log('Privileges: '+ emp.startDate)
    }
}
printEmployeeInformation(e1);
printEmployeeInformation({name: 'erererere', startDate: new Date()})

class Car {
    drive() {
        console.log('Driving...')
    }
}

class Truck {
    drive() {
        console.log('Driving a truck...');
    }
    loadCargo(amount: number) {
        console.log('loading this much cargo: ', amount);
    }
}

type Vehicle = Car | Truck;

const v1 = new Car();
const v2 = new Truck();

function useVehicle(vehicle: Vehicle) {
    vehicle.drive();
    if (vehicle instanceof Truck) { //if you were using an interface this would not work -- interfaces are not compiles to any javascript code
        vehicle.loadCargo(1000);
    }
}

useVehicle(v1);
useVehicle(v2);

interface Bird {
    type: 'bird'; //not a value, literal type (uuhh)
    flyingSpeed: number;
}

interface Horse {
    type: 'horse';
    runningSpeed: number;
}

type Animal = Bird | Horse;

function moveAnimal(animal: Animal) {
    let speed;
    switch (animal.type) {
        case "bird":
            speed = animal.flyingSpeed;
            break;
        case "horse":
            speed = animal.runningSpeed;
            break;
    }
}
moveAnimal({type: "horse", runningSpeed: 40});


