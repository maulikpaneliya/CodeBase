// JavaScript Map, Set, and ES6+ Features

// 1. Map
// ------

// Creating a Map
const map = new Map();

// Setting key-value pairs
map.set('name', 'John');
map.set(42, 'answer');
map.set({}, 'object');

// Chaining set operations
map.set('a', 1)
   .set('b', 2)
   .set('c', 3);

// Getting values
console.log(map.get('name')); // 'John'
console.log(map.get(42));     // 'answer'

// Checking existence
console.log(map.has('name')); // true

// Size
console.log(map.size); // 5

// Deleting entries
map.delete('name');

// Iterating over Map
for (const [key, value] of map) {
    console.log(`${key} = ${value}`);
}

// Map methods
map.forEach((value, key) => {
    console.log(`${key} => ${value}`);
});

// Converting to arrays
const keys = [...map.keys()];
const values = [...map.values()];
const entries = [...map.entries()];

// Clearing the map
map.clear();

// 2. WeakMap
// ---------

// WeakMap only accepts objects as keys
const weakMap = new WeakMap();
const key1 = {};
const key2 = {};

weakMap.set(key1, 'value1');
weakMap.set(key2, 'value2');

console.log(weakMap.get(key1)); // 'value1'

// 3. Set
// -----

// Creating a Set
const set = new Set([1, 2, 3, 3, 4]); // Duplicates are automatically removed

// Adding values
set.add(5);
set.add(6).add(7); // Chainable

// Checking existence
console.log(set.has(1)); // true

// Size
console.log(set.size); // 7

// Deleting values
set.delete(1);

// Iterating over Set
for (const value of set) {
    console.log(value);
}

// Set methods
set.forEach(value => {
    console.log(value);
});

// Converting to array
const setArray = [...set];

// Set operations
const set1 = new Set([1, 2, 3]);
const set2 = new Set([2, 3, 4]);

// Union
const union = new Set([...set1, ...set2]);

// Intersection
const intersection = new Set([...set1].filter(x => set2.has(x)));

// Difference
const difference = new Set([...set1].filter(x => !set2.has(x)));

// 4. WeakSet
// ---------

const weakSet = new WeakSet();
const obj1 = {};
const obj2 = {};

weakSet.add(obj1);
weakSet.add(obj2);

console.log(weakSet.has(obj1)); // true

// 5. ES6+ Features
// -------------

// Destructuring
const person = { name: 'John', age: 30 };
const { name, age } = person;

const numbers = [1, 2, 3];
const [first, second, third] = numbers;

// Rest parameters
function sum(...numbers) {
    return numbers.reduce((total, num) => total + num, 0);
}

// Spread operator
const arr1 = [1, 2, 3];
const arr2 = [4, 5, 6];
const combined = [...arr1, ...arr2];

const obj3 = { x: 1, y: 2 };
const cloned = { ...obj3 };

// Arrow functions
const add = (a, b) => a + b;
const square = x => x * x;

// Default parameters
function greet(name = 'Guest') {
    return `Hello, ${name}!`;
}

// Object property shorthand
const x = 1, y = 2;
const point = { x, y };

// Method shorthand
const methods = {
    sayHi() {
        console.log('Hi!');
    }
};

// Computed property names
const propName = 'dynamic';
const dynamic = {
    [propName]: 42
};

// Class syntax
class Animal {
    constructor(name) {
        this.name = name;
    }

    speak() {
        console.log(`${this.name} makes a sound.`);
    }
}

// Class inheritance
class Dog extends Animal {
    speak() {
        console.log(`${this.name} barks.`);
    }
}

// Private class fields (ES2022)
class Counter {
    #count = 0;

    increment() {
        this.#count++;
    }

    get value() {
        return this.#count;
    }
}

// Optional chaining
const user = {
    address: {
        street: 'Main St'
    }
};
console.log(user?.address?.street);

// Nullish coalescing
const value = null ?? 'default';

// Promise.allSettled
Promise.allSettled([
    Promise.resolve(1),
    Promise.reject('error')
]).then(results => console.log(results));

// BigInt
const bigNumber = 9007199254740991n;

// Logical assignment operators
let a = null;
a ??= 42; // a is now 42

let b = 0;
b ||= 42; // b is now 42

let c = 1;
c &&= 42; // c is now 42

// Best Practices:
// 1. Use Map when keys can be any type
// 2. Use Set for unique value collections
// 3. Use WeakMap/WeakSet for better memory management
// 4. Prefer destructuring for cleaner code
// 5. Use spread operator for immutable operations
// 6. Utilize arrow functions for cleaner syntax
// 7. Take advantage of modern class syntax
// 8. Use optional chaining for safe property access
// 9. Consider nullish coalescing for default values
// 10. Use private class fields when encapsulation is needed
