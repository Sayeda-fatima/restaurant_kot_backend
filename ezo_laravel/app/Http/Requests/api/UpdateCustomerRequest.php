<?php

namespace App\Http\Requests\api;

use Illuminate\Foundation\Http\FormRequest;

class UpdateCustomerRequest extends FormRequest
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
            'name' => 'required|string|max:255', // Name is required, string, and max 255 characters
            'phone_no' => 'required|string|max:20|unique:customers,phone_no', // Phone number is required, string, max 20 characters, and unique within the customers table
            'category' => 'nullable|string|max:255', // Category is optional, string, and max 255 characters
            'billing_address' => 'required|string|max:255', // Billing address is required, string, and max 255 characters
            'billing_province' => 'required|string|max:255', // Billing province is required, string, and max 255 characters
            'billing_postal_code' => 'required|string|max:255', // Billing postal code is required, string, and max 255 characters
            'delivery_address' => 'nullable|string|max:255', // Delivery address is optional, string, and max 255 characters
            'delivery_province' => 'nullable|string|max:255', // Delivery province is optional, string, and max 255 characters
            'delivery_postal_code' => 'nullable|string|max:255', // Delivery postal code is optional, string, and max 255 characters
            'gst_number' => 'nullable|string|max:255', // Gst number is optional, string, and max 255 characters 
            'billing_term' => 'nullable|string|max:255', // Billing term is optional, string, and max 255 characters
            'billing_type' => 'nullable|string|max:255', // Billing type is optional, string, and max 255 characters
            'date_of_birth' => 'nullable|date', // date of birth is optional, date
            'whatsapp_alert' => 'nullable|in:Y,N', // whatsapp alert is optional, enum with yes, no
        ];
    }
}
