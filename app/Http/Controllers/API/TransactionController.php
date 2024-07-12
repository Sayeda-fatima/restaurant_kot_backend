<?php

namespace App\Http\Controllers\API;

use Illuminate\Support\Facades\DB;
use Illuminate\Http\Request;
use App\Models\Transaction;
use App\Http\Requests\api\StoreTransactionRequest;
use App\Http\Requests\api\UpdateTransactionRequest;
use App\Http\Controllers\Controller;

class TransactionController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        //display transactions
        $transactions = DB::table('transactions')
                            ->select('customer_name', 'id', 'total_price', 'mode_of_payment', 'created_at')
                            ->orderby('created_at')
                            ->get();
        
        return response()->json([
            'message' => 'success',
            'transactions'=>$transactions
        ]);
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
    public function store(StoreTransactionRequest $request)
    {
        try{
            $transaction = Transaction::create([
                'total_price' => $request->total_price,
                'amount_received' => $request->amount_received,
                'change_amount' => ($request->total_price - $request->amount_received)
            ]);
            return response()->json([
                'message' => 'success',
                'data' => $transaction
            ]);
        }catch (\Exception $e) {
            error_log('Error creating transaction: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to create transaction', 'error' => $e->getMessage()], 500);
        }

    }

    /**
     * Display the specified resource.
     */
    public function show(Transaction $transaction)
    {
        //
    }

    /**
     * Show the form for editing the specified resource.
     */
    public function edit(Transaction $transaction)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(UpdateTransactionRequest $request, Transaction $transaction)
    {
        try{
            $transaction = Transaction::update([
                'total_price' => $request->total_price,
                'amount_received' => $request->amount_received,
                'change_amount' => ($request->total_price - $request->amount_received)
            ]);
            return response()->json([
                'message' => 'success',
                'data' => $transaction
            ]);
        }catch (\Exception $e) {
            error_log('Error updating transaction: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to update transaction', 'error' => $e->getMessage()], 500);
        }
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(Transaction $transaction)
    {
        try{
            $transaction -> delete();
            return response()->json([
                'message' => 'success',
                'data' => $transaction
            ],200);
        }
        catch(\Exception $e){
            error_log('Error deleting transaction: ' . $e->getMessage());

            return response()->json(['message' => 'Failed to delete transaction', 'error' => $e->getMessage()], 500);
        }
    }
    // transactions -> money out report
    public function moneyOutReport(Request $request){
        $date_from = $request->date_from;
        $date_to = $request->date_to;

        $query = DB::select('SELECT date(created_at), name, mode_of_payment, id, total_price from transactions where transaction_type=purchase and date(created_at) between ? and ?',[$date_from, $date_to]);

        return response()->json([
            'message' => 'success',
            'data' => $query
        ]);
    }
     // item report -> item report
     public function productReport(Request $request){
        $product_id = $request->product_id;
        $date_from = $request->date_from;
        $date_to = $request->date_to;

        $query = DB::select('SELECT date(transactions.created_at) as date, 
                transactions.transaction_type, 
                transactions.id as transaction_id, 
                customers.customer_name, 
                transactions.product_name, 
                transactions.product_quantity, 
                transactions.product_price, 
                products.purchase_price, 
                transactions.total_price 
            from transactions 
            left join products on transactions.product_id=products.id 
            left join customers on transactions.customer_id=customers.id
            where product_id=? and date(transactions.created_at) between ? and ?', [$product_id, $date_from, $date_to]);

        return response()->json([
            'message' => 'success',
            'data' => $query
        ]);
    }
}
