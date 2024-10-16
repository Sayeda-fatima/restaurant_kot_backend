<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\HasMany;
use Illuminate\Database\Eloquent\Relations\HasManyThrough;
use Illuminate\Database\Eloquent\Relations\HasOne;

class Product extends Model
{
    use HasFactory;

    protected $fillable = [
        'organization_id',
        'name',
        'image',
        'sell_price',
        'measuring_unit',
        'category',
        'quantity',
        'mrp',
        'purchase_price',
        'ac_sale_price',
        'non_ac_sale_price',
        'online_delivery_sell_price',
        'online_sell_price',
        'tax',
        'price_with_tax',
        'cess',
        'hsn_code',
        'description',
        'low_stock_alert',
        'storage_location',
        'bulk_purchase_unit',
        'retail_sale_unit_per_bulk_purchase',
        'bulk_purchase_unit_per_retail_sale',
        'expiry_date',
        'show_product_online_store'
    ];

    public function productImage(): HasMany{
        return $this->hasMany(ProductImage::class);
    }
    public function productStock(): HasMany {
        return $this->hasMany(ProductStock::class);
    }
    public function invoice(): HasManyThrough{
        return $this->hasManyThrough(Invoice::class, InvoiceDetails::class);
    }

    public function productCategory(): BelongsTo{
        return $this->belongsTo(ProductCategory::class);
    }

    public function invoiceDetails(){
        return $this->belongsToMany(InvoiceDetails::class);
    }
}
