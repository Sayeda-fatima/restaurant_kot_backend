<?php

namespace App\Http\Requests\api;

use Illuminate\Foundation\Http\FormRequest;

class StoreInvoiceRequest extends FormRequest
{
    /**
     * Determine if the user is authorized to make this request.
     */
    public function authorize(): bool
    {
        return true;
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array<string, \Illuminate\Contracts\Validation\ValidationRule|array<mixed>|string>
     */
    public function rules(): array
    {
        return [
            'customer_id' => 'required|integer|exists:customers,id',
            //'customer_name' => 'string|max:255',
            //'total_price' => 'required|decimal:0,2|min:0',
            //'customer_billing_address' => 'required|string|max:255',
            'mode_of_payment' => 'required|in:bank,cash,cheque',
        ];
    }
}
