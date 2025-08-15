// Objects and Classes in JavaScript - Code Examples and Best Practices

// 1. Object Literals
// -----------------
// Basic object literal
const person = {
    firstName: 'John',
    lastName: 'Doe',
    age: 30,
    hobbies: ['reading', 'music'],
    greet() {
        return `Hello, I'm ${this.firstName}`;
    }
};

// Object with computed property names
const propertyName = 'email';
const user = {
    [propertyName]: 'john@example.com'
};


// 2. Object Methods
// ---------------
// Object.keys()
const keys = Object.keys(person);

// Object.values()
const values = Object.values(person);

// Object.entries()
const entries = Object.entries(person);

// Object.freeze() - Make object immutable
const frozenObject = Object.freeze({
    prop: 42
});

// Object.assign() - Merge objects
const merged = Object.assign({}, person, user);


// 3. Classes
// ---------
class Animal {
    constructor(name) {
        this.name = name;
    }

    speak() {
        return `${this.name} makes a sound`;
    }

    // Static method
    static isAnimal(obj) {
        return obj instanceof Animal;
    }
}

// Class inheritance
class Dog extends Animal {
    constructor(name, breed) {
        super(name);
        this.breed = breed;
    }

    speak() {
        return `${this.name} barks`;
    }
}

// Using classes
const dog = new Dog('Rex', 'German Shepherd');
console.log(dog.speak()); // Rex barks


// 4. Getters and Setters
// ---------------------
class Circle {
    constructor(radius) {
        this._radius = radius;
    }

    get radius() {
        return this._radius;
    }

    set radius(value) {
        if (value <= 0) {
            throw new Error('Radius must be positive');
        }
        this._radius = value;
    }

    get area() {
        return Math.PI * this._radius ** 2;
    }
}

const circle = new Circle(5);
console.log(circle.area); // ~78.54


// 5. Private Class Features
// -----------------------
class BankAccount {
    #balance = 0;  // Private field

    constructor(initialBalance) {
        this.#balance = initialBalance;
    }

    get balance() {
        return this.#balance;
    }

    deposit(amount) {
        if (amount > 0) {
            this.#balance += amount;
            return true;
        }
        return false;
    }

    #validateWithdrawal(amount) {  // Private method
        return amount > 0 && amount <= this.#balance;
    }

    withdraw(amount) {
        if (this.#validateWithdrawal(amount)) {
            this.#balance -= amount;
            return true;
        }
        return false;
    }
}


// 6. Factory Functions
// ------------------
function createPerson(firstName, lastName) {
    return {
        firstName,
        lastName,
        fullName() {
            return `${this.firstName} ${this.lastName}`;
        }
    };
}

const person1 = createPerson('John', 'Doe');
const person2 = createPerson('Jane', 'Smith');


// 7. Prototype Methods
// ------------------
function Person(firstName, lastName) {
    this.firstName = firstName;
    this.lastName = lastName;
}

Person.prototype.getFullName = function() {
    return `${this.firstName} ${this.lastName}`;
};

const person3 = new Person('Alice', 'Johnson');


// 8. Object Composition
// -------------------
const hasName = (name) => ({
    getName: () => name
});

const canSpeak = () => ({
    speak: () => 'Speaking...'
});

const canEat = () => ({
    eat: () => 'Eating...'
});

function createCreature(name) {
    return {
        ...hasName(name),
        ...canSpeak(),
        ...canEat()
    };
}

const creature = createCreature('Creature');


// Best Practices:
// 1. Use class syntax for object-oriented programming
// 2. Keep classes focused and single-responsibility
// 3. Use private fields and methods when possible
// 4. Prefer composition over inheritance
// 5. Use getters and setters for computed properties
// 6. Use factory functions when appropriate
// 7. Follow proper encapsulation principles
// 8. Use meaningful and descriptive names
// 9. Document complex class hierarchies
