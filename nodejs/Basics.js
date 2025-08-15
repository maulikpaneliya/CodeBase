// Node.js Basics - Core Concepts and Examples

// 1. Modules System
// ----------------

// Custom module export (math.js)
const add = (a, b) => a + b;
const subtract = (a, b) => a - b;
const multiply = (a, b) => a * b;
const divide = (a, b) => b !== 0 ? a / b : null;

module.exports = {
    add,
    subtract,
    multiply,
    divide
};

// Importing custom modules
const math = require('./math');
console.log(math.add(5, 3));  // Output: 8

// ES Modules syntax (requires "type": "module" in package.json)
/*
import { add, subtract } from './math.js';
console.log(add(5, 3));  // Output: 8
*/

// 2. Global Objects
// ---------------

// Process object
console.log(process.env.NODE_ENV);  // Environment variables
console.log(process.argv);          // Command line arguments
console.log(process.cwd());         // Current working directory

// Global timing functions
setTimeout(() => {
    console.log('Delayed message');
}, 1000);

setInterval(() => {
    console.log('Recurring message');
}, 1000);

// 3. Event Loop
// -----------

// Example of non-blocking I/O
const fs = require('fs');

console.log('Start reading file...');

fs.readFile('example.txt', 'utf8', (err, data) => {
    if (err) {
        console.error('Error reading file:', err);
        return;
    }
    console.log('File contents:', data);
});

console.log('Continue executing...');

// 4. Async Patterns
// --------------

// Callbacks
function fetchData(callback) {
    setTimeout(() => {
        callback(null, { data: 'Sample data' });
    }, 1000);
}

fetchData((err, result) => {
    if (err) {
        console.error('Error:', err);
        return;
    }
    console.log('Result:', result);
});

// Promises
function fetchDataPromise() {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            resolve({ data: 'Sample data' });
        }, 1000);
    });
}

fetchDataPromise()
    .then(result => console.log('Result:', result))
    .catch(err => console.error('Error:', err));

// Async/Await
async function fetchDataAsync() {
    try {
        const result = await fetchDataPromise();
        console.log('Result:', result);
    } catch (err) {
        console.error('Error:', err);
    }
}

// 5. Error Handling
// --------------

// Try-catch blocks
try {
    throw new Error('Sample error');
} catch (err) {
    console.error('Caught error:', err.message);
}

// Uncaught exception handler
process.on('uncaughtException', (err) => {
    console.error('Uncaught Exception:', err);
    // Perform cleanup
    process.exit(1);
});

// Unhandled promise rejection handler
process.on('unhandledRejection', (reason, promise) => {
    console.error('Unhandled Rejection:', reason);
});

// 6. Events
// --------

const EventEmitter = require('events');

class MyEmitter extends EventEmitter {}
const myEmitter = new MyEmitter();

// Event listener
myEmitter.on('event', (arg) => {
    console.log('Event occurred:', arg);
});

// Emit event
myEmitter.emit('event', 'test data');

// 7. Buffers
// --------

// Creating buffers
const buf1 = Buffer.from('Hello');
const buf2 = Buffer.alloc(5);
buf2.write('World');

console.log(buf1.toString());  // Output: Hello
console.log(buf2.toString());  // Output: World

// 8. Environment Variables
// --------------------

// Using dotenv for environment variables
require('dotenv').config();

const dbConfig = {
    host: process.env.DB_HOST || 'localhost',
    port: process.env.DB_PORT || 5432,
    username: process.env.DB_USER,
    password: process.env.DB_PASS
};

// Best Practices:
// 1. Always handle errors appropriately
// 2. Use async/await for cleaner async code
// 3. Properly manage environment variables
// 4. Use event emitters for custom events
// 5. Handle process events for stability
// 6. Use proper module systems
// 7. Follow the Node.js error-first callback pattern
// 8. Use Buffer for binary data
// 9. Implement proper logging
// 10. Use proper security practices
