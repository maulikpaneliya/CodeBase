// jQuery Quick Reference
// This file contains commonly used jQuery patterns and snippets

// Document Ready
$(document).ready(function() {
    // Code here runs when DOM is ready
});

// Short version
$(function() {
    // Code here
});

// Selectors
// Basic selectors
$('#myId')              // Select by ID
$('.myClass')          // Select by class
$('div')               // Select by tag
$('div.myClass')       // Combine selectors
$('div > p')           // Child selector
$('div:first')         // First element
$('div:last')          // Last element
$('div:eq(2)')         // Element at index
$('div:gt(2)')         // Elements after index
$('div:lt(2)')         // Elements before index
$('input:checked')     // Checked inputs
$('a[href^="https"]')  // Attribute starts with
$('a[href$=".pdf"]')   // Attribute ends with
$('a[href*="example"]')// Attribute contains

// Events
// Click events
$('#myButton').click(function(e) {
    e.preventDefault();
    // Handle click
});

// Other common events
$('#element').on('mouseover', function() {});
$('#element').on('mouseout', function() {});
$('#element').hover(
    function() { /* mouseenter */ },
    function() { /* mouseleave */ }
);
$('#myForm').submit(function(e) {
    e.preventDefault();
});
$('#myInput').change(function() {});
$('#myInput').focus(function() {});
$('#myInput').blur(function() {});

// DOM Manipulation
// Adding/Removing Content
$('#element').html('<p>New content</p>');
$('#element').text('New text');
$('#element').append('<p>Added at end</p>');
$('#element').prepend('<p>Added at beginning</p>');
$('#element').after('<p>Added after</p>');
$('#element').before('<p>Added before</p>');
$('#element').remove();
$('#element').empty();

// CSS & Classes
$('#element').addClass('newClass');
$('#element').removeClass('oldClass');
$('#element').toggleClass('someClass');
$('#element').hasClass('someClass');
$('#element').css('color', 'red');
$('#element').css({
    color: 'red',
    backgroundColor: 'blue',
    fontSize: '16px'
});

// Attributes
$('#element').attr('src', 'image.jpg');
$('#element').removeAttr('src');
$('#element').prop('checked', true);

// Dimensions & Position
$('#element').width();
$('#element').height();
$('#element').innerWidth();
$('#element').innerHeight();
$('#element').outerWidth();
$('#element').outerHeight();
$('#element').position();
$('#element').offset();

// Traversing
$('#element').parent();
$('#element').parents();
$('#element').children();
$('#element').siblings();
$('#element').next();
$('#element').prev();
$('#element').closest('.someClass');
$('#element').find('.someClass');

// Effects & Animation
// Show/Hide
$('#element').show();
$('#element').hide();
$('#element').toggle();

// Fade
$('#element').fadeIn(1000);
$('#element').fadeOut(1000);
$('#element').fadeToggle(1000);
$('#element').fadeTo(1000, 0.5);

// Slide
$('#element').slideDown(1000);
$('#element').slideUp(1000);
$('#element').slideToggle(1000);

// Custom Animation
$('#element').animate({
    opacity: 0.5,
    left: '+=50',
    height: 'toggle'
}, 1000);

// Ajax
// GET Request
$.get('url', function(data) {
    console.log(data);
});

// POST Request
$.post('url', {
    name: 'John',
    age: 30
}, function(data) {
    console.log(data);
});

// Ajax with Options
$.ajax({
    url: 'api/endpoint',
    method: 'POST',
    data: {
        name: 'John',
        age: 30
    },
    dataType: 'json',
    success: function(response) {
        console.log(response);
    },
    error: function(xhr, status, error) {
        console.error(error);
    }
});

// Utilities
$.each(array, function(index, value) {});
$.map(array, function(value, index) {});
$.grep(array, function(value, index) {});
$.extend({}, obj1, obj2);
$.trim(string);
$.isArray(obj);
$.isFunction(obj);
$.parseJSON(string);

// Event Delegation
$(document).on('click', '.dynamicElement', function() {
    // Handles elements added after page load
});

// Forms
$('#myForm').serialize();
$('#myForm').serializeArray();

// Chaining
$('#element')
    .addClass('highlight')
    .fadeIn(1000)
    .delay(2000)
    .fadeOut(1000);
