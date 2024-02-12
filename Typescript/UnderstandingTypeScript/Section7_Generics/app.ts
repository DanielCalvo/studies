console.log('Hello world!')

const names = ['Max', 'Manu', 'Bob'];
const a = [];

const names2: Array<string> = []; //same as string[]

names[0].split(' '); //TS knows its a string. The array can work better if you provide information about what type goes in it

//The main type here is promise, but a promise works together with other types
const promise: Promise<string> = new Promise((resolve, reject) => {
    setTimeout(() => {
        console.log("This is done!")
    }, 2000)
});

//Lets create our own generic function, neat!

function merge<T extends object, U extends object> (objA : T, objB: U) {
    return Object.assign(objA, objB); //strange error, it works!
}

const mergeObj = merge({name: 'Aaaaa'},{hobby: 'yes'})
console.log(mergeObj.hobby);
console.log(mergeObj.name);
console.log(mergeObj);

interface Lenghty {
    length: number;
}
function countAndDescribe<T extends Lenghty> (element: T): [T, string]{
    let descriptionText = 'Got no valid value';
    if (element.length > 1) {
        descriptionText = ('Got ' + element.length + ' elements')
    }
    return [element, descriptionText];
}

console.log('Count and describe with string: ',countAndDescribe('Hi there!'));
console.log('Count and describe with array: ',countAndDescribe(['one','two']))

function extractAndConvert<T extends object, U extends keyof T>(obj: T, key: U) {
    return obj[key]
}

// This only works well for primitive types -- objects might need a more refined function to manage them as they're objects that work by reference
class DataStorage<T extends string | number | boolean> {
    private data: T[] = [];
    addItem(item: T) {
        this.data.push(item);
    }
    removeItem(item: T) {
        this.data.splice(this.data.indexOf(item));
    }
    getItems(){
        return [...this.data];
    }
}

const textStorage = new DataStorage<string>();
textStorage.addItem('one');
textStorage.addItem('two');
textStorage.removeItem('one');

console.log(textStorage.getItems());

const numberStorage = new DataStorage<number>();
numberStorage .addItem(2);

//some example utility types -- for typescript only
interface CourseGoal {
    title: string;
    description: string;
    completeUntil: Date;
}

function createCourseGoal(title: string, description: string, date: Date): CourseGoal {
    return {title: title, description: description, completeUntil: date}
}