// PHP Fundamentals - Variables, Data Types, and Basic Syntax

<?php

// 1. Basic Syntax and Comments
// --------------------------

// Single-line comment
# Alternative single-line comment
/* Multi-line
   comment */

echo "Hello, World!"; // Basic output
print("Hello again!"); // Alternative output

// 2. Variables and Data Types
// ------------------------

// Variable declaration
$name = "John";
$age = 25;
$height = 1.85;
$isStudent = true;
$hobbies = ["reading", "gaming"];
$person = null;

// Variable type checking
var_dump($name);      // string
var_dump($age);       // integer
var_dump($height);    // float/double
var_dump($isStudent); // boolean
var_dump($hobbies);   // array
var_dump($person);    // NULL

// Type casting
$number = "42";
$intNumber = (int)$number;
$floatNumber = (float)$number;
$stringNumber = (string)$age;
$boolValue = (bool)$number;

// Variable scope
$globalVar = "I'm global";

function scopeExample() {
    global $globalVar;  // Accessing global variable
    $localVar = "I'm local";
    static $staticVar = 0;  // Static variable
    $staticVar++;
    
    echo $globalVar;   // Accessible with 'global'
    echo $localVar;    // Local scope only
    echo $staticVar;   // Maintains value between calls
}

// 3. Constants
// ----------

// Define constant
define('PI', 3.14159);
const MAX_USERS = 100;

echo PI;
echo MAX_USERS;

// Magic constants
echo __LINE__;    // Current line number
echo __FILE__;    // Current file path
echo __DIR__;     // Current directory
echo __FUNCTION__; // Current function name
echo __CLASS__;   // Current class name
echo __METHOD__;  // Current method name

// 4. Operators
// ----------

// Arithmetic operators
$a = 10;
$b = 3;

echo $a + $b;  // Addition
echo $a - $b;  // Subtraction
echo $a * $b;  // Multiplication
echo $a / $b;  // Division
echo $a % $b;  // Modulus
echo $a ** $b; // Exponentiation

// Assignment operators
$x = 5;
$x += 3;   // Add and assign
$x -= 2;   // Subtract and assign
$x *= 4;   // Multiply and assign
$x /= 2;   // Divide and assign
$x %= 3;   // Modulus and assign
$x .= "!"; // Concatenate and assign

// Comparison operators
$p = 5;
$q = "5";

var_dump($p == $q);   // Equal (value)
var_dump($p === $q);  // Identical (value and type)
var_dump($p != $q);   // Not equal
var_dump($p !== $q);  // Not identical
var_dump($p < $q);    // Less than
var_dump($p <= $q);   // Less than or equal
var_dump($p > $q);    // Greater than
var_dump($p >= $q);   // Greater than or equal
var_dump($p <=> $q);  // Spaceship operator

// Logical operators
$t = true;
$f = false;

var_dump($t && $f);  // Logical AND
var_dump($t || $f);  // Logical OR
var_dump(!$t);       // Logical NOT
var_dump($t and $f); // Alternative AND
var_dump($t or $f);  // Alternative OR
var_dump($t xor $f); // Exclusive OR

// String operators
$str1 = "Hello";
$str2 = "World";
echo $str1 . " " . $str2;  // Concatenation

// 5. Type Juggling
// -------------

// Automatic type conversion
$num = 5 + "10";     // 15 (integer)
$concat = "5" . "10"; // "510" (string)
$sum = 1 + "2.5";    // 3.5 (float)

// Type checking functions
is_int($age);
is_float($height);
is_string($name);
is_bool($isStudent);
is_array($hobbies);
is_null($person);
is_numeric($number);
is_object($obj);

// 6. Error Handling
// --------------

// Error reporting
error_reporting(E_ALL);
ini_set('display_errors', 1);

// Try-catch block
try {
    // Some code that might throw an exception
    throw new Exception("Something went wrong!");
} catch (Exception $e) {
    echo "Caught exception: " . $e->getMessage();
} finally {
    echo "This always executes";
}

// Custom error handler
function customErrorHandler($errno, $errstr, $errfile, $errline) {
    echo "Error [$errno]: $errstr in $errfile on line $errline";
    return true;
}
set_error_handler("customErrorHandler");

// Best Practices:
// 1. Use meaningful variable names
// 2. Be consistent with naming conventions
// 3. Always initialize variables
// 4. Use appropriate data types
// 5. Use proper error handling
// 6. Comment your code appropriately
// 7. Follow PHP coding standards
// 8. Be careful with type juggling
// 9. Use constants for fixed values
// 10. Handle errors gracefully
