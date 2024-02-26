const person: {
    name: string;
    age: number;
} = {
    name: 'daniel',
    age: 30
};

//Apparently this syntax is better
const person2: {
    name: string;
    age: number;
    hobbies: string[];
    role: [number, string]; //first element should be a number, second element should be a string, and I should have exactly two tyes
} = {
    name: "joe",
    age: 100,
    hobbies: ['sports', 'cooking'],
    role: [2, 'author']
};

//not enums
const ADMIN = 0;
const READ_ONLY = 1;
const AUTHOR = 2;
const person3 = {
    role: AUTHOR
};
if (person3.role === AUTHOR) {
    console.log("is author!");
}

// actual enum
// you can assign your own numbers or values to enums, anything goes, strings work too!
enum Role {ADMIN, READ_ONLY, AUTHOR}
const person4 = {
    role: Role.ADMIN
};
console.log(person4.role);

let favoriteActivities: string[];
let anarchy: any[]; //you're back in JS world -- you can use any value
anarchy[0] = 0;
anarchy[1] = "4";
anarchy[2] = true;
console.log(anarchy);

for (const hobby of person2.hobbies) {
    console.log(hobby)
}

console.log(person.name);
console.log(person2.name, person2.age);