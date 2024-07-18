<?php

namespace App\Http\Controllers\API;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Facades\Gate;
use Illuminate\Support\Collection;
use App\Models\Invoice;
use App\Models\Customer;
use App\Http\Requests\api\StoreInvoiceRequest;
use App\Http\Requests\api\UpdateInvoiceRequest;
use App\Http\Requests\api\StoreInvoiceDetailsRequest;
use App\Http\Controllers\Controller;
use App\Events\NewInvoice;

class InvoiceController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        Gate::authorize('viewAny', Invoice::class);
        $invoice = DB::table('invoices')
                    ->select('id', 'customer_id', 'total_price', 'mode_of_payment')
                    ->get()
                    ->paginate(25);
        return response()->json([
            'message' => 'success',
            'data' => $invoice
        ]);
    }

    /**
     * Show the form for creating a new resource.
     */
    public function create()
    {
        return view('invoice.create');
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(StoreInvoiceRequest $request)
    {
        Gate::authorize('create', Invoice::class);
        $customer = Customer::find($request->customer_id);
        try{
            $invoice = Invoice::create([
                'customer_id' => $request->customer_id,
                'customer_name' => $customer->customer_name,
                //'total_price' => $request->total_price,
                //'total_price' => DB::raw('Select sum(total_product_price) from invoice_details group by '. $request->order_id),
                'customer_billing_address' => $customer->customer_billing_address,
                'mode_of_payment' => $request->mode_of_payment,
                'created_by' => $request->created_by
            ]);
            // possible way to update price
            //$invoice->updateTotalPrice();
            event(new NewInvoice($invoice, $invoice->status));

            return response()->json([
                'message' => 'success',
                'data' => $invoice
            ],201);
        }
        catch(\Exception $e){
            error_log('Error creating invoice: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to create invoice', 'error' => $e->getMessage()], 500);
        }
        //return Redirect::route('invoice.index')->with('success', 'Invoice generated successfully');
    }

    /**
     * Display the specified resource.
     */
    public function show(Invoice $invoice)
    {
        //
    }

    /**
     * Show the form for editing the specified resource.
     */
    public function edit(Invoice $invoice)
    {
        return view('invoice.edit');
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(UpdateInvoiceRequest $request, Invoice $invoice)
    {
        // specify access type -> ADMIN, STAFF, SALES
        Gate::authorize('update', $invoice);
        try{
            $customer = Customer::find($request->customer_id);
            $invoice->update([
                'customer_id' => $request->customer_id,
                'customer_name' => $customer->customer_name,
                'customer_billing_address' => $customer->customer_billing_address,
                'mode_of_payment' => $request->mode_of_payment
            ]);
            return response()->json([
                'message' => 'success',
                'data' => $invoice->fresh()
            ],200);
        }
        catch(\Exception $e){
            error_log('Error updating invoice: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to update invoice', 'error' => $e->getMessage()], 500);
        }
        
        //return Redirect::route('invoice.index')->with('success', 'Invoice edited successfully');
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(Invoice $invoice)
    {
        // specify access type -> ADMIN, STAFF, SALES
        Gate::authorize('update', $invoice); 
        try{
            $invoice -> delete();
            return response()->json([
                'message' => 'success',
                'data' => $invoice
            ],200);
        }
        catch(\Exception $e){
            error_log('Error deleting invoice: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to delete invoice', 'error' => $e->getMessage()], 500);
        }
        
    }

    // transaction report -> sale report
    public function saleReport(Request $request){
        // specify access type -> ADMIN
        Gate::authorize('view', Invoice::class);

        $date_from = $request->date_from;
        $date_to = $request->date_to;

        $query = DB::select('SELECT date(invoices.created_at) as date, 
            invoices.id as invoice_no, 
            invoices.customer_name, 
            customers.customer_phone_no, 
            invoices.total_price, 
            (select sum(invoice_details.quantity) from invoice_details where invoice_details.invoice_id=invoices.id group by invoice_details.invoice_id) as total_quantity
            from invoices 
            left join customers on invoices.customer_id = customers.id 
            where date(invoices.created_at) between ? and ?', [$date_from, $date_to]);

        $total_sales = DB::select('SELECT count(id) as total_sales from invoices where date(invoices.created_at) between ? and ?', [$date_from, $date_to]);

        $total_sales_quantity = DB::select('SELECT sum(quantity) as total_sales_quantity from invoice_details where date(created_at) between ? and ?', [$date_from, $date_to]);

        $total_sales_amount = DB::select('SELECT sum(total_price) as total_sales_amount from invoices where date(created_at) between ? and ?', [$date_from, $date_to]);
        
        return response()->json([
            'message' => 'success',
            'data' => $query,
            'total_sales' => $total_sales,
            'total_sales_quantity' => $total_sales_quantity,
            'total_sales_amount' => $total_sales_amount
        ],200);
    }

    // transaction report -> sale wise profit and loss statement
    public function saleProfitReport(Request $request){
        // specify access type -> ADMIN
        Gate::authorize('view', Invoice::class);

        $date_from = $request->date_from;
        $date_to = $request->date_to;

        $query = DB::select('SELECT date(i.created_at) as date, 
                   i.id as invoice_no, 
                   i.customer_name, 
                   i.total_price,
                   SUM(p.purchase_price * id.quantity) AS purchase_price, 
                   (i.total_price - SUM(p.purchase_price * id.quantity)) AS profit
            FROM invoices AS i
            LEFT JOIN invoice_details AS id ON id.invoice_id = i.id
            LEFT JOIN products AS p ON p.id = id.product_id
            WHERE date(i.created_at) BETWEEN ? AND ?
            GROUP BY i.id;', [$date_from, $date_to]);

        $total_data = DB::select('SELECT SUM(i.total_price) as total_sales_amount,
            SUM(p.purchase_price * id.quantity) AS total_purchase_price, 
            (SUM(i.total_price) - SUM(p.purchase_price * id.quantity)) AS total_profit
            FROM invoices AS i
            LEFT JOIN invoice_details AS id ON id.invoice_id = i.id
            LEFT JOIN products AS p ON p.id = id.product_id
            WHERE date(i.created_at) BETWEEN ? AND ?;', [$date_from, $date_to]);

        return response()->json([
            'message' => 'success',
            'data' => $query,
            'total_data' => $total_data
        ],200);
    }

    // transaction report -> money in report
    public function moneyInReport(Request $request){
        // specify access type -> ADMIN
        Gate::authorize('view', Invoice::class);

        $date_from = $request->date_from;
        $date_to = $request->date_to;

        $query = DB::select('SELECT date(created_at) as date, 
            customer_name, 
            mode_of_payment, 
            total_price, 
            id as sale_invoice
            from invoices
            where date(created_at) between ? and ?', [$date_from, $date_to]);
        
        $total_data = DB::select('SELECT count(id) as total_moneyIn, sum(total_price) as total_moneyIn_amount from invoices where date(created_at) between ? and ?', [$date_from, $date_to]);
        
        return response()->json([
            'message' => 'success',
            'data' => $query,
            'total_data' => $total_data
        ],200);
    }

    public function endDayReport(Request $request){
        // specify access type -> ADMIN
        Gate::authorize('view', Invoice::class);
        
        $date_from = $request->date_from;
        $date_to = $request->date_to;

        $data['sale'] = DB::select('SELECT sum(total_price) as sale from invoices where date(created_at) between ? and ?', [$date_from, $date_to]);

        $data['moneyIn'] = array_values(DB::select('SELECT sum(total_price) as moneyIn from invoices where date(created_at) between ? and ?', [$date_from, $date_to]));

        $data['moneyOut'] = DB::select("SELECT sum(total_price) as moneyOut from transactions where transaction_type='purchase' and date(created_at) between ? and ?",[$date_from, $date_to]);

        $data['purchase'] = DB::select("SELECT sum(total_price) as purchase from transactions where transaction_type='purchase' and date(created_at) between ? and ?",[$date_from, $date_to]);

        return response()->json([
            'message' => 'success',
            'data' => $data
        ]);
    }
}
