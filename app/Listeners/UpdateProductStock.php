<?php

namespace App\Listeners;

use App\Events\NewInvoice;
use App\Models\Invoice;
use App\Models\InvoiceDetails;
use App\Models\ProductStock;
use App\Models\Product;
use App\Models\Transaction;
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
        
        $invoice = $event->invoice;

        foreach ($invoice->invoiceDetails as $detail) {
            $product = Product::find($detail->product_id);
            
            if ($product && $product->product_quantity >= $detail->quantity) {
                // Record stock change
                ProductStock::create([
                    'invoice_details_id' => $detail->id,
                    'product_id' => $detail->product_id,
                    'product_name' => $product->product_name,
                    'product_stock_before_update' => $product->product_quantity,
                    'product_update_quantity' => $detail->quantity,
                    'product_update_type' => 'sale',
                    'product_stock_after_update' => $product->product_quantity - $detail->quantity
                ]);

                // update the stock in product table as well
                    $product->product_quantity -= $detail->quantity;
                    $product->save();

                // store in transaction after successful invoice generation
                Transaction::create([
                    'name' => $invoice->customer_name,
                    'product_id' => $detail->product_id,
                    'type' => 'customer',
                    'customer_id' => $invoice->customer_id,
                    'product_name' => $product->product_name,
                    'product_quantity' => $detail->quantity,
                    'product_price' => $product->mrp,
                    'total_price' => ($product->mrp * $detail->quantity),
                    'mode_of_payment' => $invoice->mode_of_payment,
                    'transaction_type' => 'sale'
                ]);
            } else {
                // Handle insufficient stock or product not found
                return error_log('Error! Stock insufficient or product not found');

            }
        }
    }

    private function updateTransaction(NewInvoice $event){
        // store in transaction after successful invoice generation
    }
}
