<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Supplier extends Model
{
    use HasFactory;

    protected $fillable = [
        'organization_id',
        'name',
        'phone_no',
        'category',
        'billing_address',
        'billing_province',
        'billing_postal_code',
        'delivery_address',
        'delivery_province',
        'delivery_postal_code',
        'gst_number',
        'billing_term',
        'billing_type',
        'date_of_birth',
        'whatsapp_alert'
    ];

    public function invoice(){
        return $this->hasMany(Invoice::class);
    }
    public function transaction(){
        return $this->hasMany(Transaction::class);
    }
}
