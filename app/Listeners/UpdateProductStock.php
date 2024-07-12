<?php

namespace App\Listeners;

use App\Events\NewInvoice;
use App\Models\Invoice;
use App\Models\InvoiceDetails;
use App\Models\ProductStock;
use App\Models\Product;
use Illuminate\Contracts\Queue\ShouldQueue;
use Illuminate\Queue\InteractsWithQueue;

class UpdateProductStock
{
    /**
     * Create the event listener.
     */
    public function __construct()
    {
        //
    }

    /**
     * Handle the event.
     */
    public function handle(NewInvoice $event): void
    {
        if($event == "PENDING"){
            //$this->updateStatus($event->invoice, $event->status);

        }
        if($event->status == "SUCCESS"){
            $invoiceId = $event->invoice->id;
            $product = $event->invoiceDetails->product_id;
        //updating the inventory for each item purchased
       // $productStock = $event->
        }
        
    }

    public function updateStatus(Invoice $invoice, string $status){
        try{
            $invoice->update(['status'=>$status]);
            return response()->json([
                'message' => 'success',
                'data' => $invoice->fresh()
            ],200);
        }
        catch(\Exception $e){
            error_log('Error updating invoice status: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to update invoice status', 'error' => $e->getMessage()], 500);
        }
    }

    private function updateStock(Invoice $invoice, InvoiceDetails $invoiceDetails, Product $product, ProductStock $productStock){
        if($invoice->status == 'SUCCESS'){
            $product->product_quantity -= $invoiceDetails->quantity;
            $productStock->product_update_quantity = $invoiceDetails->quantity;
            $productStock->product_stock_after_update -= $invoiceDetails->quantity;
        }
    }
}
