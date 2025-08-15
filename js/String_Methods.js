// JavaScript String Methods and Manipulation

// 1. String Creation
// ---------------

// String literals
const str1 = 'Hello';
const str2 = "World";
const str3 = `Hello ${str2}`; // Template literal

// String constructor
const str4 = new String('Hello'); // Not recommended

// 2. Basic String Methods
// -------------------

const text = 'Hello World';

// Length
console.log(text.length); // 11

// Case methods
console.log(text.toLowerCase()); // 'hello world'
console.log(text.toUpperCase()); // 'HELLO WORLD'

// Trim methods
const paddedText = '  Hello World  ';
console.log(paddedText.trim());      // 'Hello World'
console.log(paddedText.trimLeft());  // 'Hello World  '
console.log(paddedText.trimRight()); // '  Hello World'

// 3. String Search Methods
// --------------------

const sentence = 'The quick brown fox jumps over the lazy dog';

// indexOf and lastIndexOf
console.log(sentence.indexOf('quick'));     // 4
console.log(sentence.lastIndexOf('the'));   // 31

// includes, startsWith, endsWith
console.log(sentence.includes('fox'));      // true
console.log(sentence.startsWith('The'));    // true
console.log(sentence.endsWith('dog'));      // true

// search with regex
console.log(sentence.search(/[A-Z]/));      // 0

// match and matchAll
const matches = sentence.match(/the/gi);
const allMatches = [...sentence.matchAll(/the/gi)];

// 4. String Extraction Methods
// ------------------------

// slice
console.log(text.slice(0, 5));     // 'Hello'
console.log(text.slice(-5));       // 'World'

// substring
console.log(text.substring(6, 11)); // 'World'

// split
console.log(text.split(' '));      // ['Hello', 'World']
console.log(text.split(''));       // ['H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd']

// 5. String Modification Methods
// --------------------------

// replace and replaceAll
const modified = text.replace('Hello', 'Hi');
const allReplaced = text.replaceAll('l', 'L');

// concat
const combined = str1.concat(' ', str2);

// repeat
console.log('Ha'.repeat(3)); // 'HaHaHa'

// padStart and padEnd
console.log('5'.padStart(3, '0'));  // '005'
console.log('5'.padEnd(3, '0'));    // '500'

// 6. Regular Expressions
// ------------------

const email = 'user@example.com';

// Email validation
const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
console.log(emailRegex.test(email)); // true

// Pattern matching
const pattern = /\b\w+@\w+\.\w+\b/;
console.log(pattern.exec(email));

// Global search
const globalPattern = /\w+/g;
let match;
while ((match = globalPattern.exec(text)) !== null) {
    console.log(match[0]);
}

// 7. Unicode and Special Characters
// ----------------------------

// Unicode escape sequences
const heart = '\u2764';
const smile = '😊';

// Character codes
const char = String.fromCharCode(65); // 'A'
const code = 'A'.charCodeAt(0);       // 65

// Unicode code points
const symbol = String.fromCodePoint(0x2764);
const codePoint = smile.codePointAt(0);

// 8. Template Literals
// ----------------

const name = 'John';
const age = 30;

// Basic interpolation
const greeting = `Hello, ${name}!`;

// Multi-line strings
const html = `
    <div>
        <h1>${name}</h1>
        <p>Age: ${age}</p>
    </div>
`;

// Tagged templates
function highlight(strings, ...values) {
    return strings.reduce((result, str, i) => {
        return result + str + (values[i] ? `<span class="highlight">${values[i]}</span>` : '');
    }, '');
}

const highlighted = highlight`Hello ${name}, you are ${age} years old`;

// 9. String Performance
// -----------------

// Efficient string concatenation
const parts = ['Hello', 'World', '!'];
const efficient = parts.join(' ');

// Using StringBuilder pattern
class StringBuilder {
    constructor() {
        this.parts = [];
    }

    append(text) {
        this.parts.push(text);
        return this;
    }

    toString() {
        return this.parts.join('');
    }
}

// 10. Localization
// -------------

const number = 123456.789;
const date = new Date();

// Number formatting
console.log(number.toLocaleString('de-DE'));
console.log(number.toLocaleString('ar-EG'));

// Date formatting
console.log(date.toLocaleDateString('en-US'));
console.log(date.toLocaleDateString('ja-JP'));

// Best Practices:
// 1. Use template literals for string interpolation
// 2. Use appropriate string methods for the task
// 3. Consider using regex for complex pattern matching
// 4. Be careful with string mutations
// 5. Use efficient string concatenation methods
// 6. Handle Unicode properly
// 7. Consider localization requirements
// 8. Use string validation when needed
// 9. Handle special characters appropriately
// 10. Consider performance for large string operations
