<?php

namespace App\Events;

use App\Models\Invoice;
use App\Models\InvoiceDetails;
use App\Models\Product;
use App\Models\ProductStock;
use Illuminate\Broadcasting\Channel;
use Illuminate\Broadcasting\InteractsWithSockets;
use Illuminate\Broadcasting\PresenceChannel;
use Illuminate\Broadcasting\PrivateChannel;
use Illuminate\Contracts\Broadcasting\ShouldBroadcast;
use Illuminate\Foundation\Events\Dispatchable;
use Illuminate\Queue\SerializesModels;

class NewInvoice
{
    use Dispatchable, InteractsWithSockets, SerializesModels;
    public Invoice $invoice;
    public InvoiceDetails $invoiceDetails;
    public Product $product;
    public ProductStock $productStock;
    public string $status;

    /**
     * Create a new event instance.
     */
    public function __construct($invoice, $invoiceDetails, $product, $productStock, $status)
    {
        $this->invoice = $invoice;
        $this->invoiceDetails = $invoiceDetails;
        $this->product = $product;
        $this->productStock = $productStock;
        $this->status = $status;
    }

    /**
     * Get the channels the event should broadcast on.
     *
     * @return array<int, \Illuminate\Broadcasting\Channel>
     */
    public function broadcastOn(): array
    {
        return [
            new PrivateChannel('channel-name'),
        ];
    }
}
