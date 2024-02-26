console.log("Hello world from section 5!");

// Sometimes you don't want the option to allow for a method to be overwriten, you want to force developers to write implementation on children objects
// You do this when you want a method to be available on all classes that are based on some base class (in this case, department)
// To do this, add abstract before class. Also, abstract classes cannot be instantiated themselves! It is now only to be inherited
abstract class Department {

    // private name: string;
    // private means that employees now is a property only accessible from inside the class!
    // private employees: string[] = []; //uh-oh this is different

    // protected is like private, but available for classes that inherit/extend this class
    protected employees: string[] = []; //uh-oh this is different

    // a function tied to this class, executed when the object is created
    // shorthand initialization!
    // Also: You can only use `readonly` once: during initialization, after that it is indeed readonly

    constructor(protected readonly id: string, public name: string) {
        // this.name = n;
    }
    abstract describe(this: Department): void;

    addEmployee(employee: string) {
        this.employees.push(employee);
    }
    printEmployeeInformation(){
        console.log(this.employees.length);
        console.log(this.employees);
    }
    //a static method that you can use without instantiating this class
    static createEmployee(name: string) {
        return {name: name};
    }
    //You can also have a static value
    //However, you cannot access static `things` from inside the class
    static fiscalYear = 2030;
}


// constructor is called when called with the `new` keyword
// const accounting = new Department('FF', 'Accounting');
// accounting.addEmployee("Joe");
// accounting.addEmployee("Rock");
//
// accounting.describe();
// accounting.printEmployeeInformation();

// const accountingCopy = {name: 's', describe: accounting.describe};
// accountingCopy.describe();

//extends means inherit! Can only inherit from one class
class ITDepartment extends Department {
    // But we can have our own constructor here!
    // whenever you write your own constructor in a class that inherits from another class, you have to use the super keyword
    constructor(id: string, public admins: string[]) {
        super(id, 'IT'); //super calls the constructor of the parent class, and therefor takes the argument of the parent class
        //if you wanna use the `this` keyword, you have to do it after calling super
    };
    describe() {
        console.log('IT Department!', this.id);
    }

}
// Works as we're inherting things from the Department class
const itDept = new ITDepartment('EE', ['IT Department']);
console.log(itDept);

class AccountingDepartment extends Department {
    private lastReport: string;
    private static instance: AccountingDepartment;

    get mostRecentReport() {
        if (this.lastReport){
            return this.lastReport;
        }
        throw new Error('No report found');
    }
    set mostRecentReport(value: string){
        if (!value) {
            throw new Error('Please pass in a valid value!')
        }
        this.addReport(value);
    }
    describe() {
        console.log("Accounting department ID:", this.id)
    }

   //we return the instance we might already have, and if we don't have one, we create a new one. This is to we can have this as a singleton
   static getInstance(){
        if (AccountingDepartment.instance) {
            return this.instance;
        }
        this.instance = new AccountingDepartment('d2',[]);
        return this.instance;
   }
   private constructor(id: string, private reports: string[]) {
        super(id, 'Accounting');
        this.lastReport = reports[0];
    }
    //we want to have a dedicated logic to the accounting department to add an employee and overwrite what we inherited
    //interesting note: private properties are only accessible in the class they were defined, not in classes that inherit from it
    addEmployee(name: string) {
        if (name === 'Joe'){
            return;
        }
        this.employees.push(name)
    }

    addReport(text: string) {
        this.reports.push(text);
        this.lastReport = text;
    }
    printReports(){
        console.log(this.reports);
    }
}

// To instantiate an accounting department as a singleton, we then do it like this:
const accounting2 = AccountingDepartment.getInstance();
accounting2.addEmployee('Dani');
// Interesting: You don't call the getter as a function, you call it as a property
console.log(accounting2.mostRecentReport)
// Same for the setter
accounting2.mostRecentReport = 'Financial report';

console.log(Department.createEmployee('aaa'));
console.log(Department.fiscalYear);

accounting2.describe()