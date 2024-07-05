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
        Schema::create('invoice_details', function (Blueprint $table) {
           // $table->id('invoice_id');
            $table->unsignedBigInteger('invoice_id');
            $table->foreign('invoice_id')->references('id')->on('invoices');
            $table->unsignedBigInteger('product_id')->primary();
            $table->foreign('product_id')->references('id')->on('products');
            $table->integer('quantity');
            $table->decimal('unit_product_price', 8,2);
            $table->foreign('unit_product_price')->references('mrp')->on('products');
            $table->decimal('total_product_price', 8,2);
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        Schema::dropIfExists('invoice_details');
    }
};
