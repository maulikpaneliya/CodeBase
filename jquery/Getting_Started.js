// jQuery Getting Started Guide

// 1. Including jQuery
/*
Method 1: CDN (Recommended)
<script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>

Method 2: Local File
<script src="path/to/jquery.min.js"></script>
*/

// 2. Document Ready
// Wait for DOM to be fully loaded
$(document).ready(function() {
    // Your code here
});

// Shorter version
$(function() {
    // Your code here
});

// 3. jQuery Syntax Basics
// Basic syntax: $(selector).action()

// Select by ID
$("#myElement").hide();

// Select by class
$(".myClass").show();

// Select by element type
$("p").text("Hello World!");

// Multiple selectors
$("h1, h2, h3").css("color", "blue");

// 4. Noconflict Mode
// Avoid conflicts with other libraries
var $j = jQuery.noConflict();
$j(document).ready(function() {
    $j("#myElement").hide();
});

// Pass jQuery as $ to ready handler
jQuery.noConflict();
jQuery(function($) {
    // Use $ here
    $("#myElement").show();
});

// 5. Basic Methods

// Getters and Setters
// Get content
var text = $("#myElement").text();
var html = $("#myElement").html();
var value = $("#myInput").val();

// Set content
$("#myElement").text("New text");
$("#myElement").html("<strong>New HTML</strong>");
$("#myInput").val("New value");

// 6. Chaining
$("#myElement")
    .hide()
    .css("color", "red")
    .fadeIn()
    .addClass("highlight");

// 7. Basic Examples

// Example 1: Simple Click Handler
$(function() {
    $("#clickMe").click(function() {
        alert("Button clicked!");
    });
});

// Example 2: Toggle Element
$(function() {
    $("#toggleButton").click(function() {
        $("#targetElement").toggle();
    });
});

// Example 3: Change Content
$(function() {
    $("#changeContent").click(function() {
        $(".content").text("Content changed!");
    });
});

// 8. jQuery Objects vs DOM Elements

// jQuery object
var $element = $("#myElement");

// DOM element
var element = $("#myElement")[0];
// or
var element = $("#myElement").get(0);

// Converting DOM to jQuery
var $jqElement = $(element);

// 9. Multiple Elements Handling

// Select multiple elements
$("p").each(function(index) {
    $(this).addClass("paragraph-" + index);
});

// Filter elements
$("p").filter(".special").css("color", "red");

// Find elements
$("#container").find(".item").hide();

// 10. Error Handling

// Try-Catch with jQuery
try {
    $("#nonexistent").doSomething();
} catch (error) {
    console.error("jQuery error:", error);
}

// Check if element exists
if ($("#myElement").length > 0) {
    // Element exists, safe to proceed
    $("#myElement").show();
}

// 11. jQuery Version Check
if (jQuery.fn.jquery) {
    console.log("jQuery version:", jQuery.fn.jquery);
}

// 12. Common Patterns

// Pattern 1: Plugin Template
(function($) {
    $.fn.myPlugin = function(options) {
        // Merge default and user options
        var settings = $.extend({
            color: "red",
            background: "white"
        }, options);

        return this.each(function() {
            $(this).css({
                color: settings.color,
                backgroundColor: settings.background
            });
        });
    };
}(jQuery));

// Pattern 2: Module Pattern
var myModule = (function($) {
    // Private variables and functions
    var privateVar = "I'm private";
    
    function privateFunction() {
        return "Private function";
    }

    // Public API
    return {
        publicFunction: function() {
            $("#element").text(privateFunction());
        },
        init: function() {
            this.publicFunction();
        }
    };
}(jQuery));

// Pattern 3: Event Delegation
$(document).on("click", ".dynamic-element", function() {
    // Handles elements added after page load
});

// 13. Best Practices

// Cache jQuery selections
var $element = $("#myElement");
$element.hide();
$element.show();

// Use specific selectors
$("#myId").hide(); // Better than $("[id=myId]")

// Chain when possible
$("#myElement")
    .addClass("active")
    .removeClass("inactive")
    .show();

// Use proper event delegation
$("#container").on("click", ".item", function() {
    // Better than $(".item").click()
});

// 14. Example HTML Structure
/*
<!DOCTYPE html>
<html>
<head>
    <script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>
    <script src="getting_started.js"></script>
</head>
<body>
    <div id="container">
        <button id="clickMe">Click Me</button>
        <button id="toggleButton">Toggle</button>
        <button id="changeContent">Change Content</button>
        
        <div id="targetElement">
            Toggle this content
        </div>
        
        <div class="content">
            Original content
        </div>
        
        <p class="special">Special paragraph</p>
        <p>Regular paragraph</p>
    </div>
</body>
</html>
*/

// 15. Common Mistakes and Solutions

// Mistake 1: Not checking if element exists
// Bad
$("#nonexistent").show();

// Good
var $element = $("#nonexistent");
if ($element.length) {
    $element.show();
}

// Mistake 2: Excessive DOM queries
// Bad
$("#element").hide();
$("#element").show();

// Good
var $element = $("#element");
$element.hide();
$element.show();

// Mistake 3: Not using event delegation
// Bad
$(".dynamic-item").click(function() {});

// Good
$(document).on("click", ".dynamic-item", function() {});

// 16. Debugging Tips

// Log jQuery object
console.log($('#element'));

// Check if jQuery is loaded
if (typeof jQuery != 'undefined') {
    console.log('jQuery is loaded');
}

// Debug events
$("#element").on("click", function(e) {
    console.log("Event:", e);
    console.log("Target:", e.target);
    console.log("Current Target:", e.currentTarget);
});
