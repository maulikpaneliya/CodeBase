// PHP Object-Oriented Programming

<?php

// 1. Basic Class Definition
// ----------------------

class User {
    // Properties
    public string $name;
    protected int $age;
    private string $email;
    
    // Constants
    const MINIMUM_AGE = 18;
    
    // Constructor
    public function __construct(string $name, int $age, string $email) {
        $this->name = $name;
        $this->age = $age;
        $this->email = $email;
    }
    
    // Methods
    public function getInfo(): string {
        return "Name: {$this->name}, Age: {$this->age}";
    }
    
    // Getter and Setter
    public function getEmail(): string {
        return $this->email;
    }
    
    public function setEmail(string $email): void {
        if (filter_var($email, FILTER_VALIDATE_EMAIL)) {
            $this->email = $email;
        }
    }
}

// 2. Inheritance
// -----------

class Admin extends User {
    private array $permissions;
    
    public function __construct(string $name, int $age, string $email, array $permissions) {
        parent::__construct($name, $age, $email);
        $this->permissions = $permissions;
    }
    
    public function hasPermission(string $permission): bool {
        return in_array($permission, $this->permissions);
    }
}

// 3. Abstract Classes
// ----------------

abstract class Shape {
    abstract public function getArea(): float;
    abstract public function getPerimeter(): float;
    
    public function getInfo(): string {
        return "Area: {$this->getArea()}, Perimeter: {$this->getPerimeter()}";
    }
}

class Circle extends Shape {
    private float $radius;
    
    public function __construct(float $radius) {
        $this->radius = $radius;
    }
    
    public function getArea(): float {
        return pi() * pow($this->radius, 2);
    }
    
    public function getPerimeter(): float {
        return 2 * pi() * $this->radius;
    }
}

// 4. Interfaces
// ----------

interface PaymentProcessor {
    public function process(float $amount): bool;
    public function refund(float $amount): bool;
}

class StripeProcessor implements PaymentProcessor {
    public function process(float $amount): bool {
        // Implementation
        return true;
    }
    
    public function refund(float $amount): bool {
        // Implementation
        return true;
    }
}

// 5. Traits
// -------

trait Timestampable {
    private DateTime $createdAt;
    private ?DateTime $updatedAt = null;
    
    public function getCreatedAt(): DateTime {
        return $this->createdAt;
    }
    
    public function setUpdatedAt(): void {
        $this->updatedAt = new DateTime();
    }
}

class Article {
    use Timestampable;
    
    private string $title;
    private string $content;
    
    public function __construct(string $title, string $content) {
        $this->title = $title;
        $this->content = $content;
        $this->createdAt = new DateTime();
    }
}

// 6. Static Members
// --------------

class Database {
    private static ?Database $instance = null;
    private PDO $connection;
    
    private function __construct() {
        // Private constructor to prevent direct creation
    }
    
    public static function getInstance(): Database {
        if (self::$instance === null) {
            self::$instance = new self();
        }
        return self::$instance;
    }
    
    public static function connect(string $dsn, string $username, string $password): void {
        self::getInstance()->connection = new PDO($dsn, $username, $password);
    }
}

// 7. Magic Methods
// -------------

class Product {
    private array $data = [];
    
    public function __get(string $name) {
        return $this->data[$name] ?? null;
    }
    
    public function __set(string $name, $value): void {
        $this->data[$name] = $value;
    }
    
    public function __isset(string $name): bool {
        return isset($this->data[$name]);
    }
    
    public function __unset(string $name): void {
        unset($this->data[$name]);
    }
    
    public function __toString(): string {
        return json_encode($this->data);
    }
    
    public function __call(string $name, array $arguments) {
        // Handle dynamic method calls
    }
    
    public static function __callStatic(string $name, array $arguments) {
        // Handle dynamic static method calls
    }
}

// 8. Dependency Injection
// --------------------

interface Logger {
    public function log(string $message): void;
}

class FileLogger implements Logger {
    public function log(string $message): void {
        file_put_contents('app.log', $message . PHP_EOL, FILE_APPEND);
    }
}

class UserService {
    private Logger $logger;
    
    public function __construct(Logger $logger) {
        $this->logger = $logger;
    }
    
    public function createUser(array $data): void {
        // Create user
        $this->logger->log('User created: ' . $data['email']);
    }
}

// 9. Namespaces
// -----------

namespace App\Services;

use App\Interfaces\PaymentProcessor;
use App\Models\User;

class PaymentService {
    private PaymentProcessor $processor;
    
    public function __construct(PaymentProcessor $processor) {
        $this->processor = $processor;
    }
    
    public function processPayment(User $user, float $amount): bool {
        return $this->processor->process($amount);
    }
}

// Best Practices:
// 1. Follow SOLID principles
// 2. Use type declarations
// 3. Implement proper encapsulation
// 4. Use meaningful class and method names
// 5. Keep classes focused (Single Responsibility)
// 6. Use dependency injection
// 7. Follow PSR standards
// 8. Use proper error handling
// 9. Document your code
// 10. Use interfaces for contracts
