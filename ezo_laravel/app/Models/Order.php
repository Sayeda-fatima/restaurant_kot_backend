<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Order extends Model
{
    use HasFactory;

    protected $fillable = [
        'organization_id',
        'customer_id',
        'total_price',
        'customer_billing_address',
        'mode_of_payment',
    ];

    public function orderItem(){
        return $this->hasMany(OrderItem::class);
    }
}
