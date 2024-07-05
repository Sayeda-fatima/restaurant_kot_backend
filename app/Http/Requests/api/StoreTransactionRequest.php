<?php

namespace App\Http\Requests\api;

use Illuminate\Foundation\Http\FormRequest;

class StoreTransactionRequest extends FormRequest
{
    /**
     * Determine if the user is authorized to make this request.
     */
    public function authorize(): bool
    {
        return false;
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array<string, \Illuminate\Contracts\Validation\ValidationRule|array<mixed>|string>
     */
    public function rules(): array
    {
        return [
            //
            'transaction_id' =>'required|numeric',
            'created_on' => 'required|datetime',
            'customer_name' => 'required|string|max:255',
            'product_quantity' => 'required|integer|min:1',
            'total_price' => 'required|decimal',
            'mode_of_payment' => 'required|enum|[cash, cheque, bank]',
            'amount_received' => 'required|decimal|min:0',
            'change_amount' => 'required|decimal|min:0'
        ];
    }
}
