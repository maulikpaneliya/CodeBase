// JavaScript Functions - Code Examples and Best Practices

// 1. Function Declarations
// -----------------------
// Basic function declaration
function greet(name) {
    return `Hello, ${name}!`;
}

// Function with default parameters
function welcome(name = 'Guest', greeting = 'Hello') {
    return `${greeting}, ${name}!`;
}

// Function with multiple parameters
function add(a, b) {
    return a + b;
}


// 2. Function Expressions
// ----------------------
// Anonymous function expression
const multiply = function(a, b) {
    return a * b;
};

// Named function expression
const divide = function division(a, b) {
    if (b === 0) throw new Error('Division by zero');
    return a / b;
};


// 3. Arrow Functions
// ----------------
// Basic arrow function
const square = (x) => x * x;

// Arrow function with multiple parameters
const calculateArea = (width, height) => width * height;

// Arrow function with function body
const getRandomNumber = (min, max) => {
    const random = Math.random();
    return Math.floor(random * (max - min + 1)) + min;
};

// Arrow function returning object
const createPerson = (name, age) => ({
    name,
    age,
    greet: () => `Hello, I'm ${name}`
});


// 4. Parameters and Arguments
// -------------------------
// Rest parameters
function sum(...numbers) {
    return numbers.reduce((total, num) => total + num, 0);
}

// Spread operator with functions
const numbers = [1, 2, 3];
console.log(sum(...numbers));

// Function with object parameter destructuring
function printUserInfo({ name, age, email = 'N/A' }) {
    console.log(`Name: ${name}, Age: ${age}, Email: ${email}`);
}


// 5. Closures
// ----------
function counter() {
    let count = 0;
    return {
        increment: () => ++count,
        decrement: () => --count,
        getCount: () => count
    };
}

const myCounter = counter();
console.log(myCounter.increment()); // 1
console.log(myCounter.increment()); // 2
console.log(myCounter.decrement()); // 1


// 6. Higher-Order Functions
// -----------------------
// Function that returns a function
function multiply(factor) {
    return (number) => number * factor;
}

const double = multiply(2);
const triple = multiply(3);

console.log(double(5));  // 10
console.log(triple(5));  // 15

// Function that takes a function as argument
function operate(a, b, operation) {
    return operation(a, b);
}

console.log(operate(5, 3, add));      // 8
console.log(operate(5, 3, multiply)); // 15


// 7. Recursion
// -----------
function factorial(n) {
    if (n <= 1) return 1;
    return n * factorial(n - 1);
}

// Tail-recursive version
function factorialTail(n, accumulator = 1) {
    if (n <= 1) return accumulator;
    return factorialTail(n - 1, n * accumulator);
}


// 8. Pure Functions
// ---------------
// Pure function example
function addNumbers(a, b) {
    return a + b;
}

// Impure function example (avoid)
let total = 0;
function addToTotal(value) {
    total += value;  // Side effect
    return total;
}


// Best Practices:
// 1. Use descriptive function names that explain what the function does
// 2. Keep functions small and focused on a single task
// 3. Use arrow functions for short callbacks
// 4. Use default parameters instead of checking undefined
// 5. Avoid side effects in functions when possible
// 6. Use pure functions when possible
// 7. Use parameter destructuring for cleaner function signatures
// 8. Document complex functions with JSDoc comments
