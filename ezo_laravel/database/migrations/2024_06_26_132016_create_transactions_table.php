<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;
use App\Models\Customer;
use App\Models\Product;

return new class extends Migration
{
    /**
     * Run the migrations.
     */
    public function up(): void
    {
        Schema::create('transactions', function (Blueprint $table) {
            $table->id();
            $table->unsignedBigInteger('organization_id');
            $table->foreign('organization_id')->references('id')->on('organizations');  
            $table->string('name');   //pick from customers table or supplier table 
            $table->set('type', ['customer', 'supplier']);
            $table->unsignedBigInteger('customer_id');
            $table->foreign('customer_id')->references('id')->on('customers')->nullable();  //customers table
            $table->unsignedBigInteger('supplier_id');
            $table->foreign('supplier_id')->references('id')->on('suppliers')->nullable();         // supplier table    
            $table->unsignedBigInteger('product_id');
            $table->foreign('product_id')->references('id')->on('products');   //products table
            $table->string('product_name');    //product table
            $table->integer('product_quantity');
            $table->decimal('product_price',8,2);   //product table
            $table->decimal('total_price',8,2);
            $table->set('mode_of_payment', ['Cash', 'Cheque', 'Bank']);
            //$table->decimal('amount_received',8,2);
            //$table->decimal('change_amount',8,2);
            $table->set('transaction_type', ['sale', 'purchase']);
            $table->boolean('is_deleted')->default(0);
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        Schema::dropIfExists('transactions');
    }
};
