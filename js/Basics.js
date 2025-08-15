// JavaScript Basics - Code Examples and Best Practices

// 1. Variables
// -----------------
// Using let for variables that will be reassigned
let age = 25;
age = 26; // can be reassigned

// Using const for values that won't change
const PI = 3.14159;
const MAX_USERS = 100;

// Using var (older way, function-scoped)
var name = "John"; // Not recommended in modern JavaScript


// 2. Data Types
// -----------------
// Numbers
const integer = 42;
const float = 3.14;
const scientific = 2.998e8;

// Strings
const single = 'Single quotes';
const double = "Double quotes";
const template = `Template literal: ${integer}`; // String interpolation

// Booleans
const isTrue = true;
const isFalse = false;

// Arrays
const fruits = ['apple', 'banana', 'orange'];
const mixed = [1, 'two', true, { id: 1 }];

// Objects
const person = {
    firstName: 'John',
    lastName: 'Doe',
    age: 30,
    hobbies: ['reading', 'music']
};

// Null and Undefined
const nullValue = null;
let undefinedValue;


// 3. Operators
// -----------------
// Arithmetic
const sum = 5 + 3;
const difference = 10 - 5;
const product = 4 * 2;
const quotient = 15 / 3;
const remainder = 17 % 5;

// Comparison
const isEqual = 5 === 5;        // strict equality
const isNotEqual = 5 !== '5';   // strict inequality
const isGreater = 10 > 5;
const isLessOrEqual = 5 <= 5;

// Logical
const andOperator = true && false;  // false
const orOperator = true || false;   // true
const notOperator = !true;          // false


// 4. Control Flow
// -----------------
// If statements
if (age >= 18) {
    console.log('Adult');
} else if (age >= 13) {
    console.log('Teenager');
} else {
    console.log('Child');
}

// Switch statement
const day = 'Monday';
switch (day) {
    case 'Monday':
        console.log('Start of work week');
        break;
    case 'Friday':
        console.log('End of work week');
        break;
    default:
        console.log('Regular day');
}

// Loops
// For loop
for (let i = 0; i < 5; i++) {
    console.log(`Iteration ${i}`);
}

// While loop
let counter = 0;
while (counter < 3) {
    console.log(`Count: ${counter}`);
    counter++;
}

// For...of loop (arrays)
for (const fruit of fruits) {
    console.log(fruit);
}

// For...in loop (objects)
for (const key in person) {
    console.log(`${key}: ${person[key]}`);
}


// 5. Type Conversion
// -----------------
// String to Number
const stringNumber = "42";
const convertedNumber = Number(stringNumber);    // 42
const parsedInt = parseInt("42.9");             // 42
const parsedFloat = parseFloat("42.9");         // 42.9

// Number to String
const numberString = String(42);                // "42"
const numberToString = (42).toString();         // "42"

// Boolean Conversion
const truthyValue = Boolean(1);                 // true
const falsyValue = Boolean(0);                  // false

// Best Practices:
// 1. Always use const by default, and only use let if reassignment is needed
// 2. Use meaningful variable names
// 3. Use strict equality (===) instead of loose equality (==)
// 4. Always use semicolons at the end of statements
// 5. Use modern ES6+ features when possible
// 6. Be consistent with your coding style
