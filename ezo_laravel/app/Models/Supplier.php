<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Supplier extends Model
{
    use HasFactory;

    protected $fillable = [
        'supplier_name',
        'supplier_phone_no',
        'supplier_category',
        'supplier_billing_address',
        'supplier_billing_province',
        'supplier_billing_postal_code',
        'supplier_delivery_address',
        'supplier_delivery_province',
        'supplier_delivery_postal_code',
        'supplier_gst_number',
        'supplier_billing_term',
        'supplier_billing_type',
        'supplier_date_of_birth',
        'supplier_whatsapp_alert'
    ];

    public function invoice(){
        return $this->hasMany(Invoice::class);
    }
    public function transaction(){
        return $this->hasMany(Transaction::class);
    }
}
