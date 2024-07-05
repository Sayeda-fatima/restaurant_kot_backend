<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Transaction extends Model
{
    use HasFactory;

    public function customer(){
        return $this->belongsTo(Customer::class);
    }

    public function supplier(){
        return $this->belongsTo(Supplier::class);
    }
}
