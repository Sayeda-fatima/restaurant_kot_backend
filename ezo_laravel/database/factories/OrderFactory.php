<?php

namespace Database\Factories;

use App\Models\User;
use App\Models\Organization;
use App\Models\Customer;
use Illuminate\Database\Eloquent\Factories\Factory;

/**
 * @extends \Illuminate\Database\Eloquent\Factories\Factory<\App\Models\Order>
 */
class OrderFactory extends Factory
{
    /**
     * Define the model's default state.
     *
     * @return array<string, mixed>
     */
    public function definition(): array
    {
        return [
            'organization_id' => Organization::factory("id"),
            'customer_id' => Customer::factory(),
            'total_price' => $this->faker->numberBetween('100', '30000000'),
            'customer_billing_address' => Customer::factory("customer_billing_address"),
            'mode_of_payment' => 'cash',//$this->faker->words('bank', 'cash'),
            'created_by' => User::factory()
        ];
    }
}
