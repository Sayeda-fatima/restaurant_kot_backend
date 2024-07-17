<?php

namespace App\Models;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;
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
        'customer_billing_address',
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

    public function updateTotalPrice(){
        /*$totalPrice = 0;
         // Access invoice items through the relationship
         $invoiceItems = $this->invoiceDetails;
         print_r($invoiceItems);

         // Loop through invoice items and accumulate total
         foreach ($invoiceItems as $item) {
             $totalPrice += $item->total_product_price;//$item->quantity * $item->unit_product_price;
         }
        return $totalPrice; 
        //return $this->DB::raw();
        //return $this->invoiceDetails()->sum(DB::raw('quantity*unit_product_price'));
        $totalPrice = DB::raw('SELECT invoice_id, sum(total_product_price) as total_price from invoice_details group by invoice_id;
');
        //return $totalPrice; */
        $totalPrice= $this->invoiceDetails->sum('total_product_price');
        $this->total_amount = $totalPrice;
        $this->save();

    }
}
