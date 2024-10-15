<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Customer extends Model
{
    use HasFactory;

    protected $fillable =[
        'customer_name',
        'customer_phone_no',
        'customer_category',
        'customer_billing_address',
        'customer_billing_province',
        'customer_billing_postal_code',
        'customer_delivery_address',
        'customer_delivery_province',
        'customer_delivery_postal_code',
        'customer_gst_number',
        'customer_billing_term',
        'customer_billing_type',
        'customer_date_of_birth',
        'whatsapp_alert'
    ];
    public function invoice(){
        return $this->hasMany(Invoice::class);
    }
    public function transaction(){
        return $this->hasMany(Transaction::class);
    }
}
