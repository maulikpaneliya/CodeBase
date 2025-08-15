// PHP Control Structures and Functions

<?php

// 1. Conditional Statements
// ----------------------

// If-else statement
$age = 20;

if ($age < 13) {
    echo "Child";
} elseif ($age < 20) {
    echo "Teenager";
} else {
    echo "Adult";
}

// Switch statement
$dayOfWeek = "Monday";

switch ($dayOfWeek) {
    case "Monday":
        echo "Start of work week";
        break;
    case "Friday":
        echo "End of work week";
        break;
    case "Saturday":
    case "Sunday":
        echo "Weekend";
        break;
    default:
        echo "Regular work day";
}

// Match expression (PHP 8+)
$result = match($dayOfWeek) {
    "Monday" => "Start of work week",
    "Friday" => "End of work week",
    "Saturday", "Sunday" => "Weekend",
    default => "Regular work day",
};

// Ternary operator
$isAdult = ($age >= 18) ? "Adult" : "Minor";

// Null coalescing operator
$name = $_GET['name'] ?? 'Anonymous';

// 2. Loops
// -------

// For loop
for ($i = 0; $i < 5; $i++) {
    echo "Iteration $i";
}

// While loop
$counter = 0;
while ($counter < 5) {
    echo "Count: $counter";
    $counter++;
}

// Do-while loop
$num = 1;
do {
    echo "Number: $num";
    $num++;
} while ($num <= 5);

// Foreach loop
$colors = ['red', 'blue', 'green'];
foreach ($colors as $color) {
    echo $color;
}

// Foreach with key
$person = [
    'name' => 'John',
    'age' => 25,
    'city' => 'New York'
];

foreach ($person as $key => $value) {
    echo "$key: $value";
}

// 3. Functions
// ----------

// Basic function
function greet($name) {
    return "Hello, $name!";
}

// Function with default parameters
function calculateTotal($price, $taxRate = 0.1) {
    return $price * (1 + $taxRate);
}

// Type declarations (PHP 7+)
function add(int $a, int $b): int {
    return $a + $b;
}

// Nullable types (PHP 7.1+)
function processUser(?string $name): ?string {
    return $name ? "Processing $name" : null;
}

// Variable-length argument lists
function sum(...$numbers) {
    return array_sum($numbers);
}

// Arrow functions (PHP 7.4+)
$multiply = fn($a, $b) => $a * $b;

// Closure
$greeting = function($name) use ($defaultGreeting) {
    return $defaultGreeting . ", " . $name;
};

// Return type declarations
function divide(float $a, float $b): ?float {
    return $b != 0 ? $a / $b : null;
}

// Generator function
function range_generator($start, $end) {
    for ($i = $start; $i <= $end; $i++) {
        yield $i;
    }
}

// 4. Include and Require
// -------------------

// Including files
include 'header.php';
include_once 'config.php';
require 'functions.php';
require_once 'database.php';

// 5. Namespaces
// -----------

namespace App\Helpers;

// Using namespaces
use App\Models\User;
use App\Services\{
    EmailService,
    PaymentService
};

// 6. Error Control
// -------------

// Error suppression operator
$value = @file_get_contents('non_existent_file.txt');

// Custom error handling
function handleError($errno, $errstr) {
    echo "Error [$errno]: $errstr";
    return true;
}
set_error_handler('handleError');

// Exception handling
try {
    validateAge(-5);
} catch (InvalidArgumentException $e) {
    echo "Invalid age: " . $e->getMessage();
} catch (Exception $e) {
    echo "General error: " . $e->getMessage();
} finally {
    echo "Validation complete";
}

// 7. Alternative Syntax
// ------------------

// Alternative if
if ($condition):
    echo "True";
else:
    echo "False";
endif;

// Alternative while
while ($i < 5):
    echo $i;
    $i++;
endwhile;

// Alternative for
for ($i = 0; $i < 5; $i++):
    echo $i;
endfor;

// Alternative foreach
foreach ($array as $value):
    echo $value;
endforeach;

// Best Practices:
// 1. Use clear and consistent control structures
// 2. Keep functions small and focused
// 3. Use type declarations when possible
// 4. Handle errors appropriately
// 5. Use meaningful function names
// 6. Document complex functions
// 7. Use proper namespacing
// 8. Follow PSR standards
// 9. Avoid deep nesting
// 10. Use proper error handling
