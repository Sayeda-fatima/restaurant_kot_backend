<?php

namespace App\Http\Controllers\API;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Facades\Gate;
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
        Gate::authorize('viewAny', Customer::class);
        try{
        // display list of customers
        $customers = DB::table('customers')
                    ->select('name', 'phone_no', 'billing_type')
                    ->orderBy('category', 'asc')
                    ->orderBy('name')
                    ->get();
        
        //return view('customers.index', ['customer'=>$customers]);
        return response()->json([
            'message' => 'success',
            'data' => $customers
        ],201);
        }catch(\Exception $e){
            error_log('Error displaying customers: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to display customers', 'error' => $e->getMessage()], 500);
        }

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
        Gate::authorize('create', Customer::class);
        try{
            $customer = Customer::create([
                'organization_id' => $request->input('organization_id'),
                'name' => $request->input('name'),
                'phone_no' => $request->input('phone_no'),
                'category' => $request->input('category'),
                'billing_address' => $request->input('billing_address'),
                'billing_province' => $request->input('billing_province'),
                'billing_postal_code' => $request->input('billing_postal_code'),
                'delivery_address' => $request->input('delivery_address'),
                'delivery_province' => $request->input('delivery_province'),
                'delivery_postal_code' => $request->input('delivery_postal_code'),
                'gst_number' => $request->input('gst_number'),
                'billing_term' => $request->input('billing_term'),
                'billing_type' => $request->input('billing_type'),
                'date_of_birth' => $request->input('date_of_birth'),
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
        Gate::authorize('update', $customer);
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
        Gate::authorize('delete', $customer);
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

    public function allCustomers(Request $request){
        // display customers in for invoice
        $organization_id = $request->organization_id;
        $customer = Customer::select('name', 'phone_no')
                            ->whereRaw('organization_id=?',[$organization_id] )
                            ->orderBy('id', 'DESC')
                            ->orderby('category')
                            ->get();
        
        return response()->json([
            'customer' => $customer
        ],200);
    }

    public function searchCustomer(Request $request){
        $organization_id = $request->organization_id;
        $search = $request->get('search_term');
        if($search!=NULL){
            $customer = Customer::where('id', 'LIKE', "%$search%")
                                ->orWhere('name', 'LIKE', "%$search%")
                                ->orWhere('phone_no', 'LIKE', "$search%")
                                ->havingRaw('organization_id=?',[$organization_id])
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

        Gate::authorize('view', Customer::class);
        $organization_id = $request->organization_id;
        $date_from = $request->date_from;
        $date_to = $request->date_to;

        $query = DB::select('SELECT name, phone_no, category, billing_address, billing_type, billing_province, billing_postal_code, delivery_address, delivery_province, delivery_postal_code, billing_term, date_of_birth from customers where organization_id=? and date(created_at) between ? and ?;', [$organization_id, $date_from, $date_to]);

        return response()->json([
            'message' => 'success',
            'data' => $query
        ],200);
    }

    // customer report -> customer ledger report
    public function invoiceReport(Request $request){
        // specify access
        Gate::authorize('view', Customer::class);

        $organization_id = $request->organization_id;
        $date_from = $request->date_from;
        $date_to = $request->date_to;
        $customer_id = $request->customer_id;
        $query = DB::select('SELECT date(invoices.created_at) as date, 
                invoices.id as invoice_no, 
                invoices.total_price as credit 
            from invoices 
            where organization_id=? and customer_id=? and date(invoices.created_at) between ? and ?
            group by invoices.id', [$organization_id, $customer_id, $date_from, $date_to]);

        return response()->json([
            'message' => 'success',
            'data' => $query
        ]);
    }

    // customer report -> customer invoice ledger report
    public function invoiceDetailReport(Request $request){
        // specify access type 
        Gate::authorize('view', Customer::class);
        
        $organization_id = $request->organization_id;
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
            where invoices.organization_id=? and invoices.customer_id=? and invoice_details.invoice_id=? and date(invoices.created_at) between ? and ?', [$organization_id, $customer_id, $invoice_id, $date_from, $date_to]);

        return response()->json([
            'message' => 'success',
            'data' => $query
        ]);
    }
}
