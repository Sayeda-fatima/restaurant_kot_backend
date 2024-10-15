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
            $table->string('customer_name');
            $table->string('customer_phone_no')->unique();
            $table->string('customer_category');
            $table->string('customer_billing_address');
            // optional 
            $table->string('customer_billing_province')->nullable();
            $table->string('customer_billing_postal_code')->nullable();
            $table->string('customer_delivery_address')->nullable();
            $table->string('customer_delivery_province')->nullable();
            $table->string('customer_delivery_postal_code')->nullable();
            $table->string('customer_gst_number')->nullable();
            $table->string('customer_billing_term')->nullable();
            $table->string('customer_billing_type')->nullable();
            $table->date('customer_date_of_birth')->nullable();
            $table->enum('whatsapp_alert', ['Y','N'])->nullable();
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
