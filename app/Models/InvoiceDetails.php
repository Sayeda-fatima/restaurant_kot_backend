<?php

namespace App\Models;

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
        'quantity',
        'unit_product_price',
        'total_product_price',
    ];

    public function product(): HasMany {
        return $this->hasMany(Invoice::class);
    }

    public function invoice(): BelongsTo {
        return $this->belongsTo(Invoice::class);
    }
}
