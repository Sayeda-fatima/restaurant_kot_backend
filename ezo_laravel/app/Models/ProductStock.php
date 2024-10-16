<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class ProductStock extends Model
{
    use HasFactory;

    protected $fillable =[
        'organization_id',
        'invoice_details_id',
        'product_id',
        'product_name',
        'product_stock_before_update',
        'product_update_quantity',
        'product_update_type',
        'product_stock_after_update'
    ];
    
    public function product(){
        return $this->belongsTo(Product::class);
    }
}
