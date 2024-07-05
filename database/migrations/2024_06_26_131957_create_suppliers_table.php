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
        Schema::create('suppliers', function (Blueprint $table) {
            $table->id();
            $table->string('supplier_name');
            $table->string('supplier_phone_no');
            $table->string('supplier_category');
            $table->string('supplier_billing_address');
            // optional 
            $table->string('supplier_billing_province')->nullable();
            $table->string('supplier_billing_postal_code')->nullable();
            $table->string('supplier_delivery_address')->nullable();
            $table->string('supplier_delivery_province')->nullable();
            $table->string('supplier_delivery_postal_code')->nullable();
            $table->string('supplier_gst_number')->nullable();
            $table->string('supplier_billing_term')->nullable();
            $table->string('supplier_billing_type')->nullable();
            $table->date('supplier_date_of_birth')->nullable();
            $table->enum('supplier_whatsapp_alert', ['Y','N'])->nullable();
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        Schema::dropIfExists('suppliers');
    }
};
