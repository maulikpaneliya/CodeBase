// Modern JavaScript Features (ES6+)

// 1. Arrow Functions
// ---------------
const add = (a, b) => a + b;
const multiply = (a, b) => {
    return a * b;
};

// With default parameters
const greet = (name = 'Guest') => `Hello, ${name}!`;

// Arrow function with object return
const createUser = (name, age) => ({ name, age });

// 2. Destructuring
// -------------

// Array destructuring
const numbers = [1, 2, 3, 4, 5];
const [first, second, ...rest] = numbers;

// Object destructuring
const user = {
    name: 'John',
    age: 30,
    address: {
        city: 'New York',
        country: 'USA'
    }
};

const { name, age, address: { city } } = user;

// Parameter destructuring
const printUserInfo = ({ name, age }) => {
    console.log(`${name} is ${age} years old`);
};

// 3. Template Literals
// ----------------

// Basic usage
const message = `Hello, ${name}!`;

// Multi-line strings
const html = `
    <div>
        <h1>${title}</h1>
        <p>${content}</p>
    </div>
`;

// Tagged templates
const highlight = (strings, ...values) => {
    return strings.reduce((result, str, i) => {
        return result + str + (values[i] ? `<span class="highlight">${values[i]}</span>` : '');
    }, '');
};

// 4. Spread/Rest Operators
// ---------------------

// Spread with arrays
const arr1 = [1, 2, 3];
const arr2 = [4, 5, 6];
const combined = [...arr1, ...arr2];

// Spread with objects
const baseConfig = { api: 'http://api.example.com', timeout: 5000 };
const extendedConfig = { ...baseConfig, debug: true };

// Rest parameters
const sum = (...numbers) => numbers.reduce((total, num) => total + num, 0);

// 5. Object Enhancements
// -------------------

// Shorthand properties
const x = 10, y = 20;
const point = { x, y };

// Computed property names
const propName = 'age';
const person = {
    name: 'John',
    [propName]: 30
};

// Method shorthand
const calculator = {
    add(a, b) {
        return a + b;
    },
    multiply(a, b) {
        return a * b;
    }
};

// 6. Array Methods
// -------------

// map
const doubled = numbers.map(num => num * 2);

// filter
const evenNumbers = numbers.filter(num => num % 2 === 0);

// reduce
const sum = numbers.reduce((total, num) => total + num, 0);

// find
const found = numbers.find(num => num > 3);

// some/every
const hasNegative = numbers.some(num => num < 0);
const allPositive = numbers.every(num => num > 0);

// flatMap
const sentences = ['Hello world', 'How are you'];
const words = sentences.flatMap(sentence => sentence.split(' '));

// 7. Optional Chaining
// -----------------

// Object optional chaining
const cityName = user?.address?.city;

// Array optional chaining
const firstItem = array?.[0];

// Method optional chaining
const result = object.method?.();

// 8. Nullish Coalescing
// ------------------

const value = null;
const defaultValue = value ?? 'default';

// 9. Modules
// --------

// Named exports
export const helper = () => {};
export class MyClass {}

// Default export
export default class MainClass {}

// Import syntax
import MainClass, { helper, MyClass } from './module';

// 10. Promises and Async/Await
// ------------------------

// Promise creation
const delay = ms => new Promise(resolve => setTimeout(resolve, ms));

// Promise chaining
fetchUser()
    .then(user => fetchUserPosts(user.id))
    .then(posts => displayPosts(posts))
    .catch(error => console.error(error));

// Async/await
async function loadData() {
    try {
        const user = await fetchUser();
        const posts = await fetchUserPosts(user.id);
        return posts;
    } catch (error) {
        console.error(error);
    }
}

// Promise.all
const promises = [fetchUsers(), fetchPosts(), fetchComments()];
Promise.all(promises)
    .then(([users, posts, comments]) => {
        // Handle all results
    });

// 11. Classes and Inheritance
// -----------------------

class Animal {
    constructor(name) {
        this.name = name;
    }

    speak() {
        return `${this.name} makes a sound`;
    }

    static create(name) {
        return new this(name);
    }
}

class Dog extends Animal {
    constructor(name, breed) {
        super(name);
        this.breed = breed;
    }

    speak() {
        return `${this.name} barks`;
    }
}

// 12. Map and Set
// ------------

// Map
const userMap = new Map();
userMap.set('john', { name: 'John', age: 30 });
userMap.set('jane', { name: 'Jane', age: 25 });

// Set
const uniqueNumbers = new Set([1, 2, 2, 3, 3, 4]);

// WeakMap and WeakSet
const weakMap = new WeakMap();
const weakSet = new WeakSet();

// Best Practices:
// 1. Use const by default, let when needed
// 2. Prefer arrow functions for callbacks
// 3. Use destructuring for cleaner code
// 4. Utilize template literals for string interpolation
// 5. Use spread/rest operators for flexible functions
// 6. Leverage modern array methods
// 7. Use optional chaining for safe property access
// 8. Implement proper error handling with async/await
// 9. Use modules for better code organization
// 10. Understand and use proper data structures (Map/Set)
