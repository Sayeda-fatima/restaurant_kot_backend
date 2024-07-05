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
        Schema::create('invoices', function (Blueprint $table) {
            $table->id();
            $table->unsignedBigInteger('customer_id');
            $table->foreign('customer_id')->references('id')->on('customers');
            //$table->datetime('created_on');
            $table->string('customer_name');
            //$table->foreignId('product_id')->references('product_id')->on('products');
            //$table->integer('product_quantity');
            //$table->decimal('unit_price',8,2);
            $table->decimal('total_price',8,2);
            $table->string('billing_address');
            $table->set('mode_of_payment', ['bank', 'cash', 'cheque']);
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        Schema::dropIfExists('invoices');
    }
};
