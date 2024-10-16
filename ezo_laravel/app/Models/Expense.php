<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Expense extends Model
{
    use HasFactory;

    protected $fillable = [
        'organization_id',
        'supplier_id',
        'supplier_name',
        'expense_category',
        'total_amount',
        'amount_paid',
        'amount_due'
    ];
}
