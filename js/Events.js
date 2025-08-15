// JavaScript Events and DOM Manipulation

// 1. Event Handling
// --------------

// Basic event listener
document.getElementById('myButton').addEventListener('click', function(event) {
    console.log('Button clicked!', event);
});

// Event delegation
document.getElementById('parent').addEventListener('click', function(event) {
    if (event.target.matches('.child')) {
        console.log('Child element clicked!');
    }
});

// Remove event listener
const handler = (event) => console.log(event);
element.addEventListener('click', handler);
element.removeEventListener('click', handler);

// 2. Event Types
// -----------

// Mouse events
element.addEventListener('click', (e) => {});
element.addEventListener('dblclick', (e) => {});
element.addEventListener('mousedown', (e) => {});
element.addEventListener('mouseup', (e) => {});
element.addEventListener('mousemove', (e) => {});
element.addEventListener('mouseover', (e) => {});
element.addEventListener('mouseout', (e) => {});
element.addEventListener('mouseenter', (e) => {});
element.addEventListener('mouseleave', (e) => {});

// Keyboard events
document.addEventListener('keydown', (e) => {
    console.log('Key pressed:', e.key);
});
document.addEventListener('keyup', (e) => {});
document.addEventListener('keypress', (e) => {});

// Form events
form.addEventListener('submit', (e) => {
    e.preventDefault();
    console.log('Form submitted');
});
input.addEventListener('focus', (e) => {});
input.addEventListener('blur', (e) => {});
input.addEventListener('change', (e) => {});
input.addEventListener('input', (e) => {});

// Document/Window events
window.addEventListener('load', (e) => {});
window.addEventListener('DOMContentLoaded', (e) => {});
window.addEventListener('resize', (e) => {});
window.addEventListener('scroll', (e) => {});

// 3. Event Object Properties
// ----------------------

function handleEvent(event) {
    // Basic properties
    console.log(event.type);          // Type of event
    console.log(event.target);        // Element that triggered the event
    console.log(event.currentTarget); // Element handling the event
    console.log(event.timeStamp);     // Time event occurred
    
    // Mouse event properties
    console.log(event.clientX, event.clientY);  // Coordinates relative to viewport
    console.log(event.pageX, event.pageY);      // Coordinates relative to document
    console.log(event.screenX, event.screenY);  // Coordinates relative to screen
    
    // Keyboard event properties
    console.log(event.key);           // Key pressed
    console.log(event.keyCode);       // Key code (deprecated)
    console.log(event.code);          // Physical key code
    console.log(event.altKey);        // Alt key pressed
    console.log(event.ctrlKey);       // Ctrl key pressed
    console.log(event.shiftKey);      // Shift key pressed
}

// 4. Custom Events
// -------------

// Creating custom event
const customEvent = new CustomEvent('myEvent', {
    detail: {
        message: 'Hello from custom event'
    },
    bubbles: true,
    cancelable: true
});

// Dispatching custom event
element.dispatchEvent(customEvent);

// Listening for custom event
element.addEventListener('myEvent', (e) => {
    console.log(e.detail.message);
});

// 5. Event Propagation
// -----------------

// Event bubbling
child.addEventListener('click', (e) => {
    console.log('Child clicked');
    // e.stopPropagation(); // Stop event bubbling
});

parent.addEventListener('click', (e) => {
    console.log('Parent clicked');
});

// Event capturing
parent.addEventListener('click', (e) => {
    console.log('Capturing phase');
}, true);

// 6. Event Delegation Pattern
// -----------------------

document.getElementById('todoList').addEventListener('click', (e) => {
    if (e.target.matches('.delete-btn')) {
        e.target.closest('li').remove();
    } else if (e.target.matches('.edit-btn')) {
        // Handle edit
    }
});

// 7. Form Event Handling
// -------------------

const form = document.getElementById('myForm');

form.addEventListener('submit', (e) => {
    e.preventDefault();
    
    const formData = new FormData(form);
    const data = Object.fromEntries(formData.entries());
    
    // Form validation
    if (validateForm(data)) {
        submitForm(data);
    }
});

function validateForm(data) {
    const errors = [];
    
    if (!data.email) {
        errors.push('Email is required');
    }
    
    if (data.password.length < 6) {
        errors.push('Password must be at least 6 characters');
    }
    
    return errors.length === 0;
}

// 8. Throttling and Debouncing
// -------------------------

// Debounce function
function debounce(func, delay) {
    let timeoutId;
    
    return function (...args) {
        clearTimeout(timeoutId);
        
        timeoutId = setTimeout(() => {
            func.apply(this, args);
        }, delay);
    };
}

// Throttle function
function throttle(func, limit) {
    let inThrottle;
    
    return function (...args) {
        if (!inThrottle) {
            func.apply(this, args);
            inThrottle = true;
            
            setTimeout(() => {
                inThrottle = false;
            }, limit);
        }
    };
}

// Usage
const handleScroll = throttle(() => {
    console.log('Scrolled');
}, 1000);

window.addEventListener('scroll', handleScroll);

// 9. Mobile Events
// -------------

element.addEventListener('touchstart', (e) => {});
element.addEventListener('touchmove', (e) => {});
element.addEventListener('touchend', (e) => {});
element.addEventListener('touchcancel', (e) => {});

// 10. Intersection Observer
// ---------------------

const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
        if (entry.isIntersecting) {
            entry.target.classList.add('visible');
        }
    });
});

document.querySelectorAll('.lazy-load').forEach(el => observer.observe(el));

// Best Practices:
// 1. Use event delegation for dynamic elements
// 2. Always remove event listeners when no longer needed
// 3. Implement proper error handling
// 4. Use appropriate event types
// 5. Implement throttling/debouncing for performance
// 6. Consider touch events for mobile
// 7. Use custom events for component communication
// 8. Handle event propagation appropriately
// 9. Implement proper form validation
// 10. Use modern APIs like IntersectionObserver when appropriate
