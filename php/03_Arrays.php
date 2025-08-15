// PHP Arrays and Data Manipulation

<?php

// 1. Array Types
// ------------

// Indexed arrays
$fruits = ['apple', 'banana', 'orange'];
$numbers = [1, 2, 3, 4, 5];

// Associative arrays
$person = [
    'name' => 'John Doe',
    'age' => 30,
    'city' => 'New York'
];

// Multidimensional arrays
$users = [
    ['name' => 'John', 'age' => 30],
    ['name' => 'Jane', 'age' => 25],
    ['name' => 'Bob', 'age' => 35]
];

$matrix = [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9]
];

// 2. Array Operations
// ----------------

// Adding elements
$fruits[] = 'grape';                // Add to end
array_push($fruits, 'mango');       // Add to end
array_unshift($fruits, 'pear');     // Add to beginning

// Removing elements
array_pop($fruits);                 // Remove from end
array_shift($fruits);               // Remove from beginning
unset($fruits[1]);                  // Remove specific element

// Merging arrays
$allFruits = array_merge($fruits, ['kiwi', 'melon']);
$combined = [...$fruits, ...$numbers]; // Spread operator (PHP 7.4+)

// 3. Array Functions
// ---------------

// Checking arrays
is_array($fruits);              // Check if variable is array
empty($fruits);                 // Check if array is empty
isset($fruits[0]);             // Check if element exists
in_array('apple', $fruits);    // Check if value exists
array_key_exists('name', $person); // Check if key exists

// Array information
count($fruits);                // Count elements
sizeof($fruits);               // Alternative to count
array_count_values($fruits);   // Count frequency of values

// Array manipulation
sort($fruits);                 // Sort indexed array
rsort($fruits);                // Reverse sort
asort($person);               // Sort associative array by values
ksort($person);               // Sort associative array by keys
shuffle($fruits);             // Random shuffle

// Array transformation
array_map(fn($n) => $n * 2, $numbers);  // Transform elements
array_filter($numbers, fn($n) => $n > 2); // Filter elements
array_reduce($numbers, fn($carry, $n) => $carry + $n, 0); // Reduce to single value

// 4. Array Iteration
// ---------------

// Basic foreach
foreach ($fruits as $fruit) {
    echo $fruit;
}

// Foreach with key
foreach ($person as $key => $value) {
    echo "$key: $value";
}

// Array walking
array_walk($fruits, function($value, $key) {
    echo "Key: $key, Value: $value";
});

// 5. Array Search and Extraction
// --------------------------

// Searching
array_search('apple', $fruits);     // Find value's index
array_keys($person);                // Get all keys
array_values($person);              // Get all values

// Extracting
array_slice($fruits, 1, 2);        // Extract portion
array_splice($fruits, 1, 1, ['new']); // Remove and replace
array_chunk($fruits, 2);           // Split into chunks

// 6. Array Set Operations
// --------------------

$array1 = [1, 2, 3];
$array2 = [3, 4, 5];

array_diff($array1, $array2);       // Values in array1 but not array2
array_intersect($array1, $array2);  // Values in both arrays
array_unique(array_merge($array1, $array2)); // Unique values

// 7. Practical Examples
// -----------------

// Grouping data
$students = [
    ['name' => 'John', 'grade' => 'A'],
    ['name' => 'Jane', 'grade' => 'B'],
    ['name' => 'Bob', 'grade' => 'A']
];

$groupedByGrade = array_reduce($students, function($result, $student) {
    $result[$student['grade']][] = $student['name'];
    return $result;
}, []);

// Creating lookup table
$users = [
    ['id' => 1, 'name' => 'John'],
    ['id' => 2, 'name' => 'Jane']
];

$userLookup = array_column($users, null, 'id');

// Filtering and mapping
$numbers = range(1, 10);
$evenSquares = array_map(
    fn($n) => $n * $n,
    array_filter($numbers, fn($n) => $n % 2 === 0)
);

// 8. Advanced Array Functions
// -----------------------

// Array column
array_column($users, 'name');           // Extract single column
array_column($users, 'name', 'id');     // Extract with keys

// Array combine
$keys = ['a', 'b', 'c'];
$values = [1, 2, 3];
array_combine($keys, $values);          // Create associative array

// Array flip
array_flip($person);                    // Swap keys and values

// Array replace
array_replace($array1, $array2);        // Replace elements

// Best Practices:
// 1. Use meaningful array keys
// 2. Choose appropriate array type
// 3. Use type hints when possible
// 4. Handle array operations safely
// 5. Consider memory usage with large arrays
// 6. Use array functions instead of loops when possible
// 7. Validate array inputs
// 8. Document complex array structures
// 9. Use consistent naming conventions
// 10. Consider using collections for complex operations
