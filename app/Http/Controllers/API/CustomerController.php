<?php

namespace App\Http\Controllers\API;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Facades\Redirect;
use App\Models\Customer;
use App\Http\Requests\api\StoreCustomerRequest;
use App\Http\Requests\api\UpdateCustomerRequest;
use App\Http\Controllers\Controller;

class CustomerController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        // display list of customers
        $customers = DB::table('customers')
                    ->select('customer_name', 'customer_phone_no', 'customer_billing_type')
                    ->orderBy('customer_category', 'asc')
                    ->orderBy('customer_name')
                    ->get();
        
        //return view('customers.index', ['customer'=>$customers]);
        return response()->json([
            'message' => 'success',
            'data' => $customers
        ],201);
    }

    /**
     * Show the form for creating a new resource.
     */
    public function create()
    {
        //
        return view('customer.create');
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(StoreCustomerRequest $request)
    {
        try{
            $customer = Customer::create([
                'customer_name' => $request->input('customer_name'),
                'customer_phone_no' => $request->input('customer_phone_no'),
                'customer_category' => $request->input('customer_category'),
                'customer_billing_address' => $request->input('customer_billing_address'),
                'customer_billing_province' => $request->input('customer_billing_province'),
                'customer_billing_postal_code' => $request->input('customer_billing_postal_code'),
                'customer_delivery_address' => $request->input('customer_delivery_address'),
                'customer_delivery_province' => $request->input('customer_delivery_province'),
                'customer_delivery_postal_code' => $request->input('customer_delivery_postal_code'),
                'customer_gst_number' => $request->input('customer_gst_number'),
                'customer_billing_term' => $request->input('customer_billing_term'),
                'customer_billing_type' => $request->input('customer_billing_type'),
                'customer_date_of_birth' => $request->input('customer_date_of_birth'),
                'whatsapp_alert' => $request->input('whatsapp_alert')
            ]);
            return response()->json([
                'message' => 'success',
                'data' => $customer
            ],201);
        }
        catch(\Exception $e){
            error_log('Error adding customer: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to add customer', 'error' => $e->getMessage()], 500);
        }
        

        //return Redirect::route('customer.index')->with('success', 'Customer Added Successfully');
        
    }

    /**
     * Display the specified resource.
     */
    public function show(Customer $customer)
    {
        //request customer transactions
        
    }

    /**
     * Show the form for editing the specified resource.
     */
    public function edit(Customer $customer)
    {
        //edit customer details 
        return view('customer.edit', ['customer'=> $customer]);
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(UpdateCustomerRequest $request, Customer $customer)
    {
        try{
            $data = $request->all();
            $customer -> update($data);

            //return Redirect::route('customer.index')->with('success', 'Customer updated successfully');
            return response()->json([
                'message' => 'success',
                'data' => $customer->fresh()
            ],200);
        }

        catch(\Exception $e){
            error_log('Error updating customer: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to update customer', 'error' => $e->getMessage()], 500);
        }
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(Customer $customer)
    {
        try{
            $customer->delete();
            return response()->json([
                'message' => 'success',
                'data' => $customer
            ],200);
        }
        catch(\Exception $e){
            error_log('Error deleting customer: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to delete customer', 'error' => $e->getMessage()], 500);
        }
        //return Redirect::route('customer.index')->with('success', 'Customer deleted successfully');
        
    }

    public function allCustomers(){
        // display customers in for invoice
        $customer = Customer::orderBy('id', 'DESC')
                            ->groupby('customer_category')
                            ->get();
        
        return response()->json([
            'customer' => $customer
        ],200);
    }

    public function searchCustomer(Request $request){
        $search = $request->get('search_term');
        if($search!=NULL){
            $customer = Customer::where('customer_id', 'LIKE', "%$search%")
                                ->orWhere('customer_name', 'LIKE', "%$search%")
                                ->orWhere('customer_phone_no', 'LIKE', "$search%")
                                ->get();
            return response()->json([
                'data' => $customer,
                'message' => 'Customer found!'
            ],200);
        }
        return response()->json([
            'message' => 'Customer not found'
        ],404);
    }

    // customer report -> customer details report
    public function detailReport(Request $request){
        $date_from = $request->date_from;
        $date_to = $request->date_to;

        $query = DB::select('SELECT customer_name, customer_phone_no, customer_category, customer_billing_address, customer_billing_type,customer_billing_province, customer_billing_postal_code, customer_delivery_address, customer_delivery_province, customer_delivery_postal_code, customer_billing_term, customer_date_of_birth from customers where date(created_at) between ? and ?;', [$date_from, $date_to]);

        return response()->json([
            'message' => 'success',
            'data' => $query
        ],200);
    }

    // customer report -> customer ledger report
    public function invoiceReport(Request $request){
        $date_from = $request->date_from;
        $date_to = $request->date_to;
        $customer_id = $request->customer_id;
        $query = DB::select('SELECT date(invoices.created_at) as date, invoices.id as invoice_no, invoices.total_price as credit from invoices where customer_id=? and date(invoices.created_at) between ? and ?
        group by invoices.id', [$customer_id, $date_from, $date_to]);

        return response()->json([
            'message' => 'success',
            'data' => $query
        ]);
    }

    // customer report -> customer invoice ledger report
    public function invoiceDetailReport(Request $request){
        $date_from = $request->date_from;
        $date_to = $request->date_to;
        $customer_id = $request->customer_id;
        $invoice_id = $request->invoice_id;

        $query = DB::select('SELECT invoices.id, 
                invoice_details.product_name, 
                invoice_details.quantity, 
                invoice_details.unit_product_price, 
                invoice_details.total_product_price 
            from invoices right join invoice_details on invoices.id=invoice_details.invoice_id 
            where invoices.customer_id=? and invoice_details.invoice_id=? and date(invoices.created_at) between ? and ?', [$customer_id, $invoice_id, $date_from, $date_to]);

        return response()->json([
            'message' => 'success',
            'data' => $query
        ]);
    }
}
