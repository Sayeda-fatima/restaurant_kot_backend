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
        Schema::create('products', function (Blueprint $table) {
            $table->id();
            $table->unsignedBigInteger('organization_id');
            $table->foreign('organization_id')->references('id')->on('organizations');
            $table->string('image')->nullable();
            $table->string('name')->required;
            $table->decimal('sell_price', 8, 2)->required;
            $table->string('measuring_unit');
            $table->string('category');
            $table->foreign('category')->references('product_category')->on('product_categories');
            $table->integer('quantity')->required;
            $table->decimal('mrp', 8,2)->required;
            $table->decimal('purchase_price',8,2);
            $table->decimal('ac_sale_price',8,2);
            $table->decimal('non_ac_sale_price',8,2);
            $table->decimal('online_delivery_sell_price');
            $table->decimal('online_sell_price');
            // gst and tax (optional)
            $table->enum('tax', ['Non-GST-Supplies', 'Exempted', 'GST@0%', 'GST@0.25%', 'GST@1%', 'GST@3%', 'GST@5%', 'GST@12%', 'GST@18%', 'GST@28%'])->nullable();
            $table->enum('price_with_tax', ['Y', 'N'])->nullable();
            $table->integer('cess');
            $table->string('hsn_code');
            $table->text('description');
            //inventory details (optional)
            $table->integer('low_stock_alert');
            $table->string('storage_location');
            $table->string('bulk_purchase_unit');
            $table->decimal('retail_sale_unit_per_bulk_purchase');
            $table->decimal('bulk_purchase_unit_per_retail_sale');
            $table->date('expiry_date');
            // product display (optional)
            $table->enum('show_product_online_store', ['Yes', 'No']);
            $table->boolean('is_deleted')->default(0);
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        Schema::dropIfExists('products');
    }
};
