// Advanced Array Methods and Data Manipulation

// 1. Array Creation and Initialization
// --------------------------------

// Array literal
const fruits = ['apple', 'banana', 'orange'];

// Array constructor
const numbers = new Array(1, 2, 3, 4, 5);

// Array.from
const arrayFromString = Array.from('hello'); // ['h', 'e', 'l', 'l', 'o']
const arrayFromSet = Array.from(new Set([1, 2, 2, 3])); // [1, 2, 3]

// Array.of
const arrayOf = Array.of(1, 2, 3); // [1, 2, 3]

// Fill array
const filledArray = new Array(3).fill(0); // [0, 0, 0]

// 2. Basic Array Methods
// -------------------

// Adding/Removing Elements
const stack = [];
stack.push(1, 2);         // Add to end
stack.pop();              // Remove from end
stack.unshift(0);         // Add to beginning
stack.shift();            // Remove from beginning

// Slicing and Splicing
const original = [1, 2, 3, 4, 5];
const slice = original.slice(1, 3);  // [2, 3]
original.splice(1, 2, 'a', 'b');     // Replace elements

// 3. Array Transformation
// --------------------

const numbers = [1, 2, 3, 4, 5];

// Map
const doubled = numbers.map(num => num * 2);
const objects = numbers.map(num => ({ value: num }));

// Filter
const evenNumbers = numbers.filter(num => num % 2 === 0);
const positiveNumbers = numbers.filter(num => num > 0);

// Reduce
const sum = numbers.reduce((acc, curr) => acc + curr, 0);
const product = numbers.reduce((acc, curr) => acc * curr, 1);

// ReduceRight
const flattenedRight = [[0, 1], [2, 3], [4, 5]].reduceRight(
    (acc, curr) => [...acc, ...curr],
    []
);

// 4. Array Searching and Testing
// --------------------------

// Find and FindIndex
const users = [
    { id: 1, name: 'John' },
    { id: 2, name: 'Jane' },
    { id: 3, name: 'Bob' }
];

const user = users.find(user => user.id === 2);
const userIndex = users.findIndex(user => user.id === 2);

// Includes and IndexOf
const includes = numbers.includes(3);
const index = numbers.indexOf(3);
const lastIndex = numbers.lastIndexOf(3);

// Some and Every
const hasEven = numbers.some(num => num % 2 === 0);
const allPositive = numbers.every(num => num > 0);

// 5. Array Flattening and Mapping
// ---------------------------

// Flat
const nested = [1, [2, 3], [[4, 5], 6]];
const flattened = nested.flat();      // [1, 2, 3, [4, 5], 6]
const fullyFlattened = nested.flat(Infinity); // [1, 2, 3, 4, 5, 6]

// FlatMap
const sentences = ['Hello world', 'How are you'];
const words = sentences.flatMap(sentence => sentence.split(' '));

// 6. Array Sorting and Reversing
// --------------------------

// Sort
const unsorted = [3, 1, 4, 1, 5];
const sorted = [...unsorted].sort((a, b) => a - b);

// Custom sort
const users = [
    { name: 'John', age: 30 },
    { name: 'Jane', age: 25 },
    { name: 'Bob', age: 35 }
];

const sortedByAge = [...users].sort((a, b) => a.age - b.age);
const sortedByName = [...users].sort((a, b) => a.name.localeCompare(b.name));

// Reverse
const reversed = [...numbers].reverse();

// 7. Array Filling and Copying
// -------------------------

// Fill
const filled = new Array(3).fill(0);
const partiallyFilled = [1, 2, 3, 4].fill(0, 1, 3); // [1, 0, 0, 4]

// CopyWithin
const copied = [1, 2, 3, 4, 5].copyWithin(0, 3); // [4, 5, 3, 4, 5]

// 8. Advanced Array Operations
// ------------------------

// Removing duplicates
const uniqueValues = [...new Set([1, 2, 2, 3, 3, 4])];

// Chunking array
const chunk = (arr, size) => 
    Array.from({ length: Math.ceil(arr.length / size) }, (_, i) =>
        arr.slice(i * size, i * size + size)
    );

// Array intersection
const intersection = (arr1, arr2) =>
    arr1.filter(item => arr2.includes(item));

// Array difference
const difference = (arr1, arr2) =>
    arr1.filter(item => !arr2.includes(item));

// Array union
const union = (arr1, arr2) =>
    [...new Set([...arr1, ...arr2])];

// 9. Array Performance Optimization
// ----------------------------

// Pre-allocating array length
const optArray = new Array(1000);
for (let i = 0; i < 1000; i++) {
    optArray[i] = i;
}

// Using appropriate methods
const numbers = [1, 2, 3, 4, 5];

// Good - single iteration
const result = numbers.reduce((acc, num) => {
    if (num % 2 === 0) {
        acc.push(num * 2);
    }
    return acc;
}, []);

// Bad - multiple iterations
const result2 = numbers.filter(num => num % 2 === 0).map(num => num * 2);

// 10. Array Type Checking
// -------------------

// Check if variable is array
Array.isArray([]);        // true
Array.isArray({});        // false
Array.isArray('array');   // false

// Check if array-like object
const isArrayLike = obj =>
    obj != null && typeof obj[Symbol.iterator] === 'function';

// Best Practices:
// 1. Use appropriate array methods for the task
// 2. Consider performance implications for large arrays
// 3. Use immutable operations when needed
// 4. Properly handle edge cases
// 5. Use type checking when necessary
// 6. Consider using Set for unique values
// 7. Use array methods over loops when possible
// 8. Chain array methods responsibly
// 9. Consider memory usage with large arrays
// 10. Use appropriate error handling
