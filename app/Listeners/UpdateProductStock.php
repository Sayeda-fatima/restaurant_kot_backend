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
        if($event->status == "PENDING"){
            //$this->updateStatus($event->invoice, $event->status);

        }
        if($event->status == "SUCCESS"){
            $this->updateStatus($event->invoice, $event->status);
            $this->updateStock($event);
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

    private function updateStock(NewInvoice $event){
        /*if($invoice->status == 'SUCCESS'){
            $product->product_quantity -= $invoiceDetails->quantity;
            $productStock->product_update_quantity = $invoiceDetails->quantity;
            $productStock->product_stock_after_update -= $invoiceDetails->quantity; 
        } */
        $invoice = $event->invoice;

        foreach ($invoice->invoiceDetails as $detail) {
            $product = Product::find($detail->product_id);

            if ($product && $product->product_quantity >= $detail->quantity) {
                $product->product_quantity -= $detail->quantity;
                $product->save();

                // Record stock change
                ProductStock::create([
                    'product_id' => $product->id,
                    'product_name' => $product->product_name,
                    'product_update_quantity' => $detail->quantity,
                    'product_update_type' => 'sale',
                    'product_stock_after_update' => $product->product_quantity - $detail->quantity
                ]);
            } else {
                // Handle insufficient stock or product not found

            }
        }
    }
}
