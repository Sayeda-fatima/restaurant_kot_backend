<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Transaction extends Model
{
    use HasFactory;

    protected $fillable = [
        'organization_id',
        'customer_id',
        'customer_name',
        'product_id',
        'product_name',
        'product_quantity',
        'product_price',
        'total_price',
        'transaction_type'
    ];

    public function customer(){
        return $this->belongsTo(Customer::class);
    }

    public function supplier(){
        return $this->belongsTo(Supplier::class);
    }
}
