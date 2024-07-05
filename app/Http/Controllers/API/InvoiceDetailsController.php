<?php

namespace App\Http\Controllers\API;

use Illuminate\Support\Facades\DB;
use App\Models\InvoiceDetails;
use App\Http\Requests\api\StoreInvoiceDetailsRequest;
use App\Http\Requests\api\UpdateInvoiceDetailsRequest;
use App\Http\Controllers\Controller;

class InvoiceDetailsController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        $invoiceDetails = DB::table('invoice_details')
                            ->join('products', 'invoice_details.product_id', '=', 'products.id')
                            ->select('invoice_details.invoice_id', 'products.product_name', 'invoice_details.quantity', 'products.product_sell_price'); //select sum(total_product_price) as total 
        
        $invoice = DB::raw('select invoice_details.invoice_id, products.product_name, invoice_details.quantity, products.product_sell_price from invoice_details join products on invoice_details.product_id=products.id');
    }

    /**
     * Show the form for creating a new resource.
     */
    public function create()
    {
        //
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(StoreInvoiceDetailsRequest $request)
    {
        //
    }

    /**
     * Display the specified resource.
     */
    public function show(InvoiceDetails $invoiceDetails)
    {
        //
    }

    /**
     * Show the form for editing the specified resource.
     */
    public function edit(InvoiceDetails $invoiceDetails)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(UpdateInvoiceDetailsRequest $request, InvoiceDetails $invoiceDetails)
    {
        //
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(InvoiceDetails $invoiceDetails)
    {
        //
    }
}
