<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration
{
    /**
     * Run the migrations.
     */
    public function up(): void
    {
        Schema::create('customers', function (Blueprint $table) {
            $table->id();
            $table->unsignedBigInteger('organization_id');
            $table->foreign('organization_id')->references('id')->on('organizations');
            $table->string('name');
            $table->string('phone_no')->unique();
            $table->string('category');
            $table->string('billing_address');
            // optional 
            $table->string('billing_province')->nullable();
            $table->string('billing_postal_code')->nullable();
            $table->string('delivery_address')->nullable();
            $table->string('delivery_province')->nullable();
            $table->string('delivery_postal_code')->nullable();
            $table->string('gst_number')->nullable();
            $table->string('billing_term')->nullable();
            $table->string('billing_type')->nullable();
            $table->date('date_of_birth')->nullable();
            $table->enum('whatsapp_alert', ['Y','N'])->nullable();
            $table->boolean('is_deleted')->default(0);
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        Schema::dropIfExists('customers');
    }
};
