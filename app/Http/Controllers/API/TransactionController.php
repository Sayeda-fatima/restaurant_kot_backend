<?php

namespace App\Http\Controllers\API;

use Illuminate\Support\Facades\DB;
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
                            ->select('customer_name', 'transaction_id', 'total_price', 'mode_of_payment', 'created_at')
                            ->orderby('created_at')
                            ->orderby('product_category')
                            ->get()
                            ->paginate(25);
        
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
        //
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
        //
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(Transaction $transaction)
    {
        //
    }
}
