// jQuery Selectors Complete Guide

// 1. Basic Selectors
// ----------------------------------------

// Element Selector
$("p")                     // Selects all <p> elements
$("div")                   // Selects all <div> elements
$("span")                  // Selects all <span> elements

// ID Selector
$("#myId")                 // Selects element with id="myId"
$("#header")               // Selects element with id="header"

// Class Selector
$(".myClass")             // Selects all elements with class="myClass"
$(".highlight")           // Selects all elements with class="highlight"

// Multiple Selectors
$("p, div, span")         // Selects all <p>, <div>, and <span> elements
$(".class1, .class2")     // Selects elements with class1 or class2
$("#id1, #id2")           // Selects elements with id1 or id2

// Universal Selector
$("*")                    // Selects all elements

// 2. Hierarchy Selectors
// ----------------------------------------

// Descendant Selector
$("div span")             // Selects all <span> elements inside <div>
$("ul li")                // Selects all <li> elements inside <ul>

// Child Selector
$("parent > child")       // Selects direct children only
$("ul > li")              // Selects <li> elements that are direct children of <ul>

// Next Adjacent Selector
$("label + input")        // Selects input immediately after label
$("h2 + p")               // Selects <p> immediately after <h2>

// Next Siblings Selector
$("h2 ~ p")               // Selects all <p> elements after <h2>

// 3. Basic Filters
// ----------------------------------------

// First Selector
$("p:first")              // Selects first <p> element
$(".item:first")          // Selects first element with class="item"

// Last Selector
$("p:last")               // Selects last <p> element
$(".item:last")           // Selects last element with class="item"

// Even/Odd Selector
$("tr:even")              // Selects even rows (zero-based)
$("tr:odd")               // Selects odd rows (zero-based)

// Eq Selector
$("tr:eq(0)")            // Selects first row
$("tr:eq(-1)")           // Selects last row

// Greater Than/Less Than
$("tr:gt(3)")            // Selects rows after index 3
$("tr:lt(3)")            // Selects rows before index 3

// 4. Content Filters
// ----------------------------------------

// Contains Selector
$("p:contains('Hello')")  // Selects <p> elements containing "Hello"
$("p:contains('World')")  // Selects <p> elements containing "World"

// Empty Selector
$("p:empty")              // Selects <p> elements with no children

// Has Selector
$("p:has(span)")         // Selects <p> elements containing <span>
$("ul:has(li)")          // Selects <ul> elements containing <li>

// Parent Selector
$("p:parent")            // Selects <p> elements that have children

// 5. Visibility Filters
// ----------------------------------------

// Hidden Selector
$(":hidden")             // Selects all hidden elements
$("p:hidden")            // Selects hidden <p> elements

// Visible Selector
$(":visible")            // Selects all visible elements
$("p:visible")           // Selects visible <p> elements

// 6. Attribute Selectors
// ----------------------------------------

// Has Attribute
$("[href]")              // Selects elements with href attribute
$("[target]")            // Selects elements with target attribute

// Attribute Equals
$("[href='url']")        // Selects elements with href="url"
$("[type='text']")       // Selects elements with type="text"

// Attribute Contains
$("[href*='google']")    // Selects elements with href containing "google"
$("[class*='btn']")      // Selects elements with class containing "btn"

// Attribute Starts With
$("[href^='http']")      // Selects elements with href starting with "http"
$("[class^='col-']")     // Selects elements with class starting with "col-"

// Attribute Ends With
$("[href$='.pdf']")      // Selects elements with href ending in ".pdf"
$("[class$='-primary']") // Selects elements with class ending in "-primary"

// Attribute Not Equal
$("[href!='url']")       // Selects elements with href not equal to "url"

// Multiple Attributes
$("[type='text'][name='username']")  // AND condition

// 7. Child Filters
// ----------------------------------------

// First Child
$("p:first-child")       // Selects <p> elements that are first child
$("li:first-child")      // Selects <li> elements that are first child

// Last Child
$("p:last-child")        // Selects <p> elements that are last child
$("li:last-child")       // Selects <li> elements that are last child

// Nth Child
$("p:nth-child(2)")      // Selects <p> elements that are second child
$("tr:nth-child(odd)")   // Selects odd rows
$("tr:nth-child(even)")  // Selects even rows

// Only Child
$("p:only-child")        // Selects <p> elements that are only child

// 8. Forms Filters
// ----------------------------------------

// Input Types
$(":text")               // Selects input type="text"
$(":password")           // Selects input type="password"
$(":radio")              // Selects input type="radio"
$(":checkbox")           // Selects input type="checkbox"
$(":submit")             // Selects input type="submit"
$(":reset")              // Selects input type="reset"
$(":button")             // Selects input type="button" and <button>
$(":file")               // Selects input type="file"

// Input States
$(":enabled")            // Selects enabled form elements
$(":disabled")           // Selects disabled form elements
$(":checked")            // Selects checked radio/checkbox
$(":selected")           // Selects selected options

// 9. Advanced Selectors
// ----------------------------------------

// Not Selector
$("p:not(.exclude)")     // Selects <p> elements without class "exclude"
$("input:not(:checked)") // Selects unchecked inputs

// Multiple Conditions
$("input[type='text'][required]")  // Text inputs that are required

// Complex Selectors
$("form input:not([type='submit']):visible")  // Visible form inputs except submit

// 10. Performance Tips
// ----------------------------------------

// ID is fastest
$("#myId")               // Fastest selector

// Cache jQuery objects
var $myElement = $("#myElement");  // Cache for reuse

// Be specific with selectors
$("#myForm input.special")  // More specific than $(".special")

// Avoid universal selectors
$(".myClass > *")  // Slow
$(".myClass").children()  // Better

// Use find() for descendants
$("#parent").find(".child")  // Better than $("#parent .child")

// 11. Examples and Use Cases
// ----------------------------------------

// Example 1: Table Row Selection
$("table tr:even")        // Select even rows
    .addClass("even")     // Add class
    .end()               // Back to original selection
    .filter(":odd")      // Select odd rows
    .addClass("odd");    // Add class

// Example 2: Form Validation
$("form input[required]:not(:valid)")  // Find invalid required fields

// Example 3: Dynamic Content
$("#content").on("click", "a:not(.processed)", function() {
    // Handle dynamically added links
});

// Example 4: Complex Navigation
$("nav > ul > li:has(ul)")  // Find menu items with submenus

// 12. Common Patterns
// ----------------------------------------

// Pattern 1: Toggle Buttons
$(".toggle-btn").click(function() {
    $(this).toggleClass("active")
           .siblings()
           .removeClass("active");
});

// Pattern 2: Form Processing
$("#myForm :input:not(:submit)")  // Select all form inputs except submit

// Pattern 3: Dynamic Lists
$("ul li:has(> ul)")  // Find list items with sublists

// Pattern 4: Table Processing
$("table tbody tr")
    .filter(":odd")
    .addClass("odd")
    .end()
    .filter(":even")
    .addClass("even");

// 13. Best Practices
// ----------------------------------------

// 1. Cache selectors
var $form = $("#myForm");
var $inputs = $form.find("input");

// 2. Use chaining
$("div")
    .addClass("active")
    .find("p")
    .addClass("highlight");

// 3. Be specific
$("form#myForm input.special")  // More specific

// 4. Use proper event delegation
$(document).on("click", ".dynamic-element", function() {});

// Example HTML Reference
/*
<!DOCTYPE html>
<html>
<head>
    <script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>
</head>
<body>
    <div id="container">
        <p class="highlight">Highlighted paragraph</p>
        <p>Regular paragraph</p>
        
        <ul>
            <li>First item
                <ul>
                    <li>Sub item</li>
                </ul>
            </li>
            <li>Second item</li>
        </ul>
        
        <form id="myForm">
            <input type="text" required>
            <input type="checkbox" checked>
            <select>
                <option selected>Option 1</option>
                <option>Option 2</option>
            </select>
        </form>
        
        <table>
            <tr><td>Row 1</td></tr>
            <tr><td>Row 2</td></tr>
            <tr><td>Row 3</td></tr>
        </table>
    </div>
</body>
</html>
*/
