<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;
use Illuminate\Database\Eloquent\Factories\HasFactory;

class Post extends Model
{
    use HasFactory, SoftDeletes;

    /**
     * The attributes that are mass assignable.
     */
    protected $fillable = [
        'title',
        'content',
        'user_id',
        'category_id',
        'status',
        'published_at'
    ];

    /**
     * The attributes that should be cast.
     */
    protected $casts = [
        'published_at' => 'datetime',
        'is_featured' => 'boolean'
    ];

    /**
     * The relationships that should always be loaded.
     */
    protected $with = ['category'];

    /**
     * Get the user that owns the post
     */
    public function user()
    {
        return $this->belongsTo(User::class);
    }

    /**
     * Get the category that owns the post
     */
    public function category()
    {
        return $this->belongsTo(Category::class);
    }

    /**
     * Get the comments for the post
     */
    public function comments()
    {
        return $this->hasMany(Comment::class);
    }

    /**
     * Get the tags for the post
     */
    public function tags()
    {
        return $this->belongsToMany(Tag::class);
    }

    /**
     * Scope a query to only include published posts
     */
    public function scopePublished($query)
    {
        return $query->whereNotNull('published_at')
                    ->where('published_at', '<=', now())
                    ->where('status', 'published');
    }

    /**
     * Scope a query to only include featured posts
     */
    public function scopeFeatured($query)
    {
        return $query->where('is_featured', true);
    }

    /**
     * Get the post's excerpt
     */
    public function getExcerptAttribute()
    {
        return str_limit($this->content, 150);
    }

    /**
     * Set the post's title and slug
     */
    public function setTitleAttribute($value)
    {
        $this->attributes['title'] = $value;
        $this->attributes['slug'] = str_slug($value);
    }

    /**
     * Check if post is published
     */
    public function isPublished()
    {
        return !is_null($this->published_at) && $this->published_at <= now();
    }

    /**
     * Publish the post
     */
    public function publish()
    {
        $this->update([
            'status' => 'published',
            'published_at' => now()
        ]);
    }

    /**
     * Unpublish the post
     */
    public function unpublish()
    {
        $this->update([
            'status' => 'draft',
            'published_at' => null
        ]);
    }

    /**
     * Get related posts
     */
    public function getRelatedPosts($limit = 3)
    {
        return static::published()
            ->where('id', '!=', $this->id)
            ->whereCategoryId($this->category_id)
            ->limit($limit)
            ->get();
    }

    /**
     * Boot the model
     */
    protected static function boot()
    {
        parent::boot();

        // Auto-set user_id on creation
        static::creating(function ($post) {
            if (!$post->user_id && auth()->check()) {
                $post->user_id = auth()->id();
            }
        });

        // Clean up relationships when deleting
        static::deleting(function ($post) {
            $post->comments()->delete();
            $post->tags()->detach();
        });
    }
}

// Best Practices:
// 1. Use proper namespacing
// 2. Implement proper relationships
// 3. Use model events when needed
// 4. Implement scopes for common queries
// 5. Use mass assignment protection
// 6. Use proper type casting
// 7. Keep models clean and focused
// 8. Use accessors and mutators properly
// 9. Implement soft deletes when needed
// 10. Document methods with PHPDoc
