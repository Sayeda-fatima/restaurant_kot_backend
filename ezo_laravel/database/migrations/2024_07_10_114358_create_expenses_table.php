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
        Schema::create('expenses', function (Blueprint $table) {
            $table->id();
            $table->unsignedBigInteger('supplier_id');
            $table->foreign('supplier_id')->references('id')->on('suppliers');
            $table->string('supplier_name');
            $table->string('expense_category');
            $table->decimal('total_amount', 8,2);
            $table->decimal('amount_paid',8,2);
            $table->decimal('amount_due',8,2);
            $table->text('Note');
            $table->set('mode_of_payment', ['bank', 'cash', 'cheque']);
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        Schema::dropIfExists('expenses');
    }
};
