// PHP Database Operations with PDO and MySQL

<?php

// 1. Database Connection
// -------------------

class Database {
    private static ?PDO $instance = null;
    private static array $config = [
        'host' => 'localhost',
        'dbname' => 'myapp',
        'username' => 'root',
        'password' => '',
        'charset' => 'utf8mb4'
    ];

    private function __construct() {}

    public static function getInstance(): PDO {
        if (self::$instance === null) {
            try {
                $dsn = "mysql:host=" . self::$config['host'] . 
                       ";dbname=" . self::$config['dbname'] . 
                       ";charset=" . self::$config['charset'];
                
                self::$instance = new PDO(
                    $dsn,
                    self::$config['username'],
                    self::$config['password'],
                    [
                        PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION,
                        PDO::ATTR_DEFAULT_FETCH_MODE => PDO::FETCH_ASSOC,
                        PDO::ATTR_EMULATE_PREPARES => false
                    ]
                );
            } catch (PDOException $e) {
                throw new Exception("Connection failed: " . $e->getMessage());
            }
        }
        return self::$instance;
    }
}

// 2. Basic CRUD Operations
// ---------------------

class UserRepository {
    private PDO $db;
    
    public function __construct() {
        $this->db = Database::getInstance();
    }
    
    // Create
    public function create(array $data): int {
        try {
            $sql = "INSERT INTO users (name, email, password_hash, created_at) 
                    VALUES (:name, :email, :password_hash, NOW())";
            
            $stmt = $this->db->prepare($sql);
            $stmt->execute([
                'name' => $data['name'],
                'email' => $data['email'],
                'password_hash' => password_hash($data['password'], PASSWORD_DEFAULT)
            ]);
            
            return $this->db->lastInsertId();
        } catch (PDOException $e) {
            throw new Exception("Create failed: " . $e->getMessage());
        }
    }
    
    // Read
    public function find(int $id): ?array {
        try {
            $sql = "SELECT * FROM users WHERE id = :id";
            $stmt = $this->db->prepare($sql);
            $stmt->execute(['id' => $id]);
            
            $result = $stmt->fetch();
            return $result ?: null;
        } catch (PDOException $e) {
            throw new Exception("Find failed: " . $e->getMessage());
        }
    }
    
    // Update
    public function update(int $id, array $data): bool {
        try {
            $sql = "UPDATE users 
                    SET name = :name, 
                        email = :email, 
                        updated_at = NOW() 
                    WHERE id = :id";
            
            $stmt = $this->db->prepare($sql);
            return $stmt->execute([
                'id' => $id,
                'name' => $data['name'],
                'email' => $data['email']
            ]);
        } catch (PDOException $e) {
            throw new Exception("Update failed: " . $e->getMessage());
        }
    }
    
    // Delete
    public function delete(int $id): bool {
        try {
            $sql = "DELETE FROM users WHERE id = :id";
            $stmt = $this->db->prepare($sql);
            return $stmt->execute(['id' => $id]);
        } catch (PDOException $e) {
            throw new Exception("Delete failed: " . $e->getMessage());
        }
    }
}

// 3. Advanced Queries
// ----------------

class QueryBuilder {
    private PDO $db;
    private string $table;
    private array $where = [];
    private array $params = [];
    private ?string $orderBy = null;
    private ?int $limit = null;
    private ?int $offset = null;
    
    public function __construct(string $table) {
        $this->db = Database::getInstance();
        $this->table = $table;
    }
    
    public function where(string $column, $value, string $operator = '='): self {
        $this->where[] = "$column $operator :$column";
        $this->params[$column] = $value;
        return $this;
    }
    
    public function orderBy(string $column, string $direction = 'ASC'): self {
        $this->orderBy = "ORDER BY $column $direction";
        return $this;
    }
    
    public function limit(int $limit, ?int $offset = null): self {
        $this->limit = $limit;
        $this->offset = $offset;
        return $this;
    }
    
    public function get(): array {
        $sql = "SELECT * FROM {$this->table}";
        
        if (!empty($this->where)) {
            $sql .= " WHERE " . implode(' AND ', $this->where);
        }
        
        if ($this->orderBy) {
            $sql .= " {$this->orderBy}";
        }
        
        if ($this->limit) {
            $sql .= " LIMIT {$this->limit}";
            if ($this->offset) {
                $sql .= " OFFSET {$this->offset}";
            }
        }
        
        $stmt = $this->db->prepare($sql);
        $stmt->execute($this->params);
        return $stmt->fetchAll();
    }
}

// 4. Transactions
// ------------

class OrderProcessor {
    private PDO $db;
    
    public function __construct() {
        $this->db = Database::getInstance();
    }
    
    public function processOrder(array $orderData): bool {
        try {
            $this->db->beginTransaction();
            
            // Create order
            $orderStmt = $this->db->prepare(
                "INSERT INTO orders (user_id, total_amount, status) 
                 VALUES (:user_id, :total_amount, 'pending')"
            );
            $orderStmt->execute([
                'user_id' => $orderData['user_id'],
                'total_amount' => $orderData['total_amount']
            ]);
            $orderId = $this->db->lastInsertId();
            
            // Create order items
            $itemStmt = $this->db->prepare(
                "INSERT INTO order_items (order_id, product_id, quantity, price) 
                 VALUES (:order_id, :product_id, :quantity, :price)"
            );
            
            foreach ($orderData['items'] as $item) {
                $itemStmt->execute([
                    'order_id' => $orderId,
                    'product_id' => $item['product_id'],
                    'quantity' => $item['quantity'],
                    'price' => $item['price']
                ]);
                
                // Update product inventory
                $this->db->prepare(
                    "UPDATE products 
                     SET stock_quantity = stock_quantity - :quantity 
                     WHERE id = :product_id"
                )->execute([
                    'quantity' => $item['quantity'],
                    'product_id' => $item['product_id']
                ]);
            }
            
            $this->db->commit();
            return true;
        } catch (Exception $e) {
            $this->db->rollBack();
            throw new Exception("Order processing failed: " . $e->getMessage());
        }
    }
}

// Best Practices:
// 1. Always use prepared statements
// 2. Implement proper error handling
// 3. Use transactions for related operations
// 4. Follow the repository pattern
// 5. Use connection pooling
// 6. Implement proper logging
// 7. Use meaningful table and column names
// 8. Follow database normalization rules
// 9. Implement proper indexing
// 10. Handle connection errors gracefully
