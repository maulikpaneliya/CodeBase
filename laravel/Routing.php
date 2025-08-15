<?php

// Basic Routes
Route::get('/', function () {
    return view('welcome');
});

// Route with Parameters
Route::get('/user/{id}', function ($id) {
    return 'User '.$id;
});

// Route with Optional Parameters
Route::get('/user/{name?}', function ($name = 'John') {
    return 'Hello '.$name;
});

// Named Routes
Route::get('/profile', function () {
    return 'Profile';
})->name('profile');

// Route Groups - Middleware
Route::middleware(['auth'])->group(function () {
    Route::get('/dashboard', function () {
        return view('dashboard');
    });
    
    Route::get('/settings', function () {
        return view('settings');
    });
});

// Route Groups - Prefix
Route::prefix('admin')->group(function () {
    Route::get('/users', function () {
        return 'Admin Users';
    });
    
    Route::get('/settings', function () {
        return 'Admin Settings';
    });
});

// Route Groups - Name Prefix
Route::name('admin.')->group(function () {
    Route::get('/admin/users', function () {
        return 'Admin Users';
    })->name('users');
});

// Resource Routes
Route::resource('posts', PostController::class);

// API Routes
Route::prefix('api')->group(function () {
    Route::get('/users', [UserController::class, 'index']);
    Route::post('/users', [UserController::class, 'store']);
    Route::get('/users/{id}', [UserController::class, 'show']);
    Route::put('/users/{id}', [UserController::class, 'update']);
    Route::delete('/users/{id}', [UserController::class, 'destroy']);
});

// Route with Multiple HTTP Verbs
Route::match(['get', 'post'], '/submit', function () {
    // Handle both GET and POST
});

// Fallback Route
Route::fallback(function () {
    return 'Page Not Found';
});

// Route with Middleware and Parameters
Route::middleware(['auth', 'role:admin'])->group(function () {
    Route::get('/admin/post/{id}', function ($id) {
        return 'Admin Post '.$id;
    })->where('id', '[0-9]+');
});

// Route with Form Submission
Route::post('/submit-form', function (Request $request) {
    return $request->all();
})->middleware('csrf');

// Route with Response Types
Route::get('/data', function () {
    return response()->json([
        'name' => 'John Doe',
        'age' => 30
    ]);
});

// Route with File Download
Route::get('/download', function () {
    return response()->download('path/to/file.pdf');
});

// Route with Custom Status Code
Route::get('/error', function () {
    return response('Error', 500);
});

// Best Practices:
// 1. Use meaningful route names
// 2. Group related routes
// 3. Use resource routes for CRUD operations
// 4. Validate parameters using where() method
// 5. Use middleware for authentication and authorization
// 6. Keep routes organized by feature or module
// 7. Use route model binding when possible
// 8. Handle fallback routes for 404 errors
