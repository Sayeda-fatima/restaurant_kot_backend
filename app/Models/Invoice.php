<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\HasMany;

class Invoice extends Model
{
    use HasFactory;
    protected $fillable =[
        'customer_id',
        'customer_name',
        'total_price',
        'billing_address',
        'mode_of_payment'
    ];
    public function customer(){
        return $this->belongsTo(Customer::class);
    }

    public function product(){
        return $this->belongsTo(Product::class);
    }
    
    public function invoiceDetails(){
        return $this->hasMany(InvoiceDetails::class);
    }
}
