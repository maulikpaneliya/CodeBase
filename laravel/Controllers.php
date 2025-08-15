<?php

namespace App\Http\Controllers;

use App\Models\User;
use Illuminate\Http\Request;
use App\Http\Requests\UserRequest;
use App\Http\Resources\UserResource;

class UserController extends Controller
{
    /**
     * Constructor with middleware
     */
    public function __construct()
    {
        $this->middleware('auth')->except(['index', 'show']);
    }

    /**
     * Display a listing of users
     */
    public function index()
    {
        $users = User::paginate(10);
        return UserResource::collection($users);
    }

    /**
     * Show form for creating a new user
     */
    public function create()
    {
        return view('users.create');
    }

    /**
     * Store a new user
     */
    public function store(UserRequest $request)
    {
        try {
            $user = User::create($request->validated());
            
            return response()->json([
                'message' => 'User created successfully',
                'user' => new UserResource($user)
            ], 201);
        } catch (\Exception $e) {
            return response()->json([
                'message' => 'Error creating user',
                'error' => $e->getMessage()
            ], 500);
        }
    }

    /**
     * Display user details
     */
    public function show(User $user)
    {
        return new UserResource($user);
    }

    /**
     * Show form for editing user
     */
    public function edit(User $user)
    {
        $this->authorize('update', $user);
        return view('users.edit', compact('user'));
    }

    /**
     * Update user details
     */
    public function update(UserRequest $request, User $user)
    {
        try {
            $this->authorize('update', $user);
            
            $user->update($request->validated());
            
            return response()->json([
                'message' => 'User updated successfully',
                'user' => new UserResource($user)
            ]);
        } catch (\Exception $e) {
            return response()->json([
                'message' => 'Error updating user',
                'error' => $e->getMessage()
            ], 500);
        }
    }

    /**
     * Delete user
     */
    public function destroy(User $user)
    {
        try {
            $this->authorize('delete', $user);
            
            $user->delete();
            
            return response()->json([
                'message' => 'User deleted successfully'
            ]);
        } catch (\Exception $e) {
            return response()->json([
                'message' => 'Error deleting user',
                'error' => $e->getMessage()
            ], 500);
        }
    }

    /**
     * Custom method: Toggle user active status
     */
    public function toggleActive(User $user)
    {
        $this->authorize('update', $user);
        
        $user->is_active = !$user->is_active;
        $user->save();
        
        return response()->json([
            'message' => 'User status updated successfully',
            'is_active' => $user->is_active
        ]);
    }

    /**
     * Custom method: Get user statistics
     */
    public function getStatistics(User $user)
    {
        return response()->json([
            'posts_count' => $user->posts()->count(),
            'comments_count' => $user->comments()->count(),
            'last_login' => $user->last_login
        ]);
    }
}

// Best Practices:
// 1. Use Resource Controllers for CRUD operations
// 2. Implement Form Requests for validation
// 3. Use API Resources for response transformation
// 4. Implement proper error handling
// 5. Use authorization with policies
// 6. Keep controllers focused (Single Responsibility)
// 7. Use dependency injection
// 8. Implement proper status codes
// 9. Use route model binding
// 10. Document methods with PHPDoc
