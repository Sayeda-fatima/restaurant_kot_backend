<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Cart extends Model
{
    use HasFactory;

    protected $fillable =[
        'organization_id',
        'customer_id',
        'total_quantity',
    ];

    public function cartItem(){
        return $this->hasMany(CartItem::class);
    }
}
