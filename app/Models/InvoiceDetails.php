<?php

namespace App\Models;

use Illuminate\Support\Facades\DB;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\HasMany;

class InvoiceDetails extends Model
{
    use HasFactory;

    protected $fillable = [
        'invoice_id',
        'product_id',
        'product_name',
        'quantity',
        'unit_product_price',
        'total_product_price'
    ];

    public function product(): BelongsTo {
        return $this->belongsTo(Product::class);
    }

    public function invoice(): BelongsTo {
        return $this->belongsTo(Invoice::class);
    }
    
}
