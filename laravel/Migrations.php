<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreatePostsTable extends Migration
{
    /**
     * Run the migrations.
     */
    public function up()
    {
        Schema::create('posts', function (Blueprint $table) {
            // Primary Key
            $table->id();
            
            // Foreign Keys
            $table->foreignId('user_id')
                  ->constrained()
                  ->onDelete('cascade');
                  
            $table->foreignId('category_id')
                  ->constrained()
                  ->onDelete('cascade');

            // Basic Fields
            $table->string('title');
            $table->string('slug')->unique();
            $table->text('content');
            $table->string('status')->default('draft');
            $table->boolean('is_featured')->default(false);
            
            // Meta Fields
            $table->string('meta_title')->nullable();
            $table->text('meta_description')->nullable();
            $table->string('meta_keywords')->nullable();

            // Timestamps
            $table->timestamp('published_at')->nullable();
            $table->timestamps();
            $table->softDeletes();

            // Indexes
            $table->index(['status', 'published_at']);
            $table->index('is_featured');
        });

        // Create pivot table for posts and tags
        Schema::create('post_tag', function (Blueprint $table) {
            $table->id();
            $table->foreignId('post_id')->constrained()->onDelete('cascade');
            $table->foreignId('tag_id')->constrained()->onDelete('cascade');
            $table->timestamps();

            // Unique constraint
            $table->unique(['post_id', 'tag_id']);
        });
    }

    /**
     * Reverse the migrations.
     */
    public function down()
    {
        Schema::dropIfExists('post_tag');
        Schema::dropIfExists('posts');
    }
}

// Best Practices for Migrations:
// 1. Use meaningful table and column names
// 2. Add proper foreign key constraints
// 3. Include indexes for frequently queried columns
// 4. Use appropriate column types and lengths
// 5. Consider soft deletes when needed
// 6. Add timestamps for tracking
// 7. Make migrations reversible
// 8. Use meaningful constraints
// 9. Consider database performance
// 10. Document complex migrations
