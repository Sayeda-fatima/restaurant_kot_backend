<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class OrderItem extends Model
{
    use HasFactory;

    protected $fillable = [
        'organization_id',
        'order_id',
        'product_id',
        'product_quantity',
        'unit_product_price',
        'tax',
        'total_product_price'
    ];

    public function order(){
        return $this->belongsTo(Order::class);
    }
}
